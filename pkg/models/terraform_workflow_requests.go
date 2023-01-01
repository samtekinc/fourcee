package models

import "time"

type TerraformWorkflowRequestStatus string

const (
	TerraformWorkflowRequestStatusPending   TerraformWorkflowRequestStatus = "PENDING"
	TerraformWorkflowRequestStatusRunning   TerraformWorkflowRequestStatus = "RUNNING"
	TerraformWorkflowRequestStatusSucceeded TerraformWorkflowRequestStatus = "SUCCEEDED"
	TerraformWorkflowRequestStatusFailed    TerraformWorkflowRequestStatus = "FAILED"
)

type TerraformWorkflowRequest struct {
	TerraformWorkflowRequestId          string                         `json:"terraformWorkflowRequestId"`
	ModulePropagationExecutionRequestId string                         `json:"modulePropagationExecutionRequestId"`
	ModuleAccountAssociationKey         string                         `json:"moduleAccountAssociationKey"`
	PlanExecutionRequestId              *string                        `json:"planExecutionRequestId"`
	ApplyExecutionRequestId             *string                        `json:"applyExecutionRequestId"`
	RequestTime                         time.Time                      `json:"requestTime"`
	Status                              TerraformWorkflowRequestStatus `json:"status"`
	Destroy                             bool                           `json:"destroy"`
}

type TerraformWorkflowRequests struct {
	Items      []TerraformWorkflowRequest `json:"items"`
	NextCursor string                     `json:"nextCursor"`
}

type NewTerraformWorkflowRequest struct {
	ModulePropagationExecutionRequestId string `json:"modulePropagationExecutionRequestId"`
	ModuleAccountAssociationKey         string `json:"moduleAccountAssociationKey"`
	Destroy                             bool   `json:"destroy"`
}

type TerraformWorkflowRequestUpdate struct {
	Status                  *TerraformWorkflowRequestStatus `json:"status"`
	PlanExecutionRequestId  *string                         `json:"planExecutionRequestId"`
	ApplyExecutionRequestId *string                         `json:"applyExecutionRequestId"`
}
