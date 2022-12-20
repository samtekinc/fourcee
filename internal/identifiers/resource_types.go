package identifiers

type ResourceType string

const (
	ResourceTypeOrganizationalDimension ResourceType = "od"
	ResourceTypeOrganizationalUnit      ResourceType = "ou"
	ResourceTypeOrganizationalAccount   ResourceType = "oa"
	ResourceTypeModuleGroup             ResourceType = "mg"
	ResourceTypeModuleVersion           ResourceType = "mv"
)
