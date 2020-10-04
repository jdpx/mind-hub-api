resource "aws_s3_bucket" "mind_hub_api_pipeline_artifact_bucket" {
  bucket = "mind-hub-api-artifacts"
  acl    = "private"
}
