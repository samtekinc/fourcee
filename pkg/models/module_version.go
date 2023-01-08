package models

type ModuleVersion struct {
	ModuleVersionId  string
	ModuleGroupId    string
	Name             string
	RemoteSource     string
	TerraformVersion string
	Variables        []*ModuleVariable
}

type ModuleVersions struct {
	Items      []ModuleVersion
	NextCursor string
}

type NewModuleVersion struct {
	ModuleGroupId    string
	Name             string
	RemoteSource     string
	TerraformVersion string
	CloudPlatform    string
}

type ModuleVariable struct {
	Name        string
	Type        string
	Description string
	Default     *string
}

func (m *ModuleVersion) GetInternalMetadata() []Metadata {
	return []Metadata{
		{
			Name:  "id",
			Value: m.ModuleVersionId,
		},
		{
			Name:  "name",
			Value: m.Name,
		},
		{
			Name:  "remote_source",
			Value: m.RemoteSource,
		},
		{
			Name:  "terraform_version",
			Value: m.TerraformVersion,
		},
	}
}
