package models

import "time"

type ModulePropagationExecutionRequest struct {
	ModulePropagationExecutionRequestId string
	ModulePropagationId                 string
	RequestTime                         time.Time
	Status                              RequestStatus
	WorkflowExecutionId                 string
}

type ModulePropagationExecutionRequests struct {
	Items      []ModulePropagationExecutionRequest
	NextCursor string
}

type NewModulePropagationExecutionRequest struct {
	ModulePropagationId string
}

type ModulePropagationExecutionRequestUpdate struct {
	Status *RequestStatus
}
