package models

import "time"

type ApplyExecutionRequest struct {
	ApplyExecutionRequestId      string
	ModuleAssignmentId           string
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformConfigurationBase64 string
	TerraformPlanBase64          string
	AdditionalArguments          []string
	TerraformWorkflowRequestId   string // could be a tfexec or a tfsync request
	Status                       RequestStatus
	InitOutputKey                *string
	ApplyOutputKey               *string
	RequestTime                  time.Time
}

type ApplyExecutionRequests struct {
	Items      []ApplyExecutionRequest
	NextCursor string
}

type NewApplyExecutionRequest struct {
	ModuleAssignmentId           string
	TerraformVersion             string
	CallbackTaskToken            string
	TerraformWorkflowRequestId   string // could be a tfexec or a tfsync request
	TerraformConfigurationBase64 string
	TerraformPlanBase64          string
	AdditionalArguments          []string
}

type ApplyExecutionRequestUpdate struct {
	InitOutputKey  *string
	ApplyOutputKey *string
	Status         *RequestStatus
}
