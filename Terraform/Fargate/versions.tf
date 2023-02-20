terraform {
  # cloud {
  #   organization = "vettom"

  #   workspaces {
  #     name = "learn"
  #   }
  # }
  required_version = ">= 1.1"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.74"
    }
  }
}

provider "aws" {
  profile = "dvt"
  region  = var.aws_region
}