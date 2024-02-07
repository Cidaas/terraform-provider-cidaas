resource "cidaas_template" "sample" {
  locale        = "en-us"
  template_key  = "TERRAFORM_TEST"
  template_type = "SMS"
  content       = "sample content for resource cidaas template"
}
