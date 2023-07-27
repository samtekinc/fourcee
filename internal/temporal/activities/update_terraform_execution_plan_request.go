package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) UpdateTerraformExecutionPlanRequest(ctx context.Context, terraformExecutionRequestID uint, request *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	return r.apiClient.CreatePlanExecutionRequestForTerraformExecutionRequest(ctx, terraformExecutionRequestID, request)
}
