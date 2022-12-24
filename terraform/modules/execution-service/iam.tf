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
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:UpdateItem"
    ]
    resources = [
      aws_dynamodb_table.plan_execution_requests.arn,
      aws_dynamodb_table.apply_execution_requests.arn,
    ]
  }

  statement {
    effect    = "Allow"
    actions   = ["iam:PassRole"]
    resources = [aws_iam_role.ecs_task_execution_role.arn, aws_iam_role.executor_task_role.arn]
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
