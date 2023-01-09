package workflow

import (
	"context"
	"errors"

	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskCreateMissingModuleAssignments Task = "CreateMissingModuleAssignments"
)

type CreateMissingModuleAssignmentsInput struct {
	ModulePropagationId              string
	AccountsNeedingModuleAssignments []models.OrganizationalAccount
	ActiveModuleAssignments          []models.ModuleAssignment
}

type CreateMissingModuleAssignmentsOutput struct {
	ActiveModuleAssignments []models.ModuleAssignment
}

func (t *TaskHandler) CreateMissingModuleAssignments(ctx context.Context, input CreateMissingModuleAssignmentsInput) (*CreateMissingModuleAssignmentsOutput, error) {
	for _, orgAccount := range input.AccountsNeedingModuleAssignments {
		// check if there is an existing, but inactive module propagation assignment
		modulePropagationAssignment, err := t.apiClient.GetModulePropagationAssignment(ctx, input.ModulePropagationId, orgAccount.OrgAccountId)
		if errors.As(err, &helpers.NotFoundError{}) {
			// no existing module propagation assignment, create one
			_, moduleAssignment, err := t.apiClient.PutModulePropagationAssignment(ctx, &models.NewModulePropagationAssignment{
				ModulePropagationId: input.ModulePropagationId,
				OrgAccountId:        orgAccount.OrgAccountId,
			})
			if err != nil {
				return nil, err
			}
			input.ActiveModuleAssignments = append(input.ActiveModuleAssignments, *moduleAssignment)
		} else if err == nil {
			// there is an existing module propagation assignment, get the matching module assignment
			moduleAssignment, err := t.apiClient.GetModuleAssignment(ctx, modulePropagationAssignment.ModuleAssignmentId)
			if err != nil {
				return nil, err
			}
			// if the module assignment is inactive, activate it
			if moduleAssignment.Status != models.ModuleAssignmentStatusActive {
				newStatus := models.ModuleAssignmentStatusActive
				moduleAssignment, err = t.apiClient.UpdateModuleAssignment(ctx, moduleAssignment.ModuleAssignmentId, &models.ModuleAssignmentUpdate{
					Status: &newStatus,
				})
				if err != nil {
					return nil, err
				}
				input.ActiveModuleAssignments = append(input.ActiveModuleAssignments, *moduleAssignment)
			}
		} else if err != nil {
			return nil, err
		}
	}
	return &CreateMissingModuleAssignmentsOutput{
		ActiveModuleAssignments: input.ActiveModuleAssignments,
	}, nil
}
