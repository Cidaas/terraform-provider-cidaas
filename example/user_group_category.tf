resource "cidaas_user_group_category" "sample" {
  role_mode     = "no_roles"
  group_type    = "TerraformUserGroupCategory"
  description   = "terraform user group category description"
  allowed_roles = []
}
