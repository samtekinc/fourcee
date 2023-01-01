package models

import "time"

type PlanExecutionStatus string

const (
	PlanExecutionStatusPending   PlanExecutionStatus = "PENDING"
	PlanExecutionStatusRunning   PlanExecutionStatus = "RUNNING"
	PlanExecutionStatusSucceeded PlanExecutionStatus = "SUCCEEDED"
	PlanExecutionStatusFailed    PlanExecutionStatus = "FAILED"
)

type PlanExecutionRequest struct {
	PlanExecutionRequestId              string               `json:"planExecutionRequestId"`
	TerraformVersion                    string               `json:"terraformVersion"`
	CallbackTaskToken                   string               `json:"callbackTaskToken"`
	StateKey                            string               `json:"stateKey"`
	ModulePropagationExecutionRequestId string               `json:"groupingKey" dynamodbav:",omitempty"`
	ModuleAccountAssociationKey         string               `json:"moduleAccountAssociationKey" dynamodbav:",omitempty"`
	TerraformConfigurationBase64        string               `json:"terraformConfigurationBase64"`
	AdditionalArguments                 []string             `json:"additionalArguments"`
	WorkflowExecutionId                 string               `json:"-"`
	Status                              PlanExecutionStatus  `json:"status"`
	InitOutputKey                       string               `json:"-"`                                   // for internal use only
	InitOutput                          *TerraformInitOutput `json:"initOutput,omitempty" dynamodbav:"-"` // fetched from S3 on request
	PlanOutputKey                       string               `json:"-"`                                   // for internal use only
	PlanOutput                          *TerraformPlanOutput `json:"planOutput,omitempty" dynamodbav:"-"` // fetched from S3 on request
	RequestTime                         time.Time            `json:"requestTime"`
}

type PlanExecutionRequests struct {
	Items      []PlanExecutionRequest `json:"items"`
	NextCursor string                 `json:"nextCursor"`
}

type NewPlanExecutionRequest struct {
	TerraformVersion                    string   `json:"terraformVersion"`
	CallbackTaskToken                   string   `json:"callbackTaskToken"`
	StateKey                            string   `json:"stateKey"`
	ModulePropagationExecutionRequestId string   `json:"groupingKey" dynamodbav:",omitempty"`
	ModuleAccountAssociationKey         string   `json:"moduleAccountAssociationKey" dynamodbav:",omitempty"`
	TerraformConfigurationBase64        string   `json:"terraformConfigurationBase64"`
	AdditionalArguments                 []string `json:"additionalArguments"`
}

type PlanExecutionRequestUpdate struct {
	InitOutputKey *string
	PlanOutputKey *string
	Status        *PlanExecutionStatus
}
