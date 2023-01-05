package models

type OrganizationalDimension struct {
	OrgDimensionId string
	Name           string
	RootOrgUnitId  string
}

type OrganizationalDimensions struct {
	Items      []OrganizationalDimension
	NextCursor string
}

type NewOrganizationalDimension struct {
	Name string
}
