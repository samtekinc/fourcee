provider "aws" {
  region = "us-east-1"
}

module "organizations_service" {
  source = "../../modules/organizations-service"
  prefix = "tfom"

  vpc_id          = var.vpc_id
  efs_subnet_ids  = var.efs_subnet_ids
  task_subnet_ids = var.task_subnet_ids

  arm_client_id     = var.arm_client_id
  arm_client_secret = var.arm_client_secret
  arm_tenant_id     = var.arm_tenant_id
}
