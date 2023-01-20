resource "aws_lambda_function" "workflow_handler" {
  depends_on    = [aws_cloudwatch_log_group.workflow_handler_lambda]
  filename      = data.archive_file.empty.output_path
  function_name = "${var.prefix}-workflow-handler"
  description   = "Async handler Lambda for handling backend workflows"
  role          = aws_iam_role.workflow_handler_lambda_role.arn
  architectures = ["x86_64"]
  memory_size   = 4096
  handler       = "workflow-handler"
  runtime       = "go1.x"
  package_type  = "Zip"
  timeout       = 900

  environment {
    variables = {
      TFOM_WORKING_DIRECTORY = "./tmp/"
      TFOM_PREFIX            = var.prefix
      TFOM_STATE_BUCKET      = aws_s3_bucket.backends.bucket
      TFOM_STATE_REGION      = aws_s3_bucket.backends.region
      TFOM_RESULTS_BUCKET    = aws_s3_bucket.execution_service.bucket
      TFOM_ACCOUNT_ID        = data.aws_caller_identity.current.account_id
      TFOM_REGION            = data.aws_region.current.name
      TFOM_ALERTS_TOPIC      = aws_sns_topic.tfom_alerts.arn
    }
  }

  source_code_hash = data.archive_file.empty.output_base64sha256

  lifecycle {
    ignore_changes = [last_modified, filename, source_code_hash]
  }
}

resource "aws_cloudwatch_log_group" "workflow_handler_lambda" {
  name              = "/aws/lambda/${var.prefix}-workflow-handler"
  retention_in_days = 731
}

data "archive_file" "empty" {
  type        = "zip"
  output_path = "${path.module}/empty.zip"
  source {
    content  = "example"
    filename = "empty.txt"
  }
}
