resource "aws_cloudwatch_log_group" "mind_hub_api_api_logs" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.mind_hub_api.id}/${var.api_stage_name}"
  retention_in_days = 180
}