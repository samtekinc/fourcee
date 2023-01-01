resource "aws_cloudwatch_log_group" "execute_terraform_workflow" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-execute-terraform-workflow"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "execute_terraform_workflow" {
  name     = "${var.prefix}-execute-terraform-workflow"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.execute_terraform_workflow.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "CreateTerraformWorkflowRequest",
  "States": {
    "CreateTerraformWorkflowRequest": {
      "Type": "Task",
      "Resource": "arn:aws:states:::lambda:invoke",
      "OutputPath": "$.Payload",
      "Parameters": {
        "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
        "Payload": {
          "Payload.$": "$.StatePayload",
          "Task": "CreateTerraformWorkflowRequest",
          "Workflow": "ExecuteTerraformApply"
        }
      },
      "Retry": [
        {
          "ErrorEquals": [
            "Lambda.ServiceException",
            "Lambda.AWSLambdaException",
            "Lambda.SdkClientException",
            "Lambda.TooManyRequestsException"
          ],
          "IntervalSeconds": 2,
          "MaxAttempts": 6,
          "BackoffRate": 2
        }
      ],
      "TimeoutSeconds": 3660,
      "Next": "Parallel"
    },
    "Parallel": {
      "Type": "Parallel",
      "Next": "UpdateWorkflowSuccess",
      "Branches": [
        {
          "StartAt": "UpdateWorkflowRunning",
          "States": {
            "UpdateWorkflowRunning": {
              "Type": "Task",
              "Resource": "arn:aws:states:::dynamodb:updateItem",
              "Parameters": {
                "TableName": "${aws_dynamodb_table.terraform_workflow_requests.name}",
                "Key": {
                  "TerraformWorkflowRequestId": {
                    "S.$": "$.TerraformWorkflowRequestId"
                  }
                },
                "UpdateExpression": "SET #s = :status",
                "ExpressionAttributeValues": {
                  ":status": {
                    "S": "RUNNING"
                  }
                },
                "ExpressionAttributeNames": {
                  "#s": "Status"
                }
              },
              "Next": "ScheduleTerraformPlan",
              "ResultPath": null
            },
            "ScheduleTerraformPlan": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
              "Parameters": {
                "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
                "Payload": {
                  "Payload": {
                    "TerraformWorkflowRequestId.$": "$.TerraformWorkflowRequestId",
                    "TaskToken.$": "$$.Task.Token"
                  },
                  "Task": "ScheduleTerraformPlan",
                  "Workflow": "ExecuteTerraformApply"
                }
              },
              "Retry": [
                {
                  "ErrorEquals": [
                    "Lambda.ServiceException",
                    "Lambda.AWSLambdaException",
                    "Lambda.SdkClientException",
                    "Lambda.TooManyRequestsException"
                  ],
                  "IntervalSeconds": 2,
                  "MaxAttempts": 6,
                  "BackoffRate": 2
                }
              ],
              "TimeoutSeconds": 3660,
              "ResultPath": null,
              "Next": "ScheduleTerraformApply"
            },
            "ScheduleTerraformApply": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
              "Parameters": {
                "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
                "Payload": {
                  "Payload": {
                    "TerraformWorkflowRequestId.$": "$.TerraformWorkflowRequestId",
                    "TaskToken.$": "$$.Task.Token"
                  },
                  "Task": "ScheduleTerraformApply",
                  "Workflow": "ExecuteTerraformApply"
                }
              },
              "Retry": [
                {
                  "ErrorEquals": [
                    "Lambda.ServiceException",
                    "Lambda.AWSLambdaException",
                    "Lambda.SdkClientException",
                    "Lambda.TooManyRequestsException"
                  ],
                  "IntervalSeconds": 2,
                  "MaxAttempts": 6,
                  "BackoffRate": 2
                }
              ],
              "TimeoutSeconds": 3660,
              "ResultPath": null,
              "End": true
            }
          }
        }
      ],
      "Catch": [
        {
          "ErrorEquals": [
            "States.ALL"
          ],
          "Next": "UpdateWorkflowFailure",
          "ResultPath": "$.Error"
        }
      ],
      "ResultPath": null
    },
    "UpdateWorkflowSuccess": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.terraform_workflow_requests.name}",
        "Key": {
          "TerraformWorkflowRequestId": {
            "S.$": "$.TerraformWorkflowRequestId"
          }
        },
        "UpdateExpression": "SET #s = :status",
        "ExpressionAttributeValues": {
          ":status": {
            "S": "SUCCEEDED"
          }
        },
        "ExpressionAttributeNames": {
          "#s": "Status"
        }
      },
      "Next": "Success"
    },
    "Success": {
      "Type": "Succeed"
    },
    "UpdateWorkflowFailure": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.terraform_workflow_requests.name}",
        "Key": {
          "TerraformWorkflowRequestId": {
            "S.$": "$.TerraformWorkflowRequestId"
          }
        },
        "UpdateExpression": "SET #s = :status",
        "ExpressionAttributeValues": {
          ":status": {
            "S": "FAILED"
          }
        },
        "ExpressionAttributeNames": {
          "#s": "Status"
        }
      },
      "Next": "Fail"
    },
    "Fail": {
      "Type": "Fail"
    }
  }
}
EOF
}