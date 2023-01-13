package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformExecution Task = "ScheduleTerraformExecution"
)

type ScheduleTerraformExecutionInput struct {
	ModulePropagationExecutionRequestId string
	ModulePropagationId                 string
	ModuleAssignment                    models.ModuleAssignment
	Destroy                             bool
	TaskToken                           *string
}

type ScheduleTerraformExecutionOutput struct {
	TerraformExecutionRequestId string
}

func (t *TaskHandler) ScheduleTerraformExecution(ctx context.Context, input ScheduleTerraformExecutionInput) (*ScheduleTerraformExecutionOutput, error) {
	tfWorkflow, err := t.apiClient.PutTerraformExecutionRequest(ctx, &models.NewTerraformExecutionRequest{
		ModuleAssignmentId:                  input.ModuleAssignment.ModuleAssignmentId,
		ModulePropagationExecutionRequestId: &input.ModulePropagationExecutionRequestId,
		ModulePropagationId:                 &input.ModulePropagationId,
		Destroy:                             input.Destroy,
		CallbackTaskToken:                   input.TaskToken,
	})

	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformExecutionOutput{
		TerraformExecutionRequestId: tfWorkflow.TerraformExecutionRequestId,
	}, nil
}
