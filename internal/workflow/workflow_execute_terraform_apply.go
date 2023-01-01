package workflow

import "github.com/sheacloud/tfom/pkg/models"

const (
	WorkflowExecuteTerraformApply = "ExecuteTerraformApply"
)

type ExecuteTerraformApplyWorkflowPayload struct {
	ModuleAccountAssociation            models.ModuleAccountAssociation
	ModulePropagationExecutionRequestId string
	ModulePropagationId                 string
	Destroy                             bool
}
