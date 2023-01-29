resource "aws_s3_bucket" "backends" {
  bucket = "${var.prefix}-backend-states"
  acl    = "private"

  tags = {
    Name = "${var.prefix}-backend-states"
  }
}
