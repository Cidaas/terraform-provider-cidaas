# cidaas_consent_version sample for consent_type "SCOPES"
resource "cidaas_consent_version" "v1" {
  version         = 1
  consent_id      = cidaas_consent.sample.id
  consent_type    = "SCOPES"
  scopes          = ["developer"]
  required_fields = ["name"]
  consent_locales = [
    {
      content = "consent version in German"
      locale  = "de"
    },
    {
      content = "consent version in English"
      locale  = "en"
    }
  ]
}

# cidaas_consent_version sample for consent_type "URL"
resource "cidaas_consent_version" "v2" {
  version      = 2
  consent_id   = cidaas_consent.sample.id
  consent_type = "URL"
  consent_locales = [
    {
      content = "consent version in German"
      locale  = "de"
      url     = "https://cidaas.de/de"
    },
    {
      content = "consent version in English"
      locale  = "en"
      url     = "https://cidaas.de/en"
    }
  ]
}
