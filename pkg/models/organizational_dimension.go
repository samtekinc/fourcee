package models

import "gorm.io/gorm"

type OrgDimension struct {
	gorm.Model
	Name                          string
	RootOrgUnitID                 uint
	RootOrgUnitAssociation        OrgUnit             `gorm:"foreignKey:RootOrgUnitID"`
	OrgUnitsAssociation           []OrgUnit           `gorm:"foreignKey:OrgDimensionID"`
	ModulePropagationsAssociation []ModulePropagation `gorm:"foreignKey:OrgDimensionID"`
}

type NewOrgDimension struct {
	Name string
}

type OrgDimensionFilters struct {
	NameContains *string
}
