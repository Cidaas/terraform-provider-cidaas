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
- [Go](https://go.dev/doc/install) (1.21)



## Documentation

Official documentation on how to use this provider can be found on the
[Terraform Registry](https://registry.terraform.io/providers/Cidaas/cidaas/latest/docs). Detailed explanations of the resources can also be found in the [Supported Resources](#supported-resources) section.

## Example Usage

Below is a step-by-step guide to help you set up the provider, configure essential environment variables and integrate the provider into your configuration:

### 1. Terraform Provider Declaration

Begin by specifying the Cidaas provider in your `terraform` block in your Terraform configuration file:

```hcl
terraform {
    required_providers {
      cidaas = {
        version = "3.0.0"
        source  = "Cidaas/cidaas"
      }
    }
}
```

Terraform pulls the version configured of the Cidaas provider for your infrastructure.

### 2. Setup Environment Variables

To authenticate and authorize Terraform operations with Cidaas, set the necessary environment variables. These variables include your Cidaas client credentials, allowing the Terraform provider to complete the client credentials flow and generate an access_token. Execute the following commands in your terminal, replacing placeholders with your actual Cidaas client ID and client secret.

### For Linux and MacOS:
```bash
export TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID="ENTER CIDAAS CLIENT ID"
export TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET="ENTER CIDAAS CLIENT SECRET"
```

### For Windows:
```bash
Set-Item -Path env:TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID -Value “ENTER CIDAAS CLIENT ID“
Set-Item -Path env:TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET -Value “ENTER CIDAAS CLIENT SECRET“
```

You can get a set of client credentials from the Cidaas Admin UI by creating a new client. Simply go to the `Apps` > `App Settings` > `Create New App`. It's important to note that when creating the client, you must select the app type as **Non-Interactive**.

### 3. Add Cidaas Provider Configuration

Next, add the Cidaas provider configuration to your Terraform configuration file. Specify the `base_url` parameter to point to your Cidaas instance. For reference, check the example folder.

```hcl
provider "cidaas" {
  base_url = "https://cidaas.de"
}
```

**Note:** Starting from version 2.5.1, the `redirect_url` is no longer supported in the provider configuration. Ensure that you adjust your configuration accordingly.

By following these steps, you integrate the Cidaas Terraform provider, enabling you to manage your Cidaas resources with Terraform.

## Supported Resources

The Terraform provider for Cidaas supports a variety of resources that enables you to manage and configure different aspects of your Cidaas environment. These resources are designed to integrate with Terraform workflows, allowing you to define, provision and manage your Cidaas resources as code.


Explore the following resources to understand their attributes, functionalities and how to use them in your Terraform configurations:

- [cidaas_app](#cidaas_app-resource)
- [cidaas_custom_provider](#cidaas_custom_provider-resource)
- [cidaas_group_type](#cidaas_group_type-resource-previously-cidaas_user_group_category)
- [cidaas_hosted_page](#cidaas_hosted_page-resource)
- [cidaas_registration_field](#cidaas_registration_field-resource)
- [cidaas_role](#cidaas_role-resource)
- [cidaas_scope_group](#cidaas_scope_group-resource)
- [cidaas_scope](#cidaas_scope-resource)
- [cidaas_template_group](#cidaas_template_group-resource)
- [cidaas_template](#cidaas_template-resource)
- [cidaas_user_groups](#cidaas_user_groups-resource)
- [cidaas_webhook](#cidaas_webhook-resource)

# cidaas_app (Resource)

The App resource allows creation and management of clients in Cidaas system. When creating a client with a custom `client_id` and `client_secret` you can include the configuration in the resource. If not provided, Cidaas will generate a set for you. `client_secret` is sensitive data. Refer to the article [Terraform Sensitive Variables](https://developer.hashicorp.com/terraform/tutorials/configuration-language/sensitive-variables) to properly handle sensitive information.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:apps_read
- cidaas:apps_write
- cidaas:apps_delete


## V2 to V3 Migration:
If you are migrating from v2 to v3, please note the following changes in the v3 version:

### Attributes not supported in app config anymore:

- client_secret_expires_at
- client_id_issued_at
- push_config
- created_at
- updated_at
- admin_client
- deleted
- app_owner

### Change in data types of some attribute

 - social_providers
 - custom_providers
 - saml_providers
 - ad_providers

 The above attributes now has to be provided as set of objects.

 #### Example:
 ```terraform
 {
  ...
  social_providers = [
    {
        logo_url      = "https://cidaas.com/logo-url"
        provider_name = "sample-custom-provider"
        display_name  = "sample-custom-provider"
        type          = "CUSTOM_OPENID_CONNECT"
        is_provider_visible = true
        domains = ["https://cidaas.de/]
    },
    {
        logo_url      = "https://cidaas.com/logo-url"
        provider_name = "sample-custom-provider"
        display_name  = "sample-custom-provider"
        type          = "CUSTOM_OPENID_CONNECT"
        is_provider_visible = true
        domains = ["https://cidaas.de/]
    },
  ]
 }
 ```

### V3 App Resource Highlights:

- The resource app can now be set up with minimal configuration. The following parameters are the only required ones to create an app.
In the [schema](#schema) section, only client_name is shown as **required** because the other attributes can be configured in common_configs.
However, each attribute must appear either in the main configuration block or in common_configs. `client_name` cannot be part of common_configs.

  - client_name
  - client_type
  - company_name                   
  - company_address                
  - company_website                 
  - allowed_scopes                
  - redirect_uris
  - allowed_logout_urls
- Attribute `common_configs` added to share same configuration across multiple apps. Pleas check the samples in **examples** directory that demonstrates the use of `common_configs`.
- If you need to override any specific attribute for a particular resource where the same attribute is available in `common_configs`, you can supply the main configuration attribute directly within the resource block.
- If your configuration involves a single resource or if the common configuration attributes are not shared across multiple resources we do not suggest using `common_configs`.

## Example Usage(V3 configuration)

```terraform
# This sample demonstrates the use of both main configuration attributes and common_configs.
# It is important to note that configuring both simultaneously is not necessary.
# Please refer to the documentation of the app resource for more details.

resource "cidaas_app" "sample" {
  client_name                     = "Test Terraform Application" // unique
  client_display_name             = "Display Name of the app"    // unique
  content_align                   = "CENTER"                     // Default: CENTER
  post_logout_redirect_uris       = ["https://cidaas.com"]
  logo_align                      = "CENTER" // Default: CENTER
  allow_disposable_email          = false    // Default: false
  validate_phone_number           = false    // Default: false
  additional_access_token_payload = ["sample_payload"]
  required_fields                 = ["email"]
  mobile_settings = {
    team_id      = "sample-team-id"
    bundle_id    = "sample-bundle-id"
    package_name = "sample-package-name"
    key_hash     = "sample-key-hash"
  }
  // for custom client credentials use client_id and client_secret, you can leave blank if you want cidaas to create a set for you
  # client_id                       = ""
  # client_secret                   = ""
  policy_uri                          = "https://cidaas.com"
  tos_uri                             = "https://cidaas.com"
  imprint_uri                         = "https://cidaas.com"
  contacts                            = ["support@cidas.de"]
  token_endpoint_auth_method          = "client_secret_post" // Default: client_secret_post
  token_endpoint_auth_signing_alg     = "RS256"              // Default: RS256
  default_acr_values                  = ["default"]
  web_message_uris                    = ["https://cidaas.com"]
  allowed_fields                      = ["email"]
  smart_mfa                           = false // Default: false
  captcha_ref                         = "sample-captcha-ref"
  captcha_refs                        = ["sample"]
  consent_refs                        = ["sample"]
  communication_medium_verification   = "email_verification_required_on_usage"
  mobile_number_verification_required = false // Default: false
  enable_bot_detection                = false // Default: false
  allow_guest_login_groups = [{
    group_id      = "developer101"
    roles         = ["developer", "qa", "admin"]
    default_roles = ["developer"]
  }]
  is_login_success_page_enabled    = false // Default: false
  is_register_success_page_enabled = false // Default: false
  group_ids                        = ["sample"]
  is_group_login_selection_enabled = false // Default: false
  group_selection = {
    selectable_groups      = ["developer-users"]
    selectable_group_types = ["sample"]
  }
  group_types                     = ["sample"]
  logo_uri                        = "https://cidaas.com"
  initiate_login_uri              = "https://cidaas.com"
  registration_client_uri         = "https://cidaas.com"
  registration_access_token       = "registration access token"
  client_uri                      = "https://cidaas.com"
  jwks_uri                        = "https://cidaas.com"
  jwks                            = "https://cidaas.com/jwks"
  sector_identifier_uri           = "https://cidaas.com"
  subject_type                    = "sample subject type"
  id_token_signed_response_alg    = "RS256"
  id_token_encrypted_response_alg = "RS256"
  id_token_encrypted_response_enc = "example"
  userinfo_signed_response_alg    = "RS256"
  userinfo_encrypted_response_alg = "RS256"
  userinfo_encrypted_response_enc = "example"
  request_object_signing_alg      = "RS256"
  request_object_encryption_alg   = "RS256"
  request_object_encryption_enc   = "userinfo_encrypted_response_enc"
  request_uris                    = ["sample"]
  description                     = "app description"
  consent_page_group              = "sample-consent-page-group"
  password_policy_ref             = "password-policy-ref"
  blocking_mechanism_ref          = "blocking-mechanism-ref"
  sub                             = "sample-sub"
  role                            = "sample-role"
  mfa_configuration               = "sample-configuration"
  suggest_mfa                     = ["OFF"]
  login_spi = {
    oauth_client_id = "bcb-4a6b-9777-8a64abe6af"
    spi_url         = "https://cidaas.com/spi-url"
  }
  background_uri  = "https://cidaas.com"
  video_url       = "https://cidaas.com"
  bot_captcha_ref = "sample-bot-captcha-ref"
  application_meta_data = {
    status : "active"
    version : "1.0.0"
  }
  // common config starts here. The attributes from common config can be part of main config
  // if an attribute is available both common_config and main config then attribute from the main config will be considered to create an app
  common_configs = {
    client_type                       = "SINGLE_PAGE"
    accent_color                      = "#ef4923"                        // Default: #ef4923
    primary_color                     = "#ef4923"                        // Default: #f7941d
    media_type                        = "IMAGE"                          // Default: IMAGE
    allow_login_with                  = ["EMAIL", "MOBILE", "USER_NAME"] // Default: ["EMAIL", "MOBILE", "USER_NAME"]
    redirect_uris                     = ["https://cidaas.com"]
    allowed_logout_urls               = ["https://cidaas.com"]
    enable_deduplication              = true      // Default: false
    auto_login_after_register         = true      // Default: false
    enable_passwordless_auth          = false     // Default: true
    register_with_login_information   = false     // Default: false
    fds_enabled                       = false     // Default: true
    hosted_page_group                 = "default" // Default: default
    company_name                      = "Widas ID GmbH"
    company_address                   = "01"
    company_website                   = "https://cidaas.com"
    allowed_scopes                    = ["openid", "cidaas:register", "profile"]
    response_types                    = ["code", "token", "id_token"] // Default: ["code", "token", "id_token"]
    grant_types                       = ["client_credentials"]        // Default: ["implicit", "authorization_code", "password", "refresh_token"]
    login_providers                   = ["login_provider1", "login_provider2"]
    is_hybrid_app                     = true // Default: false
    allowed_web_origins               = ["https://cidaas.com"]
    allowed_origins                   = ["https://cidaas.com"]
    default_max_age                   = 86400      // Default: 86400
    token_lifetime_in_seconds         = 86400      // Default: 86400
    id_token_lifetime_in_seconds      = 86400      // Default: 86400
    refresh_token_lifetime_in_seconds = 15780000   // Default: 15780000
    template_group_id                 = "custtemp" // Default: default
    editable                          = true       // Default: true
    social_providers = [{
      provider_name = "cidaas social provider"
      social_id     = "fdc63bd0-6044-4fa0-abff"
      display_name  = "cidaas"
    }]
    custom_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-custom-provider"
      display_name  = "sample-custom-provider"
      type          = "CUSTOM_OPENID_CONNECT"
    }]
    saml_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-sampl-provider"
      display_name  = "sample-sampl-provider"
      type          = "SAMPL_IDP_PROVIDER"
    }]
    ad_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-ad-provider"
      display_name  = "sample-ad-provider"
      type          = "ADD_PROVIDER"
    }]
    jwe_enabled  = true // Default: false
    user_consent = true // Default: false
    allowed_groups = [{
      group_id      = "developer101"
      roles         = ["developer", "qa", "admin"]
      default_roles = ["developer"]
    }]
    operations_allowed_groups = [{
      group_id      = "developer101"
      roles         = ["developer", "qa", "admin"]
      default_roles = ["developer"]
    }]
    enabled                     = true // Default: true
    always_ask_mfa              = true // Default: false
    allowed_mfa                 = ["OFF"]
    email_verification_required = true // Default: true
    allowed_roles               = ["sample"]
    default_roles               = ["sample"]
    enable_classical_provider   = true     // Default: true
    is_remember_me_selected     = true     // Default: true
    bot_provider                = "CIDAAS" // Default: CIDAAS
    allow_guest_login           = true     // Default: false
    #  mfa Default:
    #  {
    #   setting = "OFF"
    #  }
    mfa = {
      setting = "OFF"
    }
    webfinger      = "no_redirection"
    default_scopes = ["sample"]
    pending_scopes = ["sample"]
  }
}
```

For more samples on common_configs, please refer to the examples folder.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `client_name` (String) Name of the client.

### Optional

- `accent_color` (String) The accent color of the client. e.g., `#f7941d`. The value must be a valid hex colorThe default is set to `#ef4923`.
- `ad_providers` (Attributes List) A list of Active Directory identity providers that users can authenticate with. (see [below for nested schema](#nestedatt--ad_providers))
- `additional_access_token_payload` (Set of String) Access token payload defination.
- `allow_disposable_email` (Boolean) Allow disposable email addresses. Default is set to `false` while creating an app.
- `allow_guest_login` (Boolean) Flag to specify whether guest users are allowed to access functionalities of the client. Default is set to `false`
- `allow_guest_login_groups` (Attributes List) (see [below for nested schema](#nestedatt--allow_guest_login_groups))
- `allow_login_with` (Set of String) allow_login_with is used to specify the preferred methods of login allowed for a client. Allowed values are EMAIL, MOBILE and USER_NAMEThe default is set to `['EMAIL', 'MOBILE', 'USER_NAME']`.
- `allowed_fields` (Set of String)
- `allowed_groups` (Attributes List) (see [below for nested schema](#nestedatt--allowed_groups))
- `allowed_logout_urls` (Set of String) Allowed logout URLs for OAuth2 client.
- `allowed_mfa` (Set of String)
- `allowed_origins` (Set of String) List of the origins allowed to access the client.
- `allowed_roles` (Set of String)
- `allowed_scopes` (Set of String) The URL of the company website. allowed_scopes is a required attribute. It must be provided in the main config or common_config
- `allowed_web_origins` (Set of String) List of the web origins allowed to access the client.
- `always_ask_mfa` (Boolean)
- `application_meta_data` (Map of String) A map to add metadata of a client.
- `auto_login_after_register` (Boolean) Automatically login after registration. Default is set to `false` while creating an app.
- `backchannel_logout_uri` (String)
- `background_uri` (String) The URL to the background image of the client.
- `blocking_mechanism_ref` (String)
- `bot_captcha_ref` (String)
- `bot_provider` (String)
- `captcha_ref` (String)
- `captcha_refs` (Set of String)
- `client_display_name` (String) The display name of the client.
- `client_id` (String) The client_id is the unqique identifier of the app. It's an optional attribute. If not provided, cidaas will gererate one for you and the state will be updated with the same
- `client_secret` (String, Sensitive) The client_id is the unqique identifier of the app. It's an optional attribute. If not provided, cidaas will gererate one for you and the state will be updated with the same
- `client_type` (String) The type of the client. The allowed values are SINGLE_PAGE, REGULAR_WEB, NON_INTERACTIVEIOS, ANDROID, WINDOWS_MOBILE, DESKTOP, MOBILE, DEVICE and THIRD_PARTY
- `client_uri` (String)
- `common_configs` (Attributes) The `common_configs` attribute is used for sharing the same configuration across multiple cidaas_app resources. It is a map of some attributes from the main configuration. Please check the list of the attributes that it supports in the common_confis section. if an attribute is available both common_config and main config then attribute from the main config will be considered to create an app (see [below for nested schema](#nestedatt--common_configs))
- `communication_medium_verification` (String)
- `company_address` (String) The company address.
- `company_name` (String) The name of the company that the client belongs to.
- `company_website` (String) The URL of the company website.
- `consent_page_group` (String)
- `consent_refs` (Set of String)
- `contacts` (Set of String) The contacts of the client.
- `content_align` (String) The alignment of the content of the client. e.g., `CENTER`. Allowed values are CENTER, LEFT and RIGHTThe default is set to `CENTER`.
- `custom_providers` (Attributes List) A list of custom identity providers that users can authenticate with. A custom provider can be created with the help of the resource cidaas_custom_provider. (see [below for nested schema](#nestedatt--custom_providers))
- `default_acr_values` (Set of String)
- `default_max_age` (Number) The default maximum age for the token in seconds. Default is 86400 seconds (24 hours).
- `default_roles` (Set of String)
- `default_scopes` (Set of String)
- `description` (String)
- `editable` (Boolean) Flag to define if your client is editable or not. Default is `true`.
- `email_verification_required` (Boolean)
- `enable_bot_detection` (Boolean)
- `enable_classical_provider` (Boolean)
- `enable_deduplication` (Boolean) Enable deduplication.
- `enable_passwordless_auth` (Boolean) Enable passwordless authentication. Default is set to `true` while creating an app.
- `enabled` (Boolean)
- `fds_enabled` (Boolean) Flag to enable or disable fraud detection system. By default, it is enabled when a client is created
- `grant_types` (Set of String) The grant types of the client. The default value is set to `['implicit','authorization_code', 'password', 'refresh_token']`
- `group_ids` (Set of String)
- `group_selection` (Attributes) (see [below for nested schema](#nestedatt--group_selection))
- `group_types` (Set of String)
- `hosted_page_group` (String) Hosted page group.
- `id_token_encrypted_response_alg` (String)
- `id_token_encrypted_response_enc` (String)
- `id_token_lifetime_in_seconds` (Number) The lifetime of the id_token in seconds. Default is 86400 seconds (24 hours).
- `id_token_signed_response_alg` (String)
- `imprint_uri` (String) The URL to the imprint page.
- `initiate_login_uri` (String)
- `is_group_login_selection_enabled` (Boolean)
- `is_hybrid_app` (Boolean) Flag to set if your app is hybrid or not. Default is set to `false`. Set to `true` to make your app hybrid.
- `is_login_success_page_enabled` (Boolean)
- `is_register_success_page_enabled` (Boolean)
- `is_remember_me_selected` (Boolean)
- `jwe_enabled` (Boolean) Flag to specify whether JSON Web Encryption (JWE) should be enabled for encrypting data.
- `jwks` (String)
- `jwks_uri` (String)
- `login_providers` (Set of String) With this attribute one can setup login provider to the client.
- `login_spi` (Attributes) A map defining the Login SPI configuration. (see [below for nested schema](#nestedatt--login_spi))
- `logo_align` (String)
- `logo_uri` (String)
- `media_type` (String) The media type of the client. e.g., `IMAGE`. Allowed values are VIDEO and IMAGEThe default is set to `IMAGE`.
- `mfa` (Attributes) Configuration settings for Multi-Factor Authentication (MFA). (see [below for nested schema](#nestedatt--mfa))
- `mfa_configuration` (String)
- `mobile_number_verification_required` (Boolean)
- `mobile_settings` (Attributes) (see [below for nested schema](#nestedatt--mobile_settings))
- `operations_allowed_groups` (Attributes List) (see [below for nested schema](#nestedatt--operations_allowed_groups))
- `password_policy_ref` (String)
- `pending_scopes` (Set of String)
- `policy_uri` (String) The URL to the policy of a client.
- `post_logout_redirect_uris` (Set of String)
- `primary_color` (String) The primary color of the client. e.g., `#ef4923`. The value must be a valid hex colorThe default is set to `#f7941d`.
- `redirect_uris` (Set of String) Redirect URIs for OAuth2 client.
- `refresh_token_lifetime_in_seconds` (Number) The lifetime of the refresh token in seconds. Default is 15780000 seconds.
- `register_with_login_information` (Boolean) Register with login information. Default is set to `false` while creating an app.
- `registration_access_token` (String)
- `registration_client_uri` (String)
- `request_object_encryption_alg` (String)
- `request_object_encryption_enc` (String)
- `request_object_signing_alg` (String)
- `request_uris` (Set of String)
- `required_fields` (Set of String) The required fields while registering to the client.
- `response_types` (Set of String) The response types of the client. The default value is set to `['code','token', 'id_token']`
- `role` (String)
- `saml_providers` (Attributes List) A list of SAML identity providers that users can authenticate with. (see [below for nested schema](#nestedatt--saml_providers))
- `sector_identifier_uri` (String)
- `smart_mfa` (Boolean)
- `social_providers` (Attributes List) A list of social identity providers that users can authenticate with. Examples: Google, Facebook etc... (see [below for nested schema](#nestedatt--social_providers))
- `sub` (String)
- `subject_type` (String)
- `suggest_mfa` (Set of String)
- `template_group_id` (String) The id of the template group to be configured for commenication. Default is set to the system default group.
- `token_endpoint_auth_method` (String)
- `token_endpoint_auth_signing_alg` (String)
- `token_lifetime_in_seconds` (Number) The lifetime of the token in seconds. Default is 86400 seconds (24 hours).
- `tos_uri` (String) The URL to the TOS of a client.
- `user_consent` (Boolean) Specifies whether user consent is required or not. Default is `false`
- `userinfo_encrypted_response_alg` (String)
- `userinfo_encrypted_response_enc` (String)
- `userinfo_signed_response_alg` (String)
- `validate_phone_number` (Boolean) if enabled, phone number is validaed. Default is set to `false` while creating an app.
- `video_url` (String) The URL to the video of the client.
- `web_message_uris` (Set of String) A list of URLs for web messages used.
- `webfinger` (String)

### Read-Only

- `id` (String) The ID of the resource.

<a id="nestedatt--ad_providers"></a>
### Nested Schema for `ad_providers`

Required:

- `provider_name` (String)
- `type` (String)

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)


<a id="nestedatt--allow_guest_login_groups"></a>
### Nested Schema for `allow_guest_login_groups`

Required:

- `group_id` (String)

Optional:

- `default_roles` (Set of String)
- `roles` (Set of String)


<a id="nestedatt--allowed_groups"></a>
### Nested Schema for `allowed_groups`

Required:

- `group_id` (String)

Optional:

- `default_roles` (Set of String)
- `roles` (Set of String)


<a id="nestedatt--common_configs"></a>
### Nested Schema for `common_configs`

Optional:

- `accent_color` (String)
- `ad_providers` (Attributes List) (see [below for nested schema](#nestedatt--common_configs--ad_providers))
- `allow_guest_login` (Boolean)
- `allow_login_with` (Set of String)
- `allowed_groups` (Attributes List) (see [below for nested schema](#nestedatt--common_configs--allowed_groups))
- `allowed_logout_urls` (Set of String)
- `allowed_mfa` (Set of String)
- `allowed_origins` (Set of String)
- `allowed_roles` (Set of String)
- `allowed_scopes` (Set of String)
- `allowed_web_origins` (Set of String)
- `always_ask_mfa` (Boolean)
- `auto_login_after_register` (Boolean)
- `bot_provider` (String)
- `client_type` (String)
- `company_address` (String)
- `company_name` (String)
- `company_website` (String)
- `custom_providers` (Attributes List) (see [below for nested schema](#nestedatt--common_configs--custom_providers))
- `default_max_age` (Number)
- `default_roles` (Set of String)
- `default_scopes` (Set of String)
- `editable` (Boolean)
- `email_verification_required` (Boolean)
- `enable_classical_provider` (Boolean)
- `enable_deduplication` (Boolean)
- `enable_passwordless_auth` (Boolean)
- `enabled` (Boolean)
- `fds_enabled` (Boolean)
- `grant_types` (Set of String)
- `hosted_page_group` (String)
- `id_token_lifetime_in_seconds` (Number)
- `is_hybrid_app` (Boolean)
- `is_remember_me_selected` (Boolean)
- `login_providers` (Set of String)
- `logo_align` (String)
- `media_type` (String)
- `mfa` (Attributes) (see [below for nested schema](#nestedatt--common_configs--mfa))
- `operations_allowed_groups` (Attributes List) (see [below for nested schema](#nestedatt--common_configs--operations_allowed_groups))
- `pending_scopes` (Set of String)
- `primary_color` (String)
- `redirect_uris` (Set of String)
- `refresh_token_lifetime_in_seconds` (Number)
- `register_with_login_information` (Boolean)
- `response_types` (Set of String)
- `saml_providers` (Attributes List) (see [below for nested schema](#nestedatt--common_configs--saml_providers))
- `social_providers` (Attributes List) (see [below for nested schema](#nestedatt--common_configs--social_providers))
- `template_group_id` (String)
- `token_lifetime_in_seconds` (Number)
- `webfinger` (String)

<a id="nestedatt--common_configs--ad_providers"></a>
### Nested Schema for `common_configs.ad_providers`

Required:

- `provider_name` (String)
- `type` (String)

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)


<a id="nestedatt--common_configs--allowed_groups"></a>
### Nested Schema for `common_configs.allowed_groups`

Required:

- `group_id` (String)

Optional:

- `default_roles` (Set of String)
- `roles` (Set of String)


<a id="nestedatt--common_configs--custom_providers"></a>
### Nested Schema for `common_configs.custom_providers`

Required:

- `provider_name` (String)
- `type` (String)

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)


<a id="nestedatt--common_configs--mfa"></a>
### Nested Schema for `common_configs.mfa`

Optional:

- `allowed_methods` (Set of String)
- `setting` (String)
- `time_interval_in_seconds` (Number)


<a id="nestedatt--common_configs--operations_allowed_groups"></a>
### Nested Schema for `common_configs.operations_allowed_groups`

Required:

- `group_id` (String)

Optional:

- `default_roles` (Set of String)
- `roles` (Set of String)


<a id="nestedatt--common_configs--saml_providers"></a>
### Nested Schema for `common_configs.saml_providers`

Required:

- `provider_name` (String)
- `type` (String)

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)


<a id="nestedatt--common_configs--social_providers"></a>
### Nested Schema for `common_configs.social_providers`

Required:

- `provider_name` (String)
- `social_id` (String)



<a id="nestedatt--custom_providers"></a>
### Nested Schema for `custom_providers`

Required:

- `provider_name` (String)
- `type` (String)

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)


<a id="nestedatt--group_selection"></a>
### Nested Schema for `group_selection`

Optional:

- `always_show_group_selection` (Boolean)
- `selectable_group_types` (Set of String)
- `selectable_groups` (Set of String)


<a id="nestedatt--login_spi"></a>
### Nested Schema for `login_spi`

Optional:

- `oauth_client_id` (String)
- `spi_url` (String)


<a id="nestedatt--mfa"></a>
### Nested Schema for `mfa`

Optional:

- `allowed_methods` (Set of String) Optional set of allowed MFA methods.
- `setting` (String) Specifies the Multi-Factor Authentication (MFA) setting. Allowed values are 'OFF', 'ALWAYS', 'SMART', 'TIME_BASED' and 'SMART_PLUS_TIME_BASED'.
- `time_interval_in_seconds` (Number) Optional time interval in seconds for time-based Multi-Factor Authentication.


<a id="nestedatt--mobile_settings"></a>
### Nested Schema for `mobile_settings`

Optional:

- `bundle_id` (String)
- `key_hash` (String)
- `package_name` (String)
- `team_id` (String)


<a id="nestedatt--operations_allowed_groups"></a>
### Nested Schema for `operations_allowed_groups`

Required:

- `group_id` (String)

Optional:

- `default_roles` (Set of String)
- `roles` (Set of String)


<a id="nestedatt--saml_providers"></a>
### Nested Schema for `saml_providers`

Required:

- `provider_name` (String)
- `type` (String)

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)


<a id="nestedatt--social_providers"></a>
### Nested Schema for `social_providers`

Required:

- `provider_name` (String)
- `social_id` (String)

## Import

Import is supported using the following syntax:

```shell
# The import identifier in this command is the client_id of the app to be imported.

terraform import cidaas_app.sample client_id
```

# cidaas_custom_provider (Resource)

This example demonstrates the configuration of a custom provider resource for interacting with Cidaas.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:providers_read
- cidaas:providers_write


### V2 to V3 Migration:
If you are migrating from v2 to v3, please note the following changes in the v3 version:

- The attribute `scopes` now has to be set as an array of objects instead of separate separate object
- `custom_fields` in userinfo_fields should be passed as object as shown in the Example Usage section

## Old configuration

```terraform
resource "cidaas_custom_provider" "sample" {
  ...
  scopes {
    recommended = true
    required    = true
    scope_name  = "email"
  }
  scopes {
    recommended = true
    required    = true
    scope_name  = "openid"
  }
  userinfo_fields = {
    custom_fields = [
      {
        key   = "zipcode"
        value = "123456"
      },
      {
        key   = "alternate_phone"
        value = "1234567890"
      }
    ]
  }
}
```


## Example Usage(V3 configuration)

```terraform
resource "cidaas_custom_provider" "sample" {
  standard_type          = "OAUTH2"
  authorization_endpoint = "https://cidaas.de/authz-srv/authz"
  token_endpoint         = "https://cidaas.de/token-srv/token"
  provider_name          = "terraform-sample"
  display_name           = "Terraform"
  logo_url               = "https://cidaas.de/logo"
  userinfo_endpoint      = "https://cidaas.de/users-srv/userinfo"
  scope_display_label    = "terraform sample scope display name"
  client_id              = "acb-4a6b-9777-8a64abe6af"
  client_secret          = "zcb-4a6b-9777-8a64abe6ay"
  domains                = ["cidaas.de", "cidaas.org"]

  scopes = [
    {
      recommended = true
      required    = true
      scope_name  = "email"
    },
    {
      recommended = true
      required    = true
      scope_name  = "openid"
    },
  ]

  userinfo_fields = {
    family_name        = "cp_family_name"
    address            = "cp_address"
    birthdate          = "01-01-2000"
    email              = "cp@cidaas.de"
    email_verified     = "true"
    gender             = "male"
    given_name         = "cp_given_name"
    locale             = "cp_locale"
    middle_name        = "cp_middle_name"
    mobile_number      = "100000000"
    phone_number       = "10000000"
    picture            = "https://cidaas.de/image.jpg"
    preferred_username = "cp_preferred_username"
    profile            = "cp_profile"
    updated_at         = "01-01-01"
    website            = "https://cidaas.de"
    zoneinfo           = "cp_zone_info"
    custom_fields = {
      zipcode         = "123456"
      alternate_phone = "1234567890"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `authorization_endpoint` (String) The URL for authorization of the provider.
- `client_id` (String) The client ID of the provider.
- `client_secret` (String, Sensitive) The client secret of the provider.
- `display_name` (String) The display name of the provider.
- `provider_name` (String) The unique identifier of the custom provider. This cannot be updated for an existing state.
- `scope_display_label` (String) Display label for the scope of the provider.
- `token_endpoint` (String) The URL to generate token with this provider.
- `userinfo_endpoint` (String) The URL to fetch user details using this provider.

### Optional

- `domains` (Set of String) The domains of the provider.
- `logo_url` (String) The URL for the provider's logo.
- `scopes` (Attributes List) List of scopes of the provider with details (see [below for nested schema](#nestedatt--scopes))
- `standard_type` (String) Type of standard. Allowed values `OAUTH2` and `OPENID_CONNECT`.
- `userinfo_fields` (Attributes) Object containing various user information fields with their values. The userinfo_fields section includes specific fields such as name, family_name, address, etc., along with custom_fields allowing additional user information customization (see [below for nested schema](#nestedatt--userinfo_fields))

### Read-Only

- `id` (String) The ID of the resource.

<a id="nestedatt--scopes"></a>
### Nested Schema for `scopes`

Optional:

- `recommended` (Boolean) Indicates if the scope is recommended.
- `required` (Boolean) Indicates if the scope is required.
- `scope_name` (String) The name of the scope, e.g., `openid`, `profile`.


<a id="nestedatt--userinfo_fields"></a>
### Nested Schema for `userinfo_fields`

Optional:

- `address` (String)
- `birthdate` (String)
- `custom_fields` (Map of String)
- `email` (String)
- `email_verified` (String)
- `family_name` (String)
- `gender` (String)
- `given_name` (String)
- `locale` (String)
- `middle_name` (String)
- `mobile_number` (String)
- `name` (String)
- `nickname` (String)
- `phone_number` (String)
- `picture` (String)
- `preferred_username` (String)
- `profile` (String)
- `updated_at` (String)
- `website` (String)
- `zoneinfo` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_custom_provider.resource_name provider_name
```

# cidaas_group_type (Resource)-Previously cidaas_user_group_category

The Group Type, managed through the `cidaas_group_type` resource in the provider defines and configures categories for user groups within the Cidaas system.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:group_type_read
- cidaas:group_type_write
- cidaas:group_type_delete


### V2 to V3 Migration:
If you are migrating from v2 to v3, please note that `cidaas_user_group_category` has been renamed to `cidaas_group_type`.
Please update your Terraform configuration files accordingly to ensure compatibility with the latest version(v3).

## Example Usage

```terraform
resource "cidaas_group_type" "sample" {
  role_mode     = "no_roles"
  group_type    = "TerraformSampleGroupType"
  description   = "terraform user group category description"
  allowed_roles = ["developer"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_type` (String) The unique identifier of the group type. This cannot be updated for an existing state.
- `role_mode` (String) Determines the role mode for the user group type. Allowed values are `any_roles`, `no_roles`, `roles_required` and `allowed_roles`

### Optional

- `allowed_roles` (Set of String) List of allowed roles in this group type.
- `description` (String) The `description` attribute provides details about the group type, explaining its purpose.

### Read-Only

- `created_at` (String) The timestamp when the resource was created.
- `id` (String) The ID of the resource.
- `updated_at` (String) The timestamp when the resource was last updated.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_group_type.resource_name group_type
```

# cidaas_hosted_page (Resource)

The Hosted Page resource in the provider allows you to define and manage hosted pages within the Cidaas system.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:hosted_pages_write
- cidaas:hosted_pages_read
- cidaas:hosted_pages_delete

## Example Usage

```terraform
resource "cidaas_hosted_page" "sample" {
  hosted_page_group_name = "terraform-sample-hosted-page"
  default_locale         = "en-IN"
  hosted_pages = [
    {
      hosted_page_id = "register_success"
      locale         = "en-US"
      url            = "https://cidaas.de/register_success_hosted_page"
      content        = "content"
    },
    {
      hosted_page_id = "register_success"
      locale         = "en-IN"
      url            = "https://cidaas.de/register_success_hosted_page"
      content        = "content"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `hosted_page_group_name` (String) The name of the hosted page group. This must be unique across the cidaas system and cannot be updated for an existing state.
- `hosted_pages` (Attributes List) List of hosted pages with their respective attributes (see [below for nested schema](#nestedatt--hosted_pages))

### Optional

- `default_locale` (String) The default locale for hosted pages e.g. `en-US`.

### Read-Only

- `created_at` (String) The timestamp when the resource was created.
- `id` (String) The ID of the resource.
- `updated_at` (String) The timestamp when the resource was last updated.

<a id="nestedatt--hosted_pages"></a>
### Nested Schema for `hosted_pages`

Required:

- `hosted_page_id` (String) The identifier for the hosted page, e.g., `register_success`.
- `url` (String) The URL for the hosted page.

Optional:

- `content` (String) The conent of the hosted page.
- `locale` (String) The locale for the hosted page, e.g., `en-US`.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_hosted_page.resource_name hosted_page_id
```
# cidaas_registration_field (Resource)

The `cidaas_registration_page_field` in the provider allows management of registration fields in the Cidaas system. This resource enables you to configure and customize the fields displayed during user registration.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:field_setup_read
- cidaas:field_setup_write
- cidaas:field_setup_delete


### V2 to V3 Migration:
If you are migrating from v2 to v3, please note that `cidaas_registration_page_field` has been renamed to `cidaas_registration_field`. Below is the list of changes in `cidaas_registration_field`:

- **Multiple Locales:** Internationalization now supported via `local_texts`.
- **Field Definition Attributes:** Added to specify maximum and minimum lengths for `TEXT` and `DATE` attributes.
- **Extended Datatypes:** It now supports the datatypes `TEXT`, `NUMBER`, `SELECT`, `MULTISELECT`, `RADIO`, `CHECKBOX`, `PASSWORD`, `DATE`, `URL`, `EMAIL`,
	`TEXTAREA`, `MOBILE`, `CONSENT`, `JSON_STRING`, `USERNAME`, `ARRAY`, `GROUPING` and `DAYDATE`.
- **Configuration Updates:** Adjustments required for the below attribute as shown below:

| old config           |      new config                                    |
|---------------------| ----------------------------------------------|
| required_msg | local_texts[i].required_msg |
| locale_text_min_length | field_defination.min_length |
| locale_text_max_length | field_defination.max_length |
| min_length_error_msg | local_texts[i].min_length_msg |
| max_length_error_msg | local_texts[i].max_length_msg |
| locale_text_language | The `language` attribute is no longer required. The provider computes and assigns the language based on the `locale` provided. |
| locale_text_locale | local_texts[i].locale |
| locale_text_name | local_texts[i].name |


#### Attribute `local_text`

| attributes         |      description                             |
|---------------------| ----------------------------------------------|
| locale | The locale of the field. example: de-DE |
| name |Then name of the field in the local configured. for exmaple: in **en-US** the name is `Sample Field` in de-DE `Beispielfeld`|
| max_length_msg | warning/error msg to show to the user when user exceeds the maximum character configured. This is applicable only for the attributes of base_data_type string |
| min_length_msg | warning/error msg to show to the user when user don't provide the minimum character required. This is applicable only for the attributes of base_data_type string |
| required_msg | When the flag required is set to true the required_msg must be provided. required_msg is shown if user does not provide a required field |
| attributes | The field attributes must be provided for the data_type SELECT, MULTISELECT and RADIO. it's an array of key value pairs. example shown below |
| consent_label | required when data_type is CONSENT. exmaple shown below |


### Example of `attributes`
```terraform
 local_texts = [
    {
      locale       = "en-US"
      name         = "Sample Field"
      required_msg = "The field is required"
      attributes = [
        {
          key   = "test_key"
          value = "test_value"
        }
      ]
    }
 ]
```

### Example of `consent_label`
```terraform
local_texts = [
    {
      locale       = "en-US"
      name         = "sample_consent_field"
      required_msg = "The field is required"
      consent_label = {
        label      = "test",
        label_text = "test label text"
      }
    }
  ]
```

### Attribute `field_definition`

| attributes         |      description                             |
|---------------------| ----------------------------------------------|
| max_length | The maximum length of a string type attribute |
| min_length |The minimum length of a string type attribute|
| min_date | applicable only for DATE attribute. example: "2024-06-28T18:30:00Z" |
| max_date | applicable only for DATE attribute. example: "2024-06-28T18:30:00Z" |
| initial_date_view | applicable only for DATE attribute. Allowed values: month, year and multi-year |
| initial_date | applicable only for DATE attribute. example: "2024-06-28T18:30:00Z" |

Ensure your Terraform configurations are updated accordingly to maintain compatibility with the latest version.

## Old configuration:
```terraform
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
  ]
  overwrite_with_null_value_from_social_provider = false
}
```

## Example Usage(V3 configuration)

```terraform
resource "cidaas_registration_field" "text" {
  data_type                                      = "TEXT"
  field_key                                      = "sample_text_field"
  field_type                                     = "CUSTOM"  // CUSTOM and SYSTEM, SYSTEM can not be created but modified
  internal                                       = true      // Default: false
  required                                       = true      // Default: false
  read_only                                      = true      // Default: false
  is_group                                       = false     // Default: false
  unique                                         = true      // Default: false
  overwrite_with_null_value_from_social_provider = false     // Default: true
  is_searchable                                  = true      // Default: true
  enabled                                        = true      // Default: true
  claimable                                      = true      // Default: true
  order                                          = 1         // Default: 1
  parent_group_id                                = "DEFAULT" // Default: DEFAULT
  scopes                                         = ["profile"]
  local_texts = [
    {
      locale         = "en-US"
      name           = "Sample Field"
      required_msg   = "The field is required"
      max_length_msg = "Maximum 99 chars allowed"
      min_length_msg = "Minimum 99 chars allowed"
    },
    {
      locale         = "de-DE"
      name           = "Beispielfeld"
      required_msg   = "Dieses Feld ist erforderlich"
      max_length_msg = "DE maximum 99 chars allowed"
      min_length_msg = "DE minimum 10 chars allowed"
    }
  ]
  field_definition = {
    max_length = 100
    min_length = 10
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `data_type` (String) The data type of the field. This cannot be modified for an existing resource. Allowed values are `TEXT`,`NUMBER`,`SELECT`,`MULTISELECT`,`RADIO`,`CHECKBOX`,`PASSWORD`,`DATE`,`URL`,`EMAIL`,`TEXTAREA`,`MOBILE`,`CONSENT`,`JSON_STRING`,`USERNAME`,`ARRAY`,`GROUPING`,`DAYDATE`,
- `field_key` (String) The unique identifier of the registration field. This cannot be modified for an existing resource.
- `local_texts` (Attributes List) The localized detail of the registration field. (see [below for nested schema](#nestedatt--local_texts))

### Optional

- `claimable` (Boolean) Flag to mark if a field is claimable. Defaults set to `true`
- `consent_refs` (Set of String) List of consents(the ids of the consent in cidaas must be passed) in registration. The data type must be `CONSENT` in this case
- `enabled` (Boolean) Flag to mark if a field is enabled. Defaults set to `true`
- `field_definition` (Attributes) (see [below for nested schema](#nestedatt--field_definition))
- `field_type` (String) Specifies whether the field type is `SYSTEM` or `CUSTOM`. Defaults to `CUSTOM`. This cannot be modified for an existing resource. `SYSTEM` fields cannot be created but can be modified. To modify an existing field import it first and then update.
- `internal` (Boolean) Flag to mark if a field is internal. Defaults set to `false`
- `is_group` (Boolean) Setting is_group to `true` creates a registration field group. Defaults set to `false` The data_type attribute must be set to TEXT when is_group is true.
- `is_list` (Boolean)
- `is_searchable` (Boolean) Flag to mark if a field is searchable. Defaults set to `true`
- `order` (Number) The attribute order is used to set the order of the Field in the UI. Defaults set to `1`
- `overwrite_with_null_value_from_social_provider` (Boolean) Set to true if you want the value should be reset by identity provider. Defaults set to `false`
- `parent_group_id` (String) The ID of the parent registration group. Defaults to `DEFAULT` if not provided.
- `read_only` (Boolean) Flag to mark if a field is read only. Defaults set to `false`
- `required` (Boolean) Flag to mark if a field is required in registration. Defaults set to `false`
- `scopes` (Set of String) The scopes of the registration field.
- `unique` (Boolean) Flag to mark if a field is unique. Defaults set to `false`

### Read-Only

- `base_data_type` (String) The base data type of the field. This is computed property.
- `id` (String) The ID of the resource

<a id="nestedatt--local_texts"></a>
### Nested Schema for `local_texts`

Required:

- `name` (String) Then name of the field in the local configured. for exmaple: in **en-US** the name is `Sample Field` in de-DE `Beispielfeld`.

Optional:

- `attributes` (Attributes List) The field attributes must be provided for the data_type SELECT, MULTISELECT and RADIO. it's an array of key value pairs. Example provided in the example section. (see [below for nested schema](#nestedatt--local_texts--attributes))
- `consent_label` (Attributes) required when data_type is CONSENT. Example provided in the example section. (see [below for nested schema](#nestedatt--local_texts--consent_label))
- `locale` (String) The locale of the field. example: de-DE.
- `max_length_msg` (String) warning/error msg to show to the user when user exceeds the maximum character configured. This is applicable only for the attributes of base_data_type string.
- `min_length_msg` (String) warning/error msg to show to the user when user don't provide the minimum character required. This is applicable only for the attributes of base_data_type string.
- `required_msg` (String) When the flag required is set to true the required_msg must be provided. required_msg is shown if user does not provide a required field.

<a id="nestedatt--local_texts--attributes"></a>
### Nested Schema for `local_texts.attributes`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--local_texts--consent_label"></a>
### Nested Schema for `local_texts.consent_label`

Required:

- `label` (String)
- `label_text` (String)



<a id="nestedatt--field_definition"></a>
### Nested Schema for `field_definition`

Optional:

- `initial_date` (String) The initial date. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.
- `initial_date_view` (String) The view of the calender. Applicable only for DATE attributes. Allowed values: `month`, `year` and `multi-year`
- `max_date` (String) The maximum date a user can select. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.
- `max_length` (Number) The maximum length of a string type attribute.
- `min_date` (String) The earliest date a user can select. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.
- `min_length` (Number) The minimum length of a string type attribute

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_registration_page_field.resource_name field_key
```

# cidaas_role (Resource)

The cidaas_role resource in Terraform facilitates the management of roles in Cidaas system. This resource allows you to configure and define custom roles to suit your application's specific access control requirements.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:roles_read
- cidaas:roles_write
- cidaas:roles_delete

## Example Usage

```terraform
resource "cidaas_role" "sample" {
  role        = "terraform_sample_role"
  name        = "Terraform Sample Role"
  description = "The sample is designed to demonstrate the configuration of the terraform cidaas_role resource."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `role` (String) The unique identifier of the role. The role name must be unique across the cidaas system and cannot be updated for an existing state.

### Optional

- `description` (String) The `description` attribute provides details about the role, explaining its purpose.
- `name` (String) The name of the role.

### Read-Only

- `id` (String) The ID of the role resource.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_role.resource_name role
```

# cidaas_scope_group (Resource)

The cidaas_scope_group resource in the provider allows to manage Scope Groups in Cidaas system. Scope Groups help organize and group related scopes for better categorization and access control.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:scopes_read
- cidaas:scopes_write
- cidaas:scopes_delete

## Example Usage

```terraform
resource "cidaas_scope_group" "sample" {
  group_name  = "terraform-sample-scope"
  description = "The sample is designed to demonstrate the configuration of the terraform cidaas_scope_group resource."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_name` (String) The name of the group. The group name must be unique across the cidaas system and cannot be updated for an existing state.

### Optional

- `description` (String) The `description` attribute provides details about the scope of the group, explaining its purpose.

### Read-Only

- `created_at` (String) The timestamp when the resource was created.
- `id` (String) The ID of th resource.
- `updated_at` (String) The timestamp when the resource was last updated.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_scope_group.resource_name group_name
```

# cidaas_scope (Resource)

The Scope resource allows to manage scopes in Cidaas system. Scopes define the level of access and permissions granted to an application (client).

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:scopes_read
- cidaas:scopes_write
- cidaas:scopes_delete


### V2 to V3 Migration:
If you are migrating from v2 to v3, please note the following changes in the v3 version:

- The `locale`, `language`, `title` and `description` attributes have been removed and replaced with a `localized_descriptions` block that supports a scope with multiple locale with better internationalization. Earlier only one locale was supported by the terraform plugin.
- `localized_descriptions` is a list of objects, each containing:
  - locale
  - title
  - description
- The `language` attribute is no longer required. The provider computes and assigns the language based on the `locale` provided.

## old configuration:
```terraform
resource "scope" "sample" {
  locale                = "en-US"
  language              = "en-US"
  title                 = "terraform title"
  description           = "terraform description"
  security_level        = "PUBLIC"
  scope_key             = "terraform-test-scope"
  required_user_consent = false
  group_name            = []
}
```

## Example Usage(V3 configuration)

```terraform
resource "cidaas_scope" "sample" {
  security_level        = "CONFIDENTIAL"
  scope_key             = "terraform-sample-scope"
  required_user_consent = false
  group_name            = []
  localized_descriptions = [
    {
      title       = "Cidaas Scope Tunisia Title"
      locale      = "ar-TN"
      description = "This is scope in local ar-TN"
    },
    {
      title       = "Cidaas Scope German Title"
      locale      = "de-DE"
      description = "This is scope in local de-DE"
    },
    {
      title       = "Cidaas Scope India Title"
      locale      = "en-IN"
      description = "This is scope in local en-IN"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `scope_key` (String) Unique identifier for the scope. This cannot be updated for an existing state.

### Optional

- `group_name` (Set of String) List of scope_groups to associate the scope with.
- `localized_descriptions` (Attributes List) (see [below for nested schema](#nestedatt--localized_descriptions))
- `required_user_consent` (Boolean) Indicates whether user consent is required for the scope.
- `scope_owner` (String) The owner of the scope. e.g. `ADMIN`
- `security_level` (String) The security level of the scope, e.g., `PUBLIC`. Allowed values are `PUBLIC` and `CONFIDENTIAL`

### Read-Only

- `id` (String) The ID of the resource.

<a id="nestedatt--localized_descriptions"></a>
### Nested Schema for `localized_descriptions`

Required:

- `title` (String) The title of the scope in the configured locale.

Optional:

- `description` (String) The description of the scope in the configured locale.
- `locale` (String) The locale for the scope, e.g., `en-US`.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_scope.resource_name scope_key
```

# cidaas_template_group (Resource)

The cidaas_template_group resource in the provider is used to define and manage templates groups within the Cidaas system. Template Groups categorize your communication templates allowing you to map preferred templates to specific clients effectively.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:templates_read
- cidaas:templates_write
- cidaas:templates_delete

## Example Usage

```terraform
// To create a template group, only the attribute group_id is required in the configuration.
// The attributes shown in sample-tg-2 are optional and can be configured as needed.
// If these properties are not configured in the .tf file, the provider/cidaas will compute
// and assign values to them.

// sample1
resource "cidaas_template_group" "sample-tg-1" {
  group_id = "sample_group"
}

// sample2
resource "cidaas_template_group" "sample-tg-2" {
  group_id = "group_another"
  email_sender_config = {
    from_email = "noreply@cidaas.de"
    from_name  = "Kube-dev"
    reply_to   = "noreply@cidaas.de"
    sender_names = [
      "SYSTEM",
    ]
  }
  ivr_sender_config = {
    sender_names = [
      "SYSTEM",
    ]
  }
  push_sender_config = {
    sender_names = [
      "SYSTEM",
    ]
  }
  sms_sender_config = {
    sender_names = [
      "SYSTEM",
    ]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The group_id of the Template Group. The group_id is used to import an existing template group. The maximum allowed length of a group_id is **15** characters.

### Optional

- `email_sender_config` (Attributes) The `email_sender_config` is used to configure your email sender. (see [below for nested schema](#nestedatt--email_sender_config))
- `ivr_sender_config` (Attributes) The configuration of the IVR sender. (see [below for nested schema](#nestedatt--ivr_sender_config))
- `push_sender_config` (Attributes) The configuration of the PUSH notification sender. (see [below for nested schema](#nestedatt--push_sender_config))
- `sms_sender_config` (Attributes) The configuration of the SMS sender. (see [below for nested schema](#nestedatt--sms_sender_config))

### Read-Only

- `id` (String) The ID of the resource

<a id="nestedatt--email_sender_config"></a>
### Nested Schema for `email_sender_config`

Optional:

- `from_email` (String) The email from address from which the emails will be sent when the specific group is configured.
- `from_name` (String) The `from_name` attribute is the display name that appears in the 'From' field of the emails.
- `reply_to` (String) The `reply_to` attribute is the email address where replies should be directed.
- `sender_names` (Set of String) The `sender_names` attribute defines the names associated with email senders.

Read-Only:

- `id` (String) The `ID` of the configured email sender.


<a id="nestedatt--ivr_sender_config"></a>
### Nested Schema for `ivr_sender_config`

Optional:

- `sender_names` (Set of String)

Read-Only:

- `id` (String)


<a id="nestedatt--push_sender_config"></a>
### Nested Schema for `push_sender_config`

Optional:

- `sender_names` (Set of String)

Read-Only:

- `id` (String)


<a id="nestedatt--sms_sender_config"></a>
### Nested Schema for `sms_sender_config`

Optional:

- `from_name` (String)
- `sender_names` (Set of String)

Read-Only:

- `id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_template_group.resource_name group_id
```

# cidaas_template (Resource)

The Template resource in the provider is used to define and manage templates within the Cidaas system. Templates are used for emails, SMS, IVR, and push notifications.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:templates_read
- cidaas:templates_write
- cidaas:templates_delete

### V2 to V3 Migration:
If you are migrating from v2 to v3, please note the changes in the format of the import identifier:

- In **v2**, the import identifier was formed by joining template_key, template_type and locale with the character `-`. For example: `TERRAFORM_TEMPLATE-SMS-en-us`.

- In **v3**, the import identifier format has been updated. The character `-` is replaced by the character `:`. For example: `TERRAFORM_TEMPLATE:SMS:en-us`.


### Managing System Templates:

- To create system templates, set the **is_system_template** flag to `true`.
By default, this value is `false` and creates custom templates when applied.
- When creating system templates validation checks are applied and suggestions are
provided in error messages to assist users in creating system templates.
- System templates cannot be imported using the standard Terraform import command. Instead, users
must create a configuration that matches the existing system template and run terraform apply.

## Example Usage

```terraform
// custom template example
resource "cidaas_template" "custom-template-1" {
  locale        = "en-in"
  template_key  = "TERRAFORM_TEMPLATE"
  template_type = "EMAIL"
  content       = "Indian sample content"
  subject       = "Email custom template subject with Indian English locale"
}

// custom template example with same template_key as custom-template-1 but different template_type and locale
resource "cidaas_template" "custom-template-2" {
  locale        = "de-de"
  template_key  = "TERRAFORM_TEMPLATE"
  template_type = "SMS"
  content       = "Sample SMS template content in German English"
}

// custom template example with same template_key and template_type as custom-template-2 but different locale
resource "cidaas_template" "custom-template-3" {
  locale        = "en-us"
  template_key  = "TERRAFORM_TEMPLATE"
  template_type = "SMS"
  content       = "Sample SMS template content in US English"
}


// System templates are created by setting the flag is_system_template to true.
// By default, this value is false and creates custom templates when applied.
// Validation checks are applied, and suggestions are provided in error messages to assist users in creating system templates.
// System templates cannot be imported using the standard Terraform import command.
// Instead, users must create a configuration that matches the existing system template and run terraform apply.

// Example of a system template for the template group "sample_group":
resource "cidaas_template" "system-template-1" {
  locale             = "en-us"
  template_key       = "VERIFY_USER"
  template_type      = "SMS"
  content            = "Hi {{name}}, here is the {{code}} to verify the user"
  is_system_template = true
  group_id           = "sample_group"
  processing_type    = "GENERAL"
  verification_type  = "SMS"
  usage_type         = "VERIFICATION_CONFIGURATION"
}

// Example of a  system template for the system default template_group "default"
resource "cidaas_template" "system-template-2" {
  locale             = "en-us"
  template_key       = "NOTIFY_COMMUNICATION_CHANGE"
  template_type      = "SMS"
  content            = "Your mobile number changed in {{account_name}}-account to {{communication_medium_value}}."
  is_system_template = true
  group_id           = "default"
  processing_type    = "GENERAL"
  usage_type         = "GENERAL"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `content` (String) The content of the template.
- `locale` (String) The locale of the template. e.g. `en-us`, `en-uk`. Ensure the locale is set in lowercase. Find the allowed locales in the Allowed Locales section below. It cannot be updated for an existing state.
- `template_key` (String) The unique name of the template. It cannot be updated for an existing state.
- `template_type` (String) The type of the template. Allowed template_types are EMAIL, SMS, IVR and PUSH. Template types are case sensitive. It cannot be updated for an existing state.

### Optional

- `group_id` (String) The `group_id` under which the configured template will be categorized. Only applicable for SYSTEM templates.
- `is_system_template` (Boolean) A boolean flag to decide between SYSTEM and CUSTOM template. When set to true the provider creates a SYSTEM template else CUSTOM
- `processing_type` (String) The processing_type attribute specifies the method by which the template information is processed and delivered. Only applicable for SYSTEM templates. It should be set to `GENERAL` when cidaas does not provide an allowed list of values.
- `subject` (String) Applicable only for template_type EMAIL. It represents the subject of an email.
- `usage_type` (String) The usage_type attribute specifies the specific use case or application for the template. Only applicable for SYSTEM templates. It should be set to `GENERAL` when cidaas does not provide an allowed list of values.
- `verification_type` (String) The verification_type attribute defines the method used for verification. Only applicable for SYSTEM templates.

### Read-Only

- `id` (String) The unique identifier of the template resource.
- `language` (String) The language based on the local provided in the configuration.
- `template_owner` (String) The template owner of the template.

## Import

Import is supported using the following syntax:

```shell
# System templates cannot be imported using the standard Terraform import command.
# Instead, users must create a configuration that matches the existing system template and run terraform apply.

# V3 Change Note: The format of the import identifier is changed in V3. In V2, the import identifier was joined by the chracter "-"
# However in V3, it is replaced by the chracter ":". Example: TERRAFORM_TEMPLATE:SMS:en-us 

# Below is the command to import a custom template
# Here, template_key:template_type:locale is a combination of template_key, template_type and locale, joined by the special character ":".
# For example, if the resource name is "sample" with template_key as "TERRAFORM_TEMPLATE", template_type as "SMS" and locale as "de-de", the import statement would be:

terraform import cidaas_template.sample TERRAFORM_TEMPLATE:SMS:de-de
```

# cidaas_user_groups (Resource)

The cidaas_user_groups resource enables the creation of user groups in the cidaas system. These groups allow users to be organized and assigned group-specific roles.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:groups_write
- cidaas:groups_read
- cidaas:groups_delete

## Example Usage

```terraform
# In the below examples, 'parent-user-group' is the top-level group, and its group_id is passed as parent_id in the 'child-user-group' resource.

resource "cidaas_user_groups" "parent-user-group" {
  group_type                     = "test_terraform"
  group_id                       = "sample-group-id"
  group_name                     = "sample-group-name"
  logo_url                       = "https://cidaas.de/logo"
  description                    = "sample parent user groups description"
  custom_fields                  = {}
  make_first_user_admin          = true
  member_profile_visibility      = "full"
  none_member_profile_visibility = "public"
}


resource "cidaas_user_groups" "child-user-group" {
  group_type  = "test_terraform"
  group_id    = "sample-child-group-id-sub"
  group_name  = "sample-child-group-name"
  logo_url    = "https://cidaas.de/logo"
  description = "sample child user groups description"
  custom_fields = {
    first_name  = "cidaas"
    family_name = "widaas"
  }
  parent_id = cidaas_user_groups.sample.group_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) Identifier for the user group.
- `group_name` (String) Name of the user group.
- `group_type` (String) Type of the user group.

### Optional

- `custom_fields` (Map of String) Custom fields for the user group.
- `description` (String) Description of the user group.
- `logo_url` (String) URL for the user group's logo
- `make_first_user_admin` (Boolean) Indicates whether the first user should be made an admin.
- `member_profile_visibility` (String) Visibility of member profiles. Allowed values `public` or `full`.
- `none_member_profile_visibility` (String) Visibility of non-member profiles. Allowed values `none` or `public`.
- `parent_id` (String) Identifier of the parent user group.

### Read-Only

- `created_at` (String) The timestamp when the resource was created.
- `id` (String) The unique identifier of the user group resource.
- `updated_at` (String) The timestamp when the resource was last updated.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_user_groups.resource_name group_id
```

# cidaas_webhook (Resource)

The Webhook resource in the provider facilitates integration of webhooks in the Cidaas system. This resource allows you to configure webhooks with different authentication options.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:webhook_read
- cidaas:webhook_write
- cidaas:webhook_delete

## Example Usage

```terraform
# This is a sample configuration for setting up a webhook with multiple authentication options.
# The available authentication types include apikey_config, totp_config, and cidaas_auth_config.

# When the auth_type is set to "APIKEY", the apikey_config is required, while totp_config and cidaas_auth_config are optional.
# These optional configurations can be removed if not needed. However, by including them, you can easily switch the auth_type 
# to other options by simply updating the auth_type value without needing to modify other parts of the configuration.

resource "cidaas_webhook" "sample_webhook" {
  auth_type = "APIKEY"
  url       = "https://cidaas.de/webhook-srv/webhook"
  events = [
    "ACCOUNT_MODIFIED"
  ]
  apikey_config = {
    key         = "api-key"
    placeholder = "test-apikey-placeholder"
    placement   = "query"
  }
  totp_config = {
    key         = "totp-key"
    placeholder = "test-totp-placeholder"
    placement   = "query"
  }
  cidaas_auth_config = {
    client_id = "ce90d6ba-9a5a-49b6-9a50-b8db759e9b90"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auth_type` (String) The attribute auth_type is to define how this url is secured from your end.The allowed values are `APIKEY`, `TOTP` and `CIDAAS_OAUTH2`
- `events` (Set of String) A set of events that trigger the webhook.
- `url` (String) The webhook url that needs to be called when an event occurs.

### Optional

- `apikey_config` (Attributes) Configuration for API key-based authentication. It's a **required** parameter when the auth_type is APIKEY. (see [below for nested schema](#nestedatt--apikey_config))
- `cidaas_auth_config` (Attributes) Configuration for Cidaas authentication. It's a **required** parameter when the auth_type is CIDAAS_OAUTH2. (see [below for nested schema](#nestedatt--cidaas_auth_config))
- `disable` (Boolean) Flag to disable the webhook.
- `totp_config` (Attributes) Configuration for TOTP based authentication.  It's a **required** parameter when the auth_type is TOTP. (see [below for nested schema](#nestedatt--totp_config))

### Read-Only

- `created_at` (String) The timestamp when the webhook was created.
- `id` (String) The unique identifier of the webhook resource.
- `updated_at` (String) The timestamp when the webhook was last updated.

<a id="nestedatt--apikey_config"></a>
### Nested Schema for `apikey_config`

Required:

- `key` (String) The API key that will be used to authenticate the webhook request.The key that will be passed in the request header or in query param as configured in the attribute `placement`
- `placeholder` (String) The attribute is the placeholder for the key which need to be passed as a query parameter or in the request header.
- `placement` (String) The placement of the API key in the request (e.g., query).The allowed value are `header` and `query`.


<a id="nestedatt--cidaas_auth_config"></a>
### Nested Schema for `cidaas_auth_config`

Required:

- `client_id` (String) The client ID for Cidaas authentication.


<a id="nestedatt--totp_config"></a>
### Nested Schema for `totp_config`

Required:

- `key` (String) The key used for TOTP authentication.
- `placeholder` (String) A placeholder value for the TOTP.
- `placement` (String) The placement of the TOTP in the request.The allowed value are `header` and `query`.

## Import

Import is supported using the following syntax:

```shell
# The import identifier in this command is the ID of the webhook to be imported.

terraform import cidaas_webhook.sample ae90d6ba-9a5a-49b6-9a50-b8db759e9b90
```

