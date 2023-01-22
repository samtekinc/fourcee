package main

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/awsclients"
	tfomConfig "github.com/sheacloud/tfom/internal/config"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
	"go.uber.org/zap"
)

var (
	TF_INSTALLATION_DIRECTORY = os.Getenv("TF_INSTALLATION_DIRECTORY")
	TF_WORKING_DIRECTORY      = os.Getenv("TF_WORKING_DIRECTORY")
	wasSuccessful             = false
)

func sendTaskFailure(ctx context.Context, sfnClient awsclients.StepFunctionsInterface, taskToken string) {
	if wasSuccessful {
		return
	}
	_, err := sfnClient.SendTaskFailure(ctx, &sfn.SendTaskFailureInput{
		TaskToken: &taskToken,
	})
	if err != nil {
		zap.L().Panic("unable to send task failure", zap.Error(err))
	}
}

func main() {
	taskToken := os.Getenv("TASK_TOKEN")
	if taskToken == "" {
		zap.L().Panic("TASK_TOKEN environment variable not set")
	}
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	sfnClient := sfn.NewFromConfig(cfg)
	defer sendTaskFailure(context.Background(), sfnClient, taskToken)

	logger, err := zap.NewProduction()
	if err != nil {
		panic("unable to create logger, " + err.Error())
	}
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	conf := tfomConfig.ConfigFromEnv()
	dbClient := conf.GetDatabaseClient(cfg)
	apiClient := conf.GetApiClient(cfg, dbClient)

	installationDirectory, err := terraform.NewTerraformInstallationDirectory(TF_INSTALLATION_DIRECTORY)
	if err != nil {
		panic("unable to create terraform installation directory, " + err.Error())
	}

	// get request type and ID from environment
	requestType := os.Getenv("REQUEST_TYPE")
	requestID := os.Getenv("REQUEST_ID")

	zap.L().Sugar().Infow("processing request", "requestType", requestType, "requestID", requestID)

	// get the request from the database
	switch requestType {
	case "plan":
		planRequest, err := apiClient.GetPlanExecutionRequest(ctx, requestID)
		if err != nil {
			zap.L().Panic("unable to get plan execution request", zap.Error(err))
		}
		// execute the plan
		err = runPlan(ctx, planRequest, apiClient, installationDirectory)
		if err != nil {
			zap.L().Panic("unable to run plan", zap.Error(err))
		}
	case "apply":
		applyRequest, err := apiClient.GetApplyExecutionRequest(ctx, requestID)
		if err != nil {
			zap.L().Panic("unable to get apply execution request", zap.Error(err))
		}
		// execute the apply
		err = runApply(ctx, applyRequest, apiClient, installationDirectory)
		if err != nil {
			zap.L().Panic("unable to run apply", zap.Error(err))
		}
	default:
		zap.L().Panic("invalid request type", zap.String("requestType", requestType))
	}

	wasSuccessful = true
	_, err = sfnClient.SendTaskSuccess(ctx, &sfn.SendTaskSuccessInput{
		TaskToken: &taskToken,
		Output:    aws.String("{}"),
	})
	if err != nil {
		zap.L().Panic("unable to send task success", zap.Error(err))
	}
}

