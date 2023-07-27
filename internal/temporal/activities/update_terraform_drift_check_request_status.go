package activities

import (
	"context"
	"time"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) UpdateTerraformDriftCheckRequestStatus(ctx context.Context, terraformDriftCheckRequestID uint, newStatus models.RequestStatus, syncStatus *models.TerraformDriftCheckStatus) (*models.TerraformDriftCheckRequest, error) {
	update := &models.TerraformDriftCheckRequestUpdate{
		Status: &newStatus,
	}

	if syncStatus != nil {
		update.SyncStatus = syncStatus
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

	return r.apiClient.UpdateTerraformDriftCheckRequest(ctx, terraformDriftCheckRequestID, update)
}
