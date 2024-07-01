terraform {
  required_providers {
    cidaas = {
      source  = "hashicorp.com/Cidaas/cidaas"
      version = "3.0.0"
    }
  }
}

provider "cidaas" {
  base_url = "https://cidaas.de"
}
