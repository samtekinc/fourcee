package models

import (
	"time"

	"gorm.io/gorm"
)

type TerraformExecutionRequest struct {
	gorm.Model
	ModuleAssignmentID                  uint
	PlanExecutionRequestID              *uint
	ApplyExecutionRequestID             *uint
	StartedAt                           *time.Time
	CompletedAt                         *time.Time
	Status                              RequestStatus
	Destroy                             bool
	CallbackTaskToken                   *string
	ModulePropagationID                 *uint
	ModulePropagationExecutionRequestID *uint
}

type NewTerraformExecutionRequest struct {
	ModuleAssignmentID                  uint
	Destroy                             bool
	CallbackTaskToken                   *string
	ModulePropagationID                 *uint
	ModulePropagationExecutionRequestID *uint
}

type TerraformExecutionRequestUpdate struct {
	Status                  *RequestStatus
	PlanExecutionRequestID  *uint
	StartedAt               *time.Time
	CompletedAt             *time.Time
	ApplyExecutionRequestID *uint
}

type TerraformExecutionRequestFilters struct {
	StartedBefore   *time.Time
	StartedAfter    *time.Time
	CompletedBefore *time.Time
	CompletedAfter  *time.Time
	Destroy         *bool
	Status          *RequestStatus
}
