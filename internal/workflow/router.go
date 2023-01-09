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
	apiClient         api.OrganizationsAPIClientInterface
	remoteStateBucket string
	remoteStateRegion string
}

func NewTaskHandler(apiClient api.OrganizationsAPIClientInterface, remoteStateBucket string, remoteStateRegion string) *TaskHandler {
	return &TaskHandler{
		apiClient:         apiClient,
		remoteStateBucket: remoteStateBucket,
		remoteStateRegion: remoteStateRegion,
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
	case TaskScheduleTerraformExecutionWorkflow:
		var taskPayload ScheduleTerraformExecutionWorkflowInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformExecutionWorkflow(ctx, taskPayload)
	case TaskScheduleTerraformDriftCheckWorkflow:
		var taskPayload ScheduleTerraformDriftCheckWorkflowInput
		if err := mapstructure.Decode(parsedInput.Payload, &taskPayload); err != nil {
			return nil, fmt.Errorf("unable to decode task payload: %w", err)
		}
		return t.ScheduleTerraformDriftCheckWorkflow(ctx, taskPayload)
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
	default:
		return nil, fmt.Errorf("unknown task: %s", parsedInput.Task)
	}
}
