resource "aws_s3_bucket" "mind_hub_api_pipeline_artifact_bucket" {
  bucket = "mind-hub-api-artifacts"
  acl    = "private"
}

data "aws_s3_bucket" "dev_tf_state_bucket" {
  bucket = "mind-hub-api-dev-tf-state"
}

data "aws_s3_bucket" "management_tf_state_bucket" {
  bucket = "mind-hub-api-management-tf-state"
}