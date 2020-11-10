# This is a _data source_ which allows us to get the internal
# ID (which AWS calls an "ARN") from AWS
data "aws_acm_certificate" "api_mind_jdpx_co_uk_cert" {
  provider = aws.us_east
  domain   = "api.${var.env}.mind.jdpx.co.uk"
  statuses = ["ISSUED"]
}

resource "aws_api_gateway_account" "mind_hub_api" {
  cloudwatch_role_arn = aws_iam_role.mind_hub_api_cloudwatch.arn
}

resource "aws_api_gateway_domain_name" "mind_hub_api_domain" {
  domain_name     = "api.${var.env}.mind.jdpx.co.uk"
  certificate_arn = data.aws_acm_certificate.api_mind_jdpx_co_uk_cert.arn
  security_policy = "TLS_1_2"
}

module "cors" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.1"

  api_id            = aws_api_gateway_rest_api.mind_hub_api.id
  api_resource_id   = aws_api_gateway_resource.mind_hub_proxy.id
  allow_credentials = true
  allow_headers = [
    "Authorization",
    "Content-Type",
    "X-Amz-Date",
    "X-Amz-Security-Token",
    "X-Api-Key",
    "x-correlation-id",
  ]
}

module authorizer {
  source                  = "amancevice/auth0-authorizer/aws"
  authorizer_audience     = var.auth0_audience
  authorizer_jwks_uri     = var.auth0_jwks_uri
  authorizer_token_issuer = var.auth0_token_issuer
}

resource "aws_api_gateway_rest_api" "mind_hub_api" {
  name = "mind_hub_api_${var.env}"
}

resource "aws_api_gateway_base_path_mapping" "base_path_mapping" {
  api_id = aws_api_gateway_rest_api.mind_hub_api.id

  domain_name = aws_api_gateway_domain_name.mind_hub_api_domain.domain_name
}

resource "aws_api_gateway_resource" "mind_hub_proxy" {
  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  parent_id   = aws_api_gateway_rest_api.mind_hub_api.root_resource_id # aws_api_gateway_resource.version.id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxy_method" {
  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  resource_id = aws_api_gateway_resource.mind_hub_proxy.id
  http_method = "ANY"


  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.authorizer.id

  request_parameters = {
    "method.request.path.proxy" = true
  }
}

resource "aws_api_gateway_method_settings" "proxy_method_settings" {
  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  stage_name  = "v1"
  method_path = "*/*"

  settings {
    metrics_enabled = true
    logging_level   = "INFO"
  }
}

resource "aws_api_gateway_integration" "mind_hub_lambda_integration" {
  rest_api_id             = aws_api_gateway_rest_api.mind_hub_api.id
  resource_id             = aws_api_gateway_resource.mind_hub_proxy.id
  http_method             = aws_api_gateway_method.proxy_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.mind_hub_api_graphql_api_lambda.invoke_arn
}

resource "aws_lambda_permission" "mind_hub_api_graphql_api_lambda_permission_proxy" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.mind_hub_api_graphql_api_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mind_hub_api.execution_arn}/*/${aws_api_gateway_method.proxy_method.http_method}${aws_api_gateway_resource.mind_hub_proxy.path}"
}

resource "aws_api_gateway_authorizer" "authorizer" {
  authorizer_credentials = module.authorizer.lambda_role_arn
  authorizer_uri         = "arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/${module.authorizer.lambda_arn}/invocations"
  name                   = "auth0"
  rest_api_id            = aws_api_gateway_rest_api.mind_hub_api.id
}

resource "aws_api_gateway_deployment" "mind_hub_api_deploy" {
  depends_on = [
    aws_api_gateway_resource.mind_hub_proxy,
    aws_api_gateway_method.proxy_method,
    aws_api_gateway_integration.mind_hub_lambda_integration
  ]

  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  stage_name  = "v1"

  lifecycle {
    create_before_destroy = true
  }
}
