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
    name = "RequestTime"
    type = "S"
  }

  global_secondary_index {
    name            = "StateKey-RequestTime-index"
    hash_key        = "StateKey"
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
    name = "RequestTime"
    type = "S"
  }

  global_secondary_index {
    name            = "StateKey-RequestTime-index"
    hash_key        = "StateKey"
    range_key       = "RequestTime"
    projection_type = "ALL"
  }
}
