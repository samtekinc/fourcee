resource "aws_s3_bucket" "backends" {
  bucket = "${var.prefix}-backend-states"

  tags = {
    Name = "${var.prefix}-backend-states"
  }
}

resource "aws_s3_bucket_acl" "backends_private" {
  bucket = aws_s3_bucket.backends.id
  acl    = "private"
}

resource "aws_s3_bucket_versioning" "versioning_backends" {
  bucket = aws_s3_bucket.backends.id
  versioning_configuration {
    status = "Enabled"
  }
}
