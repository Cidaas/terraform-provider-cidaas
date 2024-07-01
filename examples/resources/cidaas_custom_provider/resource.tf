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
