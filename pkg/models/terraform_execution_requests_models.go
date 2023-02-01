package models

import (
	"time"

	"gorm.io/gorm"
)

type TerraformExecutionRequest struct {
	gorm.Model
	ModuleAssignmentID                  uint
	PlanExecutionRequestAssociation     *PlanExecutionRequest
	ApplyExecutionRequestAssociation    *ApplyExecutionRequest
	StartedAt                           *time.Time
	CompletedAt                         *time.Time
	Status                              RequestStatus
	Destroy                             bool
	CallbackTaskToken                   *string
	ModulePropagationID                 *uint
	ModulePropagationExecutionRequestID *uint `gorm:"index"`
}

type NewTerraformExecutionRequest struct {
	ModuleAssignmentID                  uint
	Destroy                             bool
	CallbackTaskToken                   *string
	ModulePropagationID                 *uint
	ModulePropagationExecutionRequestID *uint
}

type TerraformExecutionRequestUpdate struct {
	Status      *RequestStatus
	StartedAt   *time.Time
	CompletedAt *time.Time
}

type TerraformExecutionRequestFilters struct {
	StartedBefore   *time.Time
	StartedAfter    *time.Time
	CompletedBefore *time.Time
	CompletedAfter  *time.Time
	Destroy         *bool
	Status          *RequestStatus
}
