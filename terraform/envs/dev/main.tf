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
}
