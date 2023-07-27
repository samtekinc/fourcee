package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) CreateTerraformDriftCheckRequest(ctx context.Context, moduleAssignmentID uint, destroy bool, modulePropagationID uint, modulePropagationDriftCheckRequestID uint) (*models.TerraformDriftCheckRequest, error) {
	return r.apiClient.CreateTerraformDriftCheckRequest(ctx, &models.NewTerraformDriftCheckRequest{
		ModuleAssignmentID:                   moduleAssignmentID,
		Destroy:                              destroy,
		CallbackTaskToken:                    nil,
		ModulePropagationID:                  &modulePropagationID,
		ModulePropagationDriftCheckRequestID: &modulePropagationDriftCheckRequestID,
	}, false)
}
