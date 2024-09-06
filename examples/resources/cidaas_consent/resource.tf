resource "cidaas_consent" "sample" {
  consent_group_id = cidaas_consent_group.sample.id
  name             = "sample_consent"
  enabled          = true # By default enabled is set to 'true'
}
