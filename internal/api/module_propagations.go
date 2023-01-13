package api

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetModulePropagationsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetModulePropagationsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetModulePropagation(ctx context.Context, id string) (*models.ModulePropagation, error) {
	thunk := c.modulePropagationsLoader.Load(ctx, dataloader.StringKey(id))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagation), nil
}

func (c *APIClient) GetModulePropagations(ctx context.Context, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagations(ctx, limit, cursor)
}

func (c *APIClient) GetModulePropagationsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByModuleGroupId(ctx, moduleGroupId, limit, cursor)
}

func (c *APIClient) GetModulePropagationsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByModuleVersionId(ctx, moduleVersionId, limit, cursor)
}

func (c *APIClient) GetModulePropagationsByOrgUnitId(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByOrgUnitId(ctx, orgUnitId, limit, cursor)
}

func (c *APIClient) GetModulePropagationsByOrgDimensionId(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	return c.dbClient.GetModulePropagationsByOrgDimensionId(ctx, orgDimensionId, limit, cursor)
}

func (c *APIClient) PutModulePropagation(ctx context.Context, input *models.NewModulePropagation) (*models.ModulePropagation, error) {
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

func (c *APIClient) DeleteModulePropagation(ctx context.Context, id string) error {
	return c.dbClient.DeleteModulePropagation(ctx, id)
}

func (c *APIClient) UpdateModulePropagation(ctx context.Context, modulePropagationId string, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error) {
	return c.dbClient.UpdateModulePropagation(ctx, modulePropagationId, update)
}
