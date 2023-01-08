resource "aws_ssm_parameter" "arm_client_id" {
  name        = "/${var.prefix}/arm_client_id"
  description = "Azure Service Principal Client ID"
  type        = "SecureString"
  value       = var.arm_client_id
}

resource "aws_ssm_parameter" "arm_client_secret" {
  name        = "/${var.prefix}/arm_client_secret"
  description = "Azure Service Principal Secret Value"
  type        = "SecureString"
  value       = var.arm_client_secret
}


resource "aws_ssm_parameter" "arm_tenant_id" {
  name        = "/${var.prefix}/arm_tenant_id"
  description = "Azure Service Principal Tenant ID"
  type        = "SecureString"
  value       = var.arm_tenant_id
}
