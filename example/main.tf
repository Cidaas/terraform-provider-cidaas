terraform {
  required_providers {

    cidaas = {
      version = "1.0.9"
      source  = "hashicorp.com/cidaas-public/cidaas"
    }
  }
}


provider "cidaas" {
  default_app_client_id     = "310be15f-6552-411e-9b97-167cae8bc1cb"
  default_app_client_secret = "0c2bf7a4-d3a9-4725-9f85-bee6d94946d9"
  default_app_redirect_uri  = "https://terraform-cidaas-test-free.cidaas.de/user-profile/editprofile"
  default_app_grant_type    = "password"
  default_app_auth_url      = "https://terraform-cidaas-test-free.cidaas.de/token-srv/token"
  default_app_app_url       = "https://terraform-cidaas-test-free.cidaas.de/apps-srv/clients"
  default_app_base_url      = "https://terraform-cidaas-test-free.cidaas.de"
}


resource "cidaas_app" "terraform_test_53" {

  client_type                     = "IOS"
  allow_login_with                = ["EMAIL", "MOBILE", "USER_NAME"]
  auto_login_after_register       = true
  enable_passwordless_auth        = false
  register_with_login_information = true
  hosted_page_group               = "default"
  client_name                     = "terraform-test-53"
  client_display_name             = "terraform-test-53"
  company_name                    = "Widas ID GmbH"
  company_address                 = "01"
  company_website                 = "https://cidaas.com"
  allowed_scopes                  = ["openid", "cidaas:register", "profile"]
  response_types                  = ["code", "token", "id_token"]
  grant_types                     = ["client_credentials"]
  template_group_id               = "custtemp"
  redirect_uris                   = ["https://cidaas.com"]
  allowed_logout_urls             = ["https://cidaas.com"]
}

resource "cidaas_registration_page_field" "terrxaform_test_field_11" {
  parent_group_id      = "DEFAULT"
  is_group             = false
  data_type            = "TEXT"
  field_key            = "terraform-test-field-11"
  required             = false
  enabled              = false
  read_only            = false
  internal             = false
  scopes               = []
  claimable            = true
  order                = 25
  field_type           = "CUSTOM"
  locale_text_locale   = "en-GB"
  locale_text_name     = "erraform-test-field"
  locale_text_language = "en"
}

