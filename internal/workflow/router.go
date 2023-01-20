package workflow

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/config"

	"go.uber.org/zap"
)

type Task string

type CommonTaskInput struct {
	Workflow string
	Task     Task
	Payload  map[string]interface{}
}

type TaskHandler struct {
	apiClient api.APIClientInterface
	config    *config.Config
}

func NewTaskHandler(apiClient api.APIClientInterface, config *config.Config) *TaskHandler {
	return &TaskHandler{
		apiClient: apiClient,
		config:    config,
	}
}

func (t *TaskHandler) RouteTask(ctx context.Context, input map[string]interface{}) (interface{}, error) {
	var parsedInput CommonTaskInput
	if err := mapstructure.Decode(input, &parsedInput); err != nil {
		return nil, fmt.Errorf("unable to decode task input: %w", err)
	}

	zap.L().Sugar().Infow("routing task", "workflow", parsedInput.Workflow, "task", parsedInput.Task, "payload", parsedInput.Payload)

	switch parsedInput.Task {
	case TaskCreateMissingModuleAssignments:
		var taskPayload CreateMissingModuleAssignmentsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.CreateMissingModuleAssignments(ctx, taskPayload)
	case TaskScheduleTerraformExecution:
		var taskPayload ScheduleTerraformExecutionInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformExecution(ctx, taskPayload)
	case TaskScheduleTerraformDriftCheck:
		var taskPayload ScheduleTerraformDriftCheckInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformDriftCheck(ctx, taskPayload)
	case TaskDeactivateModuleAssignment:
		var taskPayload DeactivateModuleAssignmentInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.DeactivateModuleAssignment(ctx, taskPayload)
	case TaskDetermineSyncStatus:
		var taskPayload DetermineSyncStatusInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.DetermineSyncStatus(ctx, taskPayload)
	case TaskClassifyModuleAssignments:
		var taskPayload ClassifyModuleAssignmentsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ClassifyModuleAssignments(ctx, taskPayload)
	case TaskListActiveModulePropagationAssignments:
		var taskPayload ListActiveModulePropagationAssignmentsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ListActiveModulePropagationAssignments(ctx, taskPayload)
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
	case TaskScheduleTerraformApply:
		var taskPayload ScheduleTerraformApplyInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformApply(ctx, taskPayload)
	case TaskScheduleTerraformPlan:
		var taskPayload ScheduleTerraformPlanInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformPlan(ctx, taskPayload)
	case TaskUpdateModuleAssignments:
		var taskPayload UpdateModuleAssignmentsInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.UpdateModuleAssignments(ctx, taskPayload)
	case TaskTallySyncStatus:
		var taskPayload TallySyncStatusInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.TallySyncStatus(ctx, taskPayload)
	default:
		return nil, fmt.Errorf("unknown task: %s", parsedInput.Task)
	}
}
