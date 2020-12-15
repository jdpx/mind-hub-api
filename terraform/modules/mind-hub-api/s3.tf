# resource "aws_s3_bucket" "mind_hub_api_pipeline_artifact_bucket" {
#   bucket = "mind-hub-api-artifacts-${var.env}"
#   acl    = "private"

#   lifecycle_rule {
#         id      = "quarterly_retention"
#         prefix  = "/"
#         enabled = true

#         expiration {
#             days = 7
#         }
#     }
# }

data "aws_s3_bucket" "mind_hub_api_pipeline_artifact_bucket" {
  bucket = "mind-hub-api-pipeline-artifacts"
}