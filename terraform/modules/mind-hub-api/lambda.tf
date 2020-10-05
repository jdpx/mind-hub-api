
resource "aws_lambda_function" "mind_hub_api_graphql_api_lambda" {
  function_name = "mind_hub_api_graphql_api_${var.env}"
  description   = "Lambda that contains the GraphQL API"

  s3_bucket = aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.bucket
  s3_key    = "function.zip"
  role      = aws_iam_role.mind_hub_api_graphql_api_role.arn

  handler = "main"
  runtime = "go1.x"
  timeout = 30

  environment {
    variables = {
      GRAPH_CMS_URL = "https://api-eu-central-1.graphcms.com/v2/ckftjhf769ysi01z7ari84qio/master"
    }
  }
}
