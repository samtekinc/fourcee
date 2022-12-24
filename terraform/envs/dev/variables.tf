variable "vpc_id" {
  type = string
}

variable "efs_subnet_ids" {
  type = list(string)
}

variable "task_subnet_ids" {
  type = list(string)
}
