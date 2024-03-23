resource "aws_route53_record" "api_my_eq_com" {
  zone_id = var.route_53_hosted_zone_id
  name    = "api.my-eq.com"
  type    = "A"

  alias {
    name                   = aws_lb.eq_api_lb.dns_name
    zone_id                = aws_lb.eq_api_lb.zone_id
    evaluate_target_health = true
  }
}

resource "aws_acm_certificate" "api_my_eq_com_cert" {
  domain_name       = "api.my-eq.com"
  validation_method = "DNS"
}

resource "aws_route53_record" "validation" {
  for_each = {
    for dvo in aws_acm_certificate.api_my_eq_com_cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }

  zone_id = var.route_53_hosted_zone_id
  name    = each.value.name
  type    = each.value.type
  records = [each.value.record]
  ttl     = 60
}

resource "aws_acm_certificate_validation" "api_my_eq_com_cert_validation" {
  certificate_arn         = aws_acm_certificate.api_my_eq_com_cert.arn
  validation_record_fqdns = [for dvo in aws_acm_certificate.api_my_eq_com_cert.domain_validation_options : dvo.resource_record_name]
}
