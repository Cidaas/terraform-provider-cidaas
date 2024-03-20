![Logo](https://raw.githubusercontent.com/Cidaas/terraform-provider-cidaas/master/logo.jpg)

## About cidaas:
[cidaas](https://www.cidaas.com)
 is a fast and secure Cloud Identity & Access Management solution that standardises what’s important and simplifies what’s complex.

## Feature set includes:
* Single Sign On (SSO) based on OAuth 2.0, OpenID Connect, SAML 2.0 
* Multi-Factor-Authentication with more than 14 authentication methods, including TOTP and FIDO2 
* Passwordless Authentication 
* Social Login (e.g. Facebook, Google, LinkedIn and more) as well as Enterprise Identity Provider (e.g. SAML or AD) 
* Security in Machine-to-Machine (M2M) and IoT

# Terraform Provider for Cidaas

The Terraform provider for Cidaas enables interaction with Cidaas instances that allows to perform CRUD operations on applications, custom providers, registration fields and many other functionalities. From managing applications to configuring custom providers, the Terraform provider enhances the user's capacity to define, provision and manipulate their Cidaas resources.

## Prerequisites

- Ensure Terraform is installed on your local machine. Find installation instructions for different operating systems [here](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli).


## Example Usage

Below is a step-by-step guide to help you set up the provider, configure essential environment variables and integrate the provider into your configuration:

### 1. Terraform Provider Declaration

Begin by specifying the Cidaas provider in your `terraform` block in your Terraform configuration file:

```hcl
terraform {
    required_providers {
      cidaas = {
        version = "1.0.0"
        source  = "Cidaas/cidaas"
      }
    }
}
```

Terraform pulls the version configured of the Cidaas provider for your infrastructure.

### 2. Setup Environment Variables

To authenticate and authorize Terraform operations with Cidaas, set the necessary environment variables. These variables include your Cidaas client credentials, allowing the Terraform provider to complete the client credentials flow and generate an access_token. Execute the following commands in your terminal, replacing placeholders with your actual Cidaas client ID and client secret.

```bash
export TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID="ENTER CIDAAS CLIENT ID"
export TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET="ENTER CIDAAS CLIENT SECRET"
```

### 3. Add Cidaas Provider Configuration

Next, add the Cidaas provider configuration to your Terraform configuration file. Specify the `base_url` parameter to point to your Cidaas instance. For reference, check the example folder.

```hcl
provider "cidaas" {
  base_url = "https://terraform-cidaas-test-free.cidaas.de"
}
```

**Note:** Starting from version 2.5.1, the `redirect_url` is no longer supported in the provider configuration. Ensure that you adjust your configuration accordingly.

By following these steps, you integrate the Cidaas Terraform provider, enabling you to manage your Cidaas resources with Terraform.

## Supported Resources

The Terraform provider for Cidaas supports a variety of resources that enables you to manage and configure different aspects of your Cidaas environment. These resources are designed to integrate with Terraform workflows, allowing you to define, provision and manage your Cidaas resources as code.

To prevent unintended changes to specific attributes of your Terraform resources, you can use the `ignore_changes` configuration within the `lifecycle` block. With this, Terraform ignores to update the attributes provided during `terraform apply`. For example, to avoid change in `client_id`, add the following block to the cidaas_app config file:

```hcl
resource "cidaas_app" "sample" {
  lifecycle {
    ignore_changes = [
      client_id
    ]
  }
}
```

Explore the following resources to understand their attributes, functionalities and how to use them in your Terraform configurations:


## Custom Provider

This example demonstrates the configuration of a custom provider resource for interacting with Cidaas. Before using this custom provider to perform CRUD operations on `cidaas_custom_provider`, ensure the following scopes are added to the client associated with the specified `client_id`. These scopes are essential for enabling the necessary permissions:

### Required Scopes:

- **cidaas:provider_read** : This scope grants read access. It is necessary for fetching data from the provider.

- **cidaas:provider_write** :  This scope is essential for write operations on the `cidaas_custom_provider`.

### Configuration Example:

Below is an example configuration for the custom provider resource in your Terraform files:

```hcl
resource "cidaas_custom_provider" "sample" {
  standard_type          = "OAUTH2"
  authorization_endpoint = "https://terraform-cidaas-test-free.cidaas.de/authz-srv/authz"
  token_endpoint         = "https://terraform-cidaas-test-free.cidaas.de/token-srv/token"
  provider_name          = "Terraform"
  display_name           = "Terraform"
  logo_url               = "https://terraform-cidaas-test-free.cidaas.de/logo"
  userinfo_endpoint      = "https://terraform-cidaas-test-free.cidaas.de/users-srv/userinfo"
  scope_display_label    = "Terraform Test Scope"
  client_id              = "add your client id"
  client_secret          = "add your cluient secret"

  scopes {
    recommended = false
    required    = false
    scope_name  = "openid"
  }
  scopes {
    recommended = false
    required    = false
    scope_name  = "profile"
  }
  scopes {
    recommended = false
    required    = false
    scope_name  = "email"
  }

  userinfo_fields {
    name               = "cp_name"
    family_name        = "cp_family_name"
    address            = "cp_address"
    birthdate          = "01-01-2000"
    email              = "cp@email.com"
    email_verified     = "true"
    gender             = "male"
    given_name         = "cp_given_name"
    locale             = "cp_locale"
    middle_name        = "cp_middle_name"
    mobile_number      = "100000000"
    nickname           = "cp_nickname"
    phone_number       = "10000000"
    picture            = "https://cp-picture.com/image.jpg"
    preferred_username = "cp_preferred_username"
    profile            = "cp_profile"
    updated_at         = "01-01-01"
    website            = "https://cp-website.com"
    zoneinfo           = "cp_zone_info"
    sub                = "bcb-4a6b-9777-8a64abe6af"
    custom_fields = [
      {
        key   = "terraform_test_cf"
        value = "key from the "
      }
    ]
  }
}
```
Refer to the detailed parameter descriptions provided in the table below :

| Key                     | Type | Description                                           |
|-------------------------|-----------|-------------------------------------------------------|
| standard_type         | String    | Type of standard. Allowed values OAUTH2 and OPENID_CONNECT
| authorization_endpoint| String    | URL for authorization in Cidaas                        |
| token_endpoin`        | String    | URL for token in Cidaas                                |
| provider_name         | String    | Name of the provider
| display_name          | String    | Display name of the provider
| logo_url              | String    | URL for the provider's logo                            |
| userinfo_endpoint     | String    | URL for userinfo in Cidaas                             |
| scope_display_label   | String    | Display label for the specified scope |
| client_id             | String    | Cidaas client ID                                      |
| client_secret         | String    | Cidaas client secret                                  |
| scopes                | List      | List of scopes with details (recommended, required, scope_name). Details in the next table |
| userinfo_fields       | Object    | Object containing various user information fields with their values |

Parameters in the scopes described here :

| Key | Type | Description                                      |
|--------------|-----------|--------------------------------------------------|
| recommended| Boolean   | Indicates if the scope is recommended            |
| required   | Boolean   | Indicates if the scope is required               |
| scope_name | String    | The name of the scope, e.g., "openid", "profile" |

Note: The userinfo_fields section includes specific fields such as name, family_name, address, etc., along with custom_fields allowing additional user information customization.

### Import Statement:

Use the following command to import an existing `cidaas_custom_provider` into Terraform:

```bash
terraform import custom_provider.resource_name provider_name
```

In this command, `resource_name` refers to the name assigned to your `cidaas_custom_provider` resource, and `provider_name` represents the actual provider name used in your Cidaas instance. Ensure the correct mapping between the Terraform resource and the corresponding Cidaas provider. This import statement is important for syncing Terraform state with the existing Cidaas resources.

## App

The App resource allows creation and management of clients in Cidaas system. To perform CRUD operations on these clients, it is necessary to assign specific scopes to the client with the designated client_id in the environment.

- cidaas:apps_read
- cidaas:apps_write
- cidaas:apps_delete

When creating a client with a custom `client_id` and `client_secret` you can include the configuration in the resource. If not provided, Cidaas will generate a set for you. `client_secret` is sensitive data. Refer to the article [Terraform Sensitive Variables](https://developer.hashicorp.com/terraform/tutorials/configuration-language/sensitive-variables) to properly handle sensitive information.

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_app" "terraform_app" {
  client_type                     = "SINGLE_PAGE"
  accent_color                    = "#ef4923"
  primary_color                   = "#ef4923"
  media_type                      = "IMAGE"
  content_align                   = "CENTER"
  allow_login_with                = ["EMAIL", "MOBILE", "USER_NAME"]
  redirect_uris                   = ["https://cidaas.com"]
  allowed_logout_urls             = ["https://cidaas.com"]
  enable_deduplication            = true
  auto_login_after_register       = true
  enable_passwordless_auth        = false
  register_with_login_information = true
  allow_disposable_email          = false
  validate_phone_number           = false
  fds_enabled                     = false
  hosted_page_group               = "default"
  client_name                     = "Terraform Test App"
  client_display_name             = "Test Test App Display Name"
  company_name                    = "Widas ID GmbH"
  company_address                 = "01"
  company_website                 = "https://cidaas.com"
  allowed_scopes                  = ["openid", "cidaas:register", "profile"]
  response_types                  = ["code", "token", "id_token"]
  grant_types                     = ["client_credentials"]
  login_providers                 = ["login_provider1", "login_provider2"]
  additional_access_token_payload = ["sample_payload"]
  required_fields                 = ["email"]
  is_hybrid_app                   = false
  allowed_web_origins             = ["https://cidaas.com"]
  allowed_origins                 = ["https://cidaas.com"]
  mobile_settings {
    team_id      = "sample-team-id"
    bundle_id    = "sample-bundle-id"
    package_name = "sample-package-name"
    key_hash     = "sample-key-hash"
  }
  default_max_age                   = 86400
  token_lifetime_in_seconds         = 86400
  id_token_lifetime_in_seconds      = 86400
  refresh_token_lifetime_in_seconds = 15780000
  template_group_id                 = "custtemp"
  policy_uri                        = "https://cidaas.com"
  tos_uri                           = "https://cidaas.com"
  imprint_uri                       = "https://cidaas.com"
  contacts                          = ["support@cidas.de"]
  token_endpoint_auth_method        = "client_secret_post"
  token_endpoint_auth_signing_alg   = "RS256"
  default_acr_values                = ["default"]
  editable                          = true
  web_message_uris                  = ["https://cidaas.com"]
  social_providers {
    provider_name = "cidaas social provider"
    social_id     = "fdc63bd0-6044-4fa0-abff"
    display_name  = "cidaas"
  }
  custom_providers {
    logo_url      = "https://cidaas.com/logo-url"
    provider_name = "sample-custom-provider"
    display_name  = "sample-custom-provider"
    type          = "CUSTOM_OPENID_CONNECT"
  }
   custom_providers {
    logo_url      = cidaas_custom_provider.sample.logo_url
    provider_name = cidaas_custom_provider.sample.provider_name
    display_name  = cidaas_custom_provider.sample.display_name
    type          = cidaas_custom_provider.sample.standard_type
  }
  saml_providers {
    logo_url      = "https://cidaas.com/logo-url"
    provider_name = "sample-sampl-provider"
    display_name  = "sample-sampl-provider"
    type          = "SAMPL_IDP_PROVIDER"
  }
  ad_providers {
    logo_url      = "https://cidaas.com/logo-url"
    provider_name = "sample-ad-provider"
    display_name  = "sample-ad-provider"
    type          = "ADD_PROVIDER"
  }
  app_owner    = "Cidaas"
  jwe_enabled  = false
  user_consent = false
  allowed_groups {
    group_id      = "developer101"
    roles         = ["developer", "qa", "admin"]
    default_roles = ["developer"]
  }

  operations_allowed_groups {
    group_id      = "developer101"
    roles         = ["developer", "qa", "admin"]
    default_roles = ["developer"]
  }

  deleted                             = false
  enabled                             = false
  allowed_fields                      = ["email"]
  always_ask_mfa                      = false
  smart_mfa                           = false
  allowed_mfa                         = ["OFF"]
  captcha_ref                         = "sample-captcha-ref"
  captcha_refs                        = ["sample"]
  consent_refs                        = ["sample"]
  communication_medium_verification   = "email_verification_required_on_usage"
  email_verification_required         = true
  mobile_number_verification_required = true
  allowed_roles                       = ["sample"]
  default_roles                       = ["sample"]
  enable_classical_provider           = false
  is_remember_me_selected             = false
  bot_provider                        = "CIDAAS"
  allow_guest_login_groups {
    group_id      = "developer101"
    roles         = ["developer", "qa", "admin"]
    default_roles = ["developer"]
  }
  is_login_success_page_enabled    = false
  is_register_success_page_enabled = false
  group_ids                        = ["sample"]
  admin_client                     = false
  is_group_login_selection_enabled = false
  group_selection {
    selectable_groups      = ["developer-users"]
    selectable_group_types = ["sample"]
  }
  group_types               = ["sample"]
  backchannel_logout_uri    = "https://test.com/logout"
  post_logout_redirect_uris = ["sample"]
  logo_align                = "CENTER"
  mfa {
    setting                  = "OFF"
    time_interval_in_seconds = 86400
    allowed_methods          = [""]
  }
  push_config {
    tenant_key = "cidaas-tenant"
    name       = "sample-push-config"
    vendor     = "cidaas"
    key        = "bcb-4a6b-9777-8a64abe6af"
    secret     = "bcb-4a6b-9777-8a64abe6af"
    owner      = "cidaas"
  }
  webfinger                       = "no_redirection"
  application_type                = ""
  logo_uri                        = "https://sample-logo.com/logo"
  initiate_login_uri              = "https://cidaas.com/initiate-login"
  client_secret_expires_at        = 3600
  client_id_issued_at             = 3600
  registration_client_uri         = "https://cidaas.com/registration-client-uri"
  registration_access_token       = "registration access token"
  client_uri                      = "https://cidaas.com/client-uri"
  jwks_uri                        = "https://cidaas.com/jwk-uri"
  jwks                            = "https://cidaas.com/jwks"
  sector_identifier_uri           = "https://cidaas.com/sector-identifier-uri"
  subject_type                    = "sample subject type"
  id_token_signed_response_alg    = "RS256"
  id_token_encrypted_response_alg = "RS256"
  id_token_encrypted_response_enc = ""
  userinfo_signed_response_alg    = "RS256"
  userinfo_encrypted_response_alg = "RS256"
  userinfo_encrypted_response_enc = ""
  request_object_signing_alg      = "RS256"
  request_object_encryption_alg   = "RS256"
  request_object_encryption_enc   = "userinfo_encrypted_response_enc"
  request_uris                    = ["sample"]
  description            = "it's a sample description of the client. The client supports system to system communication"
  default_scopes         = ["sample"]
  pending_scopes         = ["sample"]
  consent_page_group     = "sample-consent-page-group"
  password_policy_ref    = "password-policy-ref"
  blocking_mechanism_ref = "blocking-mechanism-ref"
  sub                    = "sample-sub"
  role                   = "sample-role"
  mfa_configuration      = "sample-configuration"
  suggest_mfa            = ["OFF"]
  login_spi {
    oauth_client_id = "bcb-4a6b-9777-8a64abe6af"
    spi_url         = "https://cidaas.com/spi-url"
  }
  video_url       = "https://cidaas.com/video-url"
  bot_captcha_ref = "sample-bot-captcha-ref"
  background_uri  = "http://cidaas.com/background-uri"
}
```

### Import Statement:

Use the following command to import an existing `cidaas_app` into Terraform:

```bash
terraform import cidaas_app.resource_name client_id
```

Here, client_id is the specific identifier of the existing app client in the Cidaas instance that you want to associate with your Terraform configuration.

## Scope

The Scope resource allows to manage scopes in Cidaas system. Scopes define the level of access and permissions granted to an application (client). To perform CRUD operations on Scopes, please add the below scopes to the client with the designated client_id in the environment:

* cidaas:scopes_read
* cidaas:scopes_write
* cidaas:scopes_delete

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_scope" "sample" {
  locale                = "en-US"
  language              = "en-US"
  description           = "terraform description"
  title                 = "terraform title"
  security_level        = "PUBLIC"
  scope_key             = "terraform-test-scope"
  required_user_consent = false
  group_name            = ["terraform-test-group"]
}
```

Refer to the detailed parameter descriptions provided in the table below :

| Key                   | Type     | Description                                                  |
|-----------------------|----------|--------------------------------------------------------------|
| locale                | string   | The locale for the scope, e.g., "en-US".                      |
| language              | string   | The language for the scope, e.g., "en-US".                    |
| description           | string   | Description providing information about the scope. |
| title                 | string   | The title for the scope.                       |
| security_level        | string   | The security level of the scope, e.g., "PUBLIC".              |
| scope_key             | string   | Unique identifier for the scope, used for internal reference. |
| required_user_consent | boolean  | Indicates whether user consent is required for the scope.     |
| group_name            | list     | List of group names to associate the scope with.              |

### Import Statement:

Use the following command to import an existing `cidaas_scope` into Terraform:

```bash
terraform import cidaas_scope.resource_name scope_key
```

## Scope Group

The cidaas_scope_group resource in Terraform allows to manage Scope Groups in Cidaas system. Scope Groups help organize and group related scopes for better categorization and access control. To enable CRUD operations on cidaas_scope_group, ensure that the client associated with the specified client_id has the following roles:

* cidaas:scopes_read
* cidaas:scopes_write
* cidaas:scopes_delete

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_scope_group" "sample" {
  description           = "terraform Scope Group description"
  group_name            = "TerraformScopeGroup"
}
```
Refer to the detailed parameter descriptions provided in the table below :

| Key           | Type   | Description                                                 |
|---------------|--------|-------------------------------------------------------------|
| description   | string | A description providing information about the scope group. |
| group_name    | string | Unique identifier for the scope group. |

### Import Statement:

Use the following command to import an existing `cidaas_scope_group` into Terraform:

```bash
terraform import cidaas_scope_group.resource_name scopeGroup_key
```

## Role

The cidaas_role resource in Terraform facilitates the management of roles in Cidaas system. This resource allows you to configure and define custom roles to suit your application's specific access control requirements. To enable CRUD operations on cidaas_role, ensure that the client associated with the specified client_id has the following roles:

* cidaas:roles_read
* cidaas:roles_write
* cidaas:roles_delete

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_role" "sample" {
  description     = "Role created using Cidaas custom Terraform Provider"
  role            = "role-terraform"
  name            = "role-terraform"
}
```

Refer to the detailed parameter descriptions provided in the table below :

| Key          | Type   | Description                                                  |
|--------------|--------|--------------------------------------------------------------|
| description  | string | A desription providing information about the role. |
| role         | string | Unique identifier for the role, used for internal reference.  |
| name         | string | name for the role. |


### Import Statement:

Use the following command to import an existing `cidaas_role` into Terraform:

```bash
terraform import cidaas_role.resource_name role
```

## Registration Page Field

The `cidaas_registration_page_field` in Terraform allows management of Registration Page Fields in the Cidaas system. This resource enables you to configure and customize the fields displayed during user registration. To enable CRUD operations on cidaas_registration_page_field, ensure that the client associated with the specified client_id has the following scopes:

* cidaas:field_setup_read
* cidaas:field_setup_write
* cidaas:field_setup_delete

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_registration_page_field" "sample" {
  claimable              = true
  data_type              = "TEXT"
  enabled                = false
  field_key              = "sample_field"
  field_type             = "CUSTOM"
  internal               = false
  is_group               = false
  locale_text_language   = "en"
  locale_text_locale     = "en-us"
  locale_text_name       = "Sample Field"
  order                  = 2
  parent_group_id        = "DEFAULT"
  read_only              = false
  required               = true
  required_msg           = "sample_field is required"
  locale_text_min_length = 10
  locale_text_max_length = 100
  min_length_error_msg   = "minimum length should be 10"
  max_length_error_msg   = "maximum length should be 100"
  scopes = [
    "profile",
    "cidaas:public_profile",
  ]
}
```

