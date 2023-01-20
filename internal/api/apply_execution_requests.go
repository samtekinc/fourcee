package api

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetApplyExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetApplyExecutionRequestsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequest, err := c.dbClient.GetApplyExecutionRequest(ctx, applyExecutionRequestId)
	if err != nil {
		return nil, err
	}

	return applyExecutionRequest, nil
}

func (c *APIClient) GetApplyExecutionRequestBatched(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error) {
	thunk := c.applyExecutionRequestsLoader.Load(ctx, dataloader.StringKey(applyExecutionRequestId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	applyExecutionRequest := result.(*models.ApplyExecutionRequest)

	return applyExecutionRequest, nil
}

func (c *APIClient) GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	requests, err := c.dbClient.GetApplyExecutionRequests(ctx, limit, cursor)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (c *APIClient) GetApplyExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	requests, err := c.dbClient.GetApplyExecutionRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (c *APIClient) PutApplyExecutionRequest(ctx context.Context, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeApplyExecutionRequest)
	if err != nil {
		return nil, err
	}

	applyExecutionRequest := models.ApplyExecutionRequest{
		ApplyExecutionRequestId:      applyExecutionRequestId.String(),
		ModuleAssignmentId:           input.ModuleAssignmentId,
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		TerraformPlanBase64:          input.TerraformPlanBase64,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.RequestStatusPending,
		RequestTime:                  time.Now().UTC(),
	}
	err = c.dbClient.PutApplyExecutionRequest(ctx, &applyExecutionRequest)
	if err != nil {
		return nil, err
	}

	// Start Workflow
	err = c.startTerraformCommandWorkflow(ctx, &TerraformExecutionInput{
		RequestType: "apply",
		RequestId:   applyExecutionRequestId.String(),
		TaskToken:   applyExecutionRequest.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}

func (c *APIClient) UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, input *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error) {
	return c.dbClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequestId, input)
}
