package activities

import (
	"context"
	"fmt"

	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
)

func (r *Activities) DeterminePlanSyncStatus(ctx context.Context, planExecutionRequestID uint) (*models.TerraformDriftCheckStatus, error) {
	planExecutionRequest, err := r.apiClient.GetPlanExecutionRequest(ctx, planExecutionRequestID)
	if err != nil {
		return nil, err
	}

	if planExecutionRequest.TerraformDriftCheckRequestID == nil {
		return nil, fmt.Errorf("plan execution request %d does not have a terraform drift check request", planExecutionRequestID)
	}

	planFile, err := terraform.TerraformPlanFileFromJSON(planExecutionRequest.PlanJSON)
	if err != nil {
		return nil, err
	}

	var syncStatus models.TerraformDriftCheckStatus
	if planFile.HasChanges() {
		syncStatus = models.TerraformDriftCheckStatusOutOfSync
	} else {
		syncStatus = models.TerraformDriftCheckStatusInSync
	}

	return &syncStatus, nil
}
