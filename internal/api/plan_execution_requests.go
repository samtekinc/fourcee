package api

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetPlanExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetPlanExecutionRequestsByIds(ctx, keys.Keys())
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: &results[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string, withOutputs bool) (*models.PlanExecutionRequest, error) {
	thunk := c.planExecutionRequestsLoader.Load(ctx, dataloader.StringKey(planExecutionRequestId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	planExecutionRequest := result.(*models.PlanExecutionRequest)

	if withOutputs {
		// fetch init and plan outputs from S3
		if planExecutionRequest.InitOutputKey != "" {
			initOutput, err := c.DownloadTerraformPlanInitResults(ctx, planExecutionRequest.InitOutputKey)
			if err != nil {
				return nil, err
			}
			planExecutionRequest.InitOutput = initOutput
		}

		if planExecutionRequest.PlanOutputKey != "" {
			planOutput, err := c.DownloadTerraformPlanResults(ctx, planExecutionRequest.PlanOutputKey)
			if err != nil {
				return nil, err
			}
			planExecutionRequest.PlanOutput = planOutput
		}
	}

	return planExecutionRequest, nil
}

func (c *APIClient) GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string, withOutputs bool) (*models.PlanExecutionRequests, error) {
	requests, err := c.dbClient.GetPlanExecutionRequests(ctx, limit, cursor)
	if err != nil {
		return nil, err
	}

	if withOutputs {
		for i := range requests.Items {
			// fetch init and plan outputs from S3
			if requests.Items[i].InitOutputKey != "" {
				initOutput, err := c.DownloadTerraformPlanInitResults(ctx, requests.Items[i].InitOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].InitOutput = initOutput
			}

			if requests.Items[i].PlanOutputKey != "" {
				planOutput, err := c.DownloadTerraformPlanResults(ctx, requests.Items[i].PlanOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].PlanOutput = planOutput
			}
		}
	}

	return requests, nil
}

func (c *APIClient) GetPlanExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string, withOutputs bool) (*models.PlanExecutionRequests, error) {
	requests, err := c.dbClient.GetPlanExecutionRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
	if err != nil {
		return nil, err
	}

	if withOutputs {
		for i := range requests.Items {
			// fetch init and plan outputs from S3
			if requests.Items[i].InitOutputKey != "" {
				initOutput, err := c.DownloadTerraformPlanInitResults(ctx, requests.Items[i].InitOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].InitOutput = initOutput
			}

			if requests.Items[i].PlanOutputKey != "" {
				planOutput, err := c.DownloadTerraformPlanResults(ctx, requests.Items[i].PlanOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].PlanOutput = planOutput
			}
		}
	}

	return requests, nil
}

func (c *APIClient) PutPlanExecutionRequest(ctx context.Context, input *models.NewPlanExecutionRequest) (*models.PlanExecutionRequest, error) {
	planExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypePlanExecutionRequest)
	if err != nil {
		return nil, err
	}

	planExecutionRequest := models.PlanExecutionRequest{
		PlanExecutionRequestId:       planExecutionRequestId.String(),
		ModuleAssignmentId:           input.ModuleAssignmentId,
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.RequestStatusPending,
		RequestTime:                  time.Now().UTC(),
	}
	err = c.dbClient.PutPlanExecutionRequest(ctx, &planExecutionRequest)
	if err != nil {
		return nil, err
	}

	// Start Workflow
	err = c.startTerraformCommandWorkflow(ctx, &TerraformExecutionInput{
		RequestType: "plan",
		RequestId:   planExecutionRequestId.String(),
		TaskToken:   planExecutionRequest.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}

func (c *APIClient) UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, input *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error) {
	return c.dbClient.UpdatePlanExecutionRequest(ctx, planExecutionRequestId, input)
}

func (c *APIClient) UploadTerraformPlanInitResults(ctx context.Context, applyExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	return c.dbClient.UploadTerraformPlanInitResults(ctx, applyExecutionRequestId, initResults)
}

func (c *APIClient) UploadTerraformPlanResults(ctx context.Context, applyExecutionRequestId string, applyResults *models.TerraformPlanOutput) (string, error) {
	return c.dbClient.UploadTerraformPlanResults(ctx, applyExecutionRequestId, applyResults)
}

func (c *APIClient) DownloadTerraformPlanInitResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformInitOutput, error) {
	return c.dbClient.DownloadTerraformPlanInitResults(ctx, applyExecutionRequestId)
}

func (c *APIClient) DownloadTerraformPlanResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformPlanOutput, error) {
	return c.dbClient.DownloadTerraformPlanResults(ctx, applyExecutionRequestId)
}
