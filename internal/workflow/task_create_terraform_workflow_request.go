package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskCreateTerraformWorkflowRequest Task = "CreateTerraformWorkflowRequest"
)

type CreateTerraformWorkflowRequestInput struct {
	ModulePropagationExecutionRequestId string
	ModulePropagationId                 string
	ModuleAccountAssociation            models.ModuleAccountAssociation
	Destroy                             bool
}

type CreateTerraformWorkflowRequestOutput struct {
	TerraformWorkflowRequestId string
}

func (t *TaskHandler) CreateTerraformWorkflowRequest(ctx context.Context, input CreateTerraformWorkflowRequestInput) (*CreateTerraformWorkflowRequestOutput, error) {
	tfWorkflow, err := t.apiClient.PutTerraformWorkflowRequest(ctx, &models.NewTerraformWorkflowRequest{
		ModulePropagationExecutionRequestId: input.ModulePropagationExecutionRequestId,
		ModuleAccountAssociationKey:         input.ModuleAccountAssociation.Key().String(),
		Destroy:                             input.Destroy,
	})

	if err != nil {
		return nil, err
	}

	return &CreateTerraformWorkflowRequestOutput{
		TerraformWorkflowRequestId: tfWorkflow.TerraformWorkflowRequestId,
	}, nil
}
