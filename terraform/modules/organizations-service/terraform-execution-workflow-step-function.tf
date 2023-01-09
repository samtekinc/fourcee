resource "aws_cloudwatch_log_group" "terraform_execution_workflow" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-terraform-execution-workflow"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "terraform_execution_workflow" {
  name     = "${var.prefix}-terraform-execution-workflow"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.terraform_execution_workflow.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "UpdateWorkflowRunning",
  "States": {
    "UpdateWorkflowRunning": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.terraform_execution_workflow_requests.name}",
        "Key": {
          "TerraformExecutionWorkflowRequestId": {
            "S.$": "$.TerraformExecutionWorkflowRequestId"
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
      "Next": "Parallel",
      "ResultPath": null
    },
    "Parallel": {
      "Type": "Parallel",
      "Next": "UpdateWorkflowSuccess",
      "Branches": [
        {
          "StartAt": "ScheduleTerraformPlan",
          "States": {
            "ScheduleTerraformPlan": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
              "Parameters": {
                "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
                "Payload": {
                  "Payload": {
                    "TerraformWorkflowRequestId.$": "$.TerraformExecutionWorkflowRequestId",
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
                    "TerraformWorkflowRequestId.$": "$.TerraformExecutionWorkflowRequestId",
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
        "TableName": "${aws_dynamodb_table.terraform_execution_workflow_requests.name}",
        "Key": {
          "TerraformExecutionWorkflowRequestId": {
            "S.$": "$.TerraformExecutionWorkflowRequestId"
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
      "Next": "SendTaskSuccess"
    },
    "SendTaskSuccess": {
      "Type": "Task",
      "Next": "Success",
      "Parameters": {
        "Output": {
          "Payload": {
            "TerraformExecutionWorkflowRequestId.$": "$$.Execution.Input.TerraformExecutionWorkflowRequestId"
          }
        },
        "TaskToken.$": "$$.Execution.Input.TaskToken"
      },
      "Resource": "arn:aws:states:::aws-sdk:sfn:sendTaskSuccess"
    },
    "Success": {
      "Type": "Succeed"
    },
    "UpdateWorkflowFailure": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.terraform_execution_workflow_requests.name}",
        "Key": {
          "TerraformExecutionWorkflowRequestId": {
            "S.$": "$.TerraformExecutionWorkflowRequestId"
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
      "Next": "SendTaskFailure"
    },
    "SendTaskFailure": {
      "Type": "Task",
      "Next": "Fail",
      "Parameters": {
        "TaskToken.$": "$$.Execution.Input.TaskToken"
      },
      "Resource": "arn:aws:states:::aws-sdk:sfn:sendTaskFailure"
    },
    "Fail": {
      "Type": "Fail"
    }
  }
}
EOF
}