Refer to the detailed parameter descriptions provided in the table below :

| Key                     | Type    | Description                                               |
|-------------------------|---------|-----------------------------------------------------------|
| claimable               | boolean | Indicates whether the field is claimable by the user.      |
| data_type               | string  | Specifies the data type of the field (e.g. "TEXT").       |
| enabled                 | boolean | Indicates whether the field is enabled.                   |
| field_key               | string  | Unique key identifier for the registration page field.    |
| field_type              | string  | Specifies the type of the field.         |
| internal                | boolean | Indicates whether the field is internal.                  |
| is_group                | boolean | Indicates whether the field is a group.                   |
| locale_text_language    | string  | Language code for localization (e.g. "en").              |
| locale_text_locale      | string  | Locale code for localization (e.g. "en-us").             |
| locale_text_name        | string  | Name for the field used in localization.  |
| order                   | number  | Specifies the order of the field.                         |
| parent_group_id         | string  | Identifier for the parent group of the field.             |
| read_only               | boolean | Indicates whether the field is read-only.                 |
| required                | boolean | Indicates whether the field is required during registration.|
| required_msg            | string  | Custom message for the required field validation.         |
| locale_text_min_length  | number  | Minimum length for locale_text_name.                         |
| locale_text_max_length  | number  | Maximum length for locale_text_name.                         |
| min_length_error_msg    | string  | Custom message for minimum length validation.             |
| max_length_error_msg    | string  | Custom message for maximum length validation.             |
| scopes                  | list    | List of scopes associated with the field.                 |

