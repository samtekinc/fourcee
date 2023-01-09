package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformDriftCheckWorkflow Task = "ScheduleTerraformDriftCheckWorkflow"
)

type ScheduleTerraformDriftCheckWorkflowInput struct {
	ModulePropagationDriftCheckRequestId string
	ModulePropagationId                  string
	ModuleAssignment                     models.ModuleAssignment
	Destroy                              bool
	TaskToken                            string
}

type ScheduleTerraformDriftCheckWorkflowOutput struct {
	TerraformDriftCheckWorkflowRequestId string
}

func (t *TaskHandler) ScheduleTerraformDriftCheckWorkflow(ctx context.Context, input ScheduleTerraformDriftCheckWorkflowInput) (*ScheduleTerraformDriftCheckWorkflowOutput, error) {
	tfWorkflow, err := t.apiClient.PutTerraformDriftCheckWorkflowRequest(ctx, &models.NewTerraformDriftCheckWorkflowRequest{
		ModuleAssignmentId:                   input.ModuleAssignment.ModuleAssignmentId,
		ModulePropagationDriftCheckRequestId: &input.ModulePropagationDriftCheckRequestId,
		ModulePropagationId:                  &input.ModulePropagationId,
		Destroy:                              input.Destroy,
		CallbackTaskToken:                    input.TaskToken,
	})

	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformDriftCheckWorkflowOutput{
		TerraformDriftCheckWorkflowRequestId: tfWorkflow.TerraformDriftCheckWorkflowRequestId,
	}, nil
}
