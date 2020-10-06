data "aws_dynamodb_table"  "tf_lock_state" {
  name = "mind-hub-ui-state-lock"
}