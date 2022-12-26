
resource "aws_cloudwatch_log_group" "execute_module_propagation" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-execute-module-propagation"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "execute_module_propagation" {
  name     = "${var.prefix}-execute-module-propagation"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.execute_module_propagation.arn}:*"
  }

  definition = <<EOF
{
  "Comment": "A description of my state machine",
  "StartAt": "UpdateExecutionRequestRunning",
  "States": {
    "UpdateExecutionRequestRunning": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.module_propagation_execution_requests.name}",
        "Key": {
          "ModulePropagationId": {
            "S.$": "$.ModulePropagationId"
          },
          "ModulePropagationExecutionRequestId": {
            "S.$": "$.ModulePropagationExecutionRequestId"
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
                          "Payload.$": "$",
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
                                  "OrgUnit.$": "$"
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
                  "StartAt": "ListActiveModuleAccountAssociations",
                  "States": {
                    "ListActiveModuleAccountAssociations": {
                      "Type": "Task",
                      "Resource": "arn:aws:states:::lambda:invoke",
                      "OutputPath": "$.Payload",
                      "Parameters": {
                        "Payload": {
                          "Payload.$": "$",
                          "Task": "ListActiveModuleAccountAssociations",
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
              "Next": "ClassifyModuleAccountAssociations"
            },
            "ClassifyModuleAccountAssociations": {
              "Type": "Task",
              "Resource": "arn:aws:states:::lambda:invoke",
              "OutputPath": "$.Payload",
              "Parameters": {
                "Payload": {
                  "Payload": {
                    "OrgAccountsPerOrgUnit.$": "$[0]",
                    "ActiveModuleAccountAssociations.$": "$[1].ModuleAccountAssociations"
                  },
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
                        "StartAt": "ExecuteTerraformApply",
                        "States": {
                          "ExecuteTerraformApply": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::states:startExecution.sync:2",
                            "Parameters": {
                              "StateMachineArn": "${aws_sfn_state_machine.execute_terraform_apply.arn}",
                              "Input": {
                                "StatePayload": {
                                  "ModuleAccountAssociation.$": "$",
                                  "ModulePropagationId.$": "$$.Execution.Input.ModulePropagationId",
                                  "ModulePropagationExecutionRequestId.$": "$$.Execution.Input.ModulePropagationExecutionRequestId"
                                },
                                "AWS_STEP_FUNCTIONS_STARTED_BY_EXECUTION_ID.$": "$$.Execution.Id"
                              }
                            },
                            "End": true
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
                        "StartAt": "ScheduleTerraformPlanDestroy-Callback",
                        "States": {
                          "ScheduleTerraformPlanDestroy-Callback": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
                            "OutputPath": "$.Payload",
                            "Parameters": {
                              "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
                              "Payload": {
                                "Payload.$": "$",
                                "Task": "ScheduleTerraformPlanDestroy",
                                "Workflow": "ExecuteModulePropagation"
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
                            "Next": "ScheduleTerraformApplyDestroy-Callback"
                          },
                          "ScheduleTerraformApplyDestroy-Callback": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
                            "OutputPath": "$.Payload",
                            "Parameters": {
                              "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
                              "Payload": {
                                "Payload.$": "$",
                                "Task": "ScheduleTerraformApplyDestroy",
                                "Workflow": "ExecuteModulePropagation"
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
                            "Next": "DeactivateModuleAccountAssociation"
                          },
                          "DeactivateModuleAccountAssociation": {
                            "Type": "Task",
                            "Resource": "arn:aws:states:::lambda:invoke",
                            "OutputPath": "$.Payload",
                            "Parameters": {
                              "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
                              "Payload": {
                                "Payload.$": "$",
                                "Task": "DeactivateModuleAccountAssociation",
                                "Workflow": "ExecuteModulePropagation"
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
                            "End": true
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
          "Next": "UpdateExecutionRequestFailed"
        }
      ],
      "ResultPath": null
    },
    "UpdateExecutionRequestCompleted": {
      "Type": "Task",
      "Resource": "arn:aws:states:::dynamodb:updateItem",
      "Parameters": {
        "TableName": "${aws_dynamodb_table.module_propagation_execution_requests.name}",
        "Key": {
          "ModulePropagationId": {
            "S.$": "$.ModulePropagationId"
          },
          "ModulePropagationExecutionRequestId": {
            "S.$": "$.ModulePropagationExecutionRequestId"
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
        "TableName": "${aws_dynamodb_table.module_propagation_execution_requests.name}",
        "Key": {
          "ModulePropagationId": {
            "S.$": "$.ModulePropagationId"
          },
          "ModulePropagationExecutionRequestId": {
            "S.$": "$.ModulePropagationExecutionRequestId"
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



