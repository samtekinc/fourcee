package models

import "time"

type PlanExecutionRequest struct {
	PlanExecutionRequestId       string
	ModuleAssignmentId           string
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformConfigurationBase64 string
	AdditionalArguments          []string
	TerraformWorkflowRequestId   string // could be a tfexec or a tfsync request
	Status                       RequestStatus
	InitOutputKey                *string
	PlanOutputKey                *string
	PlanFileKey                  *string
	PlanJSONKey                  *string
	RequestTime                  time.Time
}

type PlanExecutionRequests struct {
	Items      []PlanExecutionRequest
	NextCursor string
}

type NewPlanExecutionRequest struct {
	ModuleAssignmentId           string
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformWorkflowRequestId   string
	TerraformConfigurationBase64 string
	AdditionalArguments          []string
}

type PlanExecutionRequestUpdate struct {
	InitOutputKey *string
	PlanOutputKey *string
	PlanFileKey   *string
	PlanJSONKey   *string
	Status        *RequestStatus
}

type TerraformPlanOutput struct {
	PlanFile []byte
	PlanJSON []byte
}
