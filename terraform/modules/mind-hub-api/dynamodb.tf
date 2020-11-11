resource "aws_dynamodb_table" "mind_hub_api_course_progress" {
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

resource "aws_dynamodb_table" "mind_hub_api_course_note" {
  name              = "course_note"
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

resource "aws_dynamodb_table" "mind_hub_api_step_progress" {
  name              = "step_progress"
  read_capacity     = 5
  write_capacity    = 5
  hash_key          = "stepID"
  range_key         = "userID"
  stream_enabled    = false

  attribute {
    name = "stepID"
    type = "S"
  }

  attribute {
    name = "userID"
    type = "S"
  }
}

resource "aws_dynamodb_table" "mind_hub_api_step_note" {
  name              = "step_note"
  read_capacity     = 5
  write_capacity    = 5
  hash_key          = "stepID"
  range_key         = "userID"
  stream_enabled    = false

  attribute {
    name = "stepID"
    type = "S"
  }

  attribute {
    name = "userID"
    type = "S"
  }
}