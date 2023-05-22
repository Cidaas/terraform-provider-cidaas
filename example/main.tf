terraform {
  required_providers {
    cidaas = {
      source = "hashicorp.com/cidaas-public/cidaas"
      # replace the value with the right version
      version = "1.0.0"
    }
  }
}

provider "cidaas" {
  redirect_uri = "https://sso.id-dev.elisa.fi"
  base_url     = "https://sso.id-dev.elisa.fi"
}
