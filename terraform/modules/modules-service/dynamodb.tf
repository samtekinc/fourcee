resource "aws_dynamodb_table" "module_groups" {
  name         = "${var.prefix}-module-groups"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ModuleGroupId"

  attribute {
    name = "ModuleGroupId"
    type = "S"
  }
}

resource "aws_dynamodb_table" "module_versions" {
  name         = "${var.prefix}-module-versions"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ModuleGroupId"
  range_key    = "ModuleVersionId"

  attribute {
    name = "ModuleGroupId"
    type = "S"
  }

  attribute {
    name = "ModuleVersionId"
    type = "S"
  }
}
