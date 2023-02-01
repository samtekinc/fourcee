package models

import (
	"strconv"

	"gorm.io/gorm"
)

type ModulePropagation struct {
	gorm.Model
	ModuleVersionID                                uint `gorm:"index"`
	ModuleGroupID                                  uint `gorm:"index"`
	OrgUnitID                                      uint `gorm:"index"`
	OrgDimensionID                                 uint `gorm:"index"`
	Name                                           string
	Description                                    string
	Arguments                                      []Argument                           `gorm:"serializer:json"`
	AwsProviderConfigurations                      []AwsProviderConfiguration           `gorm:"serializer:json"`
	GcpProviderConfigurations                      []GcpProviderConfiguration           `gorm:"serializer:json"`
	ModuleAssignmentsAssociation                   []ModuleAssignment                   `gorm:"foreignKey:ModulePropagationID"`
	ModulePropagationExecutionRequestsAssociation  []ModulePropagationExecutionRequest  `gorm:"foreignKey:ModulePropagationID"`
	ModulePropagationDriftCheckRequestsAssociation []ModulePropagationDriftCheckRequest `gorm:"foreignKey:ModulePropagationID"`
}

type AwsProviderConfiguration struct {
	Region string
	Alias  *string
}

type GcpProviderConfiguration struct {
	Region string
	Alias  *string
}

type Argument struct {
	Name  string
	Value string
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
