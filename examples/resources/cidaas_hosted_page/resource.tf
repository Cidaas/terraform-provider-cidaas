resource "cidaas_hosted_page" "sample" {
  hosted_page_group_name = "terraform-sample-hosted-page"
  default_locale         = "en-IN"
  hosted_pages = [
    {
      hosted_page_id = "register_success"
      locale         = "en-US"
      url            = "https://cidaas.de/register_success_hosted_page"
      content        = "content"
    },
    {
      hosted_page_id = "register_success"
      locale         = "en-IN"
      url            = "https://tcidaas.de/register_success_hosted_page"
      content        = "content"
    }
  ]
}
