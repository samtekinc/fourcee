package api

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetPlanExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetPlanExecutionRequestsByIds(ctx, keys.Keys())
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: &results[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error) {
	planExecutionRequest, err := c.dbClient.GetPlanExecutionRequest(ctx, planExecutionRequestId)
	if err != nil {
		return nil, err
	}

	return planExecutionRequest, nil
}

func (c *APIClient) GetPlanExecutionRequestBatched(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error) {
	thunk := c.planExecutionRequestsLoader.Load(ctx, dataloader.StringKey(planExecutionRequestId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	planExecutionRequest := result.(*models.PlanExecutionRequest)

	return planExecutionRequest, nil
}

func (c *APIClient) GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	requests, err := c.dbClient.GetPlanExecutionRequests(ctx, limit, cursor)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (c *APIClient) GetPlanExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	requests, err := c.dbClient.GetPlanExecutionRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (c *APIClient) PutPlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	planExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypePlanExecutionRequest)
	if err != nil {
		return nil, err
	}

	planExecutionRequest := models.PlanExecutionRequest{
		PlanExecutionRequestId:       planExecutionRequestId.String(),
		ModuleAssignmentId:           input.ModuleAssignmentId,
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.RequestStatusPending,
		RequestTime:                  time.Now().UTC(),
	}
	err = c.dbClient.PutPlanExecutionRequest(ctx, &planExecutionRequest)
	if err != nil {
		return nil, err
	}

	// Start Workflow
	err = c.startTerraformCommandWorkflow(ctx, &TerraformExecutionInput{
		RequestType: "plan",
		RequestId:   planExecutionRequestId.String(),
		TaskToken:   planExecutionRequest.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}

func (c *APIClient) UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error) {
	return c.dbClient.UpdatePlanExecutionRequest(ctx, planExecutionRequestId, input)
}
