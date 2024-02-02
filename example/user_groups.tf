resource "cidaas_user_groups" "sample" {
  group_type            = "sample-group-type"
  group_id              = "sample-group-id"
  group_name            = "sample-group-name"
  logo_url              = "https://cidaas.de/logo"
  description           = "sample user groups description"
  make_first_user_admin = false
  custom_fields = {
    custom_field_name = "sample custom field"
  }
  member_profile_visibility      = "full"
  none_member_profile_visibility = "public"
  parent_id                      = "sample-parent-id"
}
