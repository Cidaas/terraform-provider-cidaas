# The data types "TEXT", "TEXTAREA", "URL", "PASSWORD", "EMAIL"
# "MOBILE", "JSON_STRING", "USERNAME", and "ARRAY" are configured as shown below.

resource "cidaas_registration_field" "text" {
  data_type                                      = "TEXT"
  field_key                                      = "sample_text_field"
  field_type                                     = "CUSTOM"
  internal                                       = true
  required                                       = true
  read_only                                      = true
  is_group                                       = false
  unique                                         = true
  overwrite_with_null_value_from_social_provider = false
  is_searchable                                  = true
  enabled                                        = true
  claimable                                      = true
  order                                          = 1
  parent_group_id                                = "DEFAULT"
  scopes                                         = ["profile"]
  local_texts = [
    {
      locale         = "en-US"
      name           = "Sample Field"
      required_msg   = "The field is required"
      max_length_msg = "Maximum 99 chars allowed"
      min_length_msg = "Minimum 99 chars allowed"
    },
    {
      locale         = "de-DE"
      name           = "Beispielfeld"
      required_msg   = "Dieses Feld ist erforderlich"
      max_length_msg = "DE maximum 99 chars allowed"
      min_length_msg = "DE minimum 10 chars allowed"
    }
  ]
  field_definition = {
    max_length = 100
    min_length = 10
  }
}
