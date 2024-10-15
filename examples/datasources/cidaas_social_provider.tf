data "cidaas_social_provider" "example" {
  filter {
    name   = "enabled_for_admin_portal"
    values = ["true"]
  }
  filter {
    name   = "enabled"
    values = ["true"]
  }
}
