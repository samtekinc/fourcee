package models

type OrganizationalUnit struct {
	OrgUnitId       string
	Name            string
	OrgDimensionId  string
	Hierarchy       string
	ParentOrgUnitId string `dynamodbav:",omitempty"`
}

type OrganizationalUnits struct {
	Items      []OrganizationalUnit
	NextCursor string
}

type NewOrganizationalUnit struct {
	Name            string
	OrgDimensionId  string
	ParentOrgUnitId string
}

type OrganizationalUnitUpdate struct {
	Name            *string
	ParentOrgUnitId *string
}
