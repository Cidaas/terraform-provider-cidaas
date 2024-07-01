
# A sample with only the required fields to create an app
# The setup below will create multiple resources with the common config
# The default values will be set when the app is created

variable "common_config" {
  type = object({
    client_type         = string
    company_name        = string
    company_address     = string
    company_website     = string
    allowed_scopes      = set(string)
    redirect_uris       = set(string)
    allowed_logout_urls = set(string)
  })
}

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

resource "cidaas_app" "sample_one" {
  client_name    = "Sample Terraform Application One"
  common_configs = var.common_config
}

resource "cidaas_app" "sample_two" {
  client_name    = "Sample Terraform Application Two"
  common_configs = var.common_config
}

# additional attribute client_display_name added
resource "cidaas_app" "sample_three" {
  client_name         = "Sample Terraform Application Three"
  client_display_name = "Sample terraform display name"
  common_configs      = var.common_config
}

// Here client_type will be override from SINGLE_PAGE(common_config.client_type) to IOS
resource "cidaas_app" "sample_four" {
  client_name    = "Sample Terraform Application Four"
  client_type    = "IOS"
  common_configs = var.common_config
}

