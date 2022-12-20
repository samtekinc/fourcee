resource "aws_dynamodb_table" "org_dimensions" {
  name         = "${var.prefix}-organizational-dimensions"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "DimensionId"

  attribute {
    name = "DimensionId"
    type = "S"
  }
}

resource "aws_dynamodb_table" "org_units" {
  name         = "${var.prefix}-organizational-units"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "OrgUnitId"

  attribute {
    name = "OrgUnitId"
    type = "S"
  }

  attribute {
    name = "DimensionId"
    type = "S"
  }

  attribute {
    name = "ParentOrgUnitId"
    type = "S"
  }

  attribute {
    name = "Hierarchy"
    type = "S"
  }

  global_secondary_index {
    name            = "DimensionId-Hierarchy-index"
    hash_key        = "DimensionId"
    range_key       = "Hierarchy"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "DimensionId-ParentOrgUnitId-index"
    hash_key        = "DimensionId"
    range_key       = "ParentOrgUnitId"
    projection_type = "ALL"
  }
}

resource "aws_dynamodb_table" "org_accounts" {
  name         = "${var.prefix}-organizational-accounts"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "OrgAccountId"

  attribute {
    name = "OrgAccountId"
    type = "S"
  }
}

resource "aws_dynamodb_table" "org_unit_memberships" {
  name         = "${var.prefix}-organizational-unit-memberships"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "DimensionId"
  range_key    = "OrgAccountId"

  attribute {
    name = "OrgAccountId"
    type = "S"
  }

  attribute {
    name = "DimensionId"
    type = "S"
  }

  attribute {
    name = "OrgUnitId"
    type = "S"
  }

  global_secondary_index {
    name            = "OrgUnitId-OrgAccountId-index"
    hash_key        = "OrgUnitId"
    range_key       = "OrgAccountId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "OrgAccountId-OrgUnitId-index"
    hash_key        = "OrgAccountId"
    range_key       = "OrgUnitId"
    projection_type = "ALL"
  }
}
