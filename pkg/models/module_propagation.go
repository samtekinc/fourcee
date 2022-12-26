package models

type ModulePropagation struct {
	ModulePropagationId string      `json:"modulePropagationId"`
	ModuleVersionId     string      `json:"moduleVersionId"`
	ModuleGroupId       string      `json:"moduleGroupId"`
	OrgUnitId           string      `json:"orgUnitId"`
	OrgDimensionId      string      `json:"orgDimensionId"`
	Arguments           []*Argument `json:"arguments"`
	Providers           []*Provider `json:"providers"`
	Name                string      `json:"name"`
	Description         string      `json:"description"`
}

type ModulePropagations struct {
	Items      []ModulePropagation `json:"items"`
	NextCursor string              `json:"nextCursor"`
}

type Provider struct {
	Name      string      `json:"name" dynamodbav:"name"`
	Arguments []*Argument `json:"arguments" dynamodbav:"arguments"`
}

type Argument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NewModulePropagation struct {
	ModuleVersionId string      `json:"moduleVersionId"`
	ModuleGroupId   string      `json:"moduleGroupId"`
	OrgUnitId       string      `json:"orgUnitId"`
	OrgDimensionId  string      `json:"orgDimensionId"`
	Arguments       []*Argument `json:"arguments"`
	Providers       []*Provider `json:"providers"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
}
