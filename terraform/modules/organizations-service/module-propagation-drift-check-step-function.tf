
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
              "Next": "ClassifyModuleAssignments",
              "OutputPath": "$.Output"
            },
            "ClassifyModuleAssignments": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke",
              "OutputPath": "$.Payload",
              "Parameters": {
                "Payload": {
                  "Payload.$": "$",
                  "Task": "ClassifyModuleAssignments",
                  "Workflow.$": "$$.StateMachine.Name"
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
                  "StartAt": "CreateMissingModuleAssignments",
                  "States": {
                    "CreateMissingModuleAssignments": {
                      "Type": "Task",
                      "Resource": "arn:aws:states:::lambda:invoke",
                      "OutputPath": "$.Payload",
                      "Parameters": {
                        "Payload": {
                          "Payload": {
                            "ModulePropagationId.$": "$$.Execution.Input.ModulePropagationId",
                            "AccountsNeedingModuleAssignments.$": "$.AccountsNeedingModuleAssignments",
                            "ActiveModuleAssignments.$": "$.ActiveModuleAssignments"
                          },
                          "Task": "CreateMissingModuleAssignments",
                          "Workflow.$": "$$.StateMachine.Name"
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
                      "Next": "MapActiveModuleAssignments"
                    },
                    "MapActiveModuleAssignments": {
                      "Type": "Map",
                      "ItemProcessor": {
                        "ProcessorConfig": {
                          "Mode": "INLINE"
                        },
                        "StartAt": "ScheduleTerraformDriftCheckWorkflow",
                        "States": {
                          "ScheduleTerraformDriftCheckWorkflow": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
                            "OutputPath": "$.Payload",
                            "Parameters": {
                              "Payload": {
                                "Payload": {
                                  "ModuleAssignment.$": "$",
                                  "ModulePropagationId.$": "$$.Execution.Input.ModulePropagationId",
                                  "ModulePropagationDriftCheckRequestId.$": "$$.Execution.Input.ModulePropagationDriftCheckRequestId",
                                  "Destroy": false,
                                  "TaskToken.$": "$$.Task.Token"
                                },
                                "Task": "ScheduleTerraformDriftCheckWorkflow",
                                "Workflow.$": "$$.StateMachine.Name"
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
                      "ItemsPath": "$.ActiveModuleAssignments"
                    }
                  }
                },
                {
                  "StartAt": "MapInactiveModuleAssignments",
                  "States": {
                    "MapInactiveModuleAssignments": {
                      "Type": "Map",
                      "ItemProcessor": {
                        "ProcessorConfig": {
                          "Mode": "INLINE"
                        },
                        "StartAt": "ScheduleTerraformDriftCheckWorkflowDestroy",
                        "States": {
                          "ScheduleTerraformDriftCheckWorkflowDestroy": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
                            "OutputPath": "$.Payload",
                            "Parameters": {
                              "Payload": {
                                "Payload": {
                                  "ModuleAssignment.$": "$",
                                  "ModulePropagationId.$": "$$.Execution.Input.ModulePropagationId",
                                  "ModulePropagationDriftCheckRequestId.$": "$$.Execution.Input.ModulePropagationDriftCheckRequestId",
                                  "Destroy": false,
                                  "TaskToken.$": "$$.Task.Token"
                                },
                                "Task": "ScheduleTerraformDriftCheckWorkflowDestroy",
                                "Workflow.$": "$$.StateMachine.Name"
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
                      "ItemsPath": "$.InactiveModuleAssignments"
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



