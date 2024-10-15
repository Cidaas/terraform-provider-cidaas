data "cidaas_role" "example" {
  filter {
    name   = "name"
    values = ["DEVELOPER"]
  }
}
