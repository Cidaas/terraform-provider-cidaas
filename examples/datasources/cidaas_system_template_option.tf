data "cidaas_system_template_option" "example" {
  filter {
    name   = "template_key"
    values = ["UN_REGISTER_USER_ALERT"]
  }
  filter {
    name   = "role"
    values = ["USER_CREATE"]
  }
}