func runPlan(ctx context.Context, request *models.PlanExecutionRequest, apiClient api.APIClientInterface, installDirectory *terraform.TerraformInstallationDirectory) error {
	// update the request to running
	newStatus := models.RequestStatusRunning
	request, err := apiClient.UpdatePlanExecutionRequest(ctx, request.PlanExecutionRequestId, &models.PlanExecutionRequestUpdate{
		Status: &newStatus,
	})
	if err != nil {
		return err
	}

	workingDirectory, err := terraform.NewWorkingDirectory(filepath.Join(TF_WORKING_DIRECTORY, request.PlanExecutionRequestId))
	if err != nil {
		return err
	}
	defer workingDirectory.DeleteDirectory()

	// install terraform
	executable, err := installDirectory.InstallTerraform(request.TerraformVersion, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}

	// add the TF config to the directory
	err = workingDirectory.AddFile(request.TerraformConfigurationBase64, "main.tf")
	if err != nil {
		return err
	}

	// init terraform
	initOutputKey := helpers.GetPlanInitOutputKey(request.PlanExecutionRequestId)
	initOutput, err := apiClient.GetResultObjectWriter(ctx, initOutputKey, true)
	if err != nil {
		return err
	}
	request, err = apiClient.UpdatePlanExecutionRequest(ctx, request.PlanExecutionRequestId, &models.PlanExecutionRequestUpdate{
		InitOutputKey: &initOutputKey,
	})
	if err != nil {
		return err
	}
	err = executable.TerraformInit(workingDirectory, initOutput)
	initOutput.Close()
	if err != nil {
		return err
	}

	// plan terraform
	planOutputKey := helpers.GetPlanOutputKey(request.PlanExecutionRequestId)
	planOutput, err := apiClient.GetResultObjectWriter(ctx, planOutputKey, true)
	if err != nil {
		return err
	}
	planFileKey := helpers.GetPlanFileKey(request.PlanExecutionRequestId)
	planFileOutput, err := apiClient.GetResultObjectWriter(ctx, planFileKey, false)
	if err != nil {
		return err
	}
	planJSONKey := helpers.GetPlanJSONKey(request.PlanExecutionRequestId)
	planJSONOutput, err := apiClient.GetResultObjectWriter(ctx, planJSONKey, false)
	if err != nil {
		return err
	}

	_, err = apiClient.UpdatePlanExecutionRequest(ctx, request.PlanExecutionRequestId, &models.PlanExecutionRequestUpdate{
		PlanOutputKey: &planOutputKey,
		PlanFileKey:   &planFileKey,
		PlanJSONKey:   &planJSONKey,
	})
	if err != nil {
		return err
	}

	err = executable.TerraformPlan(workingDirectory, request.AdditionalArguments, planOutput, planFileOutput, planJSONOutput)
	planOutput.Close()
	planFileOutput.Close()
	planJSONOutput.Close()
	if err != nil {
		return err
	}

	return nil
}

func runApply(ctx context.Context, request *models.ApplyExecutionRequest, apiClient api.APIClientInterface, installDirectory *terraform.TerraformInstallationDirectory) error {
	// update the request to running
	newStatus := models.RequestStatusRunning
	request, err := apiClient.UpdateApplyExecutionRequest(ctx, request.ApplyExecutionRequestId, &models.ApplyExecutionRequestUpdate{
		Status: &newStatus,
	})
	if err != nil {
		return err
	}

	workingDirectory, err := terraform.NewWorkingDirectory(filepath.Join(TF_WORKING_DIRECTORY, request.ApplyExecutionRequestId))
	if err != nil {
		return err
	}
	defer workingDirectory.DeleteDirectory()

	// install terraform
	executable, err := installDirectory.InstallTerraform(request.TerraformVersion, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}

	// add the TF config to the directory
	err = workingDirectory.AddFile(request.TerraformConfigurationBase64, "main.tf")
	if err != nil {
		return err
	}

	// init terraform
	initOutputKey := helpers.GetApplyInitOutputKey(request.ApplyExecutionRequestId)
	initOutput, err := apiClient.GetResultObjectWriter(ctx, initOutputKey, true)
	if err != nil {
		return err
	}
	request, err = apiClient.UpdateApplyExecutionRequest(ctx, request.ApplyExecutionRequestId, &models.ApplyExecutionRequestUpdate{
		InitOutputKey: &initOutputKey,
	})
	if err != nil {
		return err
	}
	err = executable.TerraformInit(workingDirectory, initOutput)
	initOutput.Close()
	if err != nil {
		return err
	}

	// apply terraform
	applyOutputKey := helpers.GetApplyOutputKey(request.ApplyExecutionRequestId)
	applyOutput, err := apiClient.GetResultObjectWriter(ctx, applyOutputKey, true)
	if err != nil {
		return err
	}
	_, err = apiClient.UpdateApplyExecutionRequest(ctx, request.ApplyExecutionRequestId, &models.ApplyExecutionRequestUpdate{
		ApplyOutputKey: &applyOutputKey,
	})
	if err != nil {
		return err
	}
	err = executable.TerraformApply(workingDirectory, request.TerraformPlanBase64, request.AdditionalArguments, applyOutput)
	applyOutput.Close()
	if err != nil {
		return err
	}

	return nil
}
