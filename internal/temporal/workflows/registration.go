package workflows

import "go.temporal.io/sdk/worker"

func RegisterWorkflows(w worker.Worker) {
	w.RegisterWorkflow(TerraformExecutionWorkflow)
	w.RegisterWorkflow(TerraformDriftCheckWorkflow)
	w.RegisterWorkflow(ModulePropagationExecutionWorkflow)
	w.RegisterWorkflow(ModulePropagationDriftCheckWorkflow)
}
