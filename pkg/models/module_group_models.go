package models

import "gorm.io/gorm"

type ModuleGroup struct {
	gorm.Model
	Name                          string
	CloudPlatform                 CloudPlatform
	ModuleVersionsAssociation     []ModuleVersion     `gorm:"foreignKey:ModuleGroupID"`
	ModulePropagationsAssociation []ModulePropagation `gorm:"foreignKey:ModuleGroupID"`
	ModuleAssignmentsAssociation  []ModuleAssignment  `gorm:"foreignKey:ModuleGroupID"`
}

type NewModuleGroup struct {
	Name          string
	CloudPlatform CloudPlatform
}

type ModuleGroupFilters struct {
	NameContains  *string
	CloudPlatform *CloudPlatform
}
