package api

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetModuleVersionsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetModuleVersionsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) (*models.ModuleVersion, error) {
	thunk := c.moduleVersionsLoader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%s:%s", moduleGroupId, moduleVersionId)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModuleVersion), nil
}

func (c *APIClient) GetModuleVersions(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleVersions, error) {
	return c.dbClient.GetModuleVersions(ctx, moduleGroupId, limit, cursor)
}

func (c *APIClient) PutModuleVersion(ctx context.Context, input *models.NewModuleVersion) (*models.ModuleVersion, error) {
	moduleVersionId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModuleVersion)
	if err != nil {
		return nil, err
	}

	variables, err := GetVariablesFromModule(input.RemoteSource, c.workingDirectory+moduleVersionId.String())
	if err != nil {
		return nil, err
	}

	moduleVersion := models.ModuleVersion{
		ModuleVersionId:  moduleVersionId.String(),
		ModuleGroupId:    input.ModuleGroupId,
		Name:             input.Name,
		RemoteSource:     input.RemoteSource,
		TerraformVersion: input.TerraformVersion,
		Variables:        variables,
	}
	err = c.dbClient.PutModuleVersion(ctx, &moduleVersion)
	if err != nil {
		return nil, err
	} else {
		return &moduleVersion, nil
	}
}

func (c *APIClient) DeleteModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) error {
	return c.dbClient.DeleteModuleVersion(ctx, moduleGroupId, moduleVersionId)
}
