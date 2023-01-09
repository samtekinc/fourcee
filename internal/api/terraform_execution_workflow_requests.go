package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequest(ctx context.Context, modulePropagationExecutionRequestId string) (*models.TerraformExecutionWorkflowRequest, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequest(ctx, modulePropagationExecutionRequestId)
}

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequestsByModulePropagationExecutionRequestId(ctx, modulePropagationExecutionRequestId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformExecutionWorkflowRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	return c.dbClient.GetTerraformExecutionWorkflowRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutTerraformExecutionWorkflowRequest(ctx context.Context, input *models.NewTerraformExecutionWorkflowRequest) (*models.TerraformExecutionWorkflowRequest, error) {
	id, err := identifiers.NewIdentifier(identifiers.ResourceTypeTerraformExecutionWorkflowRequest)
	if err != nil {
		return nil, err
	}

	terraformExecutionRequest := models.TerraformExecutionWorkflowRequest{
		TerraformExecutionWorkflowRequestId: id.String(),
		ModuleAssignmentId:                  input.ModuleAssignmentId,
		RequestTime:                         time.Now().UTC(),
		Status:                              models.RequestStatusPending,
		Destroy:                             input.Destroy,
		CallbackTaskToken:                   input.CallbackTaskToken,
		ModulePropagationId:                 input.ModulePropagationId,
		ModulePropagationExecutionRequestId: input.ModulePropagationExecutionRequestId,
	}

	err = c.dbClient.PutTerraformExecutionWorkflowRequest(ctx, &terraformExecutionRequest)
	if err != nil {
		return nil, err
	}

	workflowExecutionInput, err := json.Marshal(map[string]string{
		"TerraformExecutionWorkflowRequestId": id.String(),
		"TaskToken":                           input.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
		StateMachineArn: aws.String(c.terraformExecutionWorkflowArn),
		Input:           aws.String(string(workflowExecutionInput)),
	})
	if err != nil {
		return nil, err
	}

	return &terraformExecutionRequest, nil
}

func (c *OrganizationsAPIClient) UpdateTerraformExecutionWorkflowRequest(ctx context.Context, modulePropagationExecutionRequestId string, update *models.TerraformExecutionWorkflowRequestUpdate) (*models.TerraformExecutionWorkflowRequest, error) {
	return c.dbClient.UpdateTerraformExecutionWorkflowRequest(ctx, modulePropagationExecutionRequestId, update)
}
