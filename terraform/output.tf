output "myservice_lb" {
  value = module.webserver_cluster.clb_dns_name
}

output "myservice_db" {
  value = module.webserver_cluster.db_dns_name
}