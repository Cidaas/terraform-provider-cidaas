---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cidaas_template_group Resource - cidaas"
subcategory: ""
description: |-
  The cidaas_template_group resource in the provider is used to define and manage templates groups within the cidaas system. Template Groups categorize your communication templates allowing you to map preferred templates to specific clients effectively.
  Ensure that the below scopes are assigned to the client with the specified client_id:
  cidaas:templates_readcidaas:templates_writecidaas:templates_delete
---

# cidaas_template_group (Resource)

The cidaas_template_group resource in the provider is used to define and manage templates groups within the cidaas system. Template Groups categorize your communication templates allowing you to map preferred templates to specific clients effectively.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:

* cidaas:templates_read
* cidaas:templates_write
* cidaas:templates_delete

## Example Usage

```terraform
// To create a template group, only the attribute group_id is required in the configuration.
// The attributes shown in sample-tg-2 are optional and can be configured as needed.
// If these properties are not configured in the .tf file, the provider/cidaas will compute
// and assign values to them.

// sample1
resource "cidaas_template_group" "sample-tg-1" {
  group_id = "sample_group"
}

// sample2
resource "cidaas_template_group" "sample-tg-2" {
  group_id = "group_another"
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

* `group_id` (String) The group_id of the Template Group. The group_id is used to import an existing template group. The maximum allowed length of a group_id is **15** characters.

### Optional

* `email_sender_config` (Attributes) The `email_sender_config` is used to configure your email sender. (see [below for nested schema](#nestedatt--email_sender_config))
* `ivr_sender_config` (Attributes) The configuration of the IVR sender. (see [below for nested schema](#nestedatt--ivr_sender_config))
* `push_sender_config` (Attributes) The configuration of the PUSH notification sender. (see [below for nested schema](#nestedatt--push_sender_config))
* `sms_sender_config` (Attributes) The configuration of the SMS sender. (see [below for nested schema](#nestedatt--sms_sender_config))

### Read-Only

* `id` (String) The ID of the resource

<a id="nestedatt--email_sender_config"></a>

### Nested Schema for `email_sender_config`

Optional:

* `from_email` (String) The email from address from which the emails will be sent when the specific group is configured.
* `from_name` (String) The `from_name` attribute is the display name that appears in the 'From' field of the emails.
* `reply_to` (String) The `reply_to` attribute is the email address where replies should be directed.
* `sender_names` (Set of String) The `sender_names` attribute defines the names associated with email senders.

<a id="nestedatt--ivr_sender_config"></a>

### Nested Schema for `ivr_sender_config`

Optional:

* `sender_names` (Set of String)

<a id="nestedatt--push_sender_config"></a>

### Nested Schema for `push_sender_config`

Optional:

* `sender_names` (Set of String)

<a id="nestedatt--sms_sender_config"></a>

### Nested Schema for `sms_sender_config`

Optional:

* `from_name` (String)
* `sender_names` (Set of String)

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_template_group.resource_name group_id
```