data "cidaas_custom_provider" "example" {
  filter {
    name     = "provider_name"
    values   = ["dev"]
    match_by = "substring"
  }
}
