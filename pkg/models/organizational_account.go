package models

type OrganizationalAccount struct {
	OrgAccountId    string `json:"orgAccountId"`
	Name            string `json:"name"`
	CloudPlatform   string `json:"cloudPlatform"`
	CloudIdentifier string `json:"cloudIdentifier"`
	AssumeRoleName  string `json:"assumeRoleName"`
}

type OrganizationalAccounts struct {
	Items      []OrganizationalAccount `json:"items"`
	NextCursor string                  `json:"nextCursor"`
}

type NewOrganizationalAccount struct {
	Name            string `json:"name"`
	CloudPlatform   string `json:"cloudPlatform"`
	CloudIdentifier string `json:"cloudIdentifier"`
	AssumeRoleName  string `json:"assumeRoleName"`
}
