package models

type OrganizationalDimension struct {
	OrgDimensionId string `json:"orgDimensionId"`
	Name           string `json:"name"`
	RootOrgUnitId  string `json:"rootOrgUnitId"`
}

type OrganizationalDimensions struct {
	Items      []OrganizationalDimension `json:"items"`
	NextCursor string                    `json:"nextCursor"`
}

type NewOrganizationalDimension struct {
	Name string `json:"name"`
}
