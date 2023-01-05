package api

import (
	"context"
	"time"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequest(ctx context.Context, modulePropagationDriftCheckRequestId string) (*models.TerraformDriftCheckWorkflowRequest, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequest(ctx, modulePropagationDriftCheckRequestId)
}

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequestsByModulePropagationDriftCheckRequestId(ctx context.Context, modulePropagationDriftCheckRequestId string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequestsByModulePropagationDriftCheckRequestId(ctx, modulePropagationDriftCheckRequestId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequestsByModuleAccountAssociationKey(ctx, moduleAccountAssociationKey, limit, cursor)
}

func (c *OrganizationsAPIClient) PutTerraformDriftCheckWorkflowRequest(ctx context.Context, input *models.NewTerraformDriftCheckWorkflowRequest) (*models.TerraformDriftCheckWorkflowRequest, error) {
	modulePropagationDriftCheckRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeTerraformDriftCheckWorkflowRequest)
	if err != nil {
		return nil, err
	}

	modulePropagationDriftCheckRequest := models.TerraformDriftCheckWorkflowRequest{
		TerraformDriftCheckWorkflowRequestId: modulePropagationDriftCheckRequestId.String(),
		ModulePropagationDriftCheckRequestId: input.ModulePropagationDriftCheckRequestId,
		ModuleAccountAssociationKey:          input.ModuleAccountAssociationKey,
		RequestTime:                          time.Now().UTC(),
		Status:                               models.RequestStatusPending,
		Destroy:                              input.Destroy,
		SyncStatus:                           models.TerraformDriftCheckStatusPending,
	}

	err = c.dbClient.PutTerraformDriftCheckWorkflowRequest(ctx, &modulePropagationDriftCheckRequest)
	if err != nil {
		return nil, err
	}

	return &modulePropagationDriftCheckRequest, nil
}

func (c *OrganizationsAPIClient) UpdateTerraformDriftCheckWorkflowRequest(ctx context.Context, modulePropagationDriftCheckRequestId string, update *models.TerraformDriftCheckWorkflowRequestUpdate) (*models.TerraformDriftCheckWorkflowRequest, error) {
	return c.dbClient.UpdateTerraformDriftCheckWorkflowRequest(ctx, modulePropagationDriftCheckRequestId, update)
}
