resource "aws_codepipeline_webhook" "mind_hub_api_codepipeline_webhook" {
  authentication  = "GITHUB_HMAC"
  name            = "mind-hub-api-codepipeline-webhook"
  target_action   = "Source"
  target_pipeline = aws_codepipeline.mind_hub_api_pipeline.name

  authentication_configuration {
    secret_token = random_string.github_secret.result
  }

  filter {
    json_path    = "$.ref"
    match_equals = "refs/heads/{Branch}"
  }
  tags = {}
}

resource "github_repository_webhook" "mind_hub_api_github_webhook" {
  repository = var.repository_name
  events     = ["push"]

  configuration {
    url          = aws_codepipeline_webhook.mind_hub_api_codepipeline_webhook.url
    insecure_ssl = "0"
    content_type = "json"
    secret       = random_string.github_secret.result
  }
}

resource "random_string" "github_secret" {
  length  = 99
  special = false
}