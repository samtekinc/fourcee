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
	InitOutputKey                string
	InitOutput                   *TerraformInitOutput `dynamodbav:"-"` // fetched from S3 on request
	PlanOutputKey                string
	PlanOutput                   *TerraformPlanOutput `dynamodbav:"-"` // fetched from S3 on request
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
	Status        *RequestStatus
}
