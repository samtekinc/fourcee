package models

import (
	"time"

	"gorm.io/gorm"
)

type ModulePropagationDriftCheckRequest struct {
	gorm.Model
	ModulePropagationID                    uint `gorm:"index"`
	StartedAt                              *time.Time
	CompletedAt                            *time.Time
	Status                                 RequestStatus
	SyncStatus                             TerraformDriftCheckStatus
	TerraformDriftCheckRequestsAssociation []*TerraformDriftCheckRequest `gorm:"foreignKey:ModulePropagationDriftCheckRequestID"`
}

type NewModulePropagationDriftCheckRequest struct {
	ModulePropagationID uint
}

type ModulePropagationDriftCheckRequestUpdate struct {
	Status      *RequestStatus
	StartedAt   *time.Time
	CompletedAt *time.Time
	SyncStatus  *TerraformDriftCheckStatus
}

type ModulePropagationDriftCheckRequestFilters struct {
	StartedBefore   *time.Time
	StartedAfter    *time.Time
	CompletedBefore *time.Time
	CompletedAfter  *time.Time
	Status          *RequestStatus
	SyncStatus      *TerraformDriftCheckStatus
}