### Import Statement:

Use the following command to import an existing `cidaas_registration_page_field` into Terraform:

```bash
terraform import cidaas_registration_page_field.resource_name field_key
```

## Webhook

The Webhook resource in Terraform facilitates integration of webhooks in the Cidaas system. This resource allows you to configure webhooks with different authentication options. To enable CRUD operations on cidaas_webhook, ensure the client associated with the specified client_id has the following scopes:

* cidaas:webhook_read
* cidaas:webhook_write
* cidaas:webhook_delete

### Configuration Example:

The terraform configuration for webhook varies based on the **auth_type**. These comprehensive examples below demonstrates how to use `cidaas_webhook` configurations based on different authentication options, allowing you to integrate webhooks into the Cidaas system.

#### APIKEY
```hcl
resource "cidaas_webhook" "sample_webhook" {
  auth_type          = "APIKEY"
  url                = "https://cidaas.com/webhook-test"
  events             = ["ACCOUNT_MODIFIED"]
  apikey_placeholder = "api-test-placeholder"
  apikey_placement   = "query"
  apikey             = "api-test-key"
}
```

#### TOTP
```hcl
resource "cidaas_webhook" "sample_webhook" {
  auth_type        = "TOTP"
  url              = "https://cidaas.com/webhook-test"
  events           = ["ACCOUNT_MODIFIED"]
  totp_placeholder = "test-totp-placeholder"
  totp_placement   = "header"
  totpkey          = "totp-key"
}
```

