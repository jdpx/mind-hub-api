resource "aws_dynamodb_table" "mind_hub_api_user_table" {
  name              = "user"
  read_capacity     = 5
  write_capacity    = 5
  hash_key          = "PK"
  range_key         = "SK"
  stream_enabled    = false

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "SK"
    type = "S"
  }
}