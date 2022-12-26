package models

type OrganizationalUnit struct {
	OrgUnitId       string `json:"orgUnitId"`
	Name            string `json:"name"`
	OrgDimensionId  string `json:"orgDimensionId"`
	Hierarchy       string `json:"hierarchy"`
	ParentOrgUnitId string `json:"parentOrgUnitId" dynamodbav:",omitempty"`
}

type OrganizationalUnits struct {
	Items      []OrganizationalUnit `json:"items"`
	NextCursor string               `json:"nextCursor"`
}

type NewOrganizationalUnit struct {
	Name            string `json:"name"`
	OrgDimensionId  string `json:"orgDimensionId"`
	ParentOrgUnitId string `json:"parentOrgUnitId"`
}

type OrganizationalUnitUpdate struct {
	Name            *string `json:"name"`
	ParentOrgUnitId *string `json:"parentOrgUnitId"`
}