#### CIDAAS_OAUTH2
```hcl
resource "cidaas_webhook" "sample_webhook" {
  auth_type = "CIDAAS_OAUTH2"
  url       = "https://cidaas.com/webhook-test"
  events    = ["ACCOUNT_MODIFIED"]
  client_id = "jf1a884-8298-4431-a8k5-2f4130037i17"
}
```

Refer to the detailed parameter descriptions provided in the table below :

| Attribute Name | is optional | Description |
| ------ | ------ | ------ |
| auth_type | no | The attribute auth_type is to define how this url is secured from your end. The allowed values are APIKEY, TOTP and CIDAAS_OAUTH2|
| url | no | The webhook url that needs to be called when an event occurs |
| events | no | The events that trigger the webhook  |
| apikey_placeholder | yes |  **required** parameter when the auth_type is APIKEY. The attribute is the placeholder for the key which need to be passed as a query parameter or in the request header|
| apikey_placement | yes | **required** parameter when the auth_type is APIKEY. The allowed value are **header** and **query**. when the value is set to **header** the apikey will be passed in request header and when set to **query** the apikey is passed as a query parameter |
| apikey | yes | **required** parameter when the auth_type is APIKEY. This is the value of the key that will be passed in the request header or in query param |
| totp_placeholder | yes | **required** parameter when the auth_type is TOTP. The attribute is the placeholder for the totp which need to be passed as a query parameter or in the request header |
| totp_placement | yes | **required** parameter when the auth_type is TOTP. The allowed value are **header** and **query**. when the value is set to **header** the totpkey will be passed in request header and when set to **query** the totpkey is passed as a query parameter |
| totpkey | yes | **required** parameter when the auth_type is TOTP. This is the value of the totp that will be passed in the request header or in query param |
| client_id | yes | **required** parameter when the auth_type is CIDAAS_OAUTH2. This is the id of the client which will be used for authentication when the webhook is triggered|


