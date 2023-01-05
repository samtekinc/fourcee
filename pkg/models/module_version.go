package models

type ModuleVersion struct {
	ModuleVersionId  string
	ModuleGroupId    string
	Name             string
	RemoteSource     string
	TerraformVersion string
	CloudPlatform    string
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
