resource "cidaas_social_provider" "sample" {
  name                     = "Sample Social Provider"
  provider_name            = "google"
  enabled                  = true
  client_id                = "8d789b3d-b312"
  client_secret            = "96ae-ea2e8d8e6708"
  scopes                   = ["profile", "email"]
  enabled_for_admin_portal = true
  claims = {
    required_claims = {
      user_info = ["name"]
      id_token  = ["phone_number"]
    }
    optional_claims = {
      user_info = ["website"]
      id_token  = ["street_address"]
    }
  }
  userinfo_fields = [
    {
      inner_key       = "sample_custom_field"
      external_key    = "external_sample_cf"
      is_custom_field = true
      is_system_field = false
    },
    {
      inner_key       = "sample_system_field"
      external_key    = "external_sample_sf"
      is_custom_field = false
      is_system_field = true
    }
  ]
}
