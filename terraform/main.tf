provider "aws" {
  profile = "jds"
  region = "us-west-1"
}

provider "github" {
  token        = "${data.aws_ssm_parameter.github_organization.value}"
  organization = "${data.aws_ssm_parameter.github_token.value}"
}

data "aws_ssm_parameter" "github_organization" {
  name = "/github_organization"
}

data "aws_ssm_parameter" "github_token" {
  name = "/github_token"
}

# cloudfront
locals {
  s3_origin_id = "S3-Website-trumpmoney.john-shenk.com.s3-website-us-west-1.amazonaws.com"
}

resource "aws_cloudfront_distribution" "s3_distribution" {
  origin {
    domain_name = "${aws_s3_bucket.trumpmoney.website_endpoint}"
    origin_id   = "${local.s3_origin_id}"

    custom_origin_config {
      http_port = "80"
      https_port= "443"
      origin_protocol_policy = "http-only"
      origin_ssl_protocols = ["TLSv1", "TLSv1.1", "TLSv1.2"]
      origin_read_timeout = 30
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"
  price_class         = "PriceClass_200"

  aliases = ["trumpmoney.john-shenk.com"]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "${local.s3_origin_id}"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    compress               = true
    min_ttl                = 0
    default_ttl            = 86400
    max_ttl                = 31536000
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["US", "CA"]
    }
  }

  tags = {
    Environment = "production"
  }

  viewer_certificate {
    acm_certificate_arn = "arn:aws:acm:us-east-1:671958020402:certificate/426e437c-5f6c-4e81-ba47-a3152bd7a44d"
    ssl_support_method = "sni-only"
    minimum_protocol_version = "TLSv1.1_2016"
  }

  price_class = "PriceClass_All"
}

resource "aws_route53_record" "trumpmoney" {
  zone_id = "Z3P68RXJ4VECYX"
  name    = "trumpmoney.john-shenk.com"
  type    = "A"

  alias {
    name                   = "${aws_cloudfront_distribution.s3_distribution.domain_name}"
    zone_id                = "${aws_cloudfront_distribution.s3_distribution.hosted_zone_id}"
    evaluate_target_health = false
  }
}
