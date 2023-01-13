package workflow

import (
	"context"
	"encoding/base64"

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
	planRequest, err := t.apiClient.GetPlanExecutionRequest(ctx, *tfWorkflow.PlanExecutionRequestId, true)
	if err != nil {
		return nil, err
	}

	planBase64 := base64.StdEncoding.EncodeToString(planRequest.PlanOutput.PlanFile)

	additionalArguments := []string{}
	if tfWorkflow.Destroy {
		additionalArguments = append(additionalArguments, "-destroy")
	}

	applyRequest, err := t.apiClient.PutApplyExecutionRequest(ctx, &models.NewApplyExecutionRequest{
		ModuleAssignmentId:           moduleAssignment.ModuleAssignmentId,
		TerraformVersion:             moduleVersion.TerraformVersion,
		CallbackTaskToken:            input.TaskToken,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		TerraformConfigurationBase64: planRequest.TerraformConfigurationBase64,
		TerraformPlanBase64:          planBase64,
		AdditionalArguments:          additionalArguments,
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
