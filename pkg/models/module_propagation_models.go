package models

import (
	"strconv"

	"gorm.io/gorm"
)

type ModulePropagation struct {
	gorm.Model
	ModuleVersionID                                uint
	ModuleGroupID                                  uint
	OrgUnitID                                      uint
	OrgDimensionID                                 uint
	Name                                           string
	Description                                    string
	Arguments                                      []Argument
	AwsProviderConfigurations                      []AwsProviderConfiguration
	GcpProviderConfigurations                      []GcpProviderConfiguration
	ModuleAssignmentsAssociation                   []ModuleAssignment                   `gorm:"foreignKey:ModulePropagationID"`
	ModulePropagationExecutionRequestsAssociation  []ModulePropagationExecutionRequest  `gorm:"foreignKey:ModulePropagationID"`
	ModulePropagationDriftCheckRequestsAssociation []ModulePropagationDriftCheckRequest `gorm:"foreignKey:ModulePropagationID"`
}

type AwsProviderConfiguration struct {
	gorm.Model
	Region              string
	Alias               *string
	ModuleAssignmentID  *uint
	ModulePropagationID *uint
}

type GcpProviderConfiguration struct {
	gorm.Model
	Region              string
	Alias               *string
	ModuleAssignmentID  *uint
	ModulePropagationID *uint
}

type Argument struct {
	gorm.Model
	Name                string
	Value               string
	ModuleAssignmentID  *uint
	ModulePropagationID *uint
}

type AwsProviderConfigurationInput struct {
	Region string
	Alias  *string
}

type GcpProviderConfigurationInput struct {
	Region string
	Alias  *string
}

type ArgumentInput struct {
	Name  string
	Value string
}

type NewModulePropagation struct {
	ModuleVersionID           uint
	ModuleGroupID             uint
	OrgUnitID                 uint
	OrgDimensionID            uint
	Name                      string
	Description               string
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
}

type ModulePropagationUpdate struct {
	OrgDimensionID            *uint
	OrgUnitID                 *uint
	ModuleVersionID           *uint
	Name                      *string
	Description               *string
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
}

type ModulePropagationFilters struct {
	NameContains        *string
	DescriptionContains *string
}

func (m *ModulePropagation) GetInternalMetadata() []Metadata {
	return []Metadata{
		{
			Name:  "id",
			Value: strconv.FormatUint(uint64(m.ID), 10),
		},
		{
			Name:  "name",
			Value: m.Name,
		},
		{
			Name:  "description",
			Value: m.Description,
		},
	}
}
