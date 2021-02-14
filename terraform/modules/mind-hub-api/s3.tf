data "aws_s3_bucket" "mind_hub_api_pipeline_artifact_bucket" {
  bucket = "mind-hub-api-pipeline-artifacts"
}

data "aws_s3_bucket_object" "mind_hub_api_pipeline_artifact_bucket_object" {
  bucket = "mind-hub-api-pipeline-artifacts"
  key    = "graphql.zip"
}
