package activities

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

func (r *Activities) UpdateTerraformExecutionPlanRequest(ctx context.Context, terraformExecutionRequestID uint, request *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	return r.apiClient.CreatePlanExecutionRequestForTerraformExecutionRequest(ctx, terraformExecutionRequestID, request)
}
