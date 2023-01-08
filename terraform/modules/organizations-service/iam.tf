resource "aws_iam_role" "ecs_task_execution_role" {
  name = "${var.prefix}-ecs-task-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_policy" "ecs_task_execution_policy" {
  name        = "${var.prefix}-ecs-task-execution-policy"
  path        = "/"
  description = "TFOM ECS task execution policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ssm:GetParameters",
        ]
        Effect = "Allow"
        Resource = [
          aws_ssm_parameter.arm_client_id.arn,
          aws_ssm_parameter.arm_client_secret.arn,
          aws_ssm_parameter.arm_tenant_id.arn,
        ]
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_custom" {
  role       = aws_iam_role.ecs_task_execution_role.id
  policy_arn = aws_iam_policy.ecs_task_execution_policy.arn
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution" {
  role       = aws_iam_role.ecs_task_execution_role.id
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role" "executor_task_role" {
  name = "${var.prefix}-executor-task-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_policy" "executor_task_policy" {
  name        = "${var.prefix}-executor-task-policy"
  path        = "/"
  description = "TFOM executor task policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "dynamodb:*",
          "s3:*",
          "sqs:*"
        ]
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Action   = "sts:AssumeRole"
        Effect   = "Allow"
        Resource = "arn:aws:iam::*:role/gitlab-automation-role"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "executor_task" {
  role       = aws_iam_role.executor_task_role.id
  policy_arn = aws_iam_policy.executor_task_policy.arn
}


resource "aws_iam_role" "workflow_handler_lambda_role" {
  name = "${var.prefix}-workflow-handler-lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

data "aws_iam_policy_document" "workflow_handler_policy" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:UpdateItem",
      "dynamodb:DeleteItem",
      "dynamodb:GetItem",
      "dynamodb:Query",
      "dynamodb:Scan",
    ]
    resources = [
      "*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "states:StartExecution",
    ]
    resources = [
      "*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
      "s3:ListBucket",
    ]
    resources = [
      "*"
    ]
  }
}

resource "aws_iam_policy" "workflow_handler_policy" {
  name        = "${var.prefix}-workflow-handler-policy"
  path        = "/"
  description = "Allow Lambda to interact with DynamoDB"

  policy = data.aws_iam_policy_document.workflow_handler_policy.json
}

resource "aws_iam_role_policy_attachment" "workflow_handler_policy_attachment" {
  role       = aws_iam_role.workflow_handler_lambda_role.name
  policy_arn = aws_iam_policy.workflow_handler_policy.arn
}

resource "aws_iam_role_policy_attachment" "workflow_handler_lambda_role_policy_attachment" {
  role       = aws_iam_role.workflow_handler_lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}




resource "aws_iam_role" "step_functions_role" {
  name = "${var.prefix}-step-functions-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "states.amazonaws.com"
        }
      },
    ]
  })
}

data "aws_iam_policy_document" "step_functions_policy" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogDelivery",
      "logs:DeleteLogDelivery",
      "logs:DescribeLogGroups",
      "logs:DescribeResourcePolicies",
      "logs:GetLogDelivery",
      "logs:ListLogDeliveries",
      "logs:PutResourcePolicy",
      "logs:UpdateLogDelivery",
      "logs:DescribeLogGroups",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "xray:GetSamplingRules",
      "xray:GetSamplingTargets",
      "xray:PutTraceSegments",
      "xray:PutTelemetryRecords",
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "states:SendTaskFailure",
      "states:SendTaskSuccess",
      "states:SendTaskHeartbeat",
      "states:StartExecution",
      "states:StopExecution",
      "states:DescribeExecution",
      "states:DescribeStateMachine",
      "events:PutTargets",
      "events:PutRule",
      "events:DescribeRule",
      "sns:Publish"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:UpdateItem"
    ]
    resources = [
      aws_dynamodb_table.module_propagation_execution_requests.arn,
      aws_dynamodb_table.module_propagation_drift_check_requests.arn,
      aws_dynamodb_table.plan_execution_requests.arn,
      aws_dynamodb_table.apply_execution_requests.arn,
      aws_dynamodb_table.terraform_execution_workflow_requests.arn,
      aws_dynamodb_table.terraform_drift_check_workflow_requests.arn
    ]
  }

  statement {
    effect    = "Allow"
    actions   = ["iam:PassRole"]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "ecs:RunTask",
      "ecs:StopTask",
      "ecs:DescribeTasks"
    ]
    resources = [
      aws_ecs_task_definition.executor.arn,
      aws_ecs_cluster.tfom.arn
    ]
  }

  statement {
    effect  = "Allow"
    actions = ["lambda:InvokeFunction"]
    resources = [
      "*"
    ]
  }
}

resource "aws_iam_policy" "step_functions_policy" {
  name        = "${var.prefix}-step-functions-policy"
  path        = "/"
  description = "Allow Step Functions to execute ECS, Update DynamoDB Tables"

  policy = data.aws_iam_policy_document.step_functions_policy.json
}

resource "aws_iam_role_policy_attachment" "step_functions_policy_attachment" {
  role       = aws_iam_role.step_functions_role.name
  policy_arn = aws_iam_policy.step_functions_policy.arn
}
