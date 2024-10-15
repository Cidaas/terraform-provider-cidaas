data "cidaas_group_type" "example" {
  filter {
    name   = "role_mode"
    values = ["roles_required"]
  }
  filter {
    name   = "allowed_roles"
    values = ["DEVELOPER"]
  }
}
