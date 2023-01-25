package models

import (
	"time"

	"gorm.io/gorm"
)

type ApplyExecutionRequest struct {
	gorm.Model
	ModuleAssignmentID           uint
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformConfigurationBase64 string
	TerraformPlanBase64          string
	AdditionalArguments          *string
	TerraformExecutionRequestID  uint
	Status                       RequestStatus
	InitOutput                   []byte
	ApplyOutput                  []byte
	StartedAt                    *time.Time
	CompletedAt                  *time.Time
}

type NewApplyExecutionRequest struct {
	ModuleAssignmentID           uint
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformConfigurationBase64 string
	TerraformExecutionRequestID  uint
	TerraformPlanBase64          string
	AdditionalArguments          *string
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
