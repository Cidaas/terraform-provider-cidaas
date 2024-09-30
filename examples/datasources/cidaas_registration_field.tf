data "cidaas_registration_field" "example" {
  filter {
    name   = "field_type"
    values = ["CUSTOM"]
  }
}
