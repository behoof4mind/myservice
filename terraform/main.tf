terraform {
  backend "s3" {}
}

module "webserver_cluster" {
  source            = "github.com/behoof4mind/tf-module-myservice?ref=0.2.48"
  env_prefix        = var.env_prefix
  is_temp_env       = var.is_temp_env
  max_ec2_instances = var.max_ec2_instances
  min_ec2_instances = var.min_ec2_instances
  mysql_username    = var.mysql_username
  mysql_password    = var.mysql_password
}
