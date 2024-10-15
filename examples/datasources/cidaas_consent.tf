data "cidaas_consent" "example" {
  filter {
    name     = "consent_name"
    values   = ["terraform"]
    match_by = "substring"
  }
}
