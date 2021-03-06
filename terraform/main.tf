terraform {
  backend "s3" {}
}

module "webserver_cluster" {
  source            = "github.com/behoof4mind/tf-module-myservice?ref=0.2.51"
  env_prefix        = var.env_prefix
  app_version       = var.app_version
  is_temp_env       = var.is_temp_env
  max_ec2_instances = var.max_ec2_instances
  min_ec2_instances = var.min_ec2_instances
  mysql_username    = var.mysql_username
  mysql_password    = var.mysql_password
}
