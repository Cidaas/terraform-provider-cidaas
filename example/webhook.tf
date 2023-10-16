resource "cidaas_webhook" "sample_webhook" {
  auth_type = "APIKEY"
  url       = "https://cidaas.com/webhook-test"
  events = [
    "ACCOUNT_MODIFIED"
  ]
  apikey_placeholder = "api-test-placeholder"
  apikey_placement   = "query"
  apikey             = "api-test-key"
}
