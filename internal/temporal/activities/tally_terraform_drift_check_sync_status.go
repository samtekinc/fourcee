package activities

import (
	"context"
	"fmt"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) TallyTerraformDriftCheckSyncStatus(ctx context.Context, modulePropagationDriftCheckRequestID uint) (*models.TerraformDriftCheckStatus, error) {
	// get terraform drift check requests
	driftCheckRequests, err := r.apiClient.GetTerraformDriftCheckRequestsForModulePropagationDriftCheckRequest(ctx, modulePropagationDriftCheckRequestID, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	for _, driftCheckRequest := range driftCheckRequests {
		switch driftCheckRequest.SyncStatus {
		case models.TerraformDriftCheckStatusPending:
			return nil, fmt.Errorf("terraform drift check request %d is still pending", driftCheckRequest.ID)
		case models.TerraformDriftCheckStatusOutOfSync:
			return &driftCheckRequest.SyncStatus, nil
		default:
			continue
		}
	}

	inSync := models.TerraformDriftCheckStatusInSync
	return &inSync, nil
}
