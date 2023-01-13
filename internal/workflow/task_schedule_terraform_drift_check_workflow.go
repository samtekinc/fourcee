package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformDriftCheck Task = "ScheduleTerraformDriftCheck"
)

type ScheduleTerraformDriftCheckInput struct {
	ModulePropagationDriftCheckRequestId string
	ModulePropagationId                  string
	ModuleAssignment                     models.ModuleAssignment
	Destroy                              bool
	TaskToken                            *string
}

type ScheduleTerraformDriftCheckOutput struct {
	TerraformDriftCheckRequestId string
}

func (t *TaskHandler) ScheduleTerraformDriftCheck(ctx context.Context, input ScheduleTerraformDriftCheckInput) (*ScheduleTerraformDriftCheckOutput, error) {
	tfWorkflow, err := t.apiClient.PutTerraformDriftCheckRequest(ctx, &models.NewTerraformDriftCheckRequest{
		ModuleAssignmentId:                   input.ModuleAssignment.ModuleAssignmentId,
		ModulePropagationDriftCheckRequestId: &input.ModulePropagationDriftCheckRequestId,
		ModulePropagationId:                  &input.ModulePropagationId,
		Destroy:                              input.Destroy,
		CallbackTaskToken:                    input.TaskToken,
	})

	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformDriftCheckOutput{
		TerraformDriftCheckRequestId: tfWorkflow.TerraformDriftCheckRequestId,
	}, nil
}
