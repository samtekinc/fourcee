resource "aws_cloudwatch_log_group" "execute_terraform_apply" {
  # The /aws/vendedlogs/* path is special -- it gets policy length limit mitigation strategies.
  name              = "/aws/vendedlogs/AsyncWorkflow/${var.prefix}-execute-terraform-apply"
  retention_in_days = 731
}


resource "aws_sfn_state_machine" "execute_terraform_apply" {
  name     = "${var.prefix}-execute-terraform-apply"
  type     = "STANDARD"
  role_arn = aws_iam_role.step_functions_role.arn

  logging_configuration {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "${aws_cloudwatch_log_group.execute_terraform_apply.arn}:*"
  }

  definition = <<EOF
{
  "StartAt": "ScheduleTerraformPlan",
  "States": {
    "ScheduleTerraformPlan": {
      "Type": "Task",
      "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
      "OutputPath": "$.Payload",
      "Parameters": {
        "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
        "Payload": {
          "Payload": {
            "Input.$": "$.StatePayload",
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
      "Next": "ScheduleTerraformApply"
    },
    "ScheduleTerraformApply": {
      "Type": "Task",
      "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
      "OutputPath": "$.Payload",
      "Parameters": {
        "FunctionName": "${aws_lambda_function.workflow_handler.arn}",
        "Payload": {
          "Payload": {
            "Input.$": "$$.Execution.Input.StatePayload",
            "TaskToken.$": "$$.Task.Token",
            "PlanExecutionRequestId.$": "$.RequestId"
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
      "End": true,
      "TimeoutSeconds": 3660
    }
  }
}
EOF
}
