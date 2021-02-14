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
