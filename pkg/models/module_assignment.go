package models

type ModuleAssignment struct {
	ModuleAssignmentId        string
	ModuleVersionId           string
	ModuleGroupId             string
	OrgAccountId              string
	Name                      string
	Description               string
	RemoteStateRegion         string
	RemoteStateBucket         string
	RemoteStateKey            string
	Arguments                 []Argument
	AwsProviderConfigurations []AwsProviderConfiguration
	GcpProviderConfigurations []GcpProviderConfiguration
	ModulePropagationId       *string `dynamodbav:",omitempty"`
	Status                    ModuleAssignmentStatus
}

type ModuleAssignments struct {
	Items      []ModuleAssignment
	NextCursor string
}

type ModuleAssignmentStatus string

const (
	ModuleAssignmentStatusActive   ModuleAssignmentStatus = "ACTIVE"
	ModuleAssignmentStatusInactive ModuleAssignmentStatus = "INACTIVE"
)

type NewModuleAssignment struct {
	ModuleVersionId           string
	ModuleGroupId             string
	OrgAccountId              string
	Name                      string
	Description               string
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
	ModulePropagationId       *string
}

type ModuleAssignmentUpdate struct {
	Name                      *string
	Description               *string
	ModuleVersionId           *string
	Arguments                 []ArgumentInput
	AwsProviderConfigurations []AwsProviderConfigurationInput
	GcpProviderConfigurations []GcpProviderConfigurationInput
	Status                    *ModuleAssignmentStatus
}

type ModuleAssignmentFilters struct {
	IsPropagated *bool
}

type ModuleAssignmentFiltersInput struct {
	IsPropagated *bool
}
