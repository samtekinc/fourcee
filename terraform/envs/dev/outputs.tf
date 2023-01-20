output "TFOM_PREFIX" {
  value = "tfom"
}

output "TFOM_STATE_BUCKET" {
  value = module.organizations_service.state_bucket
}

output "TFOM_STATE_REGION" {
  value = module.organizations_service.state_region
}

output "TFOM_RESULTS_BUCKET" {
  value = module.organizations_service.results_bucket
}

output "TFOM_ALERTS_TOPIC" {
  value = module.organizations_service.alerts_topic
}

output "TFOM_ACCOUNT_ID" {
  value = data.aws_caller_identity.current.account_id
}

output "TFOM_REGION" {
  value = data.aws_region.current.name
}
