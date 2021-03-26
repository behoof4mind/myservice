output "myservice_lb" {
  value = module.webserver_cluster.clb_dns_name
}