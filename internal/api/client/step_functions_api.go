package client

// type TerraformExecutionInput struct {
// 	RequestType string
// 	RequestId   string
// 	TaskToken   string
// }

// func (c *APIClient) startTerraformCommandWorkflow(ctx context.Context, input *TerraformExecutionInput) error {
// 	workflowExecutionInput, err := json.Marshal(input)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = c.sfnClient.StartExecution(ctx, &sfn.StartExecutionInput{
// 		StateMachineArn: aws.String(c.terraformCommandWorkflowArn),
// 		Input:           aws.String(string(workflowExecutionInput)),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
