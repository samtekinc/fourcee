package models

type ModulePropagationAssignment struct {
	ModulePropagationId string
	OrgAccountID        string
	ModuleAssignmentId  string
}

type ModulePropagationAssignments struct {
	Items      []ModulePropagationAssignment
	NextCursor string
}

type NewModulePropagationAssignment struct {
	ModulePropagationId string
	OrgAccountID        string
}
