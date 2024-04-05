resource "cidaas_app" "sample" {
  # To prevent unintended updates to specific fields, incorporate the following lifecycle block. E.g. client_id will be ignored here
  lifecycle {
    ignore_changes = [
      client_id
    ]
  }

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
  client_name                     = "Terraform"
  client_display_name             = "Sample Terraform App"
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
  app_owner    = "CLIENT"
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
  description                     = "it's a sample description of the client. The client supports system to system communication"
  default_scopes                  = ["sample"]
  pending_scopes                  = ["sample"]
  consent_page_group              = "sample-consent-page-group"
  password_policy_ref             = "password-policy-ref"
  blocking_mechanism_ref          = "blocking-mechanism-ref"
  sub                             = "sample-sub"
  role                            = "sample-role"
  mfa_configuration               = "sample-configuration"
  suggest_mfa                     = ["OFF"]
  login_spi {
    oauth_client_id = "bcb-4a6b-9777-8a64abe6af"
    spi_url         = "https://cidaas.com/spi-url"
  }
  video_url       = "https://cidaas.com/video-url"
  bot_captcha_ref = "sample-bot-captcha-ref"
  background_uri  = "http://cidaas.com/background-uri"
  application_meta_data = {
    status : "active"
    version : "1.0.0"
  }
}
