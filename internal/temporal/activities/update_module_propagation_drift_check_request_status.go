package activities

import (
	"context"
	"time"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) UpdateModulePropagationDriftCheckRequestStatus(ctx context.Context, modulePropagationDriftCheckRequestID uint, newStatus models.RequestStatus, syncStatus *models.TerraformDriftCheckStatus) (*models.ModulePropagationDriftCheckRequest, error) {
	update := &models.ModulePropagationDriftCheckRequestUpdate{
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

	return r.apiClient.UpdateModulePropagationDriftCheckRequest(ctx, modulePropagationDriftCheckRequestID, update)
}
