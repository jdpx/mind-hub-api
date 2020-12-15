resource "aws_codepipeline" "mind_hub_api_pipeline" {
  name     = "mind-hub-api-pipeline"
  role_arn = aws_iam_role.mind_hub_api_pipeline_role.arn
  tags = {
    Environment = var.env
  }

  artifact_store {
    location = aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.bucket
    type     = "S3"
  }

  stage {
    name = "Source"

    action {
      category = "Source"
      configuration = {
        "OAuthToken"           = var.github_token
        "Branch"               = var.repository_branch
        "Owner"                = var.repository_owner
        "PollForSourceChanges" = "false"
        "Repo"                 = var.repository_name
      }
      input_artifacts = []
      name            = "Source"
      output_artifacts = [
        "SourceArtifact",
      ]
      owner     = "ThirdParty"
      provider  = "GitHub"
      run_order = 1
      version   = "1"
    }
  }
  stage {
    name = "Build"

    action {
      category = "Build"
      configuration = {
        "EnvironmentVariables" = jsonencode(
          [
            {
              name  = "env"
              type  = "PLAINTEXT"
              value = var.env
            },
          ]
        )
        "ProjectName" = aws_codebuild_project.mind_hub_api_build.name
      }
      input_artifacts = [
        "SourceArtifact",
      ]
      name = "Build"
      output_artifacts = [
        "BuiltUIArtifact",
      ]
      owner     = "AWS"
      provider  = "CodeBuild"
      run_order = 1
      version   = "1"
    }
  }

  stage {
    name = "Deploy"

    action {
      category = "Deploy"
      configuration = {
        "BucketName" = aws_s3_bucket.mind_hub_api_pipeline_artifact_bucket.bucket
        "Extract"    = "true"
      }
      input_artifacts = [
        "BuiltUIArtifact",
      ]
      name             = "DeployDev"
      output_artifacts = []
      owner            = "AWS"
      provider         = "S3"
      run_order        = 1
      version          = "1"
    }

    # action {
    #   category = "Build"
    #   configuration = {
    #     "EnvironmentVariables" = jsonencode(
    #       [
    #         {
    #           name  = "env"
    #           type  = "PLAINTEXT"
    #           value = "dev"
    #         },
    #       ]
    #     )
    #     "ProjectName" = aws_codebuild_project.mind_hub_api_terraform_deploy.name
    #   }
    #   input_artifacts = [
    #     "SourceArtifact",
    #   ]
    #   name      = "DevTerraformDeploy"
    #   owner     = "AWS"
    #   provider  = "CodeBuild"
    #   run_order = 2
    #   version   = "1"
    # }
  }
}
