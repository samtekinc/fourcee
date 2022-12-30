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
	Input struct {
		ModuleAccountAssociation            models.ModuleAccountAssociation
		ModulePropagationExecutionRequestId string
		ModulePropagationId                 string
	}
	TaskToken string
}

type ScheduleTerraformPlanOutput struct {
	PlanExecutionRequestId string
}

func (t *TaskHandler) ScheduleTerraformPlan(ctx context.Context, input ScheduleTerraformPlanInput) (*ScheduleTerraformPlanOutput, error) {
	// get module propagation details
	modulePropagation, err := t.apiClient.GetModulePropagation(ctx, input.Input.ModulePropagationId)
	if err != nil {
		return nil, err
	}

	// get module version details
	moduleVersion, err := t.apiClient.GetModuleVersion(ctx, modulePropagation.ModuleGroupId, modulePropagation.ModuleVersionId)
	if err != nil {
		return nil, err
	}

	// get org account details
	orgAccount, err := t.apiClient.GetOrganizationalAccount(ctx, input.Input.ModuleAccountAssociation.OrgAccountId)
	if err != nil {
		return nil, err
	}

	// TODO: generate TF config file based on module version and account details
	terraformConfig, err := terraform.GetTerraformConfigurationBase64(&terraform.TerraformConfigurationInput{
		ModuleAccountAssociation: &input.Input.ModuleAccountAssociation,
		ModulePropagation:        modulePropagation,
		ModuleVersion:            moduleVersion,
		OrgAccount:               orgAccount,
	})
	if err != nil {
		return nil, err
	}

	planRequest, err := t.apiClient.PutPlanExecutionRequest(ctx, &models.NewPlanExecutionRequest{
		TerraformVersion:                    moduleVersion.TerraformVersion,
		CallbackTaskToken:                   input.TaskToken,
		StateKey:                            input.Input.ModuleAccountAssociation.RemoteStateKey,
		ModulePropagationId:                 input.Input.ModulePropagationId,
		ModulePropagationExecutionRequestId: input.Input.ModulePropagationExecutionRequestId,
		ModuleAccountAssociationKey:         input.Input.ModuleAccountAssociation.Key(),
		TerraformConfigurationBase64:        terraformConfig,
	})
	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformPlanOutput{
		PlanExecutionRequestId: planRequest.PlanExecutionRequestId,
	}, nil
}
