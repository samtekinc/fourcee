package api

import (
	"context"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) (*models.ModuleVersion, error) {
	return c.dbClient.GetModuleVersion(ctx, moduleGroupId, moduleVersionId)
}

func (c *OrganizationsAPIClient) GetModuleVersions(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleVersions, error) {
	return c.dbClient.GetModuleVersions(ctx, moduleGroupId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModuleVersion(ctx context.Context, input *models.NewModuleVersion) (*models.ModuleVersion, error) {
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

func (c *OrganizationsAPIClient) DeleteModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) error {
	return c.dbClient.DeleteModuleVersion(ctx, moduleGroupId, moduleVersionId)
}