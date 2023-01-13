package api

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetModulePropagationAssignmentsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetModulePropagationAssignmentsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetModulePropagationAssignment(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModulePropagationAssignment, error) {
	thunk := c.modulePropagationAssignmentsLoader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%s:%s", modulePropagationId, orgAccountId)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagationAssignment), nil
}

func (c *APIClient) GetModulePropagationAssignments(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	return c.dbClient.GetModulePropagationAssignments(ctx, limit, cursor)
}

func (c *APIClient) GetModulePropagationAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	return c.dbClient.GetModulePropagationAssignmentsByModulePropagationId(ctx, modulePropagationId, limit, cursor)
}

func (c *APIClient) GetModulePropagationAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	return c.dbClient.GetModulePropagationAssignmentsByOrgAccountId(ctx, orgAccountId, limit, cursor)
}

func (c *APIClient) PutModulePropagationAssignment(ctx context.Context, input *models.NewModulePropagationAssignment) (*models.ModulePropagationAssignment, *models.ModuleAssignment, error) {
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
