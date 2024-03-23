output "load_balancer_url" {
  value = aws_lb.eq_api_lb.dns_name
}

output "api_url" {
  value = "https://api.my-eq.com"
}