### Import Statement:

Use the following command to import an existing `cidaas_webhook` into Terraform:

```bash
terraform import cidaas_webhook.resource_name webhook_id
```

## Hosted Page

This Hosted Page resource in Terraform allows you to define and manage hosted pages within the Cidaas system. Ensure that the required scopes are set to work with hosted pages in your Cidaas instance.

 - cidaas:hosted_pages_write
 - cidaas:hosted_pages_read
 - cidaas:hosted_pages_delete

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_hosted_page" "sample" {
  hosted_page_group_name = "hosted-page-sample-group"
  default_locale         = "en-US"

  hosted_pages {
    hosted_page_id = "register_success"
    locale         = "en-US"
    url            = "https://terraform-cidaas-test-free.cidaas.de/register_success_hosted_page"
  }

  hosted_pages {
    hosted_page_id = "login_success"
    locale         = "en-US"
    url            = "https://terraform-cidaas-test-free.cidaas.de/login_success_hosted_page"
  }
}
```
Refer to the detailed parameter descriptions provided in the table below :

| Key                    | Type    | Description                                                                   |
|------------------------|---------|-------------------------------------------------------------------------------|
| hosted_page_group_name| String  | The name of the hosted page group                                            |
| default_locale        | String  | The default locale for hosted pages e.g. "en-US".                           |
| hosted_pages          | List    | List of hosted pages with their respective attributes `hosted_page_id`, `locale`, `url` |

The parameter of the attribute `hosted_pages` described in the table below:

| Key             | Type   | Description                                                   |
|-----------------|--------|---------------------------------------------------------------|
| hosted_page_id| String | The identifier for the hosted page, e.g., "register_success". |
| locale        | String | The locale for the hosted page, e.g., "en-US".               |
| url           | String | The URL for the hosted page                                  |

### Import Statement:

Use the following command to import an existing Hosted Page:

```bash
terraform import cidaas_hosted_page.resource_name hosted_page_group_name
```

## User Group Category

The User Group Category, managed through the `cidaas_user_group_category` resource in Terraform, defines and configures categories for user groups within the Cidaas system.

Ensure that the below scopes are assigned to the client with the specified `client_id` to perform the desired CRUD operations on user group categories using Terraform.

1. **cidaas:group_type_read** : This scope provides read access to user group categories in Cidaas.

2. **cidaas:group_type_write** :  This scope allows users to perform write operations on user group categories in Cidaas. Users with this scope can create and update user group categories.

3. **cidaas:group_type_delete** :  This scope grants users the ability to delete user group categories in Cidaas.

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_user_group_category" "sample" {
  role_mode     = "no_roles"
  group_type    = "TerraformUserGroupCategory"
  description   = "terraform user group category description"
  allowed_roles = []
}
```
Refer to the detailed parameter descriptions provided in the table below :

