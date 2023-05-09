terraform {
  required_providers {
    cidaas = {
      source  = "Cidaas/cidaas"
      version = "1.0.18"
    }
  }
}

provider "cidaas" {
  redirect_uri = "https://kube-nightlybuild-dev.cidaas.de"
  base_url     = "https://kube-nightlybuild-dev.cidaas.de"
}
