package models

import "gorm.io/gorm"

type CloudAccessRole struct {
	gorm.Model
	Name                      string
	CloudPlatform             CloudPlatform
	AwsIamPoliciesAssociation []AwsIamPolicy `gorm:"many2many:cloud_access_roles_iam_policies;"`
	OrgUnitID                 uint           `gorm:"index"`
}

type NewCloudAccessRole struct {
	Name           string
	CloudPlatform  CloudPlatform
	AwsIamPolicies []uint
	OrgUnitID      uint
}

type CloudAccessRoleUpdate struct {
	Name           *string
	AwsIamPolicies *[]uint
	OrgUnitID      *uint
}

type CloudAccessRoleFilters struct {
	NameContains  *string
	CloudPlatform *CloudPlatform
}