| Key             | Type    | Description                                                   |
|-----------------|---------|---------------------------------------------------------------|
| role_mode     | String  | Determines the role mode for the user group category.          |
| group_type    | String  | The identifier for the user group category, e.g., "TerraformUserGroupCategory". |
| description   | String  | Description for the user group category.                       |
| allowed_roles | List    | List of allowed roles for the user group category.             |

### Import Statement:

Use the following command to import an existing User Group Category:

```bash
terraform import cidaas_user_group_category.resource_name user_group_category_name
```

## Template

The Template resource in Terraform is used to define and manage templates within the Cidaas system. Templates are used for emails, SMS, IVR, and push notifications. Below is an example configuration for the Cidaas Template resource, along with details of its attributes:

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_template" "sample" {
  locale        = "en-us"
  template_key  = "TERRAFORM_TEST"
  template_type = "SMS"
  content       = "Sample content for the Cidaas template resource."
}
```
Refer to the detailed parameter descriptions provided in the table below :

| Attribute Name | Type    | Is Optional | Description |
| -------------- | ------- | ----------- | ----------- |
| locale       | String  | No          | The locale of the template. e.g. "en-us", "en-uk". Ensure the locale is set in lowercase. |
| template_key | String  | No          | The unique name of the template. It cannot be updated for an existing state. |
| template_type| String  | No          | The type of the template. Allowed template_types are EMAIL, SMS, IVR and PUSH. Template types are case sensitive |
| content      | String  | No          | The content of the template. |
| subject      | String  | Yes         | Applicable only for template_type EMAIL. it represents the subject of the email. |


### Import Statement:

Use the following command to import an existing Cidaas Template:

```bash
terraform import cidaas_template.resource_name template_key_template_type
```

Here, `template_key_template_type` is a combination of `template_key` and `template_type`, joined by the special character "_". For example, if the resource name is "sample" with `template_key` as "foo" and `template_type` as "bar," the import statement would be:

```bash
terraform import cidaas_template.sample foo_bar
```

## User Groups

To enable CRUD operations on `cidaas_user_groups`, ensure the following scopes are added to the client with the specified `client_id` set in the environment:

1. **cidaas:groups_write** :  This scope grants write access to user groups in Cidaas. Users with this scope can create and modify user groups.

2. **cidaas:groups_read** :  This scope provides read access to user groups in Cidaas.

3. **cidaas:groups_delete** :  This scope allows users to delete user groups in Cidaas. Users with this scope can remove existing user groups.

### Configuration Example:

Below is an example configuration in your Terraform files:

```hcl
resource "cidaas_user_groups" "sample" {
  group_type            = "sample-group-type"
  group_id              = "sample-group-id"
  group_name            = "sample-group-name"
  logo_url              = "https://cidaas.de/logo"
  description           = "Sample user groups description"
  make_first_user_admin = false
  custom_fields = {
    custom_field_name = "sample custom field"
  }
  member_profile_visibility      = "full"
  none_member_profile_visibility = "public"
  parent_id                      = "sample-parent-id"
}
```

Refer to the detailed parameter descriptions provided in the table below :

| Key                           | Type       | Description                                                                      |
|-------------------------------|------------|----------------------------------------------------------------------------------|
| group_type                  | String     | Type of the user group                                |
| group_id                    | String     | Identifier for the user group                           |
| group_name                  | String     | Name of the user group                                |
| logo_url                    | String     | URL for the user group's logo                                                    |
| description                 | String     | Description of the user group            |
| make_first_user_admin       | Boolean    | Indicates whether the first user should be made an admin             |
| custom_fields               | Map(String) | Custom fields for the user group |
| member_profile_visibility   | String     | Visibility of member profiles. Allowed values `public` or `full`                              |
| none_member_profile_visibility | String  | Visibility of non-member profiles. Allowed values `none` or `public`           |
| parent_id                   | String     | Identifier of the parent user group |

### Import Statement:

Use the following command to import an existing `cidaas_user_groups` into Terraform:

```bash
terraform import cidaas_user_groups.resource_name user_groups_name
```

### Run these Terraform commands in the example directory to explore the provider:

1. Run `terraform init`: This command builds the Terraform Cidaas Provider.
2. Execute `terraform plan`: It reveals the execution plan based on the current configurations in the main Terraform file (`main.tf`).
3. Use `terraform apply`: This command triggers Terraform to execute the planned changes, provisioning the infrastructure accordingly.