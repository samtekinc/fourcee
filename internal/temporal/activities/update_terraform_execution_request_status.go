package activities

import (
	"context"
	"time"

	"github.com/sheacloud/tfom/pkg/models"
)

func (r *Activities) UpdateTerraformExecutionRequestStatus(ctx context.Context, terraformExecutionRequestID uint, newStatus models.RequestStatus) (*models.TerraformExecutionRequest, error) {
	update := &models.TerraformExecutionRequestUpdate{
		Status: &newStatus,
	}
	now := time.Now().UTC()
	switch newStatus {
	case models.RequestStatusRunning:
		update.Status = &newStatus
		update.StartedAt = &now
	case models.RequestStatusSucceeded:
		update.Status = &newStatus
		update.CompletedAt = &now
	case models.RequestStatusFailed:
		update.Status = &newStatus
		update.CompletedAt = &now
	default:
		return nil, nil
	}

	return r.apiClient.UpdateTerraformExecutionRequest(ctx, terraformExecutionRequestID, update)
}
