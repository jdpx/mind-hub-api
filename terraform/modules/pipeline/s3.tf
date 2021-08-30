resource "aws_s3_bucket" "mind_hub_api_pipeline_artifact_bucket" {
  bucket = "mind-hub-api-pipeline-artifacts"
  acl    = "private"

  versioning {
    enabled = true
  }

  lifecycle_rule {
    id      = "quarterly_retention"
    prefix  = "/"
    enabled = true

    expiration {
      days = 100
    }
  }
}

resource "aws_s3_bucket_public_access_block" "mind_hub_api_pipeline_artifact_bucket_block" {
  bucket                  = aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  restrict_public_buckets = true
  ignore_public_acls      = true
}
