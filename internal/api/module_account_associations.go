package api

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetModuleAccountAssociation(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModuleAccountAssociation, error) {
	return c.dbClient.GetModuleAccountAssociation(ctx, modulePropagationId, orgAccountId)
}

func (c *OrganizationsAPIClient) GetModuleAccountAssociations(ctx context.Context, limit int32, cursor string) (*models.ModuleAccountAssociations, error) {
	return c.dbClient.GetModuleAccountAssociations(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModuleAccountAssociationsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAccountAssociations, error) {
	return c.dbClient.GetModuleAccountAssociationsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModuleAccountAssociationsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAccountAssociations, error) {
	return c.dbClient.GetModuleAccountAssociationsByOrgAccountId(ctx, orgAccountId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModuleAccountAssociation(ctx context.Context, input *models.NewModuleAccountAssociation) (*models.ModuleAccountAssociation, error) {
	moduleAccountAssociation := models.ModuleAccountAssociation{
		ModulePropagationId: input.ModulePropagationId,
		OrgAccountId:        input.OrgAccountId,
		Status:              models.ModuleAccountAssociationStatusActive,
		RemoteStateBucket:   input.RemoteStateBucket,
		RemoteStateKey:      input.RemoteStateKey,
		RemoteStateRegion:   input.RemoteStateRegion,
	}

	err := c.dbClient.PutModuleAccountAssociation(ctx, &moduleAccountAssociation)
	if err != nil {
		return nil, err
	}

	return &moduleAccountAssociation, nil
}

func (c *OrganizationsAPIClient) UpdateModuleAccountAssociation(ctx context.Context, modulePropagationId string, orgAccountId string, update *models.ModuleAccountAssociationUpdate) (*models.ModuleAccountAssociation, error) {
	return c.dbClient.UpdateModuleAccountAssociation(ctx, modulePropagationId, orgAccountId, update)
}
