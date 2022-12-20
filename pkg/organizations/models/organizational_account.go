package models

type OrganizationalAccount struct {
	OrgAccountId string `json:"orgAccountId"`
	Name         string `json:"name"`
}

type OrganizationalAccounts struct {
	Items      []OrganizationalAccount `json:"items"`
	NextCursor string                  `json:"nextCursor"`
}

type NewOrganizationalAccount struct {
	Name string `json:"name"`
}
