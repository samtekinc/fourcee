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

func (c *APIClient) GetTerraformDriftCheckRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetTerraformDriftCheckRequestsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestId string) (*models.TerraformDriftCheckRequest, error) {
	thunk := c.terraformDriftCheckRequestsLoader.Load(ctx, dataloader.StringKey(terraformDriftCheckRequestId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.TerraformDriftCheckRequest), nil
}

func (c *APIClient) GetTerraformDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error) {
	return c.dbClient.GetTerraformDriftCheckRequests(ctx, limit, cursor)
}

func (c *APIClient) GetTerraformDriftCheckRequestsByModulePropagationDriftCheckRequestId(ctx context.Context, terraformDriftCheckRequestId string, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error) {
	return c.dbClient.GetTerraformDriftCheckRequestsByModulePropagationDriftCheckRequestId(ctx, terraformDriftCheckRequestId, limit, cursor)
}

func (c *APIClient) GetTerraformDriftCheckRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error) {
	return c.dbClient.GetTerraformDriftCheckRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
}

func (c *APIClient) PutTerraformDriftCheckRequest(ctx context.Context, input *models.NewTerraformDriftCheckRequest) (*models.TerraformDriftCheckRequest, error) {
	id, err := identifiers.NewIdentifier(identifiers.ResourceTypeTerraformDriftCheckRequest)
	if err != nil {
		return nil, err
	}

	terraformDriftCheckRequest := models.TerraformDriftCheckRequest{
		TerraformDriftCheckRequestId:         id.String(),
		ModuleAssignmentId:                   input.ModuleAssignmentId,
		RequestTime:                          time.Now().UTC(),
		Status:                               models.RequestStatusPending,
		Destroy:                              input.Destroy,
		CallbackTaskToken:                    input.CallbackTaskToken,
		SyncStatus:                           models.TerraformDriftCheckStatusPending,
		ModulePropagationId:                  input.ModulePropagationId,
		ModulePropagationDriftCheckRequestId: input.ModulePropagationDriftCheckRequestId,
	}

	err = c.dbClient.PutTerraformDriftCheckRequest(ctx, &terraformDriftCheckRequest)
	if err != nil {
		return nil, err
	}

	inputMap := map[string]string{
		"TerraformDriftCheckRequestId": id.String(),
	}
	if input.CallbackTaskToken != nil && *input.CallbackTaskToken != "" {
		inputMap["TaskToken"] = *input.CallbackTaskToken
	}
	workflowExecutionInput, err := json.Marshal(inputMap)
	if err != nil {
		return nil, err
	}

	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
		StateMachineArn: aws.String(c.terraformDriftCheckArn),
		Input:           aws.String(string(workflowExecutionInput)),
	})
	if err != nil {
		return nil, err
	}

	return &terraformDriftCheckRequest, nil
}

func (c *APIClient) UpdateTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestId string, update *models.TerraformDriftCheckRequestUpdate) (*models.TerraformDriftCheckRequest, error) {
	return c.dbClient.UpdateTerraformDriftCheckRequest(ctx, terraformDriftCheckRequestId, update)
}
