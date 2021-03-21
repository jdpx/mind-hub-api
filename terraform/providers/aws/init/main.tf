terraform {
  required_version = ">= 0.14.2"

  backend "s3" {
    bucket = "mind-hub-api-tf-init-state"
    key    = "terraform.tfstate"
    region = "eu-west-1"
  }

  required_providers {
    aws = "~> 3.21"
  }
}

provider "aws" {
  region = "eu-west-1"
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
