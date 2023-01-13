package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskUpdateModuleAssignments Task = "UpdateModuleAssignments"
)

type UpdateModuleAssignmentsInput struct {
	ModulePropagationId     string
	OrgAccountsPerOrgUnit   []orgAccountWrapper
	ActiveModuleAssignments []models.ModuleAssignment
}

type UpdateModuleAssignmentsOutput struct {
	OrgAccountsPerOrgUnit   []orgAccountWrapper
	ActiveModuleAssignments []models.ModuleAssignment
}

func (t *TaskHandler) UpdateModuleAssignments(ctx context.Context, input UpdateModuleAssignmentsInput) (*UpdateModuleAssignmentsOutput, error) {
	modulePropagation, err := t.apiClient.GetModulePropagation(ctx, input.ModulePropagationId)
	if err != nil {
		return nil, err
	}

	for i, moduleAssignment := range input.ActiveModuleAssignments {
		// check if we need to update the name, description, or module version needs updating
		update := models.ModuleAssignmentUpdate{}
		needsUpdate := false
		if moduleAssignment.Name != modulePropagation.Name {
			update.Name = &modulePropagation.Name
			needsUpdate = true
		}
		if moduleAssignment.Description != modulePropagation.Description {
			update.Description = &modulePropagation.Description
			needsUpdate = true
		}
		if moduleAssignment.ModuleVersionId != modulePropagation.ModuleVersionId {
			update.ModuleVersionId = &modulePropagation.ModuleVersionId
			needsUpdate = true
		}
		if needsUpdate {
			newModuleAssignment, err := t.apiClient.UpdateModuleAssignment(ctx, moduleAssignment.ModuleAssignmentId, &update)
			if err != nil {
				return nil, err
			}
			input.ActiveModuleAssignments[i] = *newModuleAssignment
		}
	}

	return &UpdateModuleAssignmentsOutput{
		OrgAccountsPerOrgUnit:   input.OrgAccountsPerOrgUnit,
		ActiveModuleAssignments: input.ActiveModuleAssignments,
	}, nil
}
