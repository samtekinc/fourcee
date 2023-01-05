package models

import "time"

type TerraformDriftCheckWorkflowRequest struct {
	TerraformDriftCheckWorkflowRequestId string
	ModulePropagationDriftCheckRequestId string
	ModuleAccountAssociationKey          string
	PlanExecutionRequestId               *string
	RequestTime                          time.Time
	Status                               RequestStatus
	SyncStatus                           TerraformDriftCheckStatus
	Destroy                              bool
}

type TerraformDriftCheckWorkflowRequests struct {
	Items      []TerraformDriftCheckWorkflowRequest
	NextCursor string
}

type NewTerraformDriftCheckWorkflowRequest struct {
	ModulePropagationDriftCheckRequestId string
	ModuleAccountAssociationKey          string
	Destroy                              bool
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
