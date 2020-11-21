terraform {
  backend "s3" {
    bucket         = "mind-hub-api-management-tf-state"
    key            = "terraform.tfstate"
    region         = "eu-west-1"
    dynamodb_table = "mind-hub-ui-state-lock"
  }
}

provider "aws" {
  version                 = "~> 3.0"
  region                  = "eu-west-1"
  shared_credentials_file = "~/.aws/credentials"
  profile                 = "mind-terraform"
}

provider "github" {
  version      = "~> 2.4"
  token        = var.github_token
  individual   = false
  organization = "jdpx"
}

module "mind-hub-api-pipeline" {
  source = "../../../modules/pipeline"

  env              = "management"
  repository_owner = "jdpx"
  repository_name  = "mind-hub-api"
  github_token     = var.github_token
}
