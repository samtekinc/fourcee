package models

import (
	"strconv"

	"gorm.io/gorm"
)

type ModuleVersion struct {
	gorm.Model
	ModuleGroupID                 uint
	Name                          string
	RemoteSource                  string
	TerraformVersion              string
	Variables                     []*ModuleVariable
	ModulePropagationsAssociation []ModulePropagation `gorm:"foreignKey:ModuleVersionID"`
	ModuleAssignmentsAssociation  []ModuleAssignment  `gorm:"foreignKey:ModuleVersionID"`
}

type NewModuleVersion struct {
	ModuleGroupID    uint
	Name             string
	RemoteSource     string
	TerraformVersion string
	CloudPlatform    string
}

type ModuleVariable struct {
	gorm.Model
	ModuleVersionID uint // Foreign key
	Name            string
	Type            string
	Description     string
	Default         *string
}

type ModuleVersionFilters struct {
	NameContains         *string
	RemoteSourceContains *string
	TerraformVersion     *string
}

func (m *ModuleVersion) GetInternalMetadata() []Metadata {
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
			Name:  "remote_source",
			Value: m.RemoteSource,
		},
		{
			Name:  "terraform_version",
			Value: m.TerraformVersion,
		},
	}
}
