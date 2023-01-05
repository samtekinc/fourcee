package models

type OrganizationalAccount struct {
	OrgAccountId    string
	Name            string
	CloudPlatform   string
	CloudIdentifier string
	AssumeRoleName  string
}

type OrganizationalAccounts struct {
	Items      []OrganizationalAccount
	NextCursor string
}

type NewOrganizationalAccount struct {
	Name            string
	CloudPlatform   string
	CloudIdentifier string
	AssumeRoleName  string
}
