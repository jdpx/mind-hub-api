resource "aws_acm_certificate" "api_mind_jdpx_co_uk_cert" {
  provider = aws.us_east
  
  domain_name   = "api.${var.env}.mind.jdpx.co.uk"
  validation_method = "DNS"

  tags = {
    Name        = "api.${var.env}.mind.jdpx.co.uk"
    Environment = var.env
  }

  lifecycle {
    create_before_destroy = true
  }
}