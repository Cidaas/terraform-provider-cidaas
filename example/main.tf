terraform {
  required_providers {

    cidaas = {
      version = "1.0.16"
      source  = "Cidaas/cidaas"

    }
  }
}


provider "cidaas" {
  default_app_client_id     = "310be15f-6552-411e-9b97-167cae8bc1cb"
  default_app_client_secret = "0c2bf7a4-d3a9-4725-9f85-bee6d94946d9"
  default_app_redirect_uri  = "https://terraform-cidaas-test-free.cidaas.de/user-profile/editprofile"
  default_app_auth_url      = "https://terraform-cidaas-test-free.cidaas.de/token-srv/token"
  default_app_app_url       = "https://terraform-cidaas-test-free.cidaas.de/apps-srv/clients"
  default_app_base_url      = "https://terraform-cidaas-test-free.cidaas.de"
  default_app_provider_url  = "https://terraform-cidaas-test-free.cidaas.de"
}

resource "cidaas_app" "terraform_test_4" {
  client_type                     = "SINGLE_PAGE"
  allow_login_with                = ["EMAIL", "MOBILE", "USER_NAME"]
  auto_login_after_register       = true
  enable_passwordless_auth        = false
  register_with_login_information = true
  hosted_page_group               = "default"
  client_name                     = "terraform-test-4"
  client_display_name             = "terraform-test-4"
  company_name                    = "Widas ID GmbH"
  company_address                 = "01"
  company_website                 = "https://cidaas.com"
  allowed_scopes                  = ["openid", "cidaas:register", "profile"]
  response_types                  = ["code", "token", "id_token"]
  grant_types                     = ["client_credentials"]
  template_group_id               = "custtemp"
  redirect_uris                   = ["https://cidaas.com"]
  allowed_logout_urls             = ["https://cidaas.com"]
  fds_enabled                     = false
  login_providers                 = ["test", "test2"]
}


resource "cidaas_custom_provider" "cp_test_1" {
  standard_type          = "OAUTH2"
  authorization_endpoint = "https://kube-nightlybuild-dev.cidaas.de//authz-srv/authz"
  token_endpoint         = "https://kube-nightlybuild-dev.cidaas.de/token-srv/token"
  provider_name          = "test4"
  display_name           = "terra test"
  logo_url               = "https://kube-nightlybuild-dev.cidaas.de/logo"
  client_id              = "575549e0-7806-4913-bbfe-70bf997927a9"
  client_secret          = "59186c9c-6867-4380-97b8-2f28fd71b2e8"
  userinfo_endpoint      = "https://qa.cidaas.de/users-srv/userinfo"
  username               = "test2"
}
