package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskTallySyncStatus Task = "TallySyncStatus"
)

type TallySyncStatusInput struct {
	ModulePropagationId                  string
	ModulePropagationDriftCheckRequestId string
}

type TallySyncStatusOutput struct{}

func (t *TaskHandler) TallySyncStatus(ctx context.Context, input TallySyncStatusInput) (*TallySyncStatusOutput, error) {
	// get terraform drift check requests for this module propagation request
	outOfSyncRequests := []models.TerraformDriftCheckRequest{}
	nextCursor := ""
	for {
		driftCheckRequests, err := t.apiClient.GetTerraformDriftCheckRequestsByModulePropagationDriftCheckRequestId(ctx, input.ModulePropagationDriftCheckRequestId, 1000, nextCursor)
		if err != nil {
			return nil, err
		}
		for _, driftCheckRequest := range driftCheckRequests.Items {
			if driftCheckRequest.SyncStatus == models.TerraformDriftCheckStatusOutOfSync {
				outOfSyncRequests = append(outOfSyncRequests, driftCheckRequest)
				break
			}
		}
		if driftCheckRequests.NextCursor == "" {
			break
		}
		nextCursor = driftCheckRequests.NextCursor
	}

	// update the module propagation request sync status
	var newStatus models.TerraformDriftCheckStatus
	if len(outOfSyncRequests) > 0 {
		newStatus = models.TerraformDriftCheckStatusOutOfSync
	} else {
		newStatus = models.TerraformDriftCheckStatusInSync
	}
	_, err := t.apiClient.UpdateModulePropagationDriftCheckRequest(ctx, input.ModulePropagationId, input.ModulePropagationDriftCheckRequestId, &models.ModulePropagationDriftCheckRequestUpdate{
		SyncStatus: &newStatus,
	})
	if err != nil {
		return nil, err
	}

	// send alert if there are out of sync requests
	if len(outOfSyncRequests) > 0 {
		subject := "TFOM Drift Detected"
		message := "The following module assignments are out of sync:\n\n"
		for _, outOfSyncRequest := range outOfSyncRequests {
			message += outOfSyncRequest.ModuleAssignmentId + "\n"
		}
		err = t.apiClient.SendAlert(ctx, subject, message)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
