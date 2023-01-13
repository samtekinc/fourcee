package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetModulePropagationExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetModulePropagationExecutionRequestsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string) (*models.ModulePropagationExecutionRequest, error) {
	thunk := c.modulePropagationExecutionRequestsLoader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%s:%s", modulePropagationId, modulePropagationExecutionRequestId)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagationExecutionRequest), nil
}

func (c *APIClient) GetModulePropagationExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error) {
	return c.dbClient.GetModulePropagationExecutionRequests(ctx, limit, cursor)
}

func (c *APIClient) GetModulePropagationExecutionRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error) {
	return c.dbClient.GetModulePropagationExecutionRequestsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *APIClient) PutModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error) {
	modulePropagationExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModulePropagationExecutionRequest)
	if err != nil {
		return nil, err
	}

	modulePropagation, err := c.GetModulePropagation(ctx, input.ModulePropagationId)
	if err != nil {
		return nil, err
	}

	moduleGroup, err := c.GetModuleGroup(ctx, modulePropagation.ModuleGroupId)
	if err != nil {
		return nil, err
	}

	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{
		ModulePropagationExecutionRequestId: modulePropagationExecutionRequestId.String(),
		ModulePropagationId:                 input.ModulePropagationId,
		RequestTime:                         time.Now().UTC(),
		Status:                              models.RequestStatusPending,
	}

	err = c.dbClient.PutModulePropagationExecutionRequest(ctx, &modulePropagationExecutionRequest)
	if err != nil {
		return nil, err
	}

	workflowExecutionInput, err := json.Marshal(map[string]string{
		"ModulePropagationId":                 input.ModulePropagationId,
		"ModulePropagationExecutionRequestId": modulePropagationExecutionRequestId.String(),
		"CloudPlatform":                       string(moduleGroup.CloudPlatform),
	})
	if err != nil {
		return nil, err
	}

	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
		StateMachineArn: aws.String(c.modulePropagationExecutionArn),
		Input:           aws.String(string(workflowExecutionInput)),
	})
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}

func (c *APIClient) UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error) {
	return c.dbClient.UpdateModulePropagationExecutionRequest(ctx, modulePropagationId, modulePropagationExecutionRequestId, update)
}
