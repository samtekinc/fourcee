resource "aws_efs_file_system" "executor" {
  creation_token = "${var.prefix}-executor"

  tags = {
    Name = "${var.prefix}-file-system"
  }
}

resource "aws_efs_mount_target" "mount" {
  for_each        = toset(var.efs_subnet_ids)
  file_system_id  = aws_efs_file_system.executor.id
  subnet_id       = each.value
  security_groups = [aws_security_group.executor_efs.id]
}
