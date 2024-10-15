data "cidaas_scope_group" "example" {
  filter {
    name     = "group_name"
    values   = ["terraform"]
    match_by = "substring"
  }
}
