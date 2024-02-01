resource "cidaas_hosted_page" "sample" {
  hosted_page_group_name = "hosted-page-sample-group"
  default_locale         = "en-US"

  hosted_pages {
    hosted_page_id = "register_success"
    locale         = "en-US"
    url            = "https://terraform-cidaas-test-free.cidaas.de/register_success_hosted_page"
  }

  hosted_pages {
    hosted_page_id = "login_success"
    locale         = "en-US"
    url            = "https://terraform-cidaas-test-free.cidaas.de/login_success_hosted_page"
  }
}
