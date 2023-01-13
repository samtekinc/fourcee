package api

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *APIClient) GetApplyExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	results, err := c.dbClient.GetApplyExecutionRequestsByIds(ctx, keys.Keys())
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

func (c *APIClient) GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, withOutputs bool) (*models.ApplyExecutionRequest, error) {
	thunk := c.applyExecutionRequestsLoader.Load(ctx, dataloader.StringKey(applyExecutionRequestId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	applyExecutionRequest := result.(*models.ApplyExecutionRequest)

	if withOutputs {
		// fetch init and apply outputs from S3
		if applyExecutionRequest.InitOutputKey != "" {
			initOutput, err := c.DownloadTerraformApplyInitResults(ctx, applyExecutionRequest.InitOutputKey)
			if err != nil {
				return nil, err
			}
			applyExecutionRequest.InitOutput = initOutput
		}

		if applyExecutionRequest.ApplyOutputKey != "" {
			applyOutput, err := c.DownloadTerraformApplyResults(ctx, applyExecutionRequest.ApplyOutputKey)
			if err != nil {
				return nil, err
			}
			applyExecutionRequest.ApplyOutput = applyOutput
		}
	}

	return applyExecutionRequest, nil
}

func (c *APIClient) GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string, withOutputs bool) (*models.ApplyExecutionRequests, error) {
	requests, err := c.dbClient.GetApplyExecutionRequests(ctx, limit, cursor)
	if err != nil {
		return nil, err
	}

	if withOutputs {
		for i := range requests.Items {
			// fetch init and apply outputs from S3
			if requests.Items[i].InitOutputKey != "" {
				initOutput, err := c.DownloadTerraformApplyInitResults(ctx, requests.Items[i].InitOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].InitOutput = initOutput
			}

			if requests.Items[i].ApplyOutputKey != "" {
				applyOutput, err := c.DownloadTerraformApplyResults(ctx, requests.Items[i].ApplyOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].ApplyOutput = applyOutput
			}
		}
	}

	return requests, nil
}

func (c *APIClient) GetApplyExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string, withOutputs bool) (*models.ApplyExecutionRequests, error) {
	requests, err := c.dbClient.GetApplyExecutionRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
	if err != nil {
		return nil, err
	}

	if withOutputs {
		for i := range requests.Items {
			// fetch init and apply outputs from S3
			if requests.Items[i].InitOutputKey != "" {
				initOutput, err := c.DownloadTerraformApplyInitResults(ctx, requests.Items[i].InitOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].InitOutput = initOutput
			}

			if requests.Items[i].ApplyOutputKey != "" {
				applyOutput, err := c.DownloadTerraformApplyResults(ctx, requests.Items[i].ApplyOutputKey)
				if err != nil {
					return nil, err
				}
				requests.Items[i].ApplyOutput = applyOutput
			}
		}
	}

	return requests, nil
}

func (c *APIClient) PutApplyExecutionRequest(ctx context.Context, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeApplyExecutionRequest)
	if err != nil {
		return nil, err
	}

	applyExecutionRequest := models.ApplyExecutionRequest{
		ApplyExecutionRequestId:      applyExecutionRequestId.String(),
		ModuleAssignmentId:           input.ModuleAssignmentId,
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		TerraformPlanBase64:          input.TerraformPlanBase64,
		AdditionalArguments:          input.AdditionalArguments,
		Status:                       models.RequestStatusPending,
		RequestTime:                  time.Now().UTC(),
	}
	err = c.dbClient.PutApplyExecutionRequest(ctx, &applyExecutionRequest)
	if err != nil {
		return nil, err
	}

	// Start Workflow
	err = c.startTerraformCommandWorkflow(ctx, &TerraformExecutionInput{
		RequestType: "apply",
		RequestId:   applyExecutionRequestId.String(),
		TaskToken:   applyExecutionRequest.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}

func (c *APIClient) UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, input *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error) {
	return c.dbClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequestId, input)
}

func (c *APIClient) UploadTerraformApplyInitResults(ctx context.Context, planExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	return c.dbClient.UploadTerraformApplyInitResults(ctx, planExecutionRequestId, initResults)
}

func (c *APIClient) UploadTerraformApplyResults(ctx context.Context, planExecutionRequestId string, planResults *models.TerraformApplyOutput) (string, error) {
	return c.dbClient.UploadTerraformApplyResults(ctx, planExecutionRequestId, planResults)
}

func (c *APIClient) DownloadTerraformApplyInitResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformInitOutput, error) {
	return c.dbClient.DownloadTerraformApplyInitResults(ctx, applyExecutionRequestId)
}

func (c *APIClient) DownloadTerraformApplyResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformApplyOutput, error) {
	return c.dbClient.DownloadTerraformApplyResults(ctx, applyExecutionRequestId)
}
