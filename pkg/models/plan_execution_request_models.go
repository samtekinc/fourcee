package models

import (
	"time"

	"gorm.io/gorm"
)

type PlanExecutionRequest struct {
	gorm.Model
	ModuleAssignmentID           uint
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformConfiguration       []byte
	AdditionalArguments          *string
	TerraformDriftCheckRequestID *uint
	TerraformExecutionRequestID  *uint
	Status                       RequestStatus
	InitOutput                   []byte
	PlanOutput                   []byte
	PlanFile                     []byte
	PlanJSON                     []byte
	StartedAt                    *time.Time
	CompletedAt                  *time.Time
}

type NewPlanExecutionRequest struct {
	ModuleAssignmentID           uint
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformDriftCheckRequestID *uint
	TerraformExecutionRequestID  *uint
	TerraformConfiguration       []byte
	AdditionalArguments          *string
}

type PlanExecutionRequestUpdate struct {
	InitOutput  []byte
	PlanOutput  []byte
	PlanFile    []byte
	PlanJSON    []byte
	StartedAt   *time.Time
	CompletedAt *time.Time
	Status      *RequestStatus
}

type PlanExecutionRequestFilters struct {
	StartedBefore   *time.Time
	StartedAfter    *time.Time
	CompletedBefore *time.Time
	CompletedAfter  *time.Time
	Destroy         *bool
	Status          *RequestStatus
}
