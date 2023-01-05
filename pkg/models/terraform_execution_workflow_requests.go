package models

import "time"

type TerraformExecutionWorkflowRequest struct {
	TerraformExecutionWorkflowRequestId string
	ModulePropagationExecutionRequestId string
	ModuleAccountAssociationKey         string
	PlanExecutionRequestId              *string
	ApplyExecutionRequestId             *string
	RequestTime                         time.Time
	Status                              RequestStatus
	Destroy                             bool
}

type TerraformExecutionWorkflowRequests struct {
	Items      []TerraformExecutionWorkflowRequest
	NextCursor string
}

type NewTerraformExecutionWorkflowRequest struct {
	ModulePropagationExecutionRequestId string
	ModuleAccountAssociationKey         string
	Destroy                             bool
}

type TerraformExecutionWorkflowRequestUpdate struct {
	Status                  *RequestStatus
	PlanExecutionRequestId  *string
	ApplyExecutionRequestId *string
}
