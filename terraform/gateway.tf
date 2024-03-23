resource "aws_internet_gateway" "eq_api_igw" {
  vpc_id = aws_vpc.eq_api_vpc.id

  tags = {
    Name = "eq-api-igw"
  }
}

resource "aws_lb_target_group" "eq_api_target_group" {
  name     = "eq-api-tg"
  port     = 443
  protocol = "HTTP"
  vpc_id   = aws_vpc.eq_api_vpc.id
  target_type = "ip"
}

resource "aws_lb" "eq_api_lb" {
  name               = "eq-api-lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.eq_api_sg.id]
  subnets            = [aws_subnet.eq_api_subnet_1.id, aws_subnet.eq_api_subnet_2.id]
}

resource "aws_lb_listener" "api_my_eq_com_https_listener" {
  load_balancer_arn = aws_lb.eq_api_lb.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = aws_acm_certificate.api_my_eq_com_cert.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.eq_api_target_group.arn
  }
}


