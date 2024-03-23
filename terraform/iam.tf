resource "aws_iam_policy" "ssm_parameter_access_policy" {
  name        = "SSMParameterAccessPolicy"
  description = "Allows access to parameters in eq-api namespace"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect   = "Allow"
      Action   = "ssm:GetParameters"
      Resource = "arn:aws:ssm:us-west-2:462498369025:parameter/eq-api/*"
    }]
  })
}

resource "aws_iam_role_policy_attachment" "attach_policy_to_role" {
  role       = "eq_api_ecs_task_execution_role"  # Update with your role name
  policy_arn = aws_iam_policy.ssm_parameter_access_policy.arn
}
