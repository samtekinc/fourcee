package activities

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

func (r *Activities) UpdateTerraformExecutionApplyRequest(ctx context.Context, terraformExecutionRequestID uint, request *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	return r.apiClient.CreateApplyExecutionRequestForTerraformExecutionRequest(ctx, terraformExecutionRequestID, request)
}
