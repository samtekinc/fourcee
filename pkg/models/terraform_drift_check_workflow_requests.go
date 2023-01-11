package models

import "time"

type TerraformDriftCheckWorkflowRequest struct {
	TerraformDriftCheckWorkflowRequestId string
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

type TerraformDriftCheckWorkflowRequests struct {
	Items      []TerraformDriftCheckWorkflowRequest
	NextCursor string
}

type NewTerraformDriftCheckWorkflowRequest struct {
	ModuleAssignmentId                   string
	Destroy                              bool
	CallbackTaskToken                    *string
	ModulePropagationId                  *string
	ModulePropagationDriftCheckRequestId *string
}

type TerraformDriftCheckWorkflowRequestUpdate struct {
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
