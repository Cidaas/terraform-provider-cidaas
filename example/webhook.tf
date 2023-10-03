resource "cidaas_webhook" "sample_webhook" {
  auth_type = "APIKEY"
  url       = "https://cidaas.com/webhook-test"
  events = [
    "ACCOUNT_MODIFIED"
  ]
  api_key_details = {
    apikey_placeholder = "apikey"
    apikey_placement   = "header"
    apikey             = "test-key"
  }
}
