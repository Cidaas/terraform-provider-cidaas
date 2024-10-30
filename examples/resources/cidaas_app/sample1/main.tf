provider "cidaas" {
  base_url = "https://cidaas.dev.de"
}

module "app1" {
  source = "git@github.com:Cidaas/terraform-cidaas-app.git"

  providers = {
    cidaas = cidaas
  }

  client_name    = "Demo SP App"
  common_configs = var.common_configs
}

module "app2" {
  source = "git@github.com:Cidaas/terraform-cidaas-app.git"

  providers = {
    cidaas = cidaas
  }

  client_name    = "Test ISO App"
  client_type    = "IOS"
  common_configs = var.common_configs
}
