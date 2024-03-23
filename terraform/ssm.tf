resource "aws_ssm_parameter" "openai_api_key" {
  name  = "/eq-api/openai_api_key"
  type  = "SecureString"
  value = var.openai_api_key
}

resource "aws_ssm_parameter" "database_host" {
  name  = "/eq-api/database_host"
  type  = "String"
  value = var.database_host
}

resource "aws_ssm_parameter" "database_user" {
  name  = "/eq-api/database_user"
  type  = "String"
  value = var.database_user
}

resource "aws_ssm_parameter" "database_password" {
  name  = "/eq-api/database_password"
  type  = "SecureString"
  value = var.database_password
}

resource "aws_ssm_parameter" "database_name" {
  name  = "/eq-api/database_name"
  type  = "String"
  value = var.database_name
}

resource "aws_ssm_parameter" "log_level" {
  name  = "/eq-api/log_level"
  type  = "String"
  value = var.log_level
}
