resource "aws_ecs_cluster" "tfom" {
  name = "${var.prefix}-cluster"
}

resource "aws_ecs_cluster_capacity_providers" "tfom_fargate" {
  cluster_name = aws_ecs_cluster.tfom.name

  capacity_providers = ["FARGATE"]
}

resource "aws_ecr_repository" "executor" {
  name                 = "${var.prefix}-executor"
  image_tag_mutability = "MUTABLE"
}

resource "aws_cloudwatch_log_group" "tfom" {
  name = "/ecs/${var.prefix}-cluster"
}

resource "aws_ecs_task_definition" "executor" {
  family                   = "${var.prefix}-executor"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 1024
  memory                   = 2048
  container_definitions = jsonencode([
    {
      name      = "executor"
      image     = "${aws_ecr_repository.executor.repository_url}:latest"
      essential = true
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group" : aws_cloudwatch_log_group.tfom.name
          "awslogs-region" : "us-east-1",
          "awslogs-stream-prefix" : "executor"
        }
      }
      mountPoints = [
        {
          sourceVolume  = "service-storage"
          containerPath = "/efs"
        }
      ]
      environment = [
        {
          name  = "TF_PLUGIN_CACHE_DIR"
          value = "/efs/.terraform.d/plugin-cache"
        },
        {
          name  = "TFOM_WORKING_DIRECTORY"
          value = "./tmp/"
        },
        {
          name  = "TFOM_PREFIX"
          value = var.prefix
        },
        {
          name  = "TFOM_STATE_BUCKET"
          value = aws_s3_bucket.backends.bucket
        },
        {
          name  = "TFOM_STATE_REGION"
          value = aws_s3_bucket.backends.region
        },
        {
          name  = "TFOM_RESULTS_BUCKET"
          value = aws_s3_bucket.execution_service.bucket
        },
        {
          name  = "TFOM_ACCOUNT_ID"
          value = data.aws_caller_identity.current.account_id
        },
        {
          name  = "TFOM_REGION"
          value = data.aws_region.current.name
        },
        {
          name  = "TFOM_ALERTS_TOPIC"
          value = aws_sns_topic.tfom_alerts.arn
        }
      ]
      secrets = [
        {
          name      = "ARM_CLIENT_ID"
          valueFrom = aws_ssm_parameter.arm_client_id.arn
        },
        {
          name      = "ARM_CLIENT_SECRET"
          valueFrom = aws_ssm_parameter.arm_client_secret.arn
        },
        {
          name      = "ARM_TENANT_ID"
          valueFrom = aws_ssm_parameter.arm_tenant_id.arn
        }
      ]
    },
  ])

  volume {
    name = "service-storage"

    efs_volume_configuration {
      file_system_id = aws_efs_file_system.executor.id
    }
  }

  task_role_arn      = aws_iam_role.executor_task_role.arn
  execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
}
