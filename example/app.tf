resource "cidaas_app" "terraform_app" {
  client_type                       = "SINGLE_PAGE"
  allow_login_with                  = ["EMAIL", "MOBILE", "USER_NAME"]
  auto_login_after_register         = true
  enable_passwordless_auth          = false
  register_with_login_information   = true
  hosted_page_group                 = "default"
  client_name                       = "Terraform Test"
  client_display_name               = "Terraform Test"
  company_name                      = "Widas ID GmbH"
  company_address                   = "01"
  company_website                   = "https://cidaas.com"
  allowed_scopes                    = ["openid", "cidaas:register", "profile"]
  response_types                    = ["code", "token", "id_token"]
  grant_types                       = ["client_credentials"]
  template_group_id                 = "custtemp"
  redirect_uris                     = ["https://cidaas.com"]
  allowed_logout_urls               = ["https://cidaas.com"]
  fds_enabled                       = false
  login_providers                   = ["login_provider1", "login_provider2"]
  default_max_age                   = 86400
  id_token_lifetime_in_seconds      = 86400
  token_lifetime_in_seconds         = 86400
  refresh_token_lifetime_in_seconds = 15780000
  # custom_provider_name            = cidaas_custom_provider.cp.provider_name
}


output "app" {
  value = {
    client_id                 = cidaas_app.terraform_app.client_id
    app_name                  = cidaas_app.terraform_app.client_name
    custom_provider_name      = cidaas_custom_provider.cp.provider_name
    custom_provider_id        = cidaas_custom_provider.cp._id
    custom_provider_client_id = cidaas_custom_provider.cp.client_id
  }
}
