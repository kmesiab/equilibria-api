resource "aws_ecs_task_definition" "eq_api_task" {
  family                   = "eq-api-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.eq_api_ecs_task_execution_role.arn
  cpu                      = "256"
  memory                   = "512"

  # Set the runtime_platform to use x86_64 architecture
  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }

  container_definitions = jsonencode([
    {
      name      = "eq-api"
      image     = "462498369025.dkr.ecr.us-west-2.amazonaws.com/equilibria-api:latest"
      cpu       = 256
      memory    = 512
      essential = true
      secrets   = [
        {
          name = "DATABASE_HOST"
          valueFrom = aws_ssm_parameter.database_host.arn
        },
        {
          name = "DATABASE_USER"
          valueFrom = aws_ssm_parameter.database_user.arn
        },
        {
          name = "DATABASE_PASSWORD"
          valueFrom = aws_ssm_parameter.database_password.arn
        },
        {
          name = "DATABASE_NAME"
          valueFrom = aws_ssm_parameter.database_name.arn
        },
        {
          name = "LOG_LEVEL"
          valueFrom = aws_ssm_parameter.log_level.arn
        },
        {
          name = "OPENAI_API_KEY"
          valueFrom = aws_ssm_parameter.openai_api_key.arn
        }
      ]
      portMappings = [
        {
          containerPort = 443
          hostPort      = 443
          protocol      = "tcp"
        },
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.eq_api_log_group.name
          awslogs-region        = var.region
          awslogs-stream-prefix = "ecs"
          awslogs-create-group  = "true"
        }
      }
    }
  ])
}
