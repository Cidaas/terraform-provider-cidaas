---
page_title: "cidaas_scope Data Source - cidaas"
---

# cidaas_scope (Data Source)

The data source `cidaas_scope` returns a list of scopes available in your Cidaas instance.
You can apply filters using the `filter` block in your Terraform configuration.


## Example Usage

```terraform
data "cidaas_scope" "example" {
  filter {
    name   = "security_level"
    values = ["CONFIDENTIAL"]
  }
  filter {
    name     = "scope_key"
    values   = ["cidaas"]
    match_by = "substring"
  }
  filter {
    name   = "required_user_consent"
    values = [false]
  }
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Block Set) (see [below for nested schema](#nestedblock--filter))
- `scope` (Block List) The returned list of scopes. (see [below for nested schema](#nestedblock--scope))

### Read-Only

- `id` (String) The data source's unique ID.

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- `name` (String) The name of the attribute to filter on.
- `values` (Set of String) The value(s) to be used in the filter.

Optional:

- `match_by` (String) The type of comparison to use for this filter. Allowed values `exact`, `substring` and `regex`


<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `localized_descriptions` (Attributes List) (see [below for nested schema](#nestedatt--scope--localized_descriptions))

Read-Only:

- `group_name` (Set of String) List of scope_groups associated with the scope.
- `id` (String) The ID of the scope.
- `required_user_consent` (Boolean) Indicates whether user consent is required for the scope.
- `scope_key` (String) Unique identifier(name) for the scope.
- `scope_owner` (String) The owner of the scope. e.g. `ADMIN`.
- `security_level` (String) The security level of the scope, `PUBLIC` or `CONFIDENTIAL`.

<a id="nestedatt--scope--localized_descriptions"></a>
### Nested Schema for `scope.localized_descriptions`

Read-Only:

- `description` (String) The description of the scope in the configured locale.
- `locale` (String) The locale for the scope, e.g., `en-US`.
- `title` (String) The title of the scope in the configured locale.

## Filterable Fields

* `scope_key`
* `security_level`
* `group_name`
* `required_user_consent`
