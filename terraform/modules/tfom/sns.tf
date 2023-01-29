resource "aws_sns_topic" "tfom_alerts" {
  name = "${var.prefix}-alerts"
}
