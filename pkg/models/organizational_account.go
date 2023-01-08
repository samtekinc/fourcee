package models

type OrganizationalAccount struct {
	OrgAccountId    string
	Name            string
	CloudPlatform   CloudPlatform
	CloudIdentifier string
	AssumeRoleName  string
	Metadata        []Metadata
}

type OrganizationalAccounts struct {
	Items      []OrganizationalAccount
	NextCursor string
}

type NewOrganizationalAccount struct {
	Name            string
	CloudPlatform   CloudPlatform
	CloudIdentifier string
	AssumeRoleName  string
	Metadata        []MetadataInput
}

type OrganizationalAccountUpdate struct {
	Metadata []MetadataInput
}

func (a *OrganizationalAccount) GetInternalMetadata() []Metadata {
	return []Metadata{
		{
			Name:  "id",
			Value: a.OrgAccountId,
		},
		{
			Name:  "name",
			Value: a.Name,
		},
	}
}
