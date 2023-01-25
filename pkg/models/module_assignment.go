package models

import "gorm.io/gorm"

type ModuleAssignment struct {
	gorm.Model
	ModuleVersionID                        uint
	ModuleGroupID                          uint
	OrgAccountID                           uint
	Name                                   string
	Description                            string
	RemoteStateRegion                      string
	RemoteStateBucket                      string
	RemoteStateKey                         string
	Arguments                              []Argument
	AwsProviderConfigurations              []AwsProviderConfiguration
	GcpProviderConfigurations              []GcpProviderConfiguration
	ModulePropagationID                    *uint
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
}
