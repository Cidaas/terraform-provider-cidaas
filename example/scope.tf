resource "cidaas_scope" "sample" {
  locale                = "en-US"
  language              = "en-US"
  description           = "terraform description"
  title                 = "terraform title"
  security_level        = "PUBLIC"
  scope_key             = "terraform-test-scope"
  required_user_consent = false
  group_name            = []
}
