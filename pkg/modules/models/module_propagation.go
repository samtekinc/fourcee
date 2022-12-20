package models

type ModulePropagation struct {
	ModulePropagationId string `json:"modulePropagationId"`
	ModuleVersionId     string `json:"moduleVersionId"`
	OrgUnitId           string `json:"orgUnitId"`
	OrgDimensionId      string `json:"orgDimensionId"`
}
