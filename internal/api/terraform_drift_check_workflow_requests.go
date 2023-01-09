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

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequest(ctx context.Context, modulePropagationDriftCheckRequestId string) (*models.TerraformDriftCheckWorkflowRequest, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequest(ctx, modulePropagationDriftCheckRequestId)
}

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequests(ctx, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequestsByModulePropagationDriftCheckRequestId(ctx context.Context, modulePropagationDriftCheckRequestId string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequestsByModulePropagationDriftCheckRequestId(ctx, modulePropagationDriftCheckRequestId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetTerraformDriftCheckWorkflowRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	return c.dbClient.GetTerraformDriftCheckWorkflowRequestsByModuleAssignmentId(ctx, moduleAssignmentId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutTerraformDriftCheckWorkflowRequest(ctx context.Context, input *models.NewTerraformDriftCheckWorkflowRequest) (*models.TerraformDriftCheckWorkflowRequest, error) {
	id, err := identifiers.NewIdentifier(identifiers.ResourceTypeTerraformDriftCheckWorkflowRequest)
	if err != nil {
		return nil, err
	}

	terraformDriftCheckRequest := models.TerraformDriftCheckWorkflowRequest{
		TerraformDriftCheckWorkflowRequestId: id.String(),
		ModuleAssignmentId:                   input.ModuleAssignmentId,
		RequestTime:                          time.Now().UTC(),
		Status:                               models.RequestStatusPending,
		Destroy:                              input.Destroy,
		CallbackTaskToken:                    input.CallbackTaskToken,
		SyncStatus:                           models.TerraformDriftCheckStatusPending,
		ModulePropagationId:                  input.ModulePropagationId,
		ModulePropagationDriftCheckRequestId: input.ModulePropagationDriftCheckRequestId,
	}

	err = c.dbClient.PutTerraformDriftCheckWorkflowRequest(ctx, &terraformDriftCheckRequest)
	if err != nil {
		return nil, err
	}

	workflowExecutionInput, err := json.Marshal(map[string]string{
		"TerraformDriftCheckWorkflowRequestId": id.String(),
		"TaskToken":                            input.CallbackTaskToken,
	})
	if err != nil {
		return nil, err
	}

	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
		StateMachineArn: aws.String(c.terraformDriftCheckWorkflowArn),
		Input:           aws.String(string(workflowExecutionInput)),
	})
	if err != nil {
		return nil, err
	}

	return &terraformDriftCheckRequest, nil
}

func (c *OrganizationsAPIClient) UpdateTerraformDriftCheckWorkflowRequest(ctx context.Context, modulePropagationDriftCheckRequestId string, update *models.TerraformDriftCheckWorkflowRequestUpdate) (*models.TerraformDriftCheckWorkflowRequest, error) {
	return c.dbClient.UpdateTerraformDriftCheckWorkflowRequest(ctx, modulePropagationDriftCheckRequestId, update)
}
