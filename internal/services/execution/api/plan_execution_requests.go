package api

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/execution/models"
)

func (c *ExecutionAPIClient) GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error) {
	// TODO: fetch outputs from S3
	planExecutionRequest, err := c.dbClient.GetPlanExecutionRequest(ctx, planExecutionRequestId)
	if err != nil {
		return nil, err
	}

	// fetch init and plan outputs from S3
	if planExecutionRequest.InitOutputKey != "" {
		initOutput, err := c.DownloadTerraformPlanInitResults(ctx, planExecutionRequest.InitOutputKey)
		if err != nil {
			return nil, err
		}
		planExecutionRequest.InitOutput = initOutput
	}

	if planExecutionRequest.PlanOutputKey != "" {
		planOutput, err := c.DownloadTerraformPlanResults(ctx, planExecutionRequest.PlanOutputKey)
		if err != nil {
			return nil, err
		}
		planExecutionRequest.PlanOutput = planOutput
	}

	return planExecutionRequest, nil
}

func (c *ExecutionAPIClient) GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetPlanExecutionRequests(ctx, limit, cursor)
}

func (c *ExecutionAPIClient) GetPlanExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetPlanExecutionRequestsByStateKey(ctx, stateKey, limit, cursor)
}

func (c *ExecutionAPIClient) PutPlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	planExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypePlanExecutionRequest)
	if err != nil {
		return nil, err
	}

	workflowExecutionId := uuid.New().String()

	planExecutionRequest := models.PlanExecutionRequest{
		PlanExecutionRequestId:       planExecutionRequestId.String(),
		TerraformVersion:             input.TerraformVersion,
		StateKey:                     input.StateKey,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		AdditionalArguments:          input.AdditionalArguments,
		WorkflowExecutionId:          workflowExecutionId,
		Status:                       models.PlanExecutionStatusPending,
		RequestTime:                  time.Now().UTC(),
	}
	err = c.dbClient.PutPlanExecutionRequest(ctx, &planExecutionRequest)
	if err != nil {
		return nil, err
	}

	// TODO: Start Workflow

	return &planExecutionRequest, nil
}

func (c *ExecutionAPIClient) UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error) {
	return c.dbClient.UpdatePlanExecutionRequest(ctx, planExecutionRequestId, input)
}

func (c *ExecutionAPIClient) UploadTerraformPlanInitResults(ctx context.Context, applyExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	return c.dbClient.UploadTerraformPlanInitResults(ctx, applyExecutionRequestId, initResults)
}

func (c *ExecutionAPIClient) UploadTerraformPlanResults(ctx context.Context, applyExecutionRequestId string, applyResults *models.TerraformPlanOutput) (string, error) {
	return c.dbClient.UploadTerraformPlanResults(ctx, applyExecutionRequestId, applyResults)
}

func (c *ExecutionAPIClient) DownloadTerraformPlanInitResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformInitOutput, error) {
	return c.dbClient.DownloadTerraformPlanInitResults(ctx, applyExecutionRequestId)
}

func (c *ExecutionAPIClient) DownloadTerraformPlanResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformPlanOutput, error) {
	return c.dbClient.DownloadTerraformPlanResults(ctx, applyExecutionRequestId)
}
