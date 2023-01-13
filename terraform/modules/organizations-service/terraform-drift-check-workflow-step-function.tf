resource "aws_cloudwatch_log_group" "terraform_drift_check" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-terraform-drift-check"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "terraform_drift_check" {
  name     = "${var.prefix}-terraform-drift-check"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.terraform_drift_check.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "UpdateWorkflowRunning",
  "States": {
    "UpdateWorkflowRunning": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "tfom-terraform-drift-check-requests",
        "Key": {
          "TerraformDriftCheckRequestId": {
            "S.$": "$.TerraformDriftCheckRequestId"
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
                "FunctionName": "arn:aws:lambda:us-east-1:306526781466:function:tfom-workflow-handler",
                "Payload": {
                  "Payload": {
                    "TerraformWorkflowRequestId.$": "$.TerraformDriftCheckRequestId",
                    "TaskToken.$": "$$.Task.Token"
                  },
                  "Task": "ScheduleTerraformPlan",
                  "Workflow.$": "$$.StateMachine.Name"
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
                "FunctionName": "arn:aws:lambda:us-east-1:306526781466:function:tfom-workflow-handler",
                "Payload": {
                  "Payload": {
                    "TerraformWorkflowRequestId.$": "$.TerraformDriftCheckRequestId"
                  },
                  "Task": "DetermineSyncStatus",
                  "Workflow.$": "$$.StateMachine.Name"
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
        "TableName": "tfom-terraform-drift-check-requests",
        "Key": {
          "TerraformDriftCheckRequestId": {
            "S.$": "$.TerraformDriftCheckRequestId"
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
      "Next": "SuccessTaskTokenExists"
    },
    "SuccessTaskTokenExists": {
      "Type": "Choice",
      "Choices": [
        {
          "Variable": "$$.Execution.Input.TaskToken",
          "IsPresent": true,
          "Next": "SendTaskSuccess"
        }
      ],
      "Default": "Success"
    },
    "SendTaskSuccess": {
      "Type": "Task",
      "Next": "Success",
      "Parameters": {
        "Output": {
          "Payload": {
            "TerraformDriftCheckRequestId.$": "$$.Execution.Input.TerraformDriftCheckRequestId"
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
        "TableName": "tfom-terraform-drift-check-requests",
        "Key": {
          "TerraformDriftCheckRequestId": {
            "S.$": "$.TerraformDriftCheckRequestId"
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
      "Next": "FailureTaskTokenExists"
    },
    "FailureTaskTokenExists": {
      "Type": "Choice",
      "Choices": [
        {
          "Variable": "$$.Execution.Input.TaskToken",
          "IsPresent": true,
          "Next": "SendTaskFailure"
        }
      ],
      "Default": "Fail"
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
