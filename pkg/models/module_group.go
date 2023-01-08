package models

type ModuleGroup struct {
	ModuleGroupId string
	Name          string
	CloudPlatform CloudPlatform
}

type ModuleGroups struct {
	Items      []ModuleGroup
	NextCursor string
}

type NewModuleGroup struct {
	Name          string
	CloudPlatform CloudPlatform
}
