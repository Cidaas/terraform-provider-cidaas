# resource "cidaas_custom_provider" "sample" {}

resource "cidaas_custom_provider" "sample" {
  authorization_endpoint = "https://terraform-cidaas-test-free.cidaas.de/authz-srv/authz"
  display_name           = "Terraform Name"
  id                     = "terraform"
  logo_url               = "https://terraform-cidaas-test-free.cidaas.de/logo"
  provider_name          = "terraform"
  scope_display_label    = "Terraform Test Scope"
  scope_names = [
    "cidaas",
    "cidaas-user",
  ]
  standard_type     = "OAUTH2"
  token_endpoint    = "https://terraform-cidaas-test-free.cidaas.de/token-srv/token"
  userinfo_endpoint = "https://qa.cidaas.de/users-srv/userinfo"
}


output "sample_custom_provider" {
  value = cidaas_custom_provider.sample
}
