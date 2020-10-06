# This is a _data source_ which allows us to get the internal
# ID (which AWS calls an "ARN") from AWS
data "aws_acm_certificate" "api_mind_jdpx_co_uk_cert" {
  provider = aws.us_east
  domain   = "api.${var.env}.mind.jdpx.co.uk"
  statuses = ["ISSUED"]
}

resource "aws_api_gateway_domain_name" "mind_hub_api_domain" {
  domain_name     = "api.${var.env}.mind.jdpx.co.uk"
  certificate_arn = data.aws_acm_certificate.api_mind_jdpx_co_uk_cert.arn
}

module "cors" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.1"

  api_id            = aws_api_gateway_rest_api.mind_hub_api.id
  api_resource_id   = aws_api_gateway_resource.mind_hub_api_resource.id
  allow_credentials = true
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
  api_id = "${aws_api_gateway_rest_api.mind_hub_api.id}"

  domain_name = "${aws_api_gateway_domain_name.mind_hub_api_domain.domain_name}"
}


resource "aws_api_gateway_resource" "mind_hub_api_resource" {
  path_part   = "query"
  parent_id   = "${aws_api_gateway_rest_api.mind_hub_api.root_resource_id}"
  rest_api_id = "${aws_api_gateway_rest_api.mind_hub_api.id}"
}

resource "aws_api_gateway_method" "mind_hub_api_resource_post_method" {
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.authorizer.id

  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  resource_id = aws_api_gateway_resource.mind_hub_api_resource.id
  http_method = "POST"
}

resource "aws_api_gateway_authorizer" "authorizer" {
  authorizer_credentials = module.authorizer.lambda_role_arn
  authorizer_uri         = "arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/${module.authorizer.lambda_arn}/invocations"
  name                   = "auth0"
  rest_api_id            = aws_api_gateway_rest_api.mind_hub_api.id
}

resource "aws_api_gateway_method_response" "mind_hub_api_resource_post_method_response_200" {
  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  resource_id = aws_api_gateway_resource.mind_hub_api_resource.id
  http_method = aws_api_gateway_method.mind_hub_api_resource_post_method.http_method
  status_code = 200

  /**
   * This is where the configuration for CORS enabling starts.
   * We need to enable those response parameters and in the 
   * integration response we will map those to actual values
   */
  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers"     = true,
    "method.response.header.Access-Control-Allow-Methods"     = true,
    "method.response.header.Access-Control-Allow-Origin"      = true,
    "method.response.header.Access-Control-Allow-Credentials" = true
  }
}

resource "aws_api_gateway_integration" "mind_hub_api_integration" {
  rest_api_id             = aws_api_gateway_rest_api.mind_hub_api.id
  resource_id             = aws_api_gateway_resource.mind_hub_api_resource.id
  http_method             = aws_api_gateway_method.mind_hub_api_resource_post_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.mind_hub_api_graphql_api_lambda.invoke_arn
}

resource "aws_api_gateway_integration_response" "mind_hub_api_resource_post_method_integration_response" {
  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  resource_id = aws_api_gateway_resource.mind_hub_api_resource.id
  http_method = aws_api_gateway_method.mind_hub_api_resource_post_method.http_method
  status_code = aws_api_gateway_method_response.mind_hub_api_resource_post_method_response_200.status_code

  depends_on = [aws_api_gateway_integration.mind_hub_api_integration]

  /**
   * This is second half of the CORS configuration.
   * Here we give values to each of the header parameters to ALLOW 
   * Cross-Origin requests from ALL hosts.
   **/
  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers"     = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods"     = "'GET,OPTIONS,POST,PUT'",
    "method.response.header.Access-Control-Allow-Origin"      = "'*'",
    "method.response.header.Access-Control-Allow-Credentials" = "'true'"
  }
}

resource "aws_lambda_permission" "mind_hub_api_graphql_api_lambda_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.mind_hub_api_graphql_api_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mind_hub_api.execution_arn}/*/${aws_api_gateway_method.mind_hub_api_resource_post_method.http_method}${aws_api_gateway_resource.mind_hub_api_resource.path}"
}

resource "aws_api_gateway_deployment" "mind_hub_api_deploy" {
  depends_on = [aws_api_gateway_integration.mind_hub_api_integration]

  rest_api_id = aws_api_gateway_rest_api.mind_hub_api.id
  stage_name  = "v1"
}
