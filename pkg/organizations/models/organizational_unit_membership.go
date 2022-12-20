package models

type OrganizationalUnitMembership struct {
	OrgAccountId   string `json:"orgAccountId"`
	OrgDimensionId string `json:"orgDimensionId"`
	OrgUnitId      string `json:"orgUnitId"`
}

type OrganizationalUnitMemberships struct {
	Items      []OrganizationalUnitMembership `json:"items"`
	NextCursor string                         `json:"nextCursor"`
}

type NewOrganizationalUnitMembership struct {
	OrgAccountId   string `json:"orgAccountId"`
	OrgDimensionId string `json:"orgDimensionId"`
	OrgUnitId      string `json:"orgUnitId"`
}
