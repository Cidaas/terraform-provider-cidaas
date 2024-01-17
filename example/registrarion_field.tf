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
