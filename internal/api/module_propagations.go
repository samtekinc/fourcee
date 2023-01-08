package api

import (
	"context"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetModulePropagation(ctx context.Context, id string) (*models.ModulePropagation, error) {
	return c.dbClient.GetModulePropagation(ctx, id)
}

func (c *OrganizationsAPIClient) GetModulePropagations(ctx context.Context, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagations(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByModuleGroupId(ctx, moduleGroupId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByModuleVersionId(ctx, moduleVersionId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationsByOrgUnitId(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByOrgUnitId(ctx, orgUnitId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationsByOrgDimensionId(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByOrgDimensionId(ctx, orgDimensionId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModulePropagation(ctx context.Context, input *models.NewModulePropagation) (*models.ModulePropagation, error) {
	modulePropagationId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModulePropagation)
	if err != nil {
		return nil, err
	}

	modulePropagation := models.ModulePropagation{
		ModulePropagationId:       modulePropagationId.String(),
		ModuleVersionId:           input.ModuleVersionId,
		ModuleGroupId:             input.ModuleGroupId,
		OrgUnitId:                 input.OrgUnitId,
		OrgDimensionId:            input.OrgDimensionId,
		Arguments:                 ArgumentInputsToArguments(input.Arguments),
		AwsProviderConfigurations: AwsProviderConfigurationInputsToAwsProviderConfigurations(input.AwsProviderConfigurations),
		GcpProviderConfigurations: GcpProviderConfigurationInputsToGcpProviderConfigurations(input.GcpProviderConfigurations),
		Name:                      input.Name,
		Description:               input.Description,
	}
	err = c.dbClient.PutModulePropagation(ctx, &modulePropagation)
	if err != nil {
		return nil, err
	} else {
		return &modulePropagation, nil
	}
}

func (c *OrganizationsAPIClient) DeleteModulePropagation(ctx context.Context, id string) error {
	return c.dbClient.DeleteModulePropagation(ctx, id)
}

func (c *OrganizationsAPIClient) UpdateModulePropagation(ctx context.Context, modulePropagationId string, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error) {
	return c.dbClient.UpdateModulePropagation(ctx, modulePropagationId, update)
}
