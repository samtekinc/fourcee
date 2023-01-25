package workflow

import (
	"context"
	"errors"

	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/models"
	"go.uber.org/zap"
)

const (
	TaskCreateMissingModuleAssignments Task = "CreateMissingModuleAssignments"
)

type CreateMissingModuleAssignmentsInput struct {
	ModulePropagationId              string
	AccountsNeedingModuleAssignments []models.OrgAccount
	ActiveModuleAssignments          []models.ModuleAssignment
}

type CreateMissingModuleAssignmentsOutput struct {
	ActiveModuleAssignments []models.ModuleAssignment
}

func (t *TaskHandler) CreateMissingModuleAssignments(ctx context.Context, input CreateMissingModuleAssignmentsInput) (*CreateMissingModuleAssignmentsOutput, error) {
	for _, orgAccount := range input.AccountsNeedingModuleAssignments {
		// check if there is an existing, but inactive module propagation assignment
		modulePropagationAssignment, err := t.apiClient.GetModulePropagationAssignment(ctx, input.ModulePropagationId, orgAccount.OrgAccountID)
		if errors.As(err, &helpers.NotFoundError{}) {
			zap.L().Sugar().Debugw("no existing module propagation assignment, creating one", "modulePropagationId", input.ModulePropagationId, "orgAccountID", orgAccount.OrgAccountID)
			// no existing module propagation assignment, create one
			_, moduleAssignment, err := t.apiClient.PutModulePropagationAssignment(ctx, &models.NewModulePropagationAssignment{
				ModulePropagationId: input.ModulePropagationId,
				OrgAccountID:        orgAccount.OrgAccountID,
			})
			if err != nil {
				return nil, err
			}
			input.ActiveModuleAssignments = append(input.ActiveModuleAssignments, *moduleAssignment)
		} else if err == nil {
			zap.L().Sugar().Debugw("existing module propagation assignment, activating module assignment", "modulePropagationId", input.ModulePropagationId, "orgAccountID", orgAccount.OrgAccountID, "modulePropagationAssignment", modulePropagationAssignment)
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
