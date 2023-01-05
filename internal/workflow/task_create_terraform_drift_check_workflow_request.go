package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskCreateTerraformDriftCheckWorkflowRequest Task = "CreateTerraformDriftCheckWorkflowRequest"
)

type CreateTerraformDriftCheckWorkflowRequestInput struct {
	ModulePropagationDriftCheckRequestId string
	ModulePropagationId                  string
	ModuleAccountAssociation             models.ModuleAccountAssociation
	Destroy                              bool
}

type CreateTerraformDriftCheckWorkflowRequestOutput struct {
	TerraformDriftCheckWorkflowRequestId string
}

func (t *TaskHandler) CreateTerraformDriftCheckWorkflowRequest(ctx context.Context, input CreateTerraformDriftCheckWorkflowRequestInput) (*CreateTerraformDriftCheckWorkflowRequestOutput, error) {
	tfWorkflow, err := t.apiClient.PutTerraformDriftCheckWorkflowRequest(ctx, &models.NewTerraformDriftCheckWorkflowRequest{
		ModulePropagationDriftCheckRequestId: input.ModulePropagationDriftCheckRequestId,
		ModuleAccountAssociationKey:          input.ModuleAccountAssociation.Key().String(),
		Destroy:                              input.Destroy,
	})

	if err != nil {
		return nil, err
	}

	return &CreateTerraformDriftCheckWorkflowRequestOutput{
		TerraformDriftCheckWorkflowRequestId: tfWorkflow.TerraformDriftCheckWorkflowRequestId,
	}, nil
}
