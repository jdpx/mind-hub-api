terraform {
  backend "s3" {
    bucket         = "mind-hub-api-tf-init-state"
    key            = "terraform.tfstate"
    region         = "eu-west-1"
  }
}

provider "aws" {
  version                 = "~> 3.0"
  region                  = "eu-west-1"
  shared_credentials_file = "~/.aws/credentials"
  profile                 = "mind-terraform"
}


resource "aws_s3_bucket" "management_state_bucket" {
  bucket = "mind-hub-api-management-tf-state"

  versioning {
    enabled = true
  }
}

resource "aws_s3_bucket" "dev_state_bucket" {
  bucket = "mind-hub-api-dev-tf-state"

  versioning {
    enabled = true
  }
}