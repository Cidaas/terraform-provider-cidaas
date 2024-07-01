// To create a template group, only the attribute group_id is required in the configuration.
// The attributes shown in sample-tg-2 are optional and can be configured as needed.
// If these properties are not configured in the .tf file, the provider/cidaas will compute
// and assign values to them.

// sample1
resource "cidaas_template_group" "sample-tg-1" {
  group_id = "sample_group_two"
}

// sample2
resource "cidaas_template_group" "sample-tg-2" {
  group_id = "sample_group"
  email_sender_config = {
    from_email = "noreply@cidaas.de"
    from_name  = "Kube-dev"
    reply_to   = "noreply@cidaas.de"
    sender_names = [
      "SYSTEM",
    ]
  }
  ivr_sender_config = {
    sender_names = [
      "SYSTEM",
    ]
  }
  push_sender_config = {
    sender_names = [
      "SYSTEM",
    ]
  }
  sms_sender_config = {
    sender_names = [
      "SYSTEM",
    ]
  }
}
