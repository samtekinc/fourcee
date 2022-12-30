resource "aws_s3_bucket" "execution_service" {
  bucket = "${var.prefix}-execution-results"
  acl    = "private"

  tags = {
    Name = "${var.prefix}-execution-results"
  }
}


resource "aws_s3_bucket" "backends" {
  bucket = "${var.prefix}-backend-states"
  acl    = "private"

  tags = {
    Name = "${var.prefix}-backend-states"
  }
}
