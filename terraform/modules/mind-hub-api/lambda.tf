
resource "aws_lambda_function" "mind_hub_api_graphql_api_lambda" {
  function_name = "mind_hub_api_graphql_api_${var.env}"
  description   = "Lambda that contains the GraphQL API"

  s3_bucket         = data.aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.bucket
  s3_key            = "graphql.zip"
  s3_object_version = data.aws_s3_bucket_object.mind_hub_api_pipeline_artifact_bucket_object.version_id
  role              = aws_iam_role.mind_hub_api_graphql_api_role.arn

  handler = "lambda"
  runtime = "go1.x"
  timeout = 30

  environment {
    variables = {
      GRAPH_CMS_URL = var.graph_cms_url
    }
  }
}
