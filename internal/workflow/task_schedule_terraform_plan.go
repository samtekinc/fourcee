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
	var inputModuleAssignmentId string
	var destroy bool
	switch strings.Split(input.TerraformWorkflowRequestId, "-")[0] {
	case string(identifiers.ResourceTypeTerraformExecutionWorkflowRequest):
		tfWorkflow, err := t.apiClient.GetTerraformExecutionWorkflowRequest(ctx, input.TerraformWorkflowRequestId)
		if err != nil {
			return nil, err
		}
		inputModuleAssignmentId = tfWorkflow.ModuleAssignmentId
		destroy = tfWorkflow.Destroy
	case string(identifiers.ResourceTypeTerraformDriftCheckWorkflowRequest):
		tfWorkflow, err := t.apiClient.GetTerraformDriftCheckWorkflowRequest(ctx, input.TerraformWorkflowRequestId)
		if err != nil {
			return nil, err
		}
		inputModuleAssignmentId = tfWorkflow.ModuleAssignmentId
		destroy = tfWorkflow.Destroy
	default:
		return nil, fmt.Errorf("invalid workflow request id: %s", input.TerraformWorkflowRequestId)
	}

	// get module assignment details
	moduleAssignment, err := t.apiClient.GetModuleAssignment(ctx, inputModuleAssignmentId)
	if err != nil {
		return nil, err
	}

	// get module propagation details
	var modulePropagation *models.ModulePropagation
	if moduleAssignment.ModulePropagationId != nil {
		modulePropagation, err = t.apiClient.GetModulePropagation(ctx, *moduleAssignment.ModulePropagationId)
		if err != nil {
			return nil, err
		}
	}

	// get module version details
	moduleVersion, err := t.apiClient.GetModuleVersion(ctx, moduleAssignment.ModuleGroupId, moduleAssignment.ModuleVersionId)
	if err != nil {
		return nil, err
	}

	// get org account details
	orgAccount, err := t.apiClient.GetOrganizationalAccount(ctx, moduleAssignment.OrgAccountId)
	if err != nil {
		return nil, err
	}

	terraformConfig, err := terraform.GetTerraformConfigurationBase64(&terraform.TerraformConfigurationInput{
		ModuleAssignment:  moduleAssignment,
		ModulePropagation: modulePropagation,
		ModuleVersion:     moduleVersion,
		OrgAccount:        orgAccount,
	})
	if err != nil {
		return nil, err
	}

	additionalArguments := []string{}
	if destroy {
		additionalArguments = append(additionalArguments, "-destroy")
	}

	planRequest, err := t.apiClient.PutPlanExecutionRequest(ctx, &models.NewPlanExecutionRequest{
		ModuleAssignmentId:           moduleAssignment.ModuleAssignmentId,
		TerraformVersion:             moduleVersion.TerraformVersion,
		CallbackTaskToken:            input.TaskToken,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
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
