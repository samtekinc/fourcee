package api

import (
	"context"
	"fmt"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetModulePropagationAssignment(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModulePropagationAssignment, error) {
	return c.dbClient.GetModulePropagationAssignment(ctx, modulePropagationId, orgAccountId)
}

func (c *OrganizationsAPIClient) GetModulePropagationAssignments(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	return c.dbClient.GetModulePropagationAssignments(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	return c.dbClient.GetModulePropagationAssignmentsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetModulePropagationAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	return c.dbClient.GetModulePropagationAssignmentsByOrgAccountId(ctx, orgAccountId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutModulePropagationAssignment(ctx context.Context, input *models.NewModulePropagationAssignment) (*models.ModulePropagationAssignment, *models.ModuleAssignment, error) {
	moduleAssignmentId, err := identifiers.NewIdentifier(identifiers.ResourceTypeModuleAssignment)
	if err != nil {
		return nil, nil, err
	}

	remoteStateKey := fmt.Sprintf("%s/state.tfstate", moduleAssignmentId.String())

	modulePropagation, err := c.GetModulePropagation(ctx, input.ModulePropagationId)
	if err != nil {
		return nil, nil, err
	}

	moduleAssignment := &models.ModuleAssignment{
		ModuleAssignmentId:        moduleAssignmentId.String(),
		ModuleVersionId:           modulePropagation.ModuleVersionId,
		ModuleGroupId:             modulePropagation.ModuleGroupId,
		OrgAccountId:              input.OrgAccountId,
		Name:                      modulePropagation.Name,
		Description:               modulePropagation.Description,
		RemoteStateRegion:         c.remoteStateRegion,
		RemoteStateBucket:         c.remoteStateBucket,
		RemoteStateKey:            remoteStateKey,
		Arguments:                 nil,
		AwsProviderConfigurations: nil,
		GcpProviderConfigurations: nil,
		ModulePropagationId:       &input.ModulePropagationId,
		Status:                    models.ModuleAssignmentStatusActive,
	}

	return c.dbClient.PutModulePropagationAssignment(ctx, moduleAssignment)
}
