package models

import (
	"time"

	"gorm.io/gorm"
)

type TerraformDriftCheckRequest struct {
	gorm.Model
	ModuleAssignmentID                   uint
	PlanExecutionRequestID               *uint
	StartedAt                            *time.Time
	CompletedAt                          *time.Time
	Status                               RequestStatus
	SyncStatus                           TerraformDriftCheckStatus
	Destroy                              bool
	CallbackTaskToken                    *string
	ModulePropagationID                  *uint
	ModulePropagationDriftCheckRequestID *uint
}

type NewTerraformDriftCheckRequest struct {
	ModuleAssignmentID                   uint
	Destroy                              bool
	CallbackTaskToken                    *string
	ModulePropagationID                  *uint
	ModulePropagationDriftCheckRequestID *uint
}

type TerraformDriftCheckRequestUpdate struct {
	Status                 *RequestStatus
	PlanExecutionRequestID *uint
	StartedAt              *time.Time
	CompletedAt            *time.Time
	SyncStatus             *TerraformDriftCheckStatus
}

type TerraformDriftCheckRequestFilters struct {
	StartedBefore   *time.Time
	StartedAfter    *time.Time
	CompletedBefore *time.Time
	CompletedAfter  *time.Time
	Destroy         *bool
	Status          *RequestStatus
	SyncStatus      *TerraformDriftCheckStatus
}

type TerraformDriftCheckStatus string

const (
	TerraformDriftCheckStatusPending   TerraformDriftCheckStatus = "PENDING"
	TerraformDriftCheckStatusInSync    TerraformDriftCheckStatus = "IN_SYNC"
	TerraformDriftCheckStatusOutOfSync TerraformDriftCheckStatus = "OUT_OF_SYNC"
)
