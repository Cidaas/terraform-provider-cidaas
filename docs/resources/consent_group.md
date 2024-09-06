---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cidaas_consent_group Resource - cidaas"
subcategory: ""
description: |-
  The Consent Group resource in the provider allows you to define and manage consent groups in Cidaas.
  Consent Groups are useful to organize and manage consents by grouping related consent items together.
  Ensure that the below scopes are assigned to the client with the specified client_id:
  cidaas:tenant_consent_readcidaas:tenant_consent_writecidaas:tenant_consent_delete
---

# cidaas_consent_group (Resource)

The Consent Group resource in the provider allows you to define and manage consent groups in Cidaas.
 Consent Groups are useful to organize and manage consents by grouping related consent items together.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:tenant_consent_read
- cidaas:tenant_consent_write
- cidaas:tenant_consent_delete

## Example Usage

```terraform
resource "cidaas_consent_group" "sample" {
  group_name  = "sample_consent_group"
  description = "sample description"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_name` (String) The name of the consent group.

### Optional

- `description` (String) Description of the consent group.

### Read-Only

- `created_at` (String) The timestamp when the consent group was created.
- `id` (String) The unique identifier of the consent group.
- `updated_at` (String) The timestamp when the consent group was last updated.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_consent_group.sample id
```