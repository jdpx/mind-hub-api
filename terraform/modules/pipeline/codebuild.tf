resource "aws_codebuild_project" "mind_hub_api_build" {
  name           = "mind-hub-api-build"
  badge_enabled  = false
  build_timeout  = 60
  queued_timeout = 480
  service_role   = aws_iam_role.mind_hub_api_build_role.arn

  artifacts {
    encryption_disabled    = false
    name                   = "mind_hub_api_build_artifact"
    override_artifact_name = false
    packaging              = "NONE"
    type                   = "CODEPIPELINE"
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/amazonlinux2-x86_64-standard:2.0"
    image_pull_credentials_type = "CODEBUILD"
    privileged_mode             = false
    type                        = "LINUX_CONTAINER"
  }

  logs_config {
    cloudwatch_logs {
      status = "ENABLED"
    }

    s3_logs {
      encryption_disabled = false
      status              = "DISABLED"
    }
  }

  source {
    type                = "CODEPIPELINE"
    buildspec           = "ci/pipeline_build_buildspec.yml"
    git_clone_depth     = 0
    insecure_ssl        = false
    report_build_status = false
  }
}

resource "aws_codebuild_project" "mind_hub_api_terraform_deploy" {
  name           = "mind_hub_api_terraform_deploy"
  badge_enabled  = false
  build_timeout  = 60
  queued_timeout = 480
  service_role   = aws_iam_role.mind_hub_api_terraform_deploy_role.arn

  artifacts {
    type = "CODEPIPELINE"
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/amazonlinux2-x86_64-standard:2.0"
    image_pull_credentials_type = "CODEBUILD"
    privileged_mode             = false
    type                        = "LINUX_CONTAINER"
  }

  logs_config {
    cloudwatch_logs {
      status = "ENABLED"
    }

    s3_logs {
      encryption_disabled = false
      status              = "DISABLED"
    }
  }

  source {
    type                = "CODEPIPELINE"
    buildspec           = "ci/pipeline_terraform_deploy_buildspec.yml"
    git_clone_depth     = 0
    insecure_ssl        = false
    report_build_status = false
  }
}