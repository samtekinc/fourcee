package models

import "time"

type TerraformExecutionRequest struct {
	TerraformExecutionRequestId         string
	ModuleAssignmentId                  string
	PlanExecutionRequestId              *string
	ApplyExecutionRequestId             *string
	RequestTime                         time.Time
	Status                              RequestStatus
	Destroy                             bool
	CallbackTaskToken                   *string
	ModulePropagationId                 *string `dynamodbav:",omitempty"`
	ModulePropagationExecutionRequestId *string `dynamodbav:",omitempty"`
}

type TerraformExecutionRequests struct {
	Items      []TerraformExecutionRequest
	NextCursor string
}

type NewTerraformExecutionRequest struct {
	ModuleAssignmentId                  string
	Destroy                             bool
	CallbackTaskToken                   *string
	ModulePropagationId                 *string
	ModulePropagationExecutionRequestId *string
}

type TerraformExecutionRequestUpdate struct {
	Status                  *RequestStatus
	PlanExecutionRequestId  *string
	ApplyExecutionRequestId *string
}
