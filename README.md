### Cidaas Terraform Provider



#### Installation Steps

- [Setup golang private repo import](https://gitlab.widas.de/cidaas-v2/cidaas-documentation/development-guidelines/-/wikis/how-to/How-to-import-private-GO-projects)

- Clone the GoLang repository 

  ```bash
  git clone git@gitlab.widas.de:customer-specific-projects/rehau/cidaas-terraform-provider-sandbox.git

- Install Cidaas Terraform Provider plugin

  ```bash
  ubuntu@~/root/.../cidaas-terraform-provider-sandbox make
  ```

#### Documentation of Usage

- Add Cidaas as a required provider to terraform configuration file

  ```hcl
  terraform {
    required_providers {
      cidaas = {
        version = "0.0.3"
        source  = "hashicorp.com/cidaas-public/cidaas"
      }
    }
  }
  ```

  

- Add Cidaas Provider configuration to terraform configuration file 

  ```hcl
  provider "cidaas" {
    default_app_client_id     = "Enter client-id of default app"
    default_app_client_secret = "Enter client-secret of default app"
    default_app_redirect_uri  = "Enter redirect-uri of default app"
    default_app_grant_type    = "password"
    default_app_auth_url      = "https://terraform-cidaas-test-free.cidaas.de/token-srv/token"
    default_app_app_url       = "https://terraform-cidaas-test-free.cidaas.de/apps-srv/clients"
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
resource "cidaas_app" "{Name of the app}" {
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



### Notes and Guidelines

- If an app/client attribute is modified for an Cidaas App which which is managed by terraform, the App with former attribute configuration is destroyed and a new App with updated attribute configuration is generated (with newly generated client-id and client-secret).

- Only password based auth flow is supported for automated login and usage of Cidaas terraform provider.
