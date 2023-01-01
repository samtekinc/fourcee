package workflow

import (
	"context"

	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformPlan Task = "ScheduleTerraformPlan"
)

type ScheduleTerraformPlanInput struct {
	TerraformWorkflowRequestId string
	TaskToken                  string
}

type ScheduleTerraformPlanOutput struct{}

func (t *TaskHandler) ScheduleTerraformPlan(ctx context.Context, input ScheduleTerraformPlanInput) (*ScheduleTerraformPlanOutput, error) {
	// get workflow details
	tfWorkflow, err := t.apiClient.GetTerraformWorkflowRequest(ctx, input.TerraformWorkflowRequestId)
	if err != nil {
		return nil, err
	}

	// get module account association details
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

	// get org account details
	orgAccount, err := t.apiClient.GetOrganizationalAccount(ctx, moduleAccountAssociationKey.OrgAccountId)
	if err != nil {
		return nil, err
	}

	terraformConfig, err := terraform.GetTerraformConfigurationBase64(&terraform.TerraformConfigurationInput{
		ModuleAccountAssociation: moduleAccountAssociation,
		ModulePropagation:        modulePropagation,
		ModuleVersion:            moduleVersion,
		OrgAccount:               orgAccount,
	})
	if err != nil {
		return nil, err
	}

	additionalArguments := []string{}
	if tfWorkflow.Destroy {
		additionalArguments = append(additionalArguments, "-destroy")
	}

	planRequest, err := t.apiClient.PutPlanExecutionRequest(ctx, &models.NewPlanExecutionRequest{
		TerraformVersion:                    moduleVersion.TerraformVersion,
		CallbackTaskToken:                   input.TaskToken,
		StateKey:                            moduleAccountAssociation.RemoteStateKey,
		ModulePropagationExecutionRequestId: tfWorkflow.ModulePropagationExecutionRequestId,
		ModuleAccountAssociationKey:         tfWorkflow.ModuleAccountAssociationKey,
		TerraformConfigurationBase64:        terraformConfig,
		AdditionalArguments:                 additionalArguments,
	})
	if err != nil {
		return nil, err
	}

	// update tf workflow with apply request id
	_, err = t.apiClient.UpdateTerraformWorkflowRequest(ctx, tfWorkflow.TerraformWorkflowRequestId, &models.TerraformWorkflowRequestUpdate{
		PlanExecutionRequestId: &planRequest.PlanExecutionRequestId,
	})
	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformPlanOutput{}, nil
}
