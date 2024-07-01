module "shared_settings" {
  source = "./modules"
  common_config = {
    client_type         = "SINGLE_PAGE"
    company_name        = "Widas ID GmbH"
    company_address     = "01"
    company_website     = "https://cidaas.com"
    allowed_scopes      = ["openid", "cidaas:register", "profile"]
    redirect_uris       = ["https://cidaas.com"]
    allowed_logout_urls = ["https://cidaas.com"]
  }
}
