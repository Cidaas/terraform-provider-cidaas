resource "cidaas_app" "sample" {
  client_name                     = "Test Terraform Application" // unique
  client_type                     = "SINGLE_PAGE"
  accent_color                    = "#ef4923"                        // Default: #ef4923
  primary_color                   = "#ef4923"                        // Default: #f7941d
  media_type                      = "IMAGE"                          // Default: IMAGE
  allow_login_with                = ["EMAIL", "MOBILE", "USER_NAME"] // Default: ["EMAIL", "MOBILE", "USER_NAME"]
  redirect_uris                   = ["https://cidaas.com"]
  allowed_logout_urls             = ["https://cidaas.com"]
  enable_deduplication            = true      // Default: false
  auto_login_after_register       = true      // Default: false
  enable_passwordless_auth        = false     // Default: true
  register_with_login_information = false     // Default: false
  hosted_page_group               = "default" // Default: default
  company_name                    = "Widas ID GmbH"
  company_address                 = "01"
  company_website                 = "https://cidaas.com"
  allowed_scopes                  = ["openid", "cidaas:register", "profile"]
  client_display_name             = "Display Name of the app" // unique
  content_align                   = "CENTER"                  // Default: CENTER
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
  enable_bot_detection                = false // Default: false
  allow_guest_login_groups = [{
    group_id      = "developer101"
    roles         = ["developer", "qa", "admin"]
    default_roles = ["developer"]
  }]
  is_login_success_page_enabled    = false // Default: false
  is_register_success_page_enabled = false // Default: false
  group_ids                        = ["sample"]
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
}
