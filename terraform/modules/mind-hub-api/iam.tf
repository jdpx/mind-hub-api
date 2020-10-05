
resource "aws_iam_role" "mind_hub_api_graphql_api_role" {
  name = "mind_hub_api_graphql_${var.env}_api_role"

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


resource "aws_iam_role_policy" "mind_hub_api_graphql_api_role_policy" {
  name = "mind_hub_api_graphql_api_${var.env}_role_policy"
  role = aws_iam_role.mind_hub_api_graphql_api_role.id

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
