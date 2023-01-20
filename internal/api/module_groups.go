package api

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetModuleGroupsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetModuleGroupsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetModuleGroup(ctx context.Context, id string) (*models.ModuleGroup, error) {
	return c.dbClient.GetModuleGroup(ctx, id)
}

func (c *APIClient) GetModuleGroupBatched(ctx context.Context, id string) (*models.ModuleGroup, error) {
	thunk := c.moduleGroupsLoader.Load(ctx, dataloader.StringKey(id))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModuleGroup), nil
}

func (c *APIClient) GetModuleGroups(ctx context.Context, limit int32, cursor string) (*models.ModuleGroups, error) {
	return c.dbClient.GetModuleGroups(ctx, limit, cursor)
}

func (c *APIClient) PutModuleGroup(ctx context.Context, input *models.NewModuleGroup) (*models.ModuleGroup, error) {
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

func (c *APIClient) DeleteModuleGroup(ctx context.Context, id string) error {
	return c.dbClient.DeleteModuleGroup(ctx, id)
}
