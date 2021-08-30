data "aws_s3_bucket" "mind_hub_api_pipeline_artifact_bucket" {
  bucket = "mind-hub-api-pipeline-artifacts"
}

data "aws_s3_bucket_object" "mind_hub_api_pipeline_artifact_bucket_object" {
  bucket = "mind-hub-api-pipeline-artifacts"
  key    = "graphql.zip"
}

resource "aws_s3_bucket_public_access_block" "mind_hub_api_pipeline_artifact_bucket_block" {
  bucket                  = aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  restrict_public_buckets = true
  ignore_public_acls      = true
}

resource "aws_s3_bucket_public_access_block" "mind_hub_api_pipeline_artifact_bucket_object_block" {
  bucket                  = aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket_object.id
  block_public_acls       = true
  block_public_policy     = true
  restrict_public_buckets = true
  ignore_public_acls      = true
}
