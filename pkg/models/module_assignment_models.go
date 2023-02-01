package models

import "gorm.io/gorm"

type ModuleAssignment struct {
	gorm.Model
	ModuleVersionID                        uint `gorm:"index"`
	ModuleGroupID                          uint `gorm:"index"`
	OrgAccountID                           uint `gorm:"index"`
	Name                                   string
	Description                            string
	RemoteStateRegion                      string
	RemoteStateBucket                      string
	RemoteStateKey                         string
	Arguments                              []Argument                 `gorm:"serializer:json"`
	AwsProviderConfigurations              []AwsProviderConfiguration `gorm:"serializer:json"`
	GcpProviderConfigurations              []GcpProviderConfiguration `gorm:"serializer:json"`
	ModulePropagationID                    *uint                      `gorm:"index"`
	Status                                 ModuleAssignmentStatus
	TerraformExecutionRequestsAssociation  []*TerraformExecutionRequest  `gorm:"foreignKey:ModuleAssignmentID"`
	TerraformDriftCheckRequestsAssociation []*TerraformDriftCheckRequest `gorm:"foreignKey:ModuleAssignmentID"`
}

type ModuleAssignmentStatus string

const (
	ModuleAssignmentStatusActive   ModuleAssignmentStatus = "ACTIVE"
	ModuleAssignmentStatusInactive ModuleAssignmentStatus = "INACTIVE"
)

type NewModuleAssignment struct {
	ModuleVersionID           uint
	ModuleGroupID             uint
	OrgAccountID              uint
	Name                      string
	Description               string
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
	ModulePropagationID       *uint
}

type ModuleAssignmentUpdate struct {
	Name                      *string
	Description               *string
	ModuleVersionID           *uint
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
	Status                    *ModuleAssignmentStatus
}

type ModuleAssignmentFilters struct {
	NameContains        *string
	DescriptionContains *string
	Status              *ModuleAssignmentStatus
	IsPropagated        *bool
	OrgAccountID        *uint
}
