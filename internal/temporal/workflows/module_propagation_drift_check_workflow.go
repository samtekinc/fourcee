package workflows

import (
	"time"

	"github.com/samtekinc/fourcee/internal/temporal/activities"
	"github.com/samtekinc/fourcee/pkg/models"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func ModulePropagationDriftCheckWorkflow(ctx workflow.Context, modulePropagationDriftCheckRequest *models.ModulePropagationDriftCheckRequest) error {
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
	err := workflow.ExecuteActivity(ctx, a.UpdateModulePropagationDriftCheckRequestStatus, modulePropagationDriftCheckRequest.ID, models.RequestStatusRunning, nil).Get(ctx, nil)
	if err != nil {
		return err
	}

	// get all of the org accounts for the module propagation
	var orgAccounts []*models.OrgAccount
	err = workflow.ExecuteActivity(ctx, a.ListModulePropagationOrgAccounts, modulePropagationDriftCheckRequest.ModulePropagationID).Get(ctx, &orgAccounts)
	if err != nil {
		return err
	}

	// get all of the existing module assignments for the module propagation
	var existingModuleAssignments []*models.ModuleAssignment
	err = workflow.ExecuteActivity(ctx, a.ListActiveModuleAssignmentsForPropagation, modulePropagationDriftCheckRequest.ModulePropagationID).Get(ctx, &existingModuleAssignments)
	if err != nil {
		return err
	}

	// update the module assignments
	err = workflow.ExecuteActivity(ctx, a.UpdateModulePropagationAssignments, modulePropagationDriftCheckRequest.ModulePropagationID, existingModuleAssignments).Get(ctx, &existingModuleAssignments)
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
	err = workflow.ExecuteActivity(ctx, a.CreateModuleAssignments, modulePropagationDriftCheckRequest.ModulePropagationID, modulePropagationAssignments.AccountsNeedingModuleAssignments).Get(ctx, &newModuleAssignments)
	if err != nil {
		return err
	}

	activeModuleAssignments := append(modulePropagationAssignments.ActiveModuleAssignments, newModuleAssignments...)
	var futures []workflow.ChildWorkflowFuture
	for _, moduleAssignment := range activeModuleAssignments {
		var terraformDriftCheckRequest *models.TerraformDriftCheckRequest
		err = workflow.ExecuteActivity(ctx, a.CreateTerraformDriftCheckRequest, moduleAssignment.ID, false, modulePropagationDriftCheckRequest.ModulePropagationID, modulePropagationDriftCheckRequest.ID).Get(ctx, &terraformDriftCheckRequest)
		if err != nil {
			return err
		}
		future := workflow.ExecuteChildWorkflow(ctx, TerraformDriftCheckWorkflow, terraformDriftCheckRequest)
		futures = append(futures, future)
	}
	for _, moduleAssignment := range modulePropagationAssignments.InactiveModuleAssignments {
		var terraformDriftCheckRequest *models.TerraformDriftCheckRequest
		err = workflow.ExecuteActivity(ctx, a.CreateTerraformDriftCheckRequest, moduleAssignment.ID, true, modulePropagationDriftCheckRequest.ModulePropagationID, modulePropagationDriftCheckRequest.ID).Get(ctx, &terraformDriftCheckRequest)
		if err != nil {
			return err
		}
		future := workflow.ExecuteChildWorkflow(ctx, TerraformDriftCheckWorkflow, terraformDriftCheckRequest)
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

	// tally the sync status
	var syncStatus *models.TerraformDriftCheckStatus
	err = workflow.ExecuteActivity(ctx, a.TallyTerraformDriftCheckSyncStatus, modulePropagationDriftCheckRequest.ID).Get(ctx, &syncStatus)
	if err != nil {
		return err
	}

	// update the request status to completed
	err = workflow.ExecuteActivity(ctx, a.UpdateModulePropagationDriftCheckRequestStatus, modulePropagationDriftCheckRequest.ID, models.RequestStatusSucceeded, syncStatus).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
