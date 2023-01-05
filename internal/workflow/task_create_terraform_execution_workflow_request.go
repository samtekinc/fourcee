package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskCreateTerraformExecutionWorkflowRequest Task = "CreateTerraformExecutionWorkflowRequest"
)

type CreateTerraformExecutionWorkflowRequestInput struct {
	ModulePropagationExecutionRequestId string
	ModulePropagationId                 string
	ModuleAccountAssociation            models.ModuleAccountAssociation
	Destroy                             bool
}

type CreateTerraformExecutionWorkflowRequestOutput struct {
	TerraformExecutionWorkflowRequestId string
}

func (t *TaskHandler) CreateTerraformExecutionWorkflowRequest(ctx context.Context, input CreateTerraformExecutionWorkflowRequestInput) (*CreateTerraformExecutionWorkflowRequestOutput, error) {
	tfWorkflow, err := t.apiClient.PutTerraformExecutionWorkflowRequest(ctx, &models.NewTerraformExecutionWorkflowRequest{
		ModulePropagationExecutionRequestId: input.ModulePropagationExecutionRequestId,
		ModuleAccountAssociationKey:         input.ModuleAccountAssociation.Key().String(),
		Destroy:                             input.Destroy,
	})

	if err != nil {
		return nil, err
	}

	return &CreateTerraformExecutionWorkflowRequestOutput{
		TerraformExecutionWorkflowRequestId: tfWorkflow.TerraformExecutionWorkflowRequestId,
	}, nil
}
