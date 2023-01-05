package models

type OrganizationalUnitMembership struct {
	OrgAccountId   string
	OrgDimensionId string
	OrgUnitId      string
}

type OrganizationalUnitMemberships struct {
	Items      []OrganizationalUnitMembership
	NextCursor string
}

type NewOrganizationalUnitMembership struct {
	OrgAccountId   string
	OrgDimensionId string
	OrgUnitId      string
}
