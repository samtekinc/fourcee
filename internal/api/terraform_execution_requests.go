package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetTerraformExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetTerraformExecutionRequestsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId string) (*models.TerraformExecutionRequest, error) {
	thunk := c.terraformExecutionRequestsLoader.Load(ctx, dataloader.StringKey(terraformExecutionRequestId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.TerraformExecutionRequest), nil
}

func (c *APIClient) GetTerraformExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformExecutionRequests, error) {
	return c.dbClient.GetTerraformExecutionRequests(ctx, limit, cursor)
}

func (c *APIClient) GetTerraformExecutionRequestsByModulePropagationExecutionRequestId(ctx context.Context, terraformExecutionRequestId string, limit int32, cursor string) (*models.TerraformExecutionRequests, error) {
	return c.dbClient.GetTerraformExecutionRequestsByModulePropagationExecutionRequestId(ctx, terraformExecutionRequestId, limit, cursor)
}

func (c *APIClient) GetTerraformExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformExecutionRequests, error) {
	return c.dbClient.GetTerraformExecutionRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
}

func (c *APIClient) PutTerraformExecutionRequest(ctx context.Context, input *models.NewTerraformExecutionRequest) (*models.TerraformExecutionRequest, error) {
	moduleAssignment, err := c.GetModuleAssignment(ctx, input.ModuleAssignmentId)
	if err != nil {
		return nil, err
	}
	// set module assignment to active if it is not already
	if moduleAssignment.Status != models.ModuleAssignmentStatusActive {
		newStatus := models.ModuleAssignmentStatusActive
		_, err = c.UpdateModuleAssignment(ctx, input.ModuleAssignmentId, &models.ModuleAssignmentUpdate{
			Status: &newStatus,
		})
		if err != nil {
			return nil, err
		}
	}

	id, err := identifiers.NewIdentifier(identifiers.ResourceTypeTerraformExecutionRequest)
	if err != nil {
		return nil, err
	}

	terraformExecutionRequest := models.TerraformExecutionRequest{
		TerraformExecutionRequestId:         id.String(),
		ModuleAssignmentId:                  input.ModuleAssignmentId,
		RequestTime:                         time.Now().UTC(),
		Status:                              models.RequestStatusPending,
		Destroy:                             input.Destroy,
		CallbackTaskToken:                   input.CallbackTaskToken,
		ModulePropagationId:                 input.ModulePropagationId,
		ModulePropagationExecutionRequestId: input.ModulePropagationExecutionRequestId,
	}

	err = c.dbClient.PutTerraformExecutionRequest(ctx, &terraformExecutionRequest)
	if err != nil {
		return nil, err
	}

	inputMap := map[string]interface{}{
		"TerraformExecutionRequestId": id.String(),
		"Destroy":                     input.Destroy,
	}
	if input.CallbackTaskToken != nil && *input.CallbackTaskToken != "" {
		inputMap["TaskToken"] = *input.CallbackTaskToken
	}
	workflowExecutionInput, err := json.Marshal(inputMap)
	if err != nil {
		return nil, err
	}

	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
		StateMachineArn: aws.String(c.terraformExecutionArn),
		Input:           aws.String(string(workflowExecutionInput)),
	})
	if err != nil {
		return nil, err
	}

	return &terraformExecutionRequest, nil
}

func (c *APIClient) UpdateTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId string, update *models.TerraformExecutionRequestUpdate) (*models.TerraformExecutionRequest, error) {
	return c.dbClient.UpdateTerraformExecutionRequest(ctx, terraformExecutionRequestId, update)
}
