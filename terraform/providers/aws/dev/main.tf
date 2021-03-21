terraform {
  required_version = "0.14.2"

  backend "s3" {
    bucket         = "mind-hub-api-dev-tf-state"
    key            = "terraform.tfstate"
    region         = "eu-west-1"
    dynamodb_table = "mind-hub-ui-state-lock"
    role_arn       = "arn:aws:iam::500248363656:role/cd"
  }

  required_providers {
    aws = "~> 3.21"
    auth0 = {
      source  = "alexkappa/auth0"
      version = "0.19.0"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
}

// Provider used to access the ACM SSL Cert from us-east-1
# https://github.com/hashicorp/terraform/issues/21472#issuecomment-497508239
provider "aws" {
  alias  = "us_east"
  region = "us-east-1"
}

module "auth0" {
  source = "../../../modules/auth0"

  env                         = "dev"
  auth0_client_domain         = var.auth0_client_domain
  auth0_client_id             = var.auth0_client_id
  auth0_client_secret         = var.auth0_client_secret
  auth0_user_default_password = var.auth0_user_default_password
}

module "mind-hub-ui" {
  source = "../../../modules/mind-hub-api"

  env                   = "dev"
  auth0_audience        = module.auth0.api_audience
  auth0_jwks_uri        = "https://${module.auth0.tenant_id}.eu.auth0.com/.well-known/jwks.json"
  auth0_token_issuer    = "https://${module.auth0.tenant_id}.eu.auth0.com/"
  graph_cms_url_mapping = "1ohTDUgTsyvWet5lJFCCqnA2F1S:ckftjhf769ysi01z7ari84qio/master"

  providers = {
    aws         = aws
    aws.us_east = aws.us_east
  }
}
