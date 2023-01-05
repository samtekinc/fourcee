package workflow

import (
	"context"
	"fmt"
	"strings"

	"github.com/sheacloud/tfom/internal/identifiers"
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
	var inputModuleAccountAssociationKey string
	var modulePropagationRequestId string
	var destroy bool
	switch strings.Split(input.TerraformWorkflowRequestId, "-")[0] {
	case string(identifiers.ResourceTypeTerraformExecutionWorkflowRequest):
		tfWorkflow, err := t.apiClient.GetTerraformExecutionWorkflowRequest(ctx, input.TerraformWorkflowRequestId)
		if err != nil {
			return nil, err
		}
		inputModuleAccountAssociationKey = tfWorkflow.ModuleAccountAssociationKey
		modulePropagationRequestId = tfWorkflow.ModulePropagationExecutionRequestId
		destroy = tfWorkflow.Destroy
	case string(identifiers.ResourceTypeTerraformDriftCheckWorkflowRequest):
		tfWorkflow, err := t.apiClient.GetTerraformDriftCheckWorkflowRequest(ctx, input.TerraformWorkflowRequestId)
		if err != nil {
			return nil, err
		}
		inputModuleAccountAssociationKey = tfWorkflow.ModuleAccountAssociationKey
		modulePropagationRequestId = tfWorkflow.ModulePropagationDriftCheckRequestId
		destroy = tfWorkflow.Destroy
	default:
		return nil, fmt.Errorf("invalid workflow request id: %s", input.TerraformWorkflowRequestId)
	}

	// get module account association details
	moduleAccountAssociationKey, err := models.ParseModuleAccountAssociationKey(inputModuleAccountAssociationKey)
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
	if destroy {
		additionalArguments = append(additionalArguments, "-destroy")
	}

	planRequest, err := t.apiClient.PutPlanExecutionRequest(ctx, &models.NewPlanExecutionRequest{
		TerraformVersion:             moduleVersion.TerraformVersion,
		CallbackTaskToken:            input.TaskToken,
		StateKey:                     moduleAccountAssociation.RemoteStateKey,
		ModulePropagationRequestId:   modulePropagationRequestId,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		ModuleAccountAssociationKey:  inputModuleAccountAssociationKey,
		TerraformConfigurationBase64: terraformConfig,
		AdditionalArguments:          additionalArguments,
	})
	if err != nil {
		return nil, err
	}

	switch strings.Split(input.TerraformWorkflowRequestId, "-")[0] {
	case string(identifiers.ResourceTypeTerraformExecutionWorkflowRequest):
		// update tf workflow with plan request id
		_, err = t.apiClient.UpdateTerraformExecutionWorkflowRequest(ctx, input.TerraformWorkflowRequestId, &models.TerraformExecutionWorkflowRequestUpdate{
			PlanExecutionRequestId: &planRequest.PlanExecutionRequestId,
		})
		if err != nil {
			return nil, err
		}
	case string(identifiers.ResourceTypeTerraformDriftCheckWorkflowRequest):
		// update tf workflow with plan request id
		_, err = t.apiClient.UpdateTerraformDriftCheckWorkflowRequest(ctx, input.TerraformWorkflowRequestId, &models.TerraformDriftCheckWorkflowRequestUpdate{
			PlanExecutionRequestId: &planRequest.PlanExecutionRequestId,
		})
		if err != nil {
			return nil, err
		}
	}

	return &ScheduleTerraformPlanOutput{}, nil
}
