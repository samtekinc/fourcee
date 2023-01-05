package workflow

import (
	"context"
	"fmt"

	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskDetermineSyncStatus Task = "DetermineSyncStatus"
)

type DetermineSyncStatusInput struct {
	TerraformWorkflowRequestId string
}

type DetermineSyncStatusOutput struct {
}

func (t *TaskHandler) DetermineSyncStatus(ctx context.Context, input DetermineSyncStatusInput) (*DetermineSyncStatusOutput, error) {
	tfWorkflow, err := t.apiClient.GetTerraformDriftCheckWorkflowRequest(ctx, input.TerraformWorkflowRequestId)
	if err != nil {
		return nil, err
	}

	if tfWorkflow.PlanExecutionRequestId == nil {
		return nil, fmt.Errorf("plan execution request id is nil")
	}

	plan, err := t.apiClient.GetPlanExecutionRequest(ctx, *tfWorkflow.PlanExecutionRequestId)
	if err != nil {
		return nil, err
	}

	if plan.PlanOutput == nil {
		return nil, fmt.Errorf("plan output is nil")
	}
	if plan.PlanOutput.PlanJSON == nil {
		return nil, fmt.Errorf("plan output plan json is nil")
	}

	planFile, err := terraform.TerraformPlanFileFromJSON(plan.PlanOutput.PlanJSON)
	if err != nil {
		return nil, err
	}

	var syncStatus models.TerraformDriftCheckStatus
	if planFile.HasChanges() {
		syncStatus = models.TerraformDriftCheckStatusOutOfSync
	} else {
		syncStatus = models.TerraformDriftCheckStatusInSync
	}
	_, err = t.apiClient.UpdateTerraformDriftCheckWorkflowRequest(ctx, input.TerraformWorkflowRequestId, &models.TerraformDriftCheckWorkflowRequestUpdate{
		SyncStatus: &syncStatus,
	})
	if err != nil {
		return nil, err
	}

	return &DetermineSyncStatusOutput{}, nil
}
