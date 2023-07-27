package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) UpdateTerraformDriftCheckPlanRequest(ctx context.Context, terraformDriftCheckRequestID uint, request *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	return r.apiClient.CreatePlanExecutionRequestForTerraformDriftCheckRequest(ctx, terraformDriftCheckRequestID, request)
}
