**Install Terraform in your local machine**

[Steps to install Terraform for different Operating System](https://learn.hashicorp.com/tutorials/terraform/install-cli)



**Install Go in your local machine**

[Steps to install go for different Operating System](https://golang.org/doc/install)




### Cidaas Terraform Provider



#### Installation Steps

- [Setup golang private repo import](https://gitlab.widas.de/cidaas-v2/cidaas-documentation/development-guidelines/-/wikis/how-to/How-to-import-private-GO-projects)

- Clone the GoLang repository 

  ```bash
  git clone -v git@gitlab.widas.de:cidaas-management/terraform.git

- Install Cidaas Terraform Provider plugin

  ```bash
  ubuntu@~/root/.../cidaas-terraform-provider-sandbox make
  ```

#### Documentation of Usage

- Cidaas as a required provider to terraform configuration file

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

  

- Add Cidaas Provider configuration to terraform configurations file inside the Example directory

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



- Setup Environment variables: Username and Password must be set as environment variable in order to allow the Cidaas terraform provider to complete Password flow and generate an access_token 

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

- Run Terraform commands going inside Example directory where Terraform config file main.tf is located. 

  1. terraform init : It will build the Terraform Cidaas Plugin/Provider.
  2. terraform Plan : It will show the plan that Terraform has to execute from the current config file(main.tf) configurations.
  3. terraform apply : The Terraform will execute the changes and the infrastructure will get provisioned.

### Notes and Guidelines

- If an app/client attribute is modified for an Cidaas App which which is managed by terraform, the App with former attribute configuration is destroyed and a new App with updated attribute configuration is generated (with newly generated client-id and client-secret).

- Only password based auth flow is supported for automated login and usage of Cidaas terraform provider.
