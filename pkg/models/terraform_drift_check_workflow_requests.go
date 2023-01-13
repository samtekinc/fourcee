package models

import "time"

type TerraformDriftCheckRequest struct {
	TerraformDriftCheckRequestId         string
	ModuleAssignmentId                   string
	PlanExecutionRequestId               *string
	RequestTime                          time.Time
	Status                               RequestStatus
	SyncStatus                           TerraformDriftCheckStatus
	Destroy                              bool
	CallbackTaskToken                    *string
	ModulePropagationId                  *string `dynamodbav:",omitempty"`
	ModulePropagationDriftCheckRequestId *string `dynamodbav:",omitempty"`
}

type TerraformDriftCheckRequests struct {
	Items      []TerraformDriftCheckRequest
	NextCursor string
}

type NewTerraformDriftCheckRequest struct {
	ModuleAssignmentId                   string
	Destroy                              bool
	CallbackTaskToken                    *string
	ModulePropagationId                  *string
	ModulePropagationDriftCheckRequestId *string
}

type TerraformDriftCheckRequestUpdate struct {
	Status                 *RequestStatus
	PlanExecutionRequestId *string
	SyncStatus             *TerraformDriftCheckStatus
}

type TerraformDriftCheckStatus string

const (
	TerraformDriftCheckStatusPending   TerraformDriftCheckStatus = "PENDING"
	TerraformDriftCheckStatusInSync    TerraformDriftCheckStatus = "IN_SYNC"
	TerraformDriftCheckStatusOutOfSync TerraformDriftCheckStatus = "OUT_OF_SYNC"
)
