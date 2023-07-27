package models

import "gorm.io/gorm"

type OrgUnit struct {
	gorm.Model
	Name                          string
	OrgDimensionID                uint                `gorm:"index"`
	ParentOrgUnitID               *uint               `gorm:"index"`
	Hierarchy                     string              `gorm:"index"`
	ChildOrgUnitsAssociation      []OrgUnit           `gorm:"foreignKey:ParentOrgUnitID"`
	OrgAccountsAssociation        []OrgAccount        `gorm:"many2many:org_accounts_org_units;"`
	ModulePropagationsAssociation []ModulePropagation `gorm:"foreignKey:OrgUnitID"`
	CloudAccessRolesAssociation   []CloudAccessRole   `gorm:"foreignKey:OrgUnitID"`
}

type NewOrgUnit struct {
	Name            string
	OrgDimensionID  uint
	ParentOrgUnitID uint
}

type OrgUnitUpdate struct {
	Name            *string
	ParentOrgUnitID *uint
}

type OrgUnitFilters struct {
	NameContains *string
}
