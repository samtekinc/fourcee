resource "aws_s3_bucket" "execution_service" {
  bucket = "${var.prefix}-execution-results"
  acl    = "private"

  tags = {
    Name = "${var.prefix}-execution-results"
  }
}
