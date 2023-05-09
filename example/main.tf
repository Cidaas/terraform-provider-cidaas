terraform {
  required_providers {
    cidaas = {
      version = "1.0.17"
      source  = "hashicorp.com/cidaas-public/cidaas"
    }
  }
}


provider "cidaas" {
  redirect_uri = "https://kube-nightlybuild-dev.cidaas.de"
  base_url     = "https://kube-nightlybuild-dev.cidaas.de"
}
