package api

import (
	"context"
	"time"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetTerraformWorkflowRequest(ctx context.Context, modulePropagationExecutionRequestId string) (*models.TerraformWorkflowRequest, error) {
	return c.dbClient.GetTerraformWorkflowRequest(ctx, modulePropagationExecutionRequestId)
}

func (c *OrganizationsAPIClient) GetTerraformWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformWorkflowRequests, error) {
	return c.dbClient.GetTerraformWorkflowRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformWorkflowRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformWorkflowRequests, error) {
	return c.dbClient.GetTerraformWorkflowRequestsByModulePropagationExecutionRequestId(ctx, modulePropagationExecutionRequestId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformWorkflowRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.TerraformWorkflowRequests, error) {
	return c.dbClient.GetTerraformWorkflowRequestsByModuleAccountAssociationKey(ctx, moduleAccountAssociationKey, limit, cursor)
}

func (c *OrganizationsAPIClient) PutTerraformWorkflowRequest(ctx context.Context, input *models.NewTerraformWorkflowRequest) (*models.TerraformWorkflowRequest, error) {
	modulePropagationExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeTerraformWorkflowRequest)
	if err != nil {
		return nil, err
	}

	modulePropagationExecutionRequest := models.TerraformWorkflowRequest{
		TerraformWorkflowRequestId:          modulePropagationExecutionRequestId.String(),
		ModulePropagationExecutionRequestId: input.ModulePropagationExecutionRequestId,
		ModuleAccountAssociationKey:         input.ModuleAccountAssociationKey,
		RequestTime:                         time.Now().UTC(),
		Status:                              models.TerraformWorkflowRequestStatusPending,
		Destroy:                             input.Destroy,
	}

	err = c.dbClient.PutTerraformWorkflowRequest(ctx, &modulePropagationExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}

func (c *OrganizationsAPIClient) UpdateTerraformWorkflowRequest(ctx context.Context, modulePropagationExecutionRequestId string, update *models.TerraformWorkflowRequestUpdate) (*models.TerraformWorkflowRequest, error) {
	return c.dbClient.UpdateTerraformWorkflowRequest(ctx, modulePropagationExecutionRequestId, update)
}
