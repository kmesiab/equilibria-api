variable "region" {
  type = string
  default = "us-west-2"
}

variable "route_53_hosted_zone_id" {
  type = string
  default = "Z1017894H0HIZEDKUB24"
}

variable "openai_api_key" {
  description = "The OpenAI API key"
  type        = string
}

variable "database_host" {
  description = "The hostname of the database server"
  type        = string
}

variable "database_user" {
  description = "The username for the database"
  type        = string
}

variable "database_password" {
  description = "The password for the database"
  type        = string
  sensitive   = true
}

variable "database_name" {
  description = "The name of the database"
  type        = string
}

variable "log_level" {
  description = "The log level for the application"
  type        = string
}
