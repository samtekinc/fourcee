package models

import "gorm.io/gorm"

type OrgDimension struct {
	gorm.Model
	Name                          string
	OrgUnitsAssociation           []OrgUnit
	ModulePropagationsAssociation []ModulePropagation
}

type NewOrgDimension struct {
	Name string
}

type OrgDimensionFilters struct {
	NameContains *string
}
