---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cidaas_role Resource - cidaas"
subcategory: ""
description: |-
  The cidaas_role resource in Terraform facilitates the management of roles in cidaas system. This resource allows you to configure and define custom roles to suit your application's specific access control requirements.
  Ensure that the below scopes are assigned to the client with the specified client_id:
  cidaas:roles_readcidaas:roles_writecidaas:roles_delete
---

# cidaas_role (Resource)

The cidaas_role resource in Terraform facilitates the management of roles in cidaas system. This resource allows you to configure and define custom roles to suit your application's specific access control requirements.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:

* cidaas:roles_read
* cidaas:roles_write
* cidaas:roles_delete

## Example Usage

```terraform
resource "cidaas_role" "sample" {
  role        = "terraform_sample_role"
  name        = "Terraform Sample Role"
  description = "The sample is designed to demonstrate the configuration of the terraform cidaas_role resource."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

* `role` (String) The unique identifier of the role. The role name must be unique across the cidaas system and cannot be updated for an existing state.

### Optional

* `description` (String) The `description` attribute provides details about the role, explaining its purpose.
* `name` (String) The name of the role.

### Read-Only

* `id` (String) The ID of the role resource.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_role.resource_name role
```
