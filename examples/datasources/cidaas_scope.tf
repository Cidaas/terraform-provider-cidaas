data "cidaas_scope" "example" {
  filter {
    name   = "security_level"
    values = ["CONFIDENTIAL"]
  }
  filter {
    name     = "scope_key"
    values   = ["cidaas"]
    match_by = "substring"
  }
  filter {
    name   = "required_user_consent"
    values = [false]
  }
}
