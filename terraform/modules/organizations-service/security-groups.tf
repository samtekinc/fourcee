resource "aws_security_group" "executor_ecs" {
  name        = "${var.prefix}-executor-ecs-sg"
  description = "${var.prefix}-executor-ecs-sg"
  vpc_id      = var.vpc_id

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_security_group" "executor_efs" {
  name        = "${var.prefix}-executor-efs-sg"
  description = "${var.prefix}-executor-efs-sg"
  vpc_id      = var.vpc_id

  ingress {
    description     = "NFS from ECS"
    from_port       = 2049
    to_port         = 2049
    protocol        = "tcp"
    security_groups = [aws_security_group.executor_ecs.id]
  }
}
