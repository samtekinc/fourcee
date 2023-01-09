package api

import (
	"context"
	"fmt"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetModuleAssignment(ctx context.Context, moduleAssignmentId string) (*models.ModuleAssignment, error) {
	return c.dbClient.GetModuleAssignment(ctx, moduleAssignmentId)
}

func (c *OrganizationsAPIClient) GetModuleAssignments(ctx context.Context, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignments(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModuleAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModuleAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByOrgAccountId(ctx, orgAccountId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModuleAssignmentsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByModuleVersionId(ctx, moduleVersionId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModuleAssignmentsByModuleGroupId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByModuleGroupId(ctx, moduleAssignmentId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModuleAssignment(ctx context.Context, input *models.NewModuleAssignment) (*models.ModuleAssignment, error) {
	moduleAssignmentId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModuleAssignment)
	if err != nil {
		return nil, err
	}

	remoteStateKey := fmt.Sprintf("%s/state.tfstate", moduleAssignmentId.String())

	moduleAssignment := models.ModuleAssignment{
		ModuleAssignmentId:        moduleAssignmentId.String(),
		ModuleVersionId:           input.ModuleVersionId,
		ModuleGroupId:             input.ModuleGroupId,
		OrgAccountId:              input.OrgAccountId,
		Name:                      input.Name,
		Description:               input.Description,
		RemoteStateRegion:         c.remoteStateRegion,
		RemoteStateBucket:         c.remoteStateBucket,
		RemoteStateKey:            remoteStateKey,
		Arguments:                 ArgumentInputsToArguments(input.Arguments),
		AwsProviderConfigurations: AwsProviderConfigurationInputsToAwsProviderConfigurations(input.AwsProviderConfigurations),
		GcpProviderConfigurations: GcpProviderConfigurationInputsToGcpProviderConfigurations(input.GcpProviderConfigurations),
		ModulePropagationId:       input.ModulePropagationId,
		Status:                    models.ModuleAssignmentStatusActive,
	}
	err = c.dbClient.PutModuleAssignment(ctx, &moduleAssignment)
	if err != nil {
		return nil, err
	} else {
		return &moduleAssignment, nil
	}
}

func (c *OrganizationsAPIClient) UpdateModuleAssignment(ctx context.Context, moduleAssignmentId string, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error) {
	return c.dbClient.UpdateModuleAssignment(ctx, moduleAssignmentId, update)
}
