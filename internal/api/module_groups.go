package api

import (
	"context"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetModuleGroup(ctx context.Context, id string) (*models.ModuleGroup, error) {
	return c.dbClient.GetModuleGroup(ctx, id)
}

func (c *OrganizationsAPIClient) GetModuleGroups(ctx context.Context, limit int32, cursor string) (*models.ModuleGroups, error) {
	return c.dbClient.GetModuleGroups(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModuleGroup(ctx context.Context, input *models.NewModuleGroup) (*models.ModuleGroup, error) {
	moduleGroupId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModuleGroup)
	if err != nil {
		return nil, err
	}

	moduleGroup := models.ModuleGroup{
		ModuleGroupId: moduleGroupId.String(),
		Name:          input.Name,
		CloudPlatform: input.CloudPlatform,
	}
	err = c.dbClient.PutModuleGroup(ctx, &moduleGroup)
	if err != nil {
		return nil, err
	} else {
		return &moduleGroup, nil
	}
}

func (c *OrganizationsAPIClient) DeleteModuleGroup(ctx context.Context, id string) error {
	return c.dbClient.DeleteModuleGroup(ctx, id)
}
