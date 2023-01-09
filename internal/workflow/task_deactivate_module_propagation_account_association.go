package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskDeactivateModuleAssignment Task = "DeactivateModuleAssignment"
)

type DeactivateModuleAssignmentInput struct {
	ModuleAssignment models.ModuleAssignment
}

type DeactivateModuleAssignmentOutput struct{}

func (t *TaskHandler) DeactivateModuleAssignment(ctx context.Context, input DeactivateModuleAssignmentInput) (*DeactivateModuleAssignmentOutput, error) {

	newStatus := models.ModuleAssignmentStatusInactive
	_, err := t.apiClient.UpdateModuleAssignment(ctx, input.ModuleAssignment.ModuleAssignmentId, &models.ModuleAssignmentUpdate{
		Status: &newStatus,
	})
	if err != nil {
		return nil, err
	}

	return &DeactivateModuleAssignmentOutput{}, nil
}
