package models

import "time"

type ModulePropagationExecutionRequestStatus string

const (
	ModulePropagationExecutionRequestStatusPending   ModulePropagationExecutionRequestStatus = "PENDING"
	ModulePropagationExecutionRequestStatusRunning   ModulePropagationExecutionRequestStatus = "RUNNING"
	ModulePropagationExecutionRequestStatusSucceeded ModulePropagationExecutionRequestStatus = "SUCCEEDED"
	ModulePropagationExecutionRequestStatusFailed    ModulePropagationExecutionRequestStatus = "FAILED"
)

type ModulePropagationExecutionRequest struct {
	ModulePropagationExecutionRequestId string                                  `json:"modulePropagationExecutionRequestId"`
	ModulePropagationId                 string                                  `json:"modulePropagationId"`
	RequestTime                         time.Time                               `json:"requestTime"`
	Status                              ModulePropagationExecutionRequestStatus `json:"status"`
	WorkflowExecutionId                 string                                  `json:"-"`
}

type ModulePropagationExecutionRequests struct {
	Items      []ModulePropagationExecutionRequest `json:"items"`
	NextCursor string                              `json:"nextCursor"`
}

type NewModulePropagationExecutionRequest struct {
	ModulePropagationId string `json:"modulePropagationId"`
}

type ModulePropagationExecutionRequestUpdate struct {
	Status *ModulePropagationExecutionRequestStatus `json:"status"`
}
