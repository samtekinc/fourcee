package models

import (
	"strconv"

	"gorm.io/gorm"
)

type OrgAccount struct {
	gorm.Model
	Name                         string
	CloudPlatform                CloudPlatform
	CloudIdentifier              string
	AssumeRoleName               string
	Metadata                     []Metadata
	OrgUnitsAssociation          []OrgUnit          `gorm:"many2many:org_accounts_org_units;"`
	ModuleAssignmentsAssociation []ModuleAssignment `gorm:"foreignKey:OrgAccountID"`
}

type NewOrgAccount struct {
	Name            string
	CloudPlatform   CloudPlatform
	CloudIdentifier string
	AssumeRoleName  string
	Metadata        []MetadataInput
}

type OrgAccountUpdate struct {
	Name            *string
	CloudPlatform   *CloudPlatform
	CloudIdentifier *string
	AssumeRoleName  *string
	Metadata        []MetadataInput
}

type OrgAccountFilters struct {
	NameContains    *string
	CloudPlatform   *CloudPlatform
	CloudIdentifier *string
}

func (a *OrgAccount) GetInternalMetadata() []Metadata {
	return []Metadata{
		{
			Name:  "id",
			Value: strconv.FormatUint(uint64(a.ID), 10),
		},
		{
			Name:  "name",
			Value: a.Name,
		},
		{
			Name:  "cloud_identifier",
			Value: a.CloudIdentifier,
		},
	}
}
