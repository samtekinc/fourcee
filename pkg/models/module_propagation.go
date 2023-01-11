package models

type ModulePropagation struct {
	ModulePropagationId       string
	ModuleVersionId           string
	ModuleGroupId             string
	OrgUnitId                 string
	OrgDimensionId            string
	Arguments                 []Argument
	AwsProviderConfigurations []AwsProviderConfiguration
	GcpProviderConfigurations []GcpProviderConfiguration
	Name                      string
	Description               string
}

type ModulePropagations struct {
	Items      []ModulePropagation
	NextCursor string
}

type AwsProviderConfiguration struct {
	Region string
	Alias  *string
}

type GcpProviderConfiguration struct {
	Region string
	Alias  *string
}

type AwsProviderConfigurationInput struct {
	Region string
	Alias  *string
}

type GcpProviderConfigurationInput struct {
	Region string
	Alias  *string
}

type Argument struct {
	Name  string
	Value string
}

type ArgumentInput struct {
	Name  string
	Value string
}

type NewModulePropagation struct {
	ModuleVersionId           string
	ModuleGroupId             string
	OrgUnitId                 string
	OrgDimensionId            string
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
	Name                      string
	Description               string
}

type ModulePropagationUpdate struct {
	OrgDimensionId            *string
	OrgUnitId                 *string
	Name                      *string
	Description               *string
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
}

func (m *ModulePropagation) GetInternalMetadata() []Metadata {
	return []Metadata{
		{
			Name:  "id",
			Value: m.ModulePropagationId,
		},
		{
			Name:  "name",
			Value: m.Name,
		},
		{
			Name:  "description",
			Value: m.Description,
		},
	}
}
