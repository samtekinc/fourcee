
resource "aws_cloudwatch_log_group" "module_propagation_drift_check" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-module-propagation-drift-check"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "module_propagation_drift_check" {
  name     = "${var.prefix}-module-propagation-drift-check"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.module_propagation_drift_check.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "UpdateExecutionRequestRunning",
  "States": {
    "UpdateExecutionRequestRunning": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.module_propagation_drift_check_requests.name}",
        "Key": {
          "ModulePropagationId": {
            "S.$": "$.ModulePropagationId"
          },
          "ModulePropagationDriftCheckRequestId": {
            "S.$": "$.ModulePropagationDriftCheckRequestId"
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
      "Next": "Parallel",
      "ResultPath": null
    },
    "Parallel": {
      "Type": "Parallel",
      "Next": "UpdateExecutionRequestCompleted",
      "Branches": [
        {
          "StartAt": "ListModulePropagationAccounts",
          "States": {
            "ListModulePropagationAccounts": {
              "Type": "Task",
              "Resource": "arn:aws:states:::states:startExecution.sync:2",
              "Parameters": {
                "StateMachineArn": "${aws_sfn_state_machine.list_mp_accounts.arn}",
                "Input": {
                  "StatePayload.$": "$",
                  "AWS_STEP_FUNCTIONS_STARTED_BY_EXECUTION_ID.$": "$$.Execution.Id"
                }
              },
              "Next": "ClassifyModuleAccountAssociations",
              "OutputPath": "$.Output"
            },
            "ClassifyModuleAccountAssociations": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke",
              "OutputPath": "$.Payload",
              "Parameters": {
                "Payload": {
                  "Payload.$": "$",
                  "Task": "ClassifyModuleAccountAssociations",
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
              "Next": "ExecuteTerraform"
            },
            "ExecuteTerraform": {
              "Type": "Parallel",
              "Branches": [
                {
                  "StartAt": "CreateMissingModuleAccountAssociations",
                  "States": {
                    "CreateMissingModuleAccountAssociations": {
                      "Type": "Task",
                      "Resource": "arn:aws:states:::lambda:invoke",
                      "OutputPath": "$.Payload",
                      "Parameters": {
                        "Payload": {
                          "Payload": {
                            "ModulePropagationId.$": "$$.Execution.Input.ModulePropagationId",
                            "AccountsNeedingModuleAccountAssociations.$": "$.AccountsNeedingModuleAccountAssociations",
                            "ActiveModuleAccountAssociations.$": "$.ActiveModuleAccountAssociations"
                          },
                          "Task": "CreateMissingModuleAccountAssociations",
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
                      "Next": "MapActiveModuleAccountAssociations"
                    },
                    "MapActiveModuleAccountAssociations": {
                      "Type": "Map",
                      "ItemProcessor": {
                        "ProcessorConfig": {
                          "Mode": "INLINE"
                        },
                        "StartAt": "RunTerraformDriftCheck",
                        "States": {
                          "RunTerraformDriftCheck": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::states:startExecution.sync:2",
                            "Parameters": {
                              "StateMachineArn": "${aws_sfn_state_machine.terraform_drift_check_workflow.arn}",
                              "Input": {
                                "StatePayload": {
                                  "ModuleAccountAssociation.$": "$",
                                  "ModulePropagationId.$": "$$.Execution.Input.ModulePropagationId",
                                  "ModulePropagationDriftCheckRequestId.$": "$$.Execution.Input.ModulePropagationDriftCheckRequestId",
                                  "Destroy": false
                                },
                                "AWS_STEP_FUNCTIONS_STARTED_BY_EXECUTION_ID.$": "$$.Execution.Id"
                              }
                            },
                            "End": true,
                            "Retry": [
                              {
                                "ErrorEquals": [
                                  "States.TaskFailed"
                                ],
                                "BackoffRate": 2,
                                "IntervalSeconds": 10,
                                "MaxAttempts": 3
                              }
                            ]
                          }
                        }
                      },
                      "End": true,
                      "ItemsPath": "$.ActiveModuleAccountAssociations"
                    }
                  }
                },
                {
                  "StartAt": "MapInactiveModuleAccountAssociations",
                  "States": {
                    "MapInactiveModuleAccountAssociations": {
                      "Type": "Map",
                      "ItemProcessor": {
                        "ProcessorConfig": {
                          "Mode": "INLINE"
                        },
                        "StartAt": "RunTerraformDriftCheckDestroy",
                        "States": {
                          "RunTerraformDriftCheckDestroy": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::states:startExecution.sync:2",
                            "Parameters": {
                              "StateMachineArn": "${aws_sfn_state_machine.terraform_drift_check_workflow.arn}",
                              "Input": {
                                "StatePayload": {
                                  "ModuleAccountAssociation.$": "$",
                                  "ModulePropagationId.$": "$$.Execution.Input.ModulePropagationId",
                                  "ModulePropagationDriftCheckRequestId.$": "$$.Execution.Input.ModulePropagationDriftCheckRequestId",
                                  "Destroy": true
                                },
                                "AWS_STEP_FUNCTIONS_STARTED_BY_EXECUTION_ID.$": "$$.Execution.Id"
                              }
                            },
                            "Retry": [
                              {
                                "ErrorEquals": [
                                  "States.TaskFailed"
                                ],
                                "BackoffRate": 2,
                                "IntervalSeconds": 10,
                                "MaxAttempts": 3
                              }
                            ],
                            "End": true,
                            "ResultPath": null
                          }
                        }
                      },
                      "End": true,
                      "ItemsPath": "$.InactiveModuleAccountAssociations"
                    }
                  }
                }
              ],
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
          "Next": "UpdateExecutionRequestFailed",
          "ResultPath": "$.Error"
        }
      ],
      "ResultPath": null
    },
    "UpdateExecutionRequestCompleted": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.module_propagation_drift_check_requests.name}",
        "Key": {
          "ModulePropagationId": {
            "S.$": "$.ModulePropagationId"
          },
          "ModulePropagationDriftCheckRequestId": {
            "S.$": "$.ModulePropagationDriftCheckRequestId"
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
    "UpdateExecutionRequestFailed": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.module_propagation_drift_check_requests.name}",
        "Key": {
          "ModulePropagationId": {
            "S.$": "$.ModulePropagationId"
          },
          "ModulePropagationDriftCheckRequestId": {
            "S.$": "$.ModulePropagationDriftCheckRequestId"
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
      "Next": "Fail",
      "ResultPath": null
    },
    "Fail": {
      "Type": "Fail"
    }
  }
}
EOF
}



