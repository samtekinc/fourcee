package models

type ModuleVersion struct {
	ModuleVersionId string `json:"moduleVersionId"`
	ModuleGroupId   string `json:"moduleGroupId"`
	Name            string `json:"name"`
	RemoteSource    string `json:"remoteSource"`
}

type ModuleVersions struct {
	Items      []ModuleVersion `json:"items"`
	NextCursor string          `json:"nextCursor"`
}

type NewModuleVersion struct {
	ModuleGroupId string `json:"moduleGroupId"`
	Name          string `json:"name"`
	RemoteSource  string `json:"remoteSource"`
}
