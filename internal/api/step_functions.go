package api

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

type TerraformExecutionWorkflowInput struct {
	RequestType string
	RequestId   string
	TaskToken   string
}

func (c *OrganizationsAPIClient) startTerraformCommandWorkflow(ctx context.Context, input *TerraformExecutionWorkflowInput) error {
	workflowExecutionInput, err := json.Marshal(input)
	if err != nil {
		return err
	}

	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
		StateMachineArn: aws.String(c.terraformCommandWorkflowArn),
		Input:           aws.String(string(workflowExecutionInput)),
	})
	if err != nil {
		return err
	}

	return nil
}
