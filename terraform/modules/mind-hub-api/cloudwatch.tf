resource "aws_cloudwatch_log_group" "mind_hub_api_api_logs" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.mind_hub_api.id}/${var.api_stage_name}"
  retention_in_days = 180
}

resource "aws_cloudwatch_log_group" "mind_hub_api_lambda_logs" {
  name              = "/aws/lambda/mind_hub_api_graphql_api_${var.env}"
  retention_in_days = 120
}

resource "aws_cloudwatch_log_group" "mind_hub_cloudfront_invalidation_logs" {
  name              = "/aws/codebuild/mind-hub-cloudfront-invalidation"
  retention_in_days = 120
}

resource "aws_cloudwatch_log_group" "mind_hub_api_build_logs" {
  name              = "/aws/codebuild/mind-hub-api-build"
  retention_in_days = 120
}

resource "aws_cloudwatch_log_group" "mind_hub_ui_build_logs" {
  name              = "/aws/codebuild/mind-hub-ui-build"
  retention_in_days = 120
}

resource "aws_cloudwatch_log_group" "mind_hub_api_terraform_deploy_logs" {
  name              = "/aws/codebuild/mind_hub_api_terraform_deploy"
  retention_in_days = 120
}