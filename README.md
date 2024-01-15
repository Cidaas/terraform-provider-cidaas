![Logo](logo.jpg)

## About cidaas:
[cidaas](https://www.cidaas.com)
 is a fast and secure Cloud Identity & Access Management solution that standardises what’s important and simplifies what’s complex.

## Feature set includes:
* Single Sign On (SSO) based on OAuth 2.0, OpenID Connect, SAML 2.0 
* Multi-Factor-Authentication with more than 14 authentication methods, including TOTP and FIDO2 
* Passwordless Authentication 
* Social Login (e.g. Facebook, Google, LinkedIn and more) as well as Enterprise Identity Provider (e.g. SAML or AD) 
* Security in Machine-to-Machine (M2M) and IoT

# Cidaas Provider for Terraform

The cidaas provider for terraform is used to interact with cidaas instances. It provides resources that allow you to create Apps and Registration Page Fields as part of a Terraform deployment.

### Prerequisites

- Install Terraform in your local machine. Find steps to install Terraform for different operating system [here](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)

## Example Usage

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

- Setup Environment variables: client_id and client_secret must be set as environment variable in order to allow Cidaas terraform provider to complete client credentials flow and generate an access_token

  ```bash
  export TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID="ENTER CIDAAS CLIENT ID"
  ```

  ```bash
  export TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET="ENTER CIDAAS CLIENT SECRET"
  ```

- Add Cidaas Provider configuration to terraform configuration file inside Example directory

  ```hcl
  provider "cidaas" {
    redirect_uri  = "Enter redirect-uri of default app"
    base_url      = "https://terraform-cidaas-test-free.cidaas.de"
  }
  ```

## Supported Cidaas Resources

### Cidaas Custom Provider Resource

Example custom provider resource configuration. Please add the below scopes to the client with client_id set in the env in order to perform CRUD on cidaas_custom_provider

* cidaas:identity_provider_read
* cidaas:identity_provider_write

```hcl
resource "cidaas_custom_provider" "cp" {
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
    zoneinfo           = "cp_zone_info",
    sub                = "bcb-4a6b-9777-8a64abe6af"
    custom_fields = [
      {
        key   = "terraform_test_cf"
        value = "key from the "
      }
    ],
  }
}
```

Use the command below to import an existing cidaas_custom_provider

```ssh
terraform import cidaas_custom_provider.<resource name> provider_name
```

##### Cidaas App Resource

An example of App resource configuration.

Please add the below scopes to the client with client_id set in the env in order to perform CRUD on cidaas_app

* cidaas:apps_read
* cidaas:apps_write
* cidaas:apps_delete

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

Use the command below to import an existing cidaas_app

```ssh
terraform import cidaas_app.<resource name> client_id
```


##### Cidaas Scope Resource

An example of Scope resource configuration. Please add the below scopes to the client with client_id set in the env in order to perform CRUD on cidaas_scope

* cidaas:scopes_read
* cidaas:scopes_write
* cidaas:scopes_delete

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

Use the command below to import an existing cidaas_scope

```ssh
terraform import cidaas_scope.<resource name> scope_key
```


##### Cidaas Registration Page Field Resource

An example of Registration Page Field resource configuration. Please add the below scopes to the client with client_id set in the env in order to perform CRUD on cidaas_registration_page_field

* cidaas:field_setup_read
* cidaas:field_setup_write
* cidaas:field_setup_delete

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

Use the command below to import an existing cidaas_registration_page_field

```ssh
terraform import cidaas_registration_page_field.<resource name> field_key
```



##### Cidaas Webhook Resource

Some examples of Webhook resource configuration shown below. Please add the below scopes to the client with client_id set in the env in order to perform CRUD on cidaas_webhook

* cidaas:webhook_read
* cidaas:webhook_write
* cidaas:webhook_delete

The terraform configuration for webhook varies based on the **auth_type** provided in the configuration file. Here is the detail of the attribues below

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

The example shown below are the configurations with the required parameters for each auth_type

* APIKEY
```hcl
resource "cidaas_webhook" "sample_webhook" {
  auth_type = "APIKEY"
  url       = "https://cidaas.com/webhook-test"
  events = [
    "ACCOUNT_MODIFIED"
  ]
  apikey_placeholder = "api-test-placeholder"
  apikey_placement   = "query"
  apikey             = "api-test-key"
}
```

* TOTP
```hcl
resource "cidaas_webhook" "sample_webhook" {
  auth_type = "TOTP"
  url       = "https://cidaas.com/webhook-test"
  events = [
    "ACCOUNT_MODIFIED"
  ]
  totp_placeholder   = "test-totp-placeholder"
  totp_placement     = "header"
  totpkey            = "totp-key"
}
```

* CIDAAS_OAUTH2
```hcl
resource "cidaas_webhook" "sample_webhook" {
  auth_type = "CIDAAS_OAUTH2"
  url       = "https://cidaas.com/webhook-test"
  events = [
    "ACCOUNT_MODIFIED"
  ]
  client_id          = "jf1a884-8298-4431-a8k5-2f4130037i17"
}
```

Use the command below to import an existing cidaas_webhook

```ssh
terraform import cidaas_webhook.<resource name> webhook_id
```

##### Cidaas Hosted Page Resource

Please add the below scopes to the client with client_id set in the env in order to perform CRUD on cidaas_hosted_page

* cidaas:hosted_pages_write
* cidaas:hosted_pages_read
* cidaas:hosted_pages_delete

```hcl
resource "cidaas_hosted_page" "sample" {
  hosted_page_group_name = "hosted-page-sample-group"
  default_locale         = "en-us"

  hosted_pages {
    hosted_page_id = "register_success"
    locale         = "en-us"
    url            = "https://terraform-cidaas-test-free.cidaas.de/register_success_hosted_page"
  }

  hosted_pages {
    hosted_page_id = "login_success"
    locale         = "en-us"
    url            = "https://terraform-cidaas-test-free.cidaas.de/login_success_hosted_page"
  }
}

```

Use the command below to import an existing cidaas_hosted_page

```ssh
terraform import cidaas_hosted_page.<resource name> hosted_page_group_name
```

##### To start using the provider run the Terraform commands below going inside the example directory where Terraform config files are available

  1. terraform init : It will build the Terraform Cidaas Plugin/Provider.
  2. terraform Plan : It will show the plan that Terraform has to execute from the current config file(main.tf) configurations.
  3. terraform apply : The Terraform will execute the changes and the infrastructure will get provisioned.
