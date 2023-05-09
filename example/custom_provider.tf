resource "cidaas_custom_provider" "customer_provider" {
  standard_type          = "OAUTH2"
  authorization_endpoint = "https://kube-nightlybuild-dev.cidaas.de/authz-srv/authz"
  token_endpoint         = "https://kube-nightlybuild-dev.cidaas.de/token-srv/token"
  provider_name          = "Terraform"
  display_name           = "Terraform"
  logo_url               = "https://kube-nightlybuild-dev.cidaas.de/logo"
  userinfo_endpoint      = "https://qa.cidaas.de/users-srv/userinfo"
  username               = "Terraform User"
  scope_names            = ["cidaas", "cidaas-user"]
  scope_display_label    = "Terraform Test Scope"
}
