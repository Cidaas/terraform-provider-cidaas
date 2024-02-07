terraform {
  required_providers {
    cidaas = {
      source = "Cidaas/cidaas"
      # replace the value with the right version
      version = "1.1.2"
    }
  }
}

provider "cidaas" {
  base_url = "https://terraform-cidaas-test-free.cidaas.de"
}
