package models

import "time"

type PlanExecutionRequest struct {
	PlanExecutionRequestId       string
	TerraformVersion             string
	CallbackTaskToken            string
	StateKey                     string
	ModuleAccountAssociationKey  string
	TerraformConfigurationBase64 string
	AdditionalArguments          []string
	ModulePropagationRequestId   string // could be a mpexec or an mpsync request
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
	TerraformVersion             string
	CallbackTaskToken            string
	StateKey                     string
	ModulePropagationRequestId   string
	TerraformWorkflowRequestId   string
	ModuleAccountAssociationKey  string
	TerraformConfigurationBase64 string
	AdditionalArguments          []string
}

type PlanExecutionRequestUpdate struct {
	InitOutputKey *string
	PlanOutputKey *string
	Status        *RequestStatus
}
