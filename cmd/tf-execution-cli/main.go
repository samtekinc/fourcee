package main

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/database"
	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
	"go.uber.org/zap"
)

var (
	TF_INSTALLATION_DIRECTORY = os.Getenv("TF_INSTALLATION_DIRECTORY")
	TF_WORKING_DIRECTORY      = os.Getenv("TF_WORKING_DIRECTORY")
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("unable to create logger, " + err.Error())
	}
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)
	sfnClient := sfn.NewFromConfig(cfg)

	dbInput := database.DatabaseClientInput{
		DynamoDB:              dynamodbClient,
		S3:                    s3Client,
		DimensionsTableName:   "tfom-organizational-dimensions",
		UnitsTableName:        "tfom-organizational-units",
		AccountsTableName:     "tfom-organizational-accounts",
		MembershipsTableName:  "tfom-organizational-unit-memberships",
		GroupsTableName:       "tfom-module-groups",
		VersionsTableName:     "tfom-module-versions",
		PropagationsTableName: "tfom-module-propagations",
		ModulePropagationExecutionRequestsTableName:  "tfom-module-propagation-execution-requests",
		ModulePropagationDriftCheckRequestsTableName: "tfom-module-propagation-drift-check-requests",
		ModuleAssignmentsTableName:                   "tfom-module-assignments",
		ModulePropagationAssignmentsTableName:        "tfom-module-propagation-assignments",
		TerraformExecutionRequestsTableName:          "tfom-terraform-execution-requests",
		TerraformDriftCheckRequestsTableName:         "tfom-terraform-drift-check-requests",
		PlanExecutionsTableName:                      "tfom-plan-execution-requests",
		ApplyExecutionsTableName:                     "tfom-apply-execution-requests",
		ResultsBucketName:                            "tfom-execution-results",
	}
	dbClient := database.NewDatabaseClient(&dbInput)
	apiInput := api.APIClientInput{
		DBClient:                       dbClient,
		WorkingDirectory:               "./tmp/",
		SfnClient:                      sfnClient,
		ModulePropagationExecutionArn:  "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-module-propagation-execution",
		ModulePropagationDriftCheckArn: "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-module-propagation-drift-check",
		TerraformExecutionArn:          "arn:aws:states:us-east-1:306526781466:stateMachine:tfom-terraform-execution",
		RemoteStateBucket:              "tfom-backend-states",
		RemoteStateRegion:              "us-east-1",
		DataLoaderWaitTime:             time.Millisecond * 16,
	}
	apiClient := api.NewAPIClient(&apiInput)

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
		planRequest, err := apiClient.GetPlanExecutionRequest(ctx, requestID, false)
		if err != nil {
			zap.L().Panic("unable to get plan execution request", zap.Error(err))
		}
		// execute the plan
		err = runPlan(ctx, planRequest, apiClient, installationDirectory)
		if err != nil {
			zap.L().Panic("unable to run plan", zap.Error(err))
		}
	case "apply":
		applyRequest, err := apiClient.GetApplyExecutionRequest(ctx, requestID, false)
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
}

func runPlan(ctx context.Context, request *models.PlanExecutionRequest, apiClient *api.APIClient, installDirectory *terraform.TerraformInstallationDirectory) error {
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
	initOutput := executable.TerraformInit(workingDirectory)

	// upload init results
	key, err := apiClient.UploadTerraformPlanInitResults(ctx, request.PlanExecutionRequestId, initOutput)
	if err != nil {
		return err
	}
	request, err = apiClient.UpdatePlanExecutionRequest(ctx, request.PlanExecutionRequestId, &models.PlanExecutionRequestUpdate{
		InitOutputKey: &key,
	})
	if err != nil {
		return err
	}

	// if the init failed, return the error
	if initOutput.Error != nil {
		return initOutput.Error
	}

	// plan terraform
	planOutput := executable.TerraformPlan(workingDirectory, request.AdditionalArguments)

	// upload plan results
	key, err = apiClient.UploadTerraformPlanResults(ctx, request.PlanExecutionRequestId, planOutput)
	if err != nil {
		return err
	}
	_, err = apiClient.UpdatePlanExecutionRequest(ctx, request.PlanExecutionRequestId, &models.PlanExecutionRequestUpdate{
		PlanOutputKey: &key,
	})
	if err != nil {
		return err
	}

	// if the plan failed, return the error
	if planOutput.Error != nil {
		return planOutput.Error
	}

	return nil
}

func runApply(ctx context.Context, request *models.ApplyExecutionRequest, apiClient *api.APIClient, installDirectory *terraform.TerraformInstallationDirectory) error {
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
	initOutput := executable.TerraformInit(workingDirectory)

	// upload init results
	key, err := apiClient.UploadTerraformApplyInitResults(ctx, request.ApplyExecutionRequestId, initOutput)
	if err != nil {
		return err
	}
	request, err = apiClient.UpdateApplyExecutionRequest(ctx, request.ApplyExecutionRequestId, &models.ApplyExecutionRequestUpdate{
		InitOutputKey: &key,
	})
	if err != nil {
		return err
	}

	if initOutput.Error != nil {
		return initOutput.Error
	}

	// apply terraform
	applyOutput := executable.TerraformApply(workingDirectory, request.TerraformPlanBase64, request.AdditionalArguments)

	// upload apply results
	key, err = apiClient.UploadTerraformApplyResults(ctx, request.ApplyExecutionRequestId, applyOutput)
	if err != nil {
		return err
	}
	_, err = apiClient.UpdateApplyExecutionRequest(ctx, request.ApplyExecutionRequestId, &models.ApplyExecutionRequestUpdate{
		ApplyOutputKey: &key,
	})
	if err != nil {
		return err
	}

	if applyOutput.Error != nil {
		return applyOutput.Error
	}

	return nil
}
