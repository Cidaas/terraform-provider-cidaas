---
page_title: "cidaas_app Resource - cidaas"
subcategory: ""
description: |-
  The App resource allows creation and management of clients in Cidaas system. When creating a client with a custom client_id and client_secret you can include the configuration in the resource. If not provided, Cidaas will generate a set for you. client_secret is sensitive data. Refer to the article Terraform Sensitive Variables https://developer.hashicorp.com/terraform/tutorials/configuration-language/sensitive-variables to properly handle sensitive information.
  Ensure that the below scopes are assigned to the client with the specified client_id:
  cidaas:apps_readcidaas:apps_writecidaas:apps_delete
---

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
- application_type

### Change in data types of some attributes

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
        domains = ["cidaas.de"]
    },
    {
        logo_url      = "https://cidaas.com/logo-url"
        provider_name = "sample-custom-provider"
        display_name  = "sample-custom-provider"
        type          = "CUSTOM_OPENID_CONNECT"
        is_provider_visible = true
        domains = ["cidaas.de"]
    },
  ]
 }
 ```
### Handling schema change error for existing state
If you encounter the following error message when the below specified attributes are present in the state, please follow the steps to fix the error:

```shell
Error: Unable to Read Previously Saved State for UpgradeResourceState
...
There was an error reading the saved resource state using the current resource schema.
...
AttributeName("group_selection"): invalid JSON, expected "{", got "["
```

#### Affected Attributes:
- group_selection
- login_spi
- mfa
- mobile_settings

To resolve this issue, manually update the Terraform state file by following these steps:

1. Open the state file (`terraform.tfstate`) and locate the `cidaas_app.<resource_name_in_your_config>` resource.
2. Search for the affected attributes listed above.
3. Update their types to JSON objects. Ensure they are set as objects (`{}`) and not arrays (`[]`).

##### Example:

Before:
```json
"group_selection": [
  {
    "selectable_groups" : ["developer-users"]
    "selectable_group_types" : ["sample"]
    "always_show_group_selection" : null
  }
]
```

After:
```json
"group_selection": {
  "selectable_groups" : ["developer-users"]
  "selectable_group_types" : ["sample"]
  "always_show_group_selection" : null
}
```

Alternatively, you can resolve the issue by deleting the existing state of the specific resource and importing it from Cidaas.
However, this approach can be risky, so please proceed with caution.
Ensure you only delete the specific resource from the state file that is causing the error, not the entire file or any other resources.

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
- `additional_access_token_payload` (Set of String) Access token payload definition.
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

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)
- `provider_name` (String)
- `type` (String)


<a id="nestedatt--allow_guest_login_groups"></a>
### Nested Schema for `allow_guest_login_groups`

Optional:

- `default_roles` (Set of String)
- `group_id` (String)
- `roles` (Set of String)


<a id="nestedatt--allowed_groups"></a>
### Nested Schema for `allowed_groups`

Optional:

- `default_roles` (Set of String)
- `group_id` (String)
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

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)
- `provider_name` (String)
- `type` (String)


<a id="nestedatt--common_configs--allowed_groups"></a>
### Nested Schema for `common_configs.allowed_groups`

Optional:

- `default_roles` (Set of String)
- `group_id` (String)
- `roles` (Set of String)


<a id="nestedatt--common_configs--custom_providers"></a>
### Nested Schema for `common_configs.custom_providers`

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)
- `provider_name` (String)
- `type` (String)


<a id="nestedatt--common_configs--mfa"></a>
### Nested Schema for `common_configs.mfa`

Optional:

- `allowed_methods` (Set of String)
- `setting` (String)
- `time_interval_in_seconds` (Number)


<a id="nestedatt--common_configs--operations_allowed_groups"></a>
### Nested Schema for `common_configs.operations_allowed_groups`

Optional:

- `default_roles` (Set of String)
- `group_id` (String)
- `roles` (Set of String)


<a id="nestedatt--common_configs--saml_providers"></a>
### Nested Schema for `common_configs.saml_providers`

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)
- `provider_name` (String)
- `type` (String)


<a id="nestedatt--common_configs--social_providers"></a>
### Nested Schema for `common_configs.social_providers`

Optional:

- `provider_name` (String)
- `social_id` (String)



<a id="nestedatt--custom_providers"></a>
### Nested Schema for `custom_providers`

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)
- `provider_name` (String)
- `type` (String)


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

Optional:

- `default_roles` (Set of String)
- `group_id` (String)
- `roles` (Set of String)


<a id="nestedatt--saml_providers"></a>
### Nested Schema for `saml_providers`

Optional:

- `display_name` (String)
- `domains` (Set of String)
- `is_provider_visible` (Boolean)
- `logo_url` (String)
- `provider_name` (String)
- `type` (String)


<a id="nestedatt--social_providers"></a>
### Nested Schema for `social_providers`

Optional:

- `provider_name` (String)
- `social_id` (String)

## Import

Import is supported using the following syntax:

```shell
# The import identifier in this command is the client_id of the app to be imported.

terraform import cidaas_app.sample client_id
```