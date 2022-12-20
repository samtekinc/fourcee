package models

type OrganizationalUnitMembership struct {
	OrgAccountId string `json:"orgAccountId"`
	DimensionId  string `json:"dimensionId"`
	OrgUnitId    string `json:"orgUnitId"`
}

type OrganizationalUnitMemberships struct {
	Items      []OrganizationalUnitMembership `json:"items"`
	NextCursor string                         `json:"nextCursor"`
}

type NewOrganizationalUnitMembership struct {
	OrgAccountId string `json:"orgAccountId"`
	DimensionId  string `json:"dimensionId"`
	OrgUnitId    string `json:"orgUnitId"`
}
