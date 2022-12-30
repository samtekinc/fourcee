resource "aws_dynamodb_table" "org_dimensions" {
  name         = "${var.prefix}-organizational-dimensions"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "OrgDimensionId"

  attribute {
    name = "OrgDimensionId"
    type = "S"
  }
}

resource "aws_dynamodb_table" "org_units" {
  name         = "${var.prefix}-organizational-units"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "OrgDimensionId"
  range_key    = "OrgUnitId"

  attribute {
    name = "OrgUnitId"
    type = "S"
  }

  attribute {
    name = "OrgDimensionId"
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
    name            = "OrgDimensionId-Hierarchy-index"
    hash_key        = "OrgDimensionId"
    range_key       = "Hierarchy"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "OrgDimensionId-ParentOrgUnitId-index"
    hash_key        = "OrgDimensionId"
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
  hash_key     = "OrgDimensionId"
  range_key    = "OrgAccountId"

  attribute {
    name = "OrgAccountId"
    type = "S"
  }

  attribute {
    name = "OrgDimensionId"
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

resource "aws_dynamodb_table" "module_propagations" {
  name         = "${var.prefix}-module-propagations"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ModulePropagationId"

  attribute {
    name = "ModulePropagationId"
    type = "S"
  }

  attribute {
    name = "ModuleGroupId"
    type = "S"
  }

  attribute {
    name = "ModuleVersionId"
    type = "S"
  }

  attribute {
    name = "OrgUnitId"
    type = "S"
  }

  attribute {
    name = "OrgDimensionId"
    type = "S"
  }

  global_secondary_index {
    name            = "ModuleGroupId-ModuleVersionId-index"
    hash_key        = "ModuleGroupId"
    range_key       = "ModuleVersionId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "ModuleVersionId-index"
    hash_key        = "ModuleVersionId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "OrgUnitId-index"
    hash_key        = "OrgUnitId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "OrgDimensionId-OrgUnitId-index"
    hash_key        = "OrgDimensionId"
    range_key       = "OrgUnitId"
    projection_type = "ALL"
  }
}


resource "aws_dynamodb_table" "module_propagation_execution_requests" {
  name         = "${var.prefix}-module-propagation-execution-requests"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ModulePropagationId"
  range_key    = "ModulePropagationExecutionRequestId"

  attribute {
    name = "ModulePropagationId"
    type = "S"
  }

  attribute {
    name = "ModulePropagationExecutionRequestId"
    type = "S"
  }

  attribute {
    name = "RequestTime"
    type = "S"
  }

  global_secondary_index {
    name            = "ModulePropagationId-RequestTime-index"
    hash_key        = "ModulePropagationId"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }
}

resource "aws_dynamodb_table" "module_account_associations" {
  name         = "${var.prefix}-module-account-associations"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ModulePropagationId"
  range_key    = "OrgAccountId"

  attribute {
    name = "ModulePropagationId"
    type = "S"
  }

  attribute {
    name = "OrgAccountId"
    type = "S"
  }

  global_secondary_index {
    name            = "OrgAccountId-ModulePropagationId-index"
    hash_key        = "OrgAccountId"
    range_key       = "ModulePropagationId"
    projection_type = "ALL"
  }
}


resource "aws_dynamodb_table" "plan_execution_requests" {
  name         = "${var.prefix}-plan-execution-requests"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "PlanExecutionRequestId"

  attribute {
    name = "PlanExecutionRequestId"
    type = "S"
  }

  attribute {
    name = "StateKey"
    type = "S"
  }

  attribute {
    name = "ModulePropagationExecutionRequestId"
    type = "S"
  }

  attribute {
    name = "ModuleAccountAssociationKey"
    type = "S"
  }

  attribute {
    name = "RequestTime"
    type = "S"
  }

  global_secondary_index {
    name            = "StateKey-RequestTime-index"
    hash_key        = "StateKey"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "ModulePropagationExecutionRequestId-RequestTime-index"
    hash_key        = "ModulePropagationExecutionRequestId"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "ModuleAccountAssociationKey-RequestTime-index"
    hash_key        = "ModuleAccountAssociationKey"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }
}


resource "aws_dynamodb_table" "apply_execution_requests" {
  name         = "${var.prefix}-apply-execution-requests"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ApplyExecutionRequestId"

  attribute {
    name = "ApplyExecutionRequestId"
    type = "S"
  }

  attribute {
    name = "StateKey"
    type = "S"
  }

  attribute {
    name = "ModulePropagationExecutionRequestId"
    type = "S"
  }

  attribute {
    name = "ModuleAccountAssociationKey"
    type = "S"
  }

  attribute {
    name = "RequestTime"
    type = "S"
  }

  global_secondary_index {
    name            = "StateKey-RequestTime-index"
    hash_key        = "StateKey"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "ModulePropagationExecutionRequestId-RequestTime-index"
    hash_key        = "ModulePropagationExecutionRequestId"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "ModuleAccountAssociationKey-RequestTime-index"
    hash_key        = "ModuleAccountAssociationKey"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }
}
