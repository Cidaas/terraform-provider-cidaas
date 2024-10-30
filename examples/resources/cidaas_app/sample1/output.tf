output "app1_client_id" {
  value = module.app1.client_id
}

output "app2_client_id" {
  value = module.app2.client_id
}

output "app1_client_name" {
  value = module.app1.client_name
}

output "app2_client_name" {
  value = module.app2.client_name
}

output "app1_client_secret" {
  value     = module.app1.client_secret
  sensitive = true
}

output "app2_client_secret" {
  value     = module.app2.client_secret
  sensitive = true
}
