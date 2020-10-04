


resource "aws_lambda_function" "time" {
  s3_bucket     = "mind-hub-api-artifacts"
  s3_key        = "function.zip"
  function_name = "test_function"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.iam_for_lambda.arn

  environment {
    variables = {
      GRAPH_CMS_URL = "https://api-eu-central-1.graphcms.com/v2/ckftjhf769ysi01z7ari84qio/master"
    }
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}


resource "aws_iam_role_policy" "iam_for_lambda_policy" {
  name = "iam_for_lambda_policy"
  role = aws_iam_role.iam_for_lambda.id

  policy = <<EOF
{
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ],
  "Version": "2012-10-17"
}
EOF
}
