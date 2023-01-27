package workflows

import (
	"time"

	"github.com/sheacloud/tfom/internal/temporal/activities"
	"github.com/sheacloud/tfom/pkg/models"
	"go.temporal.io/sdk/workflow"
)

func ModulePropagationExecutionWorkflow(ctx workflow.Context, modulePropagationExecutionRequest *models.ModulePropagationExecutionRequest) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 120,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var a *activities.Activities

	// update the request status to running
	err := workflow.ExecuteActivity(ctx, a.UpdateModulePropagationExecutionRequestStatus, modulePropagationExecutionRequest.ID, models.RequestStatusRunning, nil).Get(ctx, nil)
	if err != nil {
		return err
	}

	// get all of the org accounts for the module propagation
	var orgAccounts []*models.OrgAccount
	err = workflow.ExecuteActivity(ctx, a.ListModulePropagationOrgAccounts, modulePropagationExecutionRequest.ModulePropagationID).Get(ctx, &orgAccounts)
	if err != nil {
		return err
	}

	// get all of the active module assignments for the module propagation
	var moduleAssignments []*models.ModuleAssignment
	err = workflow.ExecuteActivity(ctx, a.ListActiveModuleAssignmentsForPropagation, modulePropagationExecutionRequest.ModulePropagationID).Get(ctx, &moduleAssignments)
	if err != nil {
		return err
	}

	// update the request status to completed
	err = workflow.ExecuteActivity(ctx, a.UpdateModulePropagationExecutionRequestStatus, modulePropagationExecutionRequest.ID, models.RequestStatusSucceeded).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
