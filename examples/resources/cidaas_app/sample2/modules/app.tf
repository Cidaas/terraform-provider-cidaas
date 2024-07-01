variable "common_config" {
  type = object({
    client_type                       = string
    accent_color                      = string      // Default: #ef4923
    primary_color                     = string      // Default: #f7941d
    media_type                        = string      // Default: IMAGE
    allow_login_with                  = set(string) // Default: ["EMAIL", "MOBILE", "USER_NAME"]
    redirect_uris                     = set(string)
    allowed_logout_urls               = set(string)
    enable_deduplication              = bool   // Default: false
    auto_login_after_register         = bool   // Default: false
    enable_passwordless_auth          = bool   // Default: true
    register_with_login_information   = bool   // Default: false
    fds_enabled                       = bool   // Default: true
    hosted_page_group                 = string // Default: default
    company_name                      = string
    company_address                   = string
    company_website                   = string
    allowed_scopes                    = set(string)
    response_types                    = set(string) // Default: ["code", "token", "id_token"]
    grant_types                       = set(string) // Default: ["implicit", "authorization_code", "password", "refresh_token"]
    login_providers                   = set(string)
    is_hybrid_app                     = bool // Default: false
    allowed_web_origins               = set(string)
    allowed_origins                   = set(string)
    default_max_age                   = number // Default: 86400
    token_lifetime_in_seconds         = number // Default: 86400
    id_token_lifetime_in_seconds      = number // Default: 86400
    refresh_token_lifetime_in_seconds = number // Default: 15780000
    template_group_id                 = string // Default: default
    editable                          = bool   // Default: true
    social_providers = list(object({
      provider_name = string
      social_id     = string
      display_name  = string
    }))
    custom_providers = list(object({
      logo_url      = string
      provider_name = string
      display_name  = string
      type          = string
    }))
    saml_providers = list(object({
      logo_url      = string
      provider_name = string
      display_name  = string
      type          = string
    }))
    ad_providers = list(object({
      logo_url      = string
      provider_name = string
      display_name  = string
      type          = string
    }))
    jwe_enabled  = bool // Default: false
    user_consent = bool // Default: false
    allowed_groups = list(object({
      group_id      = string
      roles         = set(string)
      default_roles = set(string)
    }))
    operations_allowed_groups = list(object({
      group_id      = string
      roles         = set(string)
      default_roles = set(string)
    }))
    enabled                     = bool // Default: true
    always_ask_mfa              = bool // Default: false
    allowed_mfa                 = set(string)
    email_verification_required = bool // Default: true
    allowed_roles               = set(string)
    default_roles               = set(string)
    enable_classical_provider   = bool   // Default: true
    is_remember_me_selected     = bool   // Default: true
    bot_provider                = string // Default: CIDAAS
    allow_guest_login           = bool   // Default: false
    #  mfa Default:
    #  {
    #   setting = "OFF"
    #  }
    mfa = object({
      setting = string
    })
    webfinger      = string
    default_scopes = set(string)
    pending_scopes = set(string)
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
// For string attributes please don't provide value as empty string otherwise you will end up getting an error like below
# When applying changes to cidaas_app.sample, provider
# "provider[\"hashicorp.com/cidaas/cidaas\"]" produced an unexpected new value:
# .consent_page_group: was cty.StringVal(""), but now null.

# This is a bug in the provider, which should be reported in the provider's own issue
# tracker.

# For an attribute with default value, if not provided in the config then default will be considered while creating the app
# The default values of the attributes are shared next to the attributes


# The config below has the list of common config and main config
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
  group_types = ["sample"]
  // api throws validation failed invalid backchanel logout url for an invalid url
  # backchannel_logout_uri              = "https://cidaas.com"
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
  common_configs = var.common_config
}
