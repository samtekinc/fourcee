package models

import "time"

type ApplyExecutionRequest struct {
	ApplyExecutionRequestId      string
	TerraformVersion             string
	CallbackTaskToken            string
	StateKey                     string
	ModuleAccountAssociationKey  string
	TerraformConfigurationBase64 string
	TerraformPlanBase64          string
	AdditionalArguments          []string
	ModulePropagationRequestId   string // could be a mpexec or an mpsync request
	TerraformWorkflowRequestId   string // could be a tfexec or a tfsync request
	Status                       RequestStatus
	InitOutputKey                string
	InitOutput                   *TerraformInitOutput `dynamodbav:"-"` // fetched from S3 on request
	ApplyOutputKey               string
	ApplyOutput                  *TerraformApplyOutput `dynamodbav:"-"` // fetched from S3 on request
	RequestTime                  time.Time
}

type ApplyExecutionRequests struct {
	Items      []ApplyExecutionRequest
	NextCursor string
}

type NewApplyExecutionRequest struct {
	TerraformVersion             string
	CallbackTaskToken            string
	StateKey                     string
	ModulePropagationRequestId   string // could be a mpexec or an mpsync request
	TerraformWorkflowRequestId   string // could be a tfexec or a tfsync request
	ModuleAccountAssociationKey  string
	TerraformConfigurationBase64 string
	TerraformPlanBase64          string
	AdditionalArguments          []string
}

type ApplyExecutionRequestUpdate struct {
	InitOutputKey  *string
	ApplyOutputKey *string
	Status         *RequestStatus
}
