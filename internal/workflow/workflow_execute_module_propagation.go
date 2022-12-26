package workflow

const (
	WorkflowExecuteModulePropagation = "ExecuteModulePropagation"
)

type ExecuteModulePropagationWorkflowPayload struct {
	ModulePropagationId                 string
	ModulePropagationExecutionRequestId string
}
