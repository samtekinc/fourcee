package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) UpdateTerraformExecutionApplyRequest(ctx context.Context, terraformExecutionRequestID uint, request *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	return r.apiClient.CreateApplyExecutionRequestForTerraformExecutionRequest(ctx, terraformExecutionRequestID, request)
}
