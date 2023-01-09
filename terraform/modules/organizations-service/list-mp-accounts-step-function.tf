
resource "aws_cloudwatch_log_group" "list_mp_accounts" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-list-mp-accounts"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "list_mp_accounts" {
  name     = "${var.prefix}-list-mp-accounts"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.list_mp_accounts.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "GatherImpactedAccounts",
  "States": {
    "GatherImpactedAccounts": {
      "Type": "Parallel",
      "Branches": [
        {
          "StartAt": "ListModulePropagationOrgUnits",
          "States": {
            "ListModulePropagationOrgUnits": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke",
              "OutputPath": "$.Payload",
              "Parameters": {
                "Payload": {
                  "Payload.$": "$.StatePayload",
                  "Task": "ListModulePropagationOrgUnits",
                  "Workflow": "ExecuteModulePropagation"
                },
                "FunctionName": "${aws_lambda_function.workflow_handler.arn}"
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
              "Next": "MapOrgUnits"
            },
            "MapOrgUnits": {
              "Type": "Map",
              "ItemProcessor": {
                "ProcessorConfig": {
                  "Mode": "INLINE"
                },
                "StartAt": "ListOrgUnitAccounts",
                "States": {
                  "ListOrgUnitAccounts": {
                    "Type": "Task",
                    "Resource": "arn:aws:states:::lambda:invoke",
                    "OutputPath": "$.Payload",
                    "Parameters": {
                      "Payload": {
                        "Payload": {
                          "OrgUnit.$": "$",
                          "CloudPlatform.$": "$$.Execution.Input.StatePayload.CloudPlatform"
                        },
                        "Task": "ListOrgUnitAccounts",
                        "Workflow": "ExecuteModulePropagation"
                      },
                      "FunctionName": "${aws_lambda_function.workflow_handler.arn}"
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
                    "End": true
                  }
                }
              },
              "End": true,
              "ItemsPath": "$.OrgUnits"
            }
          }
        },
        {
          "StartAt": "ListActiveModulePropagationAssignments",
          "States": {
            "ListActiveModulePropagationAssignments": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke",
              "OutputPath": "$.Payload",
              "Parameters": {
                "Payload": {
                  "Payload.$": "$.StatePayload",
                  "Task": "ListActiveModulePropagationAssignments",
                  "Workflow": "ExecuteModulePropagation"
                },
                "FunctionName": "${aws_lambda_function.workflow_handler.arn}"
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
              "End": true
            }
          }
        }
      ],
      "End": true,
      "ResultSelector": {
        "OrgAccountsPerOrgUnit.$": "$[0]",
        "ActiveModuleAssignments.$": "$[1].ModuleAssignments"
      }
    }
  }
}
EOF
}



