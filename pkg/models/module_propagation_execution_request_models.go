package models

import (
	"time"

	"gorm.io/gorm"
)

type ModulePropagationExecutionRequest struct {
	gorm.Model
	ModulePropagationID                   uint `gorm:"index"`
	StartedAt                             *time.Time
	CompletedAt                           *time.Time
	Status                                RequestStatus
	TerraformExecutionRequestsAssociation []*TerraformExecutionRequest `gorm:"foreignKey:ModulePropagationExecutionRequestID"`
}

type NewModulePropagationExecutionRequest struct {
	ModulePropagationID uint
}

type ModulePropagationExecutionRequestUpdate struct {
	Status      *RequestStatus
	StartedAt   *time.Time
	CompletedAt *time.Time
}

type ModulePropagationExecutionRequestFilters struct {
	StartedBefore   *time.Time
	StartedAfter    *time.Time
	CompletedBefore *time.Time
	CompletedAfter  *time.Time
	Status          *RequestStatus
}
