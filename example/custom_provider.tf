resource "cidaas_custom_provider" "sample" {
  authorization_endpoint = "https://terraform-cidaas-test-free.cidaas.de/authz"
  client_id              = "client_id"
  client_secret          = "client_secret"
  display_name           = "terraform"
  logo_url               = "https://avatars3.githubusercontent.com/Cidaas"
  provider_name          = "terraform"
  standard_type          = "OAUTH2"
  token_endpoint         = "https://terraform-cidaas-test-free.cidaas.de/token"
  userinfo_endpoint      = "https://terraform-cidaas-test-free.cidaas.de/userinfo"
  scope_display_label    = "scope_d_name"

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
  }
}
