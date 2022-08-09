terraform {
  required_providers {
    cidaas = {
      source  = "Cidaas/cidaas"
      version = "=1.0.9"
    }
  }
}

provider "cidaas" {
  default_app_client_id     = "efa47d1c-6ad0-4112-b2c0-7b937ac22972"
  default_app_client_secret = "19c9924a-932c-424c-ac6d-91473edb613f"
  default_app_redirect_uri  = "https://example.com"
  default_app_grant_type    = "client_credentials"
  default_app_auth_url      = "https://idp-dev.stackit.cloud/token-srv/token"
  default_app_app_url       = "https://idp-dev.stackit.cloud/apps-srv/clients"
  default_app_base_url      = "https://idp-dev.stackit.cloud"
}


resource "cidaas_app" "genchevm" {
  client_type                     = "SINGLE_PAGE"
  allow_login_with                = ["EMAIL", "MOBILE", "USER_NAME"]
  auto_login_after_register       = true
  enable_passwordless_auth        = false
  register_with_login_information = true
  hosted_page_group               = "default"
  client_name                     = "testMisho"
  client_display_name             = "testMisho"
  company_name                    = "testMisho"
  company_address                 = "testMisho"
  company_website                 = "https://stackit.de"
  allowed_scopes                  = ["profile", "openid", "email", "mobile"]
  response_types                  = ["token"]
  grant_types                     = ["client_credentials", "refresh_token", "password", "authorization_code", "implicit"]
  template_group_id               = "default"
  redirect_uris                   = ["https://cidaas.com"]
  allowed_logout_urls             = ["https://cidaas.com"]
}
