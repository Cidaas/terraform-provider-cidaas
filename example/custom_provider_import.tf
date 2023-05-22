resource "cidaas_custom_provider" "sample" {}

output "sample_custom_provider" {
  value = cidaas_custom_provider.sample
}
