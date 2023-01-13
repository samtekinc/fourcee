package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskDeactivateModuleAssignment Task = "DeactivateModuleAssignment"
)

type DeactivateModuleAssignmentInput struct {
	TerraformWorkflowRequestId string
}

type DeactivateModuleAssignmentOutput struct{}

func (t *TaskHandler) DeactivateModuleAssignment(ctx context.Context, input DeactivateModuleAssignmentInput) (*DeactivateModuleAssignmentOutput, error) {
	// get workflow details
	terraformWorkflow, err := t.apiClient.GetTerraformExecutionRequest(ctx, input.TerraformWorkflowRequestId)
	if err != nil {
		return nil, err
	}

	newStatus := models.ModuleAssignmentStatusInactive
	_, err = t.apiClient.UpdateModuleAssignment(ctx, terraformWorkflow.ModuleAssignmentId, &models.ModuleAssignmentUpdate{
		Status: &newStatus,
	})
	if err != nil {
		return nil, err
	}

	return &DeactivateModuleAssignmentOutput{}, nil
}
