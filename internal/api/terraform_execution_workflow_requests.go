package api

import (
	"context"
	"time"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequest(ctx context.Context, modulePropagationExecutionRequestId string) (*models.TerraformExecutionWorkflowRequest, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequest(ctx, modulePropagationExecutionRequestId)
}

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequestsByModulePropagationExecutionRequestId(ctx, modulePropagationExecutionRequestId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequestsByModuleAccountAssociationKey(ctx, moduleAccountAssociationKey, limit, cursor)
}

func (c *OrganizationsAPIClient) PutTerraformExecutionWorkflowRequest(ctx context.Context, input *models.NewTerraformExecutionWorkflowRequest) (*models.TerraformExecutionWorkflowRequest, error) {
	modulePropagationExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeTerraformExecutionWorkflowRequest)
	if err != nil {
		return nil, err
	}

	modulePropagationExecutionRequest := models.TerraformExecutionWorkflowRequest{
		TerraformExecutionWorkflowRequestId: modulePropagationExecutionRequestId.String(),
		ModulePropagationExecutionRequestId: input.ModulePropagationExecutionRequestId,
		ModuleAccountAssociationKey:         input.ModuleAccountAssociationKey,
		RequestTime:                         time.Now().UTC(),
		Status:                              models.RequestStatusPending,
		Destroy:                             input.Destroy,
	}

	err = c.dbClient.PutTerraformExecutionWorkflowRequest(ctx, &modulePropagationExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}

func (c *OrganizationsAPIClient) UpdateTerraformExecutionWorkflowRequest(ctx context.Context, modulePropagationExecutionRequestId string, update *models.TerraformExecutionWorkflowRequestUpdate) (*models.TerraformExecutionWorkflowRequest, error) {
	return c.dbClient.UpdateTerraformExecutionWorkflowRequest(ctx, modulePropagationExecutionRequestId, update)
}
