resource "aws_s3_bucket" "trumpmoney" {
  bucket = "trumpmoney.john-shenk.com"
  acl    = ""
  force_destroy = false
  website {
    index_document = "index.html"
  }
}

resource "aws_s3_bucket_policy" "trumpmoney" {
  bucket = "${aws_s3_bucket.trumpmoney.id}"
  policy = <<EOF
{
  "Version":"2012-10-17",
  "Statement":[
    {
      "Sid":"PublicReadForGetBucketObjects",
      "Effect":"Deny",
      "NotPrincipal":{
        "AWS": [
          "${aws_iam_role.trumpmoney.arn}"
        ]
      },
      "Action":"s3:GetObject",
      "Resource":"arn:aws:s3:::trumpmoney.john-shenk.com/*"
    }
  ]
}
EOF
}

resource "aws_iam_role" "trumpmoney" {
  name = "trumpmoney_pipeline_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "trumpmoney" {
  name = "trumpmoney_pipeline_policy"
  role = "${aws_iam_role.trumpmoney.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect":"Allow",
      "Action": [
        "s3:Get*",
        "s3:List*"
      ],
      "Resource": [
        "${aws_s3_bucket.trumpmoney.arn}",
        "${aws_s3_bucket.trumpmoney.arn}/index.html",
        "${aws_s3_bucket.trumpmoney.arn}/wasm_exec.js",
        "${aws_s3_bucket.trumpmoney.arn}/main.wasm"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "codebuild:BatchGetBuilds",
        "codebuild:StartBuild"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "codebuild:*"
      ],
      "Resource": "*"
    },{
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ]
    },{
      "Effect": "Allow",
      "Action": [
        "cloudfront:CreateInvalidation",
        "cloudfront:GetInvalidation",
        "cloudfront:ListInvalidations"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}


resource "aws_codebuild_project" "trumpmoney" {
  name          = "trumpmoney"
  description   = "trumpmoney"
  build_timeout = "5"
  service_role  = "${aws_iam_role.trumpmoney.arn}"

  artifacts {
    type = "NO_ARTIFACTS"
  }

  cache {
    type     = "S3"
    location = "${aws_s3_bucket.trumpmoney.bucket}"
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/standard:2.0"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"

    # environment variables go here
  }

  source {
    type            = "GITHUB"
    location        = "https://github.com/stinkyfingers/trumpmoney.git"
    git_clone_depth = 1
    buildspec       = "buildspec.yml"
  }

  tags = {
    "Environment" = "Prod"
  }
}

resource "aws_codebuild_webhook" "trumpmoney" {
  project_name = "${aws_codebuild_project.trumpmoney.name}"
}
