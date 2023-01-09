resource "aws_cloudwatch_log_group" "terraform_drift_check_workflow" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-terraform-drift-check-workflow"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "terraform_drift_check_workflow" {
  name     = "${var.prefix}-terraform-drift-check-workflow"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.terraform_drift_check_workflow.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "UpdateWorkflowRunning",
  "States": {
    "UpdateWorkflowRunning": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.terraform_drift_check_workflow_requests.name}",
        "Key": {
          "TerraformDriftCheckWorkflowRequestId": {
            "S.$": "$.TerraformDriftCheckWorkflowRequestId"
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
                    "TerraformWorkflowRequestId.$": "$.TerraformDriftCheckWorkflowRequestId",
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
              "Next": "DetermineSyncStatus"
            },
            "DetermineSyncStatus": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke",
              "Parameters": {
                "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
                "Payload": {
                  "Payload": {
                    "TerraformWorkflowRequestId.$": "$.TerraformDriftCheckWorkflowRequestId"
                  },
                  "Task": "DetermineSyncStatus",
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
        "TableName": "${aws_dynamodb_table.terraform_drift_check_workflow_requests.name}",
        "Key": {
          "TerraformDriftCheckWorkflowRequestId": {
            "S.$": "$.TerraformDriftCheckWorkflowRequestId"
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
            "TerraformDriftCheckWorkflowRequestId.$": "$$.Execution.Input.TerraformDriftCheckWorkflowRequestId"
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
        "TableName": "${aws_dynamodb_table.terraform_drift_check_workflow_requests.name}",
        "Key": {
          "TerraformDriftCheckWorkflowRequestId": {
            "S.$": "$.TerraformDriftCheckWorkflowRequestId"
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
