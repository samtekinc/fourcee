package workflow

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformApply Task = "ScheduleTerraformApply"
)

type ScheduleTerraformApplyInput struct {
	TerraformWorkflowRequestId string
	TaskToken                  string
}

type ScheduleTerraformApplyOutput struct {
	ApplyExecutionRequestId string
}

func (t *TaskHandler) ScheduleTerraformApply(ctx context.Context, input ScheduleTerraformApplyInput) (*ScheduleTerraformApplyOutput, error) {
	// get workflow details
	tfWorkflow, err := t.apiClient.GetTerraformExecutionRequest(ctx, input.TerraformWorkflowRequestId)
	if err != nil {
		return nil, err
	}

	// get module assignment
	moduleAssignment, err := t.apiClient.GetModuleAssignment(ctx, tfWorkflow.ModuleAssignmentId)
	if err != nil {
		return nil, err
	}

	// get module version details
	moduleVersion, err := t.apiClient.GetModuleVersion(ctx, moduleAssignment.ModuleGroupId, moduleAssignment.ModuleVersionId)
	if err != nil {
		return nil, err
	}

	// get plan request details
	planRequest, err := t.apiClient.GetPlanExecutionRequest(ctx, *tfWorkflow.PlanExecutionRequestId)
	if err != nil {
		return nil, err
	}

	if planRequest.PlanFileKey == nil {
		return nil, fmt.Errorf("plan file key is nil")
	}

	planFile, err := t.apiClient.DownloadResultObject(ctx, *planRequest.PlanFileKey)
	if err != nil {
		return nil, err
	}

	planBase64 := base64.StdEncoding.EncodeToString(planFile)

	additionalArguments := []string{} // no need to add the destroy flag here, it's already in the plan

	applyRequest, err := t.apiClient.PutApplyExecutionRequest(ctx, &models.NewApplyExecutionRequest{
		ModuleAssignmentId:         moduleAssignment.ModuleAssignmentId,
		TerraformVersion:           moduleVersion.TerraformVersion,
		CallbackTaskToken:          input.TaskToken,
		TerraformWorkflowRequestId: input.TerraformWorkflowRequestId,
		TerraformConfiguration:     planRequest.TerraformConfiguration,
		TerraformPlan:              planBase64,
		AdditionalArguments:        additionalArguments,
	})
	if err != nil {
		return nil, err
	}

	// update tf workflow with apply request id
	_, err = t.apiClient.UpdateTerraformExecutionRequest(ctx, input.TerraformWorkflowRequestId, &models.TerraformExecutionRequestUpdate{
		ApplyExecutionRequestId: &applyRequest.ApplyExecutionRequestId,
	})
	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformApplyOutput{
		ApplyExecutionRequestId: applyRequest.ApplyExecutionRequestId,
	}, nil
}
