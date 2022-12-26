package workflow

import (
	"context"

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

	// TODO: generate TF config file based on module version and account details

	planRequest, err := t.apiClient.PutPlanExecutionRequest(ctx, &models.NewPlanExecutionRequest{
		TerraformVersion:             moduleVersion.TerraformVersion,
		CallbackTaskToken:            input.TaskToken,
		StateKey:                     input.Input.ModuleAccountAssociation.RemoteStateKey,
		GroupingKey:                  input.Input.ModulePropagationExecutionRequestId,
		TerraformConfigurationBase64: "cHJvdmlkZXIgImF3cyIgewogIHJlZ2lvbiA9ICJ1cy1lYXN0LTEiCn0KCmRhdGEgImF3c19yZWdpb24iICJjdXJyZW50IiB7fQoKb3V0cHV0ICJyZWdpb24iIHsKICB2YWx1ZSA9IGRhdGEuYXdzX3JlZ2lvbi5jdXJyZW50Lm5hbWUKfQo=",
	})
	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformPlanOutput{
		PlanExecutionRequestId: planRequest.PlanExecutionRequestId,
	}, nil
}
