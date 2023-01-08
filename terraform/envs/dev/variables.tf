variable "vpc_id" {
  type = string
}

variable "efs_subnet_ids" {
  type = list(string)
}

variable "task_subnet_ids" {
  type = list(string)
}

variable "arm_client_id" {
  type = string
}

variable "arm_client_secret" {
  type = string
}

variable "arm_tenant_id" {
  type = string
}
