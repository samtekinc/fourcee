package workflow

import "github.com/sheacloud/tfom/pkg/models"

const (
	WorkflowExecuteTerraformApply = "ExecuteTerraformApply"
)

type ExecuteTerraformApplyWorkflowPayload struct {
	ModuleAssignment                    models.ModuleAssignment
	ModulePropagationExecutionRequestId string
	ModulePropagationId                 string
	Destroy                             bool
}
