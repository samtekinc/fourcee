
resource "aws_cloudwatch_log_group" "terraform_command" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-terraform-command"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "terraform_command" {
  name     = "${var.prefix}-terraform-command"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.terraform_command.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "RequestType",
  "States": {
    "RequestType": {
      "Type": "Choice",
      "Choices": [
        {
          "Variable": "$.RequestType",
          "StringEquals": "plan",
          "Next": "UpdatePlanRequest"
        },
        {
          "Variable": "$.RequestType",
          "StringEquals": "apply",
          "Next": "UpdateApplyRequest"
        }
      ]
    },
    "UpdatePlanRequest": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.plan_execution_requests.name}",
        "Key": {
          "PlanExecutionRequestId": {
            "S.$": "$.RequestId"
          }
        },
        "UpdateExpression": "SET WorkflowExecutionId = :execId",
        "ExpressionAttributeValues": {
          ":execId": {
            "S.$": "$$.Execution.Id"
          }
        }
      },
      "Next": "ExecuteTerraform",
      "ResultPath": null
    },
    "UpdateApplyRequest": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.apply_execution_requests.name}",
        "Key": {
          "ApplyExecutionRequestId": {
            "S.$": "$.RequestId"
          }
        },
        "UpdateExpression": "SET WorkflowExecutionId = :execId, #s = :status",
        "ExpressionAttributeValues": {
          ":execId": {
            "S.$": "$$.Execution.Id"
          },
          ":status": {
            "S": "RUNNING"
          }
        },
        "ExpressionAttributeNames": {
          "#s": "Status"
        }
      },
      "Next": "ExecuteTerraform",
      "ResultPath": null
    },
    "FailApplyRequest": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.apply_execution_requests.name}",
        "Key": {
          "ApplyExecutionRequestId": {
            "S.$": "$.RequestId"
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
    "FailPlanRequest": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.plan_execution_requests.name}",
        "Key": {
          "PlanExecutionRequestId": {
            "S.$": "$.RequestId"
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
      "Parameters": {
        "TaskToken.$": "$$.Execution.Input.TaskToken"
      },
      "Resource": "arn:aws:states:::aws-sdk:sfn:sendTaskFailure",
      "Next": "Fail"
    },
    "Fail": {
      "Type": "Fail"
    },
    "ExecuteTerraform": {
      "Type": "Task",
      "Resource": "arn:aws:states:::ecs:runTask.waitForTaskToken",
      "TimeoutSeconds": 3600,
      "Parameters": {
        "LaunchType": "FARGATE",
        "Cluster": "${aws_ecs_cluster.tfom.arn}",
        "TaskDefinition": "${aws_ecs_task_definition.executor.arn}",
        "NetworkConfiguration": {
          "AwsvpcConfiguration": {
            "AssignPublicIp": "ENABLED",
            "Subnets": ${jsonencode(var.task_subnet_ids)},
            "SecurityGroups": [
              "${aws_security_group.executor_ecs.id}"
            ]
          }
        },
        "Overrides": {
          "ContainerOverrides": [
            {
              "Name": "executor",
              "Environment": [
                {
                  "Name": "REQUEST_TYPE",
                  "Value.$": "$.RequestType"
                },
                {
                  "Name": "REQUEST_ID",
                  "Value.$": "$.RequestId"
                },
                {
                  "Name": "TF_INSTALLATION_DIRECTORY",
                  "Value": "/efs/tf-installation"
                },
                {
                  "Name": "TF_WORKING_DIRECTORY",
                  "Value": "/tmp/tf-working"
                },
                {
                  "Name": "TASK_TOKEN",
                  "Value.$": "$$.Task.Token"
                }
              ]
            }
          ]
        }
      },
      "Catch": [
        {
          "ErrorEquals": [
            "States.ALL"
          ],
          "Next": "FailureRequestType",
          "ResultPath": "$.Error"
        }
      ],
      "Next": "SuccessRequestType",
      "ResultPath": null
    },
    "FailureRequestType": {
      "Type": "Choice",
      "Choices": [
        {
          "Variable": "$.RequestType",
          "StringEquals": "plan",
          "Next": "FailPlanRequest"
        },
        {
          "Variable": "$.RequestType",
          "StringEquals": "apply",
          "Next": "FailApplyRequest"
        }
      ]
    },
    "SuccessRequestType": {
      "Type": "Choice",
      "Choices": [
        {
          "Variable": "$.RequestType",
          "StringEquals": "plan",
          "Next": "SuccessPlanRequest"
        },
        {
          "Variable": "$.RequestType",
          "StringEquals": "apply",
          "Next": "SuccessApplyRequest"
        }
      ]
    },
    "SuccessPlanRequest": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.plan_execution_requests.name}",
        "Key": {
          "PlanExecutionRequestId": {
            "S.$": "$.RequestId"
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
    "SuccessApplyRequest": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.apply_execution_requests.name}",
        "Key": {
          "ApplyExecutionRequestId": {
            "S.$": "$.RequestId"
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
      "Parameters": {
        "Output": {
          "Payload": {
            "RequestId.$": "$$.Execution.Input.RequestId"
          }
        },
        "TaskToken.$": "$$.Execution.Input.TaskToken"
      },
      "Resource": "arn:aws:states:::aws-sdk:sfn:sendTaskSuccess",
      "Next": "Success"
    },
    "Success": {
      "Type": "Succeed"
    }
  }
}
EOF
}
