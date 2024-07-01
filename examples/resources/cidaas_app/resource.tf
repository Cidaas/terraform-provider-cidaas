resource "cidaas_app" "sample" {
  client_type         = "SINGLE_PAGE"
  client_name         = "Terraform Example App"
  client_display_name = "The client Terraform Example App is a sample application designed to demonstrate the configuration of the terraform cidaas_app resource."
  redirect_uris       = ["https://cidaas.de/home"]
  allowed_logout_urls = ["https://cidaas.de/logout"]
  company_name        = "Cidaas"
  company_address     = "12 Wimsheim, Germany"
  company_website     = "https://cidaas.de"
  allowed_scopes      = ["profile", "email"]
}
