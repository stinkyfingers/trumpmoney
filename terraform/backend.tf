terraform {
  backend "s3" {
    bucket = "remotebackend"
    key    = "trumpmoney/terraform.tfstate"
    region = "us-west-1"
    profile = "jds"
  }
}

data "terraform_remote_state" "trumpmoney" {
  backend = "s3"
  config = {
    bucket  = "remotebackend"
    key     = "trumpmoney/terraform.tfstate"
    region  = "us-west-1"
    profile = "jds"
  }
}
