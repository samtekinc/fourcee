package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sheacloud/tfom/internal/services/execution/api"
	"github.com/sheacloud/tfom/internal/services/execution/database"
	"github.com/sheacloud/tfom/internal/services/execution/terraform"
	"github.com/sheacloud/tfom/pkg/execution/models"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	dbInput := database.ExecutionDatabaseClientInput{
		DynamoDB:                 dynamodbClient,
		S3:                       s3Client,
		PlanExecutionsTableName:  "tfom-exec-service-plan-execution-requests",
		ApplyExecutionsTableName: "tfom-exec-service-apply-execution-requests",
		ResultsBucketName:        "tfom-exec-service-execution-results",
	}
	execDbClient := database.NewExecutionDatabaseClient(&dbInput)
	execApiClient := api.NewExecutionAPIClient(execDbClient)

	installationDirectory, err := terraform.NewTerraformInstallationDirectory("./tf-installation")
	if err != nil {
		panic("unable to create terraform installation directory, " + err.Error())
	}

	// get request type and ID from environment
	requestType := os.Getenv("REQUEST_TYPE")
	requestID := os.Getenv("REQUEST_ID")

	// get the request from the database
	switch requestType {
	case "plan":
		planRequest, err := execApiClient.GetPlanExecutionRequest(ctx, requestID)
		if err != nil {
			panic("unable to get plan execution request, " + err.Error())
		}
		// execute the plan
		err = runPlan(ctx, planRequest, execApiClient, installationDirectory)
		if err != nil {
			panic("unable to run plan, " + err.Error())
		}
	case "apply":
		applyRequest, err := execApiClient.GetApplyExecutionRequest(ctx, requestID)
		if err != nil {
			panic("unable to get apply execution request, " + err.Error())
		}
		// execute the apply
		err = runApply(ctx, applyRequest, execApiClient, installationDirectory)
		if err != nil {
			panic("unable to run apply, " + err.Error())
		}
	default:
		panic("invalid request type")
	}
}

func runPlan(ctx context.Context, request *models.PlanExecutionRequest, apiClient *api.ExecutionAPIClient, installDirectory *terraform.TerraformInstallationDirectory) error {
	workingDirectory, err := terraform.NewWorkingDirectory("./example-tf/" + request.PlanExecutionRequestId)
	if err != nil {
		return err
	}
	defer workingDirectory.DeleteDirectory()

	// install terraform
	executable, err := installDirectory.InstallTerraform(request.TerraformVersion, "darwin", "arm64")
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
	if initOutput.Error != nil {
		return initOutput.Error
	}

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

	// plan terraform
	planOutput := executable.TerraformPlan(workingDirectory)
	if planOutput.Error != nil {
		return planOutput.Error
	}

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

	return nil
}

func runApply(ctx context.Context, request *models.ApplyExecutionRequest, apiClient *api.ExecutionAPIClient, installDirectory *terraform.TerraformInstallationDirectory) error {
	workingDirectory, err := terraform.NewWorkingDirectory("./example-tf/" + request.ApplyExecutionRequestId)
	if err != nil {
		return err
	}
	defer workingDirectory.DeleteDirectory()

	// install terraform
	executable, err := installDirectory.InstallTerraform(request.TerraformVersion, "darwin", "arm64")
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
	if initOutput.Error != nil {
		return initOutput.Error
	}

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

	// apply terraform
	applyOutput := executable.TerraformApply(workingDirectory, request.TerraformPlanBase64)
	if applyOutput.Error != nil {
		return applyOutput.Error
	}

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

	return nil
}
