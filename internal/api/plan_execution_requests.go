package api

import (
	"context"
	"time"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error) {
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

func (c *OrganizationsAPIClient) GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetPlanExecutionRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetPlanExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetPlanExecutionRequestsByStateKey(ctx, stateKey, limit, cursor)
}

func (c *OrganizationsAPIClient) GetPlanExecutionRequestsByGroupingKey(ctx context.Context, groupingKey string, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetPlanExecutionRequestsByGroupingKey(ctx, groupingKey, limit, cursor)
}

func (c *OrganizationsAPIClient) PutPlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	planExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypePlanExecutionRequest)
	if err != nil {
		return nil, err
	}

	planExecutionRequest := models.PlanExecutionRequest{
		PlanExecutionRequestId:       planExecutionRequestId.String(),
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		StateKey:                     input.StateKey,
		GroupingKey:                  input.GroupingKey,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.PlanExecutionStatusPending,
		RequestTime:                  time.Now().UTC(),
	}
	err = c.dbClient.PutPlanExecutionRequest(ctx, &planExecutionRequest)
	if err != nil {
		return nil, err
	}

	// Start Workflow
	err = c.startTerraformExecutionWorkflow(ctx, &TerraformExecutionWorkflowInput{
		RequestType: "plan",
		RequestId:   planExecutionRequestId.String(),
		TaskToken:   planExecutionRequest.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}

func (c *OrganizationsAPIClient) UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error) {
	return c.dbClient.UpdatePlanExecutionRequest(ctx, planExecutionRequestId, input)
}

func (c *OrganizationsAPIClient) UploadTerraformPlanInitResults(ctx context.Context, applyExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	return c.dbClient.UploadTerraformPlanInitResults(ctx, applyExecutionRequestId, initResults)
}

func (c *OrganizationsAPIClient) UploadTerraformPlanResults(ctx context.Context, applyExecutionRequestId string, applyResults *models.TerraformPlanOutput) (string, error) {
	return c.dbClient.UploadTerraformPlanResults(ctx, applyExecutionRequestId, applyResults)
}

func (c *OrganizationsAPIClient) DownloadTerraformPlanInitResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformInitOutput, error) {
	return c.dbClient.DownloadTerraformPlanInitResults(ctx, applyExecutionRequestId)
}

func (c *OrganizationsAPIClient) DownloadTerraformPlanResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformPlanOutput, error) {
	return c.dbClient.DownloadTerraformPlanResults(ctx, applyExecutionRequestId)
}
