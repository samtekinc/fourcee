package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformExecutionWorkflow Task = "ScheduleTerraformExecutionWorkflow"
)

type ScheduleTerraformExecutionWorkflowInput struct {
	ModulePropagationExecutionRequestId string
	ModulePropagationId                 string
	ModuleAssignment                    models.ModuleAssignment
	Destroy                             bool
	TaskToken                           string
}

type ScheduleTerraformExecutionWorkflowOutput struct {
	TerraformExecutionWorkflowRequestId string
}

func (t *TaskHandler) ScheduleTerraformExecutionWorkflow(ctx context.Context, input ScheduleTerraformExecutionWorkflowInput) (*ScheduleTerraformExecutionWorkflowOutput, error) {
	tfWorkflow, err := t.apiClient.PutTerraformExecutionWorkflowRequest(ctx, &models.NewTerraformExecutionWorkflowRequest{
		ModuleAssignmentId:                  input.ModuleAssignment.ModuleAssignmentId,
		ModulePropagationExecutionRequestId: &input.ModulePropagationExecutionRequestId,
		ModulePropagationId:                 &input.ModulePropagationId,
		Destroy:                             input.Destroy,
		CallbackTaskToken:                   input.TaskToken,
	})

	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformExecutionWorkflowOutput{
		TerraformExecutionWorkflowRequestId: tfWorkflow.TerraformExecutionWorkflowRequestId,
	}, nil
}
