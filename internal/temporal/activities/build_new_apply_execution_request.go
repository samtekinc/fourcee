package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) BuildNewApplyExecutionRequest(ctx context.Context, moduleAssignmentID uint, terraformExecutionRequestID uint) (*models.NewApplyExecutionRequest, error) {
	planExecutionRequest, err := r.apiClient.GetPlanExecutionRequestForTerraformExecutionRequest(ctx, terraformExecutionRequestID)
	if err != nil {
		return nil, err
	}

	applyExecutionRequest := &models.NewApplyExecutionRequest{
		ModuleAssignmentID:          moduleAssignmentID,
		CallbackTaskToken:           "",
		TerraformVersion:            planExecutionRequest.TerraformVersion,
		TerraformConfiguration:      planExecutionRequest.TerraformConfiguration,
		TerraformExecutionRequestID: terraformExecutionRequestID,
		TerraformPlan:               planExecutionRequest.PlanFile,
		AdditionalArguments:         nil,
	}

	return applyExecutionRequest, nil
}
