
### Changelog

### Resource `Role`

No Change

### Resource `Custom Provider`

- The attribute scopes now has to be set as an array of map/object instead of configuring separate map.object

#### old configuration:
```
resource "cidaas_custom_provider" "sample" {
  ...
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
}
```
#### new configuration:
```
resource "cidaas_custom_provider" "sample" {
  ...
  scopes [{
    recommended = false
    required    = false
    scope_name  = "openid"
  },
  {
    recommended = false
    required    = false
    scope_name  = "profile"
  },
  {
    recommended = false
    required    = false
    scope_name  = "email"
  }]
}
```

- A new attribute domains added which is an array of domains. example: domains = ["https://demo.cidaas.de","https://accounts.widas.de"]
- validations added to validate the url attributes(authorization_endpoint,token_endpoint...). checks if it has `https://`
- Now custom_fields in userinfo_fields can be passed as map/object simply as below instead of providing key value pair in array

#### old configuration:

```
userinfo_fields = {
  custom_fields = [
    {
      key   = "key1"
      value = "value1"
    },
    {
      key   = "key2"
      value = "value2"
    }
  ]
}
```

#### new configuration:
```
userinfo_fields = {
  custom_fields = {
    key1 = "value1"
    key2 = "value2"
  }
}
```

### Sample

```
resource "cidaas_custom_provider" "sample" {
  standard_type          = "OAUTH2"
  authorization_endpoint = "https://terraform-cidaas-test-free.cidaas.de/authz-srv/authz"
  token_endpoint         = "https://terraform-cidaas-test-free.cidaas.de/token-srv/token"
  provider_name          = "terraform-test1000"
  display_name           = "Terraform"
  logo_url               = "https://terraform-cidaas-test-free.cidaas.de/logo"
  userinfo_endpoint      = "https://terraform-cidaas-test-free.cidaas.de/users-srv/userinfo"
  scope_display_label    = "Terraform Test Scope"
  client_id              = "test1"
  client_secret          = "test-secret"
  domains                = ["https://demo.cidaas.de","https://accounts.widas.de"]

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
    custom_fields = {
      pincode = "123456"
    }
  }
}

```

### Resource `Scope`

- The `locale`, `language`, `title`, and `description` attributes have been removed and replaced with a `localized_descriptions` block that supports a scope with multiple locale with better internationalization. Earlier only one locale was supported by the terraform plugin.

- `localized_descriptions` is a list of objects, each containing:
  - locale
  - title
  - description

- The `language` attribute is no longer required. The provider computes and assigns the language based on the `locale` provided.

#### old configuration:
```
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

#### new configuration:
```
resource "scope" "sample" {
  security_level        = "CONFIDENTIAL"
  scope_key             = "terraform-test-scope-100"
  required_user_consent = false
  group_name            = []
  localized_descriptions = [
    {
      title       = "terraform tunisia title"
      locale      = "ar-TN"
      description = "terraform description"
    },
    {
      title       = "terraform hu-HU title"
      locale      = "hu-HU"
      description = "terraform description"
    },
    {
      title       = "terraform english india title"
      locale      = "en-IN"
      description = "terraform english india description"
    }
  ]
}
```

- The `locale` attribute default is `en-US` if not provided.
- The `required_user_consent` attribute default is `false` if not provided.
- The `security_level` attribute default is `PUBLIC` if not provided.

**Migration Steps:**

- Update your Terraform configuration file to replace the `locale`, `title`, and `description` attributes with the new `localized_descriptions` block.
- Ensure the attributes with default values are provided if the existing configuration has different values for those attributes


### Resource `Scope Group`

No Change

### Resource `App`

what changed?

What's new?
- attribute ``common_config`` added to support inheritence in resource app. The supported attributes in common config are shared below.
- Utilize `common_config`s only when there are shared configuration attributes for multiple resources.
- If you need to override any specific attribute for a particular resource where the same attribute is available in `common_config`, you can supply the main configuration attribute directly within the resource block.
- If your configuration involves a single resource or if the common configuration attributes are not shared across multiple resources we do not suggest using `common_config`s.


Attributes not supported in app config anymore:

- client_secret_expires_at
- client_id_issued_at
- push_config
- created_at
- updated_at
- admin_client
- deleted
- app_owner


List of attributes supported in `common_config`s

- client_type
- accent_color
- primary_color
- media_type
- allow_login_with
- redirect_uris
- allowed_logout_urls
- enable_deduplication
- auto_login_after_register
- enable_passwordless_auth
- register_with_login_information
- fds_enabled
- hosted_page_group
- company_name
- company_address
- company_website
- allowed_scopes
- response_types
- grant_types
- login_providers
- is_hybrid_app
- allowed_web_origins
- allowed_origins
- default_max_age
- token_lifetime_in_seconds
- id_token_lifetime_in_seconds
- refresh_token_lifetime_in_seconds
- template_group_id
- editable
- social_providers
- custom_providers
- saml_providers
- ad_providers
- jwe_enabled
- user_consent
- allowed_groups
- operations_allowed_groups
- enabled
- always_ask_mfa
- allowed_mfa
- email_verification_required
- allowed_roles
- default_roles
- enable_classical_provider
- is_remember_me_selected
- bot_provider
- allow_guest_login
- mfa
- webfinger
- default_scopes
- pending_scopes


#### change in data types of some attribute

 - social_providers
 - custom_providers
 - saml_providers
 - ad_providers

 The above attributes now has to be provided as set of objects.

 example:
 ```
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
 ```

### app resource highlights

The resource app now can be setup with very less config attribute. The below are the only required paraemeter to create an app

- client_name // can not be part of `common_config`
- client_type
- company_name                   
- company_address                
- company_website                 
- allowed_scopes                
- redirect_uris
- allowed_logout_urls

The provider created an app with other attribute with it's default values wherever applicable.

The list of default attibutes below with it's default:


| Attribute                         | Value                                        |
|-----------------------------------|----------------------------------------------|
| accent_color                      | "#ef4923"                                    |
| primary_color                     | "#f7941d"                                    |
| media_type                        | "IMAGE"                                      |
| allow_login_with                  | ["EMAIL", "MOBILE", "USER_NAME"]             |
| enable_deduplication              | false                                        |
| auto_login_after_register         | false                                        |
| enable_passwordless_auth          | true                                         |
| register_with_login_information   | false                                        |
| fds_enabled                       | true                                         |
| hosted_page_group                 | "default"                                    |
| response_types                    | ["code", "token", "id_token"]                |
| grant_types                       | ["implicit", "authorization_code", "password", "refresh_token"] |
| is_hybrid_app                     | false                                        |
| default_max_age                   | 86400                                        |
| token_lifetime_in_seconds         | 86400                                        |
| id_token_lifetime_in_seconds      | 86400                                        |
| refresh_token_lifetime_in_seconds | 15780000                                     |
| template_group_id                 | "default"                                    |
| editable                          | true                                         |
| jwe_enabled                       | false                                        |
| user_consent                      | false                                        |
| enabled                           | true                                         |
| always_ask_mfa                    | false                                        |
| email_verification_required       | true                                         |
| enable_classical_provider         | true                                         |
| is_remember_me_selected           | true                                         |
| bot_provider                      | "CIDAAS"                                     |
| allow_guest_login                 | false                                        |
| mfa                               | { setting = "OFF" }                          |