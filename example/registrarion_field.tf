resource "cidaas_registration_page_field" "sample" {
  claimable            = true
  data_type            = "TEXT"
  enabled              = false
  field_key            = "given_name_test_v21"
  field_type           = "CUSTOM"
  internal             = false
  is_group             = false
  locale_text_language = "en"
  locale_text_locale   = "en-us"
  locale_text_name     = "Given Name TEST v21"
  order                = 2
  parent_group_id      = "DEFAULT"
  read_only            = false
  required             = true
  scopes = [
    "profile",
    "cidaas:public_profile",
  ]
}
