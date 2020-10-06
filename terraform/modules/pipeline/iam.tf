resource "aws_iam_role" "mind_hub_api_pipeline_role" {
  name = "mind_hub_api_pipeline_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codepipeline.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "mind_hub_api_pipeline_policy" {
  name = "mind_hub_api_pipeline_policy"
  role = aws_iam_role.mind_hub_api_pipeline_role.id

  policy = <<EOF
{
  "Statement": [
    {
      "Action": [
        "iam:PassRole"
      ],
      "Resource": "*",
      "Effect": "Allow",
      "Condition": {
        "StringEqualsIfExists": {
          "iam:PassedToService": [
            "cloudformation.amazonaws.com",
            "elasticbeanstalk.amazonaws.com",
            "ec2.amazonaws.com",
            "ecs-tasks.amazonaws.com"
          ]
        }
      }
    },
    {
      "Action": [
        "cloudwatch:*",
        "s3:*"
      ],
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": [
        "lambda:InvokeFunction",
        "lambda:ListFunctions"
      ],
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": [
        "codebuild:BatchGetBuilds",
        "codebuild:StartBuild"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ],
  "Version": "2012-10-17"
}
EOF
}


# Create an IAM role for CodeBuild to assume
resource "aws_iam_role" "mind_hub_api_build_role" {
  name = "mind-hub-api-build_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

# Create an IAM role policy for CodeBuild to use implicitly
resource "aws_iam_role_policy" "mind_hub_api_build_role_policy" {
  name = "mind_hub_api_build_role_policy"
  role = aws_iam_role.mind_hub_api_build_role.name

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "${aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.arn}",
        "${aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.arn}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "ssm:Get*"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "${aws_iam_role.mind_hub_api_build_role.arn}"
    }
  ]
}
POLICY
}


# Create an IAM role for CodeBuild to assume
resource "aws_iam_role" "mind_hub_api_terraform_deploy_role" {
  name = "mind_hub_api_terraform_deploy_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

# Create an IAM role policy for CodeBuild to use implicitly
resource "aws_iam_role_policy" "mind_hub_api_terraform_deploy_role_policy" {
  name = "mind_hub_api_terraform_deploy_role_policy"
  role = aws_iam_role.mind_hub_api_terraform_deploy_role.name

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "${aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.arn}",
        "${aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.arn}/*",
        "${data.aws_s3_bucket.dev_tf_state_bucket.arn}",
        "${data.aws_s3_bucket.dev_tf_state_bucket.arn}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "dynamodb:GetItem",
        "dynamodb:PutItem",
        "dynamodb:DeleteItem",
        "dynamodb:Describe*",
        "dynamodb:List*"
      ],
      "Resource": "${data.aws_dynamodb_table.tf_lock_state.arn}"
    },
    {
      "Effect": "Allow",
      "Action": [
        "ssm:Get*"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "iam:Get*",
        "iam:List*",
        "iam:PassRole"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "${aws_iam_role.mind_hub_api_terraform_deploy_role.arn}"
    }
  ]
}
POLICY
}