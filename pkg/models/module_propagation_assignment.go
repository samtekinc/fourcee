package models

type ModulePropagationAssignment struct {
	ModulePropagationId string
	OrgAccountId        string
	ModuleAssignmentId  string
}

type ModulePropagationAssignments struct {
	Items      []ModulePropagationAssignment
	NextCursor string
}

type NewModulePropagationAssignment struct {
	ModulePropagationId string
	OrgAccountId        string
}
