package models

type ModulePropagation struct {
	ModulePropagationId       string                     `json:"modulePropagationId"`
	ModuleVersionId           string                     `json:"moduleVersionId"`
	ModuleGroupId             string                     `json:"moduleGroupId"`
	OrgUnitId                 string                     `json:"orgUnitId"`
	OrgDimensionId            string                     `json:"orgDimensionId"`
	Arguments                 []Argument                 `json:"arguments"`
	AwsProviderConfigurations []AwsProviderConfiguration `json:"awsProviderConfigurations"`
	Name                      string                     `json:"name"`
	Description               string                     `json:"description"`
}

type ModulePropagations struct {
	Items      []ModulePropagation `json:"items"`
	NextCursor string              `json:"nextCursor"`
}

type AwsProviderConfiguration struct {
	Region string `json:"region"`
	Alias  string `json:"alias"`
}

type AwsProviderConfigurationInput struct {
	Region string `json:"region"`
	Alias  string `json:"alias"`
}

type Argument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ArgumentInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NewModulePropagation struct {
	ModuleVersionId           string                          `json:"moduleVersionId"`
	ModuleGroupId             string                          `json:"moduleGroupId"`
	OrgUnitId                 string                          `json:"orgUnitId"`
	OrgDimensionId            string                          `json:"orgDimensionId"`
	Arguments                 []ArgumentInput                 `json:"arguments"`
	AwsProviderConfigurations []AwsProviderConfigurationInput `json:"awsProviderConfigurations"`
	Name                      string                          `json:"name"`
	Description               string                          `json:"description"`
}

type ModulePropagationUpdate struct {
	OrgDimensionId            *string                         `json:"orgDimensionId"`
	OrgUnitId                 *string                         `json:"orgUnitId"`
	Name                      *string                         `json:"name"`
	Description               *string                         `json:"description"`
	Arguments                 []ArgumentInput                 `json:"arguments"`
	AwsProviderConfigurations []AwsProviderConfigurationInput `json:"awsProviderConfigurations"`
}
