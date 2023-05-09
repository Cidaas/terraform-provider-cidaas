terraform {
  required_providers {
    cidaas = {
      source  = "Cidaas/cidaas"
      version = "1.0.0"
    }
  }
}

provider "cidaas" {
  redirect_uri = "https://terraform-cidaas-test-free.cidaas.de"
  base_url     = "https://terraform-cidaas-test-free.cidaas.de"
}
