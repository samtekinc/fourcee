package models

import "time"

type TerraformExecutionWorkflowRequest struct {
	TerraformExecutionWorkflowRequestId string
	ModuleAssignmentId                  string
	PlanExecutionRequestId              *string
	ApplyExecutionRequestId             *string
	RequestTime                         time.Time
	Status                              RequestStatus
	Destroy                             bool
	CallbackTaskToken                   string
	ModulePropagationId                 *string `dynamodbav:",omitempty"`
	ModulePropagationExecutionRequestId *string `dynamodbav:",omitempty"`
}

type TerraformExecutionWorkflowRequests struct {
	Items      []TerraformExecutionWorkflowRequest
	NextCursor string
}

type NewTerraformExecutionWorkflowRequest struct {
	ModuleAssignmentId                  string
	Destroy                             bool
	CallbackTaskToken                   string
	ModulePropagationId                 *string
	ModulePropagationExecutionRequestId *string
}

type TerraformExecutionWorkflowRequestUpdate struct {
	Status                  *RequestStatus
	PlanExecutionRequestId  *string
	ApplyExecutionRequestId *string
}
