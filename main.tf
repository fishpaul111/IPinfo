terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region  = "us-west-2"
}

resource "aws_instance" "IPinfo" {
  ami           = "ami-830c94e3"
  instance_type = "t2.micro"

  tags = {
    Name = "IPinfoInstance"
  }
}

resource "aws_lambda_function" "IPinfo" {
  function_name    = "IPinfo"
  filename         = "getIPinfo.zip"
  handler          = "getIPinfo"
  source_code_hash = sha256(filebase64("getIPinfo.zip"))
  role             = aws_iam_role.IPinfo.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 1
}

resource "aws_iam_role" "IPinfo" {
  name               = "IPinfo"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY
}
