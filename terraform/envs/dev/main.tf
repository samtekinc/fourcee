provider "aws" {
  region = "us-east-1"
}

module "organizations_service" {
  source = "../../modules/organizations-service"
  prefix = "tfom-org-service"
}

module "execution_service" {
  source = "../../modules/execution-service"
  prefix = "tfom-exec-service"

  vpc_id          = var.vpc_id
  efs_subnet_ids  = var.efs_subnet_ids
  task_subnet_ids = var.task_subnet_ids
}
