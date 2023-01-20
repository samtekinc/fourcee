output "state_bucket" {
  value = aws_s3_bucket.backends.bucket
}

output "state_region" {
  value = aws_s3_bucket.backends.region
}

output "results_bucket" {
  value = aws_s3_bucket.execution_service.bucket
}

output "alerts_topic" {
  value = aws_sns_topic.tfom_alerts.arn
}
