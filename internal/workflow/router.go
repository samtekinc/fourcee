package workflow

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/sheacloud/tfom/internal/api"

	"go.uber.org/zap"
)

type Workflow string

type Task string

type CommonTaskInput struct {
	Workflow Workflow
	Task     Task
	Payload  map[string]interface{}
}

type TaskHandler struct {
	apiClient api.OrganizationsAPIClientInterface
}

func NewTaskHandler(apiClient api.OrganizationsAPIClientInterface) *TaskHandler {
	return &TaskHandler{
		apiClient: apiClient,
	}
}

func (t *TaskHandler) RouteTask(ctx context.Context, input map[string]interface{}) (interface{}, error) {
	var parsedInput CommonTaskInput
	if err := mapstructure.Decode(input, &parsedInput); err != nil {
		return nil, fmt.Errorf("unable to decode task input: %w", err)
	}

	zap.L().Sugar().Infow("routing task", "workflow", parsedInput.Workflow, "task", parsedInput.Task, "payload", parsedInput.Payload)

	switch parsedInput.Task {
	case TaskCreateMissingModuleAccountAssociations:
		var taskPayload CreateMissingModuleAccountAssociationsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.CreateMissingModuleAccountAssociations(ctx, taskPayload)
	case TaskDeactivateModuleAccountAssociation:
		var taskPayload DeactivateModuleAccountAssociationInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.DeactivateModuleAccountAssociation(ctx, taskPayload)
	case TaskClassifyModuleAccountAssociations:
		var taskPayload ClassifyModuleAccountAssociationsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ClassifyModuleAccountAssociations(ctx, taskPayload)
	case TaskListActiveModuleAccountAssociations:
		var taskPayload ListActiveModuleAccountAssociationsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ListActiveModuleAccountAssociations(ctx, taskPayload)
	case TaskListModulePropagationOrgUnits:
		var taskPayload ListModulePropagationOrgUnitsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ListModulePropagationOrgUnits(ctx, taskPayload)
	case TaskListOrgUnitAccounts:
		var taskPayload ListOrgUnitAccountsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ListOrgUnitAccounts(ctx, taskPayload)
	case TaskScheduleTerraformApplyDestroy:
		var taskPayload ScheduleTerraformApplyDestroyInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformApplyDestroy(ctx, taskPayload)
	case TaskScheduleTerraformApply:
		var taskPayload ScheduleTerraformApplyInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformApply(ctx, taskPayload)
	case TaskScheduleTerraformPlanDestroy:
		var taskPayload ScheduleTerraformPlanDestroyInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformPlanDestroy(ctx, taskPayload)
	case TaskScheduleTerraformPlan:
		var taskPayload ScheduleTerraformPlanInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformPlan(ctx, taskPayload)
	default:
		return nil, fmt.Errorf("unknown task: %s", parsedInput.Task)
	}
}
