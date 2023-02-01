package workflows

import (
	"time"

	"github.com/sheacloud/tfom/internal/temporal/activities"
	"github.com/sheacloud/tfom/pkg/models"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func ModulePropagationExecutionWorkflow(ctx workflow.Context, modulePropagationExecutionRequest *models.ModulePropagationExecutionRequest) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 1,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
	}
	childOptions := workflow.ChildWorkflowOptions{
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	ctx = workflow.WithChildOptions(ctx, childOptions)
	logger := workflow.GetLogger(ctx)
	var a *activities.Activities

	// update the request status to running
	err := workflow.ExecuteActivity(ctx, a.UpdateModulePropagationExecutionRequestStatus, modulePropagationExecutionRequest.ID, models.RequestStatusRunning).Get(ctx, nil)
	if err != nil {
		return err
	}

	// get all of the org accounts for the module propagation
	var orgAccounts []*models.OrgAccount
	err = workflow.ExecuteActivity(ctx, a.ListModulePropagationOrgAccounts, modulePropagationExecutionRequest.ModulePropagationID).Get(ctx, &orgAccounts)
	if err != nil {
		return err
	}

	// get all of the existing module assignments for the module propagation
	var existingModuleAssignments []*models.ModuleAssignment
	err = workflow.ExecuteActivity(ctx, a.ListActiveModuleAssignmentsForPropagation, modulePropagationExecutionRequest.ModulePropagationID).Get(ctx, &existingModuleAssignments)
	if err != nil {
		return err
	}

	// update the module assignments
	err = workflow.ExecuteActivity(ctx, a.UpdateModulePropagationAssignments, modulePropagationExecutionRequest.ModulePropagationID, existingModuleAssignments).Get(ctx, &existingModuleAssignments)
	if err != nil {
		return err
	}

	// classify the module assignments
	var modulePropagationAssignments *activities.ModulePropagationAssignments
	err = workflow.ExecuteActivity(ctx, a.ClassifyModulePropagationAssignments, orgAccounts, existingModuleAssignments).Get(ctx, &modulePropagationAssignments)
	if err != nil {
		return err
	}

	// create the missing module assignments
	var newModuleAssignments []*models.ModuleAssignment
	err = workflow.ExecuteActivity(ctx, a.CreateModuleAssignments, modulePropagationExecutionRequest.ModulePropagationID, modulePropagationAssignments.AccountsNeedingModuleAssignments).Get(ctx, &newModuleAssignments)
	if err != nil {
		return err
	}

	activeModuleAssignments := append(modulePropagationAssignments.ActiveModuleAssignments, newModuleAssignments...)
	var futures []workflow.ChildWorkflowFuture
	for _, moduleAssignment := range activeModuleAssignments {
		var terraformExecutionRequest *models.TerraformExecutionRequest
		err = workflow.ExecuteActivity(ctx, a.CreateTerraformExecutionRequest, moduleAssignment.ID, false, modulePropagationExecutionRequest.ModulePropagationID, modulePropagationExecutionRequest.ID).Get(ctx, &terraformExecutionRequest)
		if err != nil {
			return err
		}
		future := workflow.ExecuteChildWorkflow(ctx, TerraformExecutionWorkflow, terraformExecutionRequest)
		futures = append(futures, future)
	}
	for _, moduleAssignment := range modulePropagationAssignments.InactiveModuleAssignments {
		var terraformExecutionRequest *models.TerraformExecutionRequest
		err = workflow.ExecuteActivity(ctx, a.CreateTerraformExecutionRequest, moduleAssignment.ID, true, modulePropagationExecutionRequest.ModulePropagationID, modulePropagationExecutionRequest.ID).Get(ctx, &terraformExecutionRequest)
		if err != nil {
			return err
		}
		future := workflow.ExecuteChildWorkflow(ctx, TerraformExecutionWorkflow, terraformExecutionRequest)
		futures = append(futures, future)
	}
	var errs []error
	// wait for futures to complete
	for _, future := range futures {
		if err := future.Get(ctx, nil); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		logger.Error("One or more child workflows failed", "errors", errs)
		return errs[0]
	}

	// update the request status to completed
	err = workflow.ExecuteActivity(ctx, a.UpdateModulePropagationExecutionRequestStatus, modulePropagationExecutionRequest.ID, models.RequestStatusSucceeded).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
