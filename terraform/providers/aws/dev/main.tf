terraform {
  required_version = ">= 0.14.2"

  backend "s3" {
    bucket         = "mind-hub-api-dev-tf-state"
    key            = "terraform.tfstate"
    region         = "eu-west-1"
    dynamodb_table = "mind-hub-ui-state-lock"
  }

  required_providers {
    aws = "~> 3.21"
  }
}

provider "aws" {
  region                  = "eu-west-1"
  shared_credentials_file = "~/.aws/credentials"
  profile                 = "mind-terraform"
}

// Provider used to access the ACM SSL Cert from us-east-1
# https://github.com/hashicorp/terraform/issues/21472#issuecomment-497508239
provider "aws" {
  alias  = "us_east"
  region = "us-east-1"
}


module "mind-hub-ui" {
  source = "../../../modules/mind-hub-api"

  env                = "dev"
  auth0_audience     = var.auth0_audience
  auth0_jwks_uri     = var.auth0_jwks_uri
  auth0_token_issuer = var.auth0_token_issuer
  graph_cms_url      = "https://api-eu-central-1.graphcms.com/v2/ckftjhf769ysi01z7ari84qio/master"

  providers = {
    aws         = aws
    aws.us_east = aws.us_east
  }
}
