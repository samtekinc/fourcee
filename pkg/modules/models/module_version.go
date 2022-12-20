package models

type ModuleVersion struct {
	ModuleVersionId  string            `json:"moduleVersionId"`
	ModuleGroupId    string            `json:"moduleGroupId"`
	Name             string            `json:"name"`
	RemoteSource     string            `json:"remoteSource"`
	TerraformVersion string            `json:"terraformVersion"`
	Variables        []*ModuleVariable `json:"variables"`
}

type ModuleVersions struct {
	Items      []ModuleVersion `json:"items"`
	NextCursor string          `json:"nextCursor"`
}

type NewModuleVersion struct {
	ModuleGroupId    string `json:"moduleGroupId"`
	Name             string `json:"name"`
	RemoteSource     string `json:"remoteSource"`
	TerraformVersion string `json:"terraformVersion"`
}

type ModuleVariable struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Default     *string `json:"default"`
}
