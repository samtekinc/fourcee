package api

import (
	"context"
	"time"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/organizations/models"
)

func (c *OrganizationsAPIClient) GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string) (*models.ModulePropagationExecutionRequest, error) {
	return c.dbClient.GetModulePropagationExecutionRequest(ctx, modulePropagationId, modulePropagationExecutionRequestId)
}

func (c *OrganizationsAPIClient) GetModulePropagationExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error) {
	return c.dbClient.GetModulePropagationExecutionRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationExecutionRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error) {
	return c.dbClient.GetModulePropagationExecutionRequestsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error) {
	modulePropagationExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModulePropagationExecutionRequest)
	if err != nil {
		return nil, err
	}

	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{
		ModulePropagationExecutionRequestId: modulePropagationExecutionRequestId.String(),
		ModulePropagationId:                 input.ModulePropagationId,
		RequestTime:                         time.Now().UTC(),
		Status:                              models.ModulePropagationExecutionRequestStatusPending,
	}

	err = c.dbClient.PutModulePropagationExecutionRequest(ctx, &modulePropagationExecutionRequest)
	if err != nil {
		return nil, err
	}

	// TODO: trigger workflow

	return &modulePropagationExecutionRequest, nil
}

func (c *OrganizationsAPIClient) UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error) {
	return c.dbClient.UpdateModulePropagationExecutionRequest(ctx, modulePropagationId, modulePropagationExecutionRequestId, update)
}
