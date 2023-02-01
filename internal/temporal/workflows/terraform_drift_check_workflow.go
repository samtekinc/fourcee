package workflows

import (
	"time"

	"github.com/sheacloud/tfom/internal/temporal/activities"
	"github.com/sheacloud/tfom/pkg/models"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func TerraformDriftCheckWorkflow(ctx workflow.Context, terraformDriftCheckWorkflowRequest *models.TerraformDriftCheckRequest) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 120,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	planCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 15,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 1,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
		},
	})

	var a *activities.Activities

	// update the terraform execution request status to running
	err := workflow.ExecuteActivity(ctx, a.UpdateTerraformDriftCheckRequestStatus, terraformDriftCheckWorkflowRequest.ID, models.RequestStatusRunning, nil).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Get the plan execution request
	var newPlanExecutionRequest *models.NewPlanExecutionRequest
	err = workflow.ExecuteActivity(ctx, a.BuildNewPlanExecutionRequest, terraformDriftCheckWorkflowRequest.ModuleAssignmentID, nil, terraformDriftCheckWorkflowRequest.ID, terraformDriftCheckWorkflowRequest.Destroy).Get(ctx, &newPlanExecutionRequest)
	if err != nil {
		return err
	}

	// update the terraform execution request with the plan execution request
	var planExecutionRequest *models.PlanExecutionRequest
	err = workflow.ExecuteActivity(ctx, a.UpdateTerraformDriftCheckPlanRequest, terraformDriftCheckWorkflowRequest.ID, newPlanExecutionRequest).Get(ctx, &planExecutionRequest)
	if err != nil {
		return err
	}

	// execute the plan
	err = workflow.ExecuteActivity(planCtx, a.TerraformPlan, planExecutionRequest).Get(ctx, &planExecutionRequest)
	if err != nil {
		return err
	}

	// determine the sync status
	var syncStatus *models.TerraformDriftCheckStatus
	err = workflow.ExecuteActivity(ctx, a.DeterminePlanSyncStatus, planExecutionRequest.ID).Get(ctx, &syncStatus)
	if err != nil {
		return err
	}

	// update the terraform execution request status to completed
	err = workflow.ExecuteActivity(ctx, a.UpdateTerraformDriftCheckRequestStatus, terraformDriftCheckWorkflowRequest.ID, models.RequestStatusSucceeded, syncStatus).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
