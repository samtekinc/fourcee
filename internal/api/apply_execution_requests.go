package api

import (
	"context"
	"time"

	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequest, err := c.dbClient.GetApplyExecutionRequest(ctx, applyExecutionRequestId)
	if err != nil {
		return nil, err
	}

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

	return applyExecutionRequest, nil
}

func (c *OrganizationsAPIClient) GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	requests, err := c.dbClient.GetApplyExecutionRequests(ctx, limit, cursor)
	if err != nil {
		return nil, err
	}

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

	return requests, nil
}

func (c *OrganizationsAPIClient) GetApplyExecutionRequestsByModulePropagationRequestId(ctx context.Context, modulePropagationRequestId string, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	requests, err := c.dbClient.GetApplyExecutionRequestsByModulePropagationRequestId(ctx, modulePropagationRequestId, limit, cursor)
	if err != nil {
		return nil, err
	}

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

	return requests, nil
}

func (c *OrganizationsAPIClient) GetApplyExecutionRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	requests, err := c.dbClient.GetApplyExecutionRequestsByModuleAccountAssociationKey(ctx, moduleAccountAssociationKey, limit, cursor)
	if err != nil {
		return nil, err
	}

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

	return requests, nil
}

func (c *OrganizationsAPIClient) PutApplyExecutionRequest(ctx context.Context, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeApplyExecutionRequest)
	if err != nil {
		return nil, err
	}

	applyExecutionRequest := models.ApplyExecutionRequest{
		ApplyExecutionRequestId:      applyExecutionRequestId.String(),
		TerraformVersion:             input.TerraformVersion,
		CallbackTaskToken:            input.CallbackTaskToken,
		StateKey:                     input.StateKey,
		ModulePropagationRequestId:   input.ModulePropagationRequestId,
		TerraformWorkflowRequestId:   input.TerraformWorkflowRequestId,
		ModuleAccountAssociationKey:  input.ModuleAccountAssociationKey,
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
	err = c.startTerraformExecutionWorkflow(ctx, &TerraformExecutionWorkflowInput{
		RequestType: "apply",
		RequestId:   applyExecutionRequestId.String(),
		TaskToken:   applyExecutionRequest.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}

func (c *OrganizationsAPIClient) UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, input *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error) {
	return c.dbClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequestId, input)
}

func (c *OrganizationsAPIClient) UploadTerraformApplyInitResults(ctx context.Context, planExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	return c.dbClient.UploadTerraformApplyInitResults(ctx, planExecutionRequestId, initResults)
}

func (c *OrganizationsAPIClient) UploadTerraformApplyResults(ctx context.Context, planExecutionRequestId string, planResults *models.TerraformApplyOutput) (string, error) {
	return c.dbClient.UploadTerraformApplyResults(ctx, planExecutionRequestId, planResults)
}

func (c *OrganizationsAPIClient) DownloadTerraformApplyInitResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformInitOutput, error) {
	return c.dbClient.DownloadTerraformApplyInitResults(ctx, applyExecutionRequestId)
}

func (c *OrganizationsAPIClient) DownloadTerraformApplyResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformApplyOutput, error) {
	return c.dbClient.DownloadTerraformApplyResults(ctx, applyExecutionRequestId)
}
