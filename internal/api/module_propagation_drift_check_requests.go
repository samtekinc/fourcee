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

func (c *APIClient) GetModulePropagationDriftCheckRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetModulePropagationDriftCheckRequestsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string) (*models.ModulePropagationDriftCheckRequest, error) {
	return c.dbClient.GetModulePropagationDriftCheckRequest(ctx, modulePropagationId, modulePropagationDriftCheckRequestId)
}

func (c *APIClient) GetModulePropagationDriftCheckRequestBatched(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string) (*models.ModulePropagationDriftCheckRequest, error) {
	thunk := c.modulePropagationDriftCheckRequestsLoader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%s:%s", modulePropagationId, modulePropagationDriftCheckRequestId)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagationDriftCheckRequest), nil
}

func (c *APIClient) GetModulePropagationDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error) {
	return c.dbClient.GetModulePropagationDriftCheckRequests(ctx, limit, cursor)
}

func (c *APIClient) GetModulePropagationDriftCheckRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error) {
	return c.dbClient.GetModulePropagationDriftCheckRequestsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *APIClient) PutModulePropagationDriftCheckRequest(ctx context.Context, input *models.NewModulePropagationDriftCheckRequest) (*models.ModulePropagationDriftCheckRequest, error) {
	modulePropagationDriftCheckRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModulePropagationDriftCheckRequest)
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

	modulePropagationDriftCheckRequest := models.ModulePropagationDriftCheckRequest{
		ModulePropagationDriftCheckRequestId: modulePropagationDriftCheckRequestId.String(),
		ModulePropagationId:                  input.ModulePropagationId,
		RequestTime:                          time.Now().UTC(),
		Status:                               models.RequestStatusPending,
		SyncStatus:                           models.TerraformDriftCheckStatusPending,
	}

	err = c.dbClient.PutModulePropagationDriftCheckRequest(ctx, &modulePropagationDriftCheckRequest)
	if err != nil {
		return nil, err
	}

	workflowSyncInput, err := json.Marshal(map[string]string{
		"ModulePropagationId":                  input.ModulePropagationId,
		"ModulePropagationDriftCheckRequestId": modulePropagationDriftCheckRequestId.String(),
		"CloudPlatform":                        string(moduleGroup.CloudPlatform),
	})
	if err != nil {
		return nil, err
	}

	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
		StateMachineArn: aws.String(c.modulePropagationDriftCheckArn),
		Input:           aws.String(string(workflowSyncInput)),
	})
	if err != nil {
		return nil, err
	}

	return &modulePropagationDriftCheckRequest, nil
}

func (c *APIClient) UpdateModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error) {
	return c.dbClient.UpdateModulePropagationDriftCheckRequest(ctx, modulePropagationId, modulePropagationDriftCheckRequestId, update)
}
