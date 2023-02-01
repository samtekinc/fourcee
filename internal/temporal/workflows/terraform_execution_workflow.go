package workflows

import (
	"time"

	"github.com/sheacloud/tfom/internal/temporal/activities"
	"github.com/sheacloud/tfom/pkg/models"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func TerraformExecutionWorkflow(ctx workflow.Context, terraformExecutionWorkflowRequest *models.TerraformExecutionRequest) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 1,
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
	applyCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 60,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 1,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
		},
	})

	var a *activities.Activities

	// update the terraform execution request status to running
	err := workflow.ExecuteActivity(ctx, a.UpdateTerraformExecutionRequestStatus, terraformExecutionWorkflowRequest.ID, models.RequestStatusRunning).Get(ctx, nil)
	if err != nil {
		return err
	}

	// if the destroy flag is not set, activate the module assignment
	if !terraformExecutionWorkflowRequest.Destroy {
		err = workflow.ExecuteActivity(ctx, a.ActivateModuleAssignment, terraformExecutionWorkflowRequest.ModuleAssignmentID).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Get the plan execution request
	var newPlanExecutionRequest *models.NewPlanExecutionRequest
	err = workflow.ExecuteActivity(ctx, a.BuildNewPlanExecutionRequest, terraformExecutionWorkflowRequest.ModuleAssignmentID, nil, terraformExecutionWorkflowRequest.ID, terraformExecutionWorkflowRequest.Destroy).Get(ctx, &newPlanExecutionRequest)
	if err != nil {
		return err
	}

	// update the terraform execution request with the plan execution request
	var planExecutionRequest *models.PlanExecutionRequest
	err = workflow.ExecuteActivity(ctx, a.UpdateTerraformExecutionPlanRequest, terraformExecutionWorkflowRequest.ID, newPlanExecutionRequest).Get(ctx, &planExecutionRequest)
	if err != nil {
		return err
	}

	// execute the plan
	err = workflow.ExecuteActivity(planCtx, a.TerraformPlan, planExecutionRequest).Get(ctx, &planExecutionRequest)
	if err != nil {
		return err
	}

	// Get the apply execution request
	var newApplyExecutionRequest *models.NewApplyExecutionRequest
	err = workflow.ExecuteActivity(ctx, a.BuildNewApplyExecutionRequest, terraformExecutionWorkflowRequest.ModuleAssignmentID, terraformExecutionWorkflowRequest.ID).Get(ctx, &newApplyExecutionRequest)
	if err != nil {
		return err
	}

	// update the terraform execution request with the apply execution request
	var applyExecutionRequest *models.ApplyExecutionRequest
	err = workflow.ExecuteActivity(ctx, a.UpdateTerraformExecutionApplyRequest, terraformExecutionWorkflowRequest.ID, newApplyExecutionRequest).Get(ctx, &applyExecutionRequest)
	if err != nil {
		return err
	}

	// execute the apply
	err = workflow.ExecuteActivity(applyCtx, a.TerraformApply, applyExecutionRequest).Get(ctx, &applyExecutionRequest)
	if err != nil {
		return err
	}

	// if the destroy flag is set, deactivate the module assignment
	if terraformExecutionWorkflowRequest.Destroy {
		err = workflow.ExecuteActivity(ctx, a.DeactivateModuleAssignment, terraformExecutionWorkflowRequest.ModuleAssignmentID).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// update the terraform execution request status to completed
	err = workflow.ExecuteActivity(ctx, a.UpdateTerraformExecutionRequestStatus, terraformExecutionWorkflowRequest.ID, models.RequestStatusSucceeded).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
