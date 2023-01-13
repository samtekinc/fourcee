package api

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetModuleAssignmentsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetModuleAssignmentsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetModuleAssignment(ctx context.Context, moduleAssignmentId string) (*models.ModuleAssignment, error) {
	thunk := c.moduleAssignmentsLoader.Load(ctx, dataloader.StringKey(moduleAssignmentId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModuleAssignment), nil
}

func (c *APIClient) GetModuleAssignments(ctx context.Context, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignments(ctx, limit, cursor)
}

func (c *APIClient) GetModuleAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *APIClient) GetModuleAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByOrgAccountId(ctx, orgAccountId, limit, cursor)
}

func (c *APIClient) GetModuleAssignmentsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByModuleVersionId(ctx, moduleVersionId, limit, cursor)
}

func (c *APIClient) GetModuleAssignmentsByModuleGroupId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	return c.dbClient.GetModuleAssignmentsByModuleGroupId(ctx, moduleAssignmentId, limit, cursor)
}

func (c *APIClient) PutModuleAssignment(ctx context.Context, input *models.NewModuleAssignment) (*models.ModuleAssignment, error) {
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

func (c *APIClient) UpdateModuleAssignment(ctx context.Context, moduleAssignmentId string, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error) {
	return c.dbClient.UpdateModuleAssignment(ctx, moduleAssignmentId, update)
}
