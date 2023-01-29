provider "aws" {
  region = "us-east-1"
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

module "organizations_service" {
  source = "../../modules/tfom"
  prefix = "tfom"
}
