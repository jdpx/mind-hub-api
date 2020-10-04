# This is a _data source_ which allows us to get the internal
# ID (which AWS calls an "ARN") from AWS
data "aws_acm_certificate" "api_dev_mind_jdpx_co_uk" {
  provider = aws.us_east
  domain   = "api.dev.mind.jdpx.co.uk"
  statuses = ["ISSUED"]
}


resource "aws_api_gateway_rest_api" "api" {
  name = "query_api"
}

resource "aws_api_gateway_resource" "resource" {
  path_part   = "query"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
  resource_id   = "${aws_api_gateway_resource.resource.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "api_response_200" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_resource.resource.id
  http_method = aws_api_gateway_method.method.http_method
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

resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = "${aws_api_gateway_rest_api.api.id}"
  resource_id             = "${aws_api_gateway_resource.resource.id}"
  http_method             = "${aws_api_gateway_method.method.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "${aws_lambda_function.time.invoke_arn}"
}

resource "aws_api_gateway_integration_response" "mockapi_integration_response" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_resource.resource.id
  http_method = aws_api_gateway_method.method.http_method
  status_code = aws_api_gateway_method_response.api_response_200.status_code

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

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.time.function_name}"
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
}

resource "aws_api_gateway_deployment" "time_deploy" {
  depends_on = [aws_api_gateway_integration.integration]

  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  stage_name  = "v1"
}

output "url" {
  value = "${aws_api_gateway_deployment.time_deploy.invoke_url}${aws_api_gateway_resource.resource.path}"
}

resource "aws_api_gateway_domain_name" "domain" {
  domain_name     = "api.dev.mind.jdpx.co.uk"
  certificate_arn = data.aws_acm_certificate.api_dev_mind_jdpx_co_uk.arn
}

resource "aws_api_gateway_base_path_mapping" "base_path_mapping" {
  api_id = "${aws_api_gateway_rest_api.api.id}"

  domain_name = "${aws_api_gateway_domain_name.domain.domain_name}"
}

module "cors" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.1"

  api_id            = aws_api_gateway_rest_api.api.id
  api_resource_id   = aws_api_gateway_resource.resource.id
  allow_credentials = true
}
