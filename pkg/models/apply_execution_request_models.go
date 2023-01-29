package models

import (
	"time"

	"gorm.io/gorm"
)

type ApplyExecutionRequest struct {
	gorm.Model
	ModuleAssignmentID          uint
	TerraformVersion            string
	CallbackTaskToken           string
	TerraformConfiguration      []byte
	TerraformPlan               []byte
	AdditionalArguments         *string
	TerraformExecutionRequestID uint
	Status                      RequestStatus
	InitOutput                  []byte
	ApplyOutput                 []byte
	StartedAt                   *time.Time
	CompletedAt                 *time.Time
}

type NewApplyExecutionRequest struct {
	ModuleAssignmentID          uint
	TerraformVersion            string
	CallbackTaskToken           string
	TerraformConfiguration      []byte
	TerraformExecutionRequestID uint
	TerraformPlan               []byte
	AdditionalArguments         *string
}

type ApplyExecutionRequestUpdate struct {
	InitOutput  []byte
	ApplyOutput []byte
	StartedAt   *time.Time
	CompletedAt *time.Time
	Status      *RequestStatus
}

type ApplyExecutionRequestFilters struct {
	StartedBefore   *time.Time
	StartedAfter    *time.Time
	CompletedBefore *time.Time
	CompletedAfter  *time.Time
	Destroy         *bool
	Status          *RequestStatus
}
