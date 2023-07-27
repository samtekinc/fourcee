package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) CreateTerraformExecutionRequest(ctx context.Context, moduleAssignmentID uint, destroy bool, modulePropagationID uint, modulePropagationExecutionRequestID uint) (*models.TerraformExecutionRequest, error) {
	return r.apiClient.CreateTerraformExecutionRequest(ctx, &models.NewTerraformExecutionRequest{
		ModuleAssignmentID:                  moduleAssignmentID,
		Destroy:                             destroy,
		CallbackTaskToken:                   nil,
		ModulePropagationID:                 &modulePropagationID,
		ModulePropagationExecutionRequestID: &modulePropagationExecutionRequestID,
	}, false)
}
