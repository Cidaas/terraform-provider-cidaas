resource "cidaas_group_type" "sample" {
  role_mode     = "no_roles"
  group_type    = "TerraformSampleGroupType"
  description   = "terraform user group category description"
  allowed_roles = ["developer"]
}
