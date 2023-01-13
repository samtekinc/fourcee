package identifiers

type ResourceType string

const (
	ResourceTypeOrganizationalDimension            ResourceType = "od"
	ResourceTypeOrganizationalUnit                 ResourceType = "ou"
	ResourceTypeOrganizationalAccount              ResourceType = "oa"
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
