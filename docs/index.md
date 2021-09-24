---
layout: "cidaas"
page_title: "Provider: cidaas"
description: |-
  The cidaas provider is used to interract with the cidaas instances and make changes inside the cidaas instances.
---

# Cidaas Provider

The cidaas provider is used to interact with cidaas instances. It provides resources that allow you to create Apps and Registration Page Fields  as part of a Terraform deployment.



## Example Usage

```hcl
terraform {
    required_providers {
      cidaas = {
        version = "1.0.7"
        source  = "hashicorp.com/cidaas-public/cidaas"
      }
    }
  }
```

~> Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system. See [Environment Variables](#environment-variables) for a better alternative .


- Add Cidaas Provider configuration to terraform configuration file inside Example directory

  ```hcl
  provider "cidaas" {
    default_app_client_id     = "Enter client-id of default app"
    default_app_client_secret = "Enter client-secret of default app"
    default_app_redirect_uri  = "Enter redirect-uri of default app"
    default_app_grant_type    = "password"
    default_app_auth_url      = "https://terraform-cidaas-test-free.cidaas.de/token-srv/token"
    default_app_app_url       = "https://terraform-cidaas-test-free.cidaas.de/apps-srv/clients"
    default_app_base_url      = "https://terraform-cidaas-test-free.cidaas.de"
  }
  ```



- Setup Environment variables: Username and Password must be set as environment variable in order to allow Cidaas terraform provider to complete Password flow and generate an access_token 

  ```bash
  export CIDAAS_USERNAME="ENTER CIDAAS USERNAME"
  ```

  ```bash
  export CIDAAS_PASSWORD="ENTER CIDAAS PASSWORD"
  ```



#### Supported Cidaas Resources

##### Cidaas App Resource

Example App resource configuration

```hcl
resource "cidaas_App" "Enter resource name for resource type App" {
  client_type                     = "NON_INTERACTIVE"
  allow_login_with                = ["EMAIL", "MOBILE", "USER_NAME"]
  auto_login_after_register       = true
  enable_passwordless_auth        = false
  register_with_login_information = true
  hosted_page_group               = "default"
  client_name                     = "Enter client name"
  client_display_name             = "Enter client display name"
  company_name                    = "Enter company name"
  company_address                 = "Enter company address"
  company_website                 = "https://cidaas.com"
  allowed_scopes                  = ["openid", "cidaas:register", "profile"]
  response_types                  = ["code", "token", "id_token"]
  grant_types                     = ["client_credentials"]
  template_group_id               = "Enter template-group-id"
}
```

```hcl
resource "cidaas_registration_page_field" "Enter resource name for resource type registration_page_fields" {
  parent_group_id      = "DEFAULT"
  is_group             = false
  data_type            = "TEXT"
  field_key            = "Enter registration_page_field name"
  required             = false
  enabled              = false
  read_only            = false
  internal             = false
  scopes               = []
  claimable            = true
  order                = 25
  field_type           = "CUSTOM"
  locale_text_locale   = "en-GB"
  locale_text_name     = "erraform-test-field"
  locale_text_language = "en"
}             
```


- Run Terraform commands going inside Example directory where Terraform config file main.tf is located . 

  1. terraform init : It will build the Terraform Cidaas Plugin/Provider.
  2. terraform Plan : It will show the plan that Terraform has to execute from the current config file(main.tf) configurations.
  3. terraform apply : The Terraform will execute the changes and the infrastructure will get provisioned.
