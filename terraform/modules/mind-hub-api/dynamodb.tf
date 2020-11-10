resource "aws_dynamodb_table" "mind_hub_api_db" {
  name              = "course_progress"
  read_capacity     = 5
  write_capacity    = 5
  hash_key          = "courseID"
  range_key         = "userID"
  stream_enabled    = false

  attribute {
    name = "courseID"
    type = "S"
  }

  attribute {
    name = "userID"
    type = "S"
  }
}