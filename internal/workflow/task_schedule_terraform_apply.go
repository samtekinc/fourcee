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
	tfWorkflow, err := t.apiClient.GetTerraformExecutionWorkflowRequest(ctx, input.TerraformWorkflowRequestId)
	if err != nil {
		return nil, err
	}

	// get module account association
	moduleAccountAssociationKey, err := models.ParseModuleAccountAssociationKey(tfWorkflow.ModuleAccountAssociationKey)
	if err != nil {
		return nil, err
	}
	moduleAccountAssociation, err := t.apiClient.GetModuleAccountAssociation(ctx, moduleAccountAssociationKey.ModulePropagationId, moduleAccountAssociationKey.OrgAccountId)
	if err != nil {
		return nil, err
	}

	// get module propagation details
	modulePropagation, err := t.apiClient.GetModulePropagation(ctx, moduleAccountAssociation.ModulePropagationId)
	if err != nil {
		return nil, err
	}

	// get module version details
	moduleVersion, err := t.apiClient.GetModuleVersion(ctx, modulePropagation.ModuleGroupId, modulePropagation.ModuleVersionId)
	if err != nil {
		return nil, err
	}

	// get plan request details
	planRequest, err := t.apiClient.GetPlanExecutionRequest(ctx, *tfWorkflow.PlanExecutionRequestId)
	if err != nil {
		return nil, err
	}

	planBase64 := base64.StdEncoding.EncodeToString(planRequest.PlanOutput.PlanFile)

	additionalArguments := []string{}
	if tfWorkflow.Destroy {
		additionalArguments = append(additionalArguments, "-destroy")
	}

	applyRequest, err := t.apiClient.PutApplyExecutionRequest(ctx, &models.NewApplyExecutionRequest{
		TerraformVersion:             moduleVersion.TerraformVersion,
		CallbackTaskToken:            input.TaskToken,
		StateKey:                     moduleAccountAssociation.RemoteStateKey,
		ModulePropagationRequestId:   tfWorkflow.ModulePropagationExecutionRequestId,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		ModuleAccountAssociationKey:  tfWorkflow.ModuleAccountAssociationKey,
		TerraformConfigurationBase64: planRequest.TerraformConfigurationBase64,
		TerraformPlanBase64:          planBase64,
		AdditionalArguments:          additionalArguments,
	})
	if err != nil {
		return nil, err
	}

	// update tf workflow with apply request id
	_, err = t.apiClient.UpdateTerraformExecutionWorkflowRequest(ctx, input.TerraformWorkflowRequestId, &models.TerraformExecutionWorkflowRequestUpdate{
		ApplyExecutionRequestId: &applyRequest.ApplyExecutionRequestId,
	})
	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformApplyOutput{
		ApplyExecutionRequestId: applyRequest.ApplyExecutionRequestId,
	}, nil
}
