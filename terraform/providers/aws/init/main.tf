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

resource "aws_s3_bucket_public_access_block" "management_state_bucket_block" {
  bucket                  = aws_s3_bucket.management_state_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  restrict_public_buckets = true
  ignore_public_acls      = true
}

resource "aws_s3_bucket_public_access_block" "dev_state_bucket_block" {
  bucket                  = aws_s3_bucket.dev_state_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  restrict_public_buckets = true
  ignore_public_acls      = true
}
