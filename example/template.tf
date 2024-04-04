resource "cidaas_template" "sample" {
  locale        = "en-in"
  template_key  = "TERRAFORM_TEMPLATE"
  template_type = "SMS"
  content       = "India English sample content"
}

resource "cidaas_template" "sample2" {
  locale        = "de-de"
  template_key  = "TERRAFORM_TEMPLATE"
  template_type = "SMS"
  content       = "German sample content"
}

resource "cidaas_template" "sample3" {
  locale        = "en-us"
  template_key  = "TERRAFORM_TEMPLATE"
  template_type = "SMS"
  content       = "US English sample content"
}
