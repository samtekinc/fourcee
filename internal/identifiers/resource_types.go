package identifiers

type ResourceType string

const (
	ResourceTypeOrgDimension                       ResourceType = "od"
	ResourceTypeOrgUnit                            ResourceType = "ou"
	ResourceTypeOrgAccount                         ResourceType = "oa"
	ResourceTypeModuleGroup                        ResourceType = "mg"
	ResourceTypeModuleVersion                      ResourceType = "mv"
	ResourceTypeModulePropagation                  ResourceType = "mp"
	ResourceTypeModuleAssignment                   ResourceType = "ma"
	ResourceTypeTerraformExecutionRequest          ResourceType = "tfexec"
	ResourceTypeTerraformDriftCheckRequest         ResourceType = "tfdrift"
	ResourceTypePlanExecutionRequest               ResourceType = "plan"
	ResourceTypeApplyExecutionRequest              ResourceType = "apply"
	ResourceTypeModulePropagationExecutionRequest  ResourceType = "mpexec"
	ResourceTypeModulePropagationDriftCheckRequest ResourceType = "mpdrift"
)
