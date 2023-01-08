package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string) (*models.ModulePropagationDriftCheckRequest, error) {
	return c.dbClient.GetModulePropagationDriftCheckRequest(ctx, modulePropagationId, modulePropagationDriftCheckRequestId)
}

func (c *OrganizationsAPIClient) GetModulePropagationDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error) {
	return c.dbClient.GetModulePropagationDriftCheckRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationDriftCheckRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error) {
	return c.dbClient.GetModulePropagationDriftCheckRequestsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModulePropagationDriftCheckRequest(ctx context.Context, input *models.NewModulePropagationDriftCheckRequest) (*models.ModulePropagationDriftCheckRequest, error) {
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
		StateMachineArn: aws.String(c.modulePropagationDriftCheckWorkflowArn),
		Input:           aws.String(string(workflowSyncInput)),
	})
	if err != nil {
		return nil, err
	}

	return &modulePropagationDriftCheckRequest, nil
}

func (c *OrganizationsAPIClient) UpdateModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error) {
	return c.dbClient.UpdateModulePropagationDriftCheckRequest(ctx, modulePropagationId, modulePropagationDriftCheckRequestId, update)
}
