package models

type ModuleGroup struct {
	ModuleGroupId string `json:"moduleGroupId"`
	Name          string `json:"name"`
}

type ModuleGroups struct {
	Items      []ModuleGroup `json:"items"`
	NextCursor string        `json:"nextCursor"`
}

type NewModuleGroup struct {
	Name string `json:"name"`
}
