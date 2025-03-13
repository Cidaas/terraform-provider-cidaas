---
page_title: "cidaas_registration_field Resource - cidaas"
subcategory: ""
description: |-
  The cidaas_registration_field in the provider allows management of registration fields in the Cidaas system. This resource enables you to configure and customize the fields displayed during user registration.
  Ensure that the below scopes are assigned to the client with the specified client_id:
  cidaas:field_setup_readcidaas:field_setup_writecidaas:field_setup_delete
---

# cidaas_registration_field (Resource)

The `cidaas_registration_field` in the provider allows management of registration fields in the Cidaas system. This resource enables you to configure and customize the fields displayed during user registration.

 Ensure that the below scopes are assigned to the client with the specified `client_id`:
- cidaas:field_setup_read
- cidaas:field_setup_write
- cidaas:field_setup_delete


### V2 to V3 Migration:
If you are migrating from v2 to v3, please note that `cidaas_registration_page_field` has been renamed to `cidaas_registration_field`. Below is the list of changes in `cidaas_registration_field`:

- **Multiple Locales:** Internationalization now supported via `local_texts`.
- **Field Definition Attributes:** Added to specify maximum and minimum lengths for `TEXT` and `DATE` attributes.
- **Extended Datatypes:** It now supports the datatypes `TEXT`, `NUMBER`, `SELECT`, `MULTISELECT`, `RADIO`, `CHECKBOX`, `PASSWORD`, `DATE`, `URL`, `EMAIL`,
	`TEXTAREA`, `MOBILE`, `CONSENT`, `JSON_STRING`, `USERNAME`, `ARRAY`, `GROUPING` and `DAYDATE`.
- **Configuration Updates:** Adjustments required for the below attribute as shown below:

| old config           |      new config                                    |
|---------------------| ----------------------------------------------|
| required_msg | local_texts[i].required_msg |
| locale_text_min_length | field_defination.min_length |
| locale_text_max_length | field_defination.max_length |
| min_length_error_msg | local_texts[i].min_length_msg |
| max_length_error_msg | local_texts[i].max_length_msg |
| locale_text_language | The `language` attribute is no longer required. The provider computes and assigns the language based on the `locale` provided. |
| locale_text_locale | local_texts[i].locale |
| locale_text_name | local_texts[i].name |


#### Attribute `local_text`

| attributes         |      description                             |
|---------------------| ----------------------------------------------|
| locale | The locale of the field. example: de-DE |
| name |The name of the field in the local configured. for example: in **en-US** the name is `Sample Field` in de-DE `Beispielfeld`|
| max_length_msg | warning/error msg to show to the user when user exceeds the maximum character configured. This is applicable only for the attributes of base_data_type string |
| min_length_msg | warning/error msg to show to the user when user don't provide the minimum character required. This is applicable only for the attributes of base_data_type string |
| required_msg | When the flag required is set to true the required_msg must be provided. required_msg is shown if user does not provide a required field |
| attributes | The field attributes must be provided for the data_type SELECT, MULTISELECT and RADIO. it's an array of key value pairs. example shown below |
| consent_label | required when data_type is CONSENT. exmaple shown below |


### Example of `attributes`
```terraform
 local_texts = [
    {
      locale       = "en-US"
      name         = "Sample Field"
      required_msg = "The field is required"
      attributes = [
        {
          key   = "test_key"
          value = "test_value"
        }
      ]
    }
 ]
```

### Example of `consent_label`
```terraform
local_texts = [
    {
      locale       = "en-US"
      name         = "sample_consent_field"
      required_msg = "The field is required"
      consent_label = {
        label      = "test",
        label_text = "test label text"
      }
    }
  ]
```

### Attribute `field_definition`

| attributes         |      description                             |
|---------------------| ----------------------------------------------|
| max_length | The maximum length of a string type attribute |
| min_length |The minimum length of a string type attribute|
| min_date | applicable only for DATE attribute. example: "2024-06-28T18:30:00Z" |
| max_date | applicable only for DATE attribute. example: "2024-06-28T18:30:00Z" |
| initial_date_view | applicable only for DATE attribute. Allowed values: month, year and multi-year |
| initial_date | applicable only for DATE attribute. example: "2024-06-28T18:30:00Z" |

Ensure your Terraform configurations are updated accordingly to maintain compatibility with the latest version.

## Old configuration:
```terraform
resource "cidaas_registration_page_field" "sample" {
  claimable              = true
  data_type              = "TEXT"
  enabled                = false
  field_key              = "sample_field"
  field_type             = "CUSTOM"
  internal               = false
  is_group               = false
  locale_text_language   = "en"
  locale_text_locale     = "en-us"
  locale_text_name       = "Sample Field"
  order                  = 2
  parent_group_id        = "DEFAULT"
  read_only              = false
  required               = true
  required_msg           = "sample_field is required"
  locale_text_min_length = 10
  locale_text_max_length = 100
  min_length_error_msg   = "minimum length should be 10"
  max_length_error_msg   = "maximum length should be 100"
  scopes = [
    "profile",
  ]
  overwrite_with_null_value_from_social_provider = false
}
```

## Example Usage(V3 configuration)

```terraform
resource "cidaas_registration_field" "text" {
  data_type                                      = "TEXT"
  field_key                                      = "sample_text_field"
  field_type                                     = "CUSTOM"  // CUSTOM and SYSTEM, SYSTEM can not be created but modified
  internal                                       = true      // Default: false
  required                                       = true      // Default: false
  read_only                                      = true      // Default: false
  is_group                                       = false     // Default: false
  unique                                         = true      // Default: false
  overwrite_with_null_value_from_social_provider = false     // Default: true
  is_searchable                                  = true      // Default: true
  enabled                                        = true      // Default: true
  claimable                                      = true      // Default: true
  order                                          = 1         // Default: 1
  parent_group_id                                = "DEFAULT" // Default: DEFAULT
  scopes                                         = ["profile"]
  local_texts = [
    {
      locale         = "en-US"
      name           = "Sample Field"
      required_msg   = "The field is required"
      max_length_msg = "Maximum 99 chars allowed"
      min_length_msg = "Minimum 99 chars allowed"
    },
    {
      locale         = "de-DE"
      name           = "Beispielfeld"
      required_msg   = "Dieses Feld ist erforderlich"
      max_length_msg = "DE maximum 99 chars allowed"
      min_length_msg = "DE minimum 10 chars allowed"
    }
  ]
  field_definition = {
    regex = "^.{10,100}$"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `data_type` (String) The data type of the field. This cannot be modified for an existing resource. Allowed values are `TEXT`,`NUMBER`,`SELECT`,`MULTISELECT`,`RADIO`,`CHECKBOX`,`PASSWORD`,`DATE`,`URL`,`EMAIL`,`TEXTAREA`,`MOBILE`,`CONSENT`,`JSON_STRING`,`USERNAME`,`ARRAY`,`GROUPING`,`DAYDATE`,
- `field_key` (String) The unique identifier of the registration field. This cannot be modified for an existing resource.
- `local_texts` (Attributes List) The localized detail of the registration field. (see [below for nested schema](#nestedatt--local_texts))

### Optional

- `claimable` (Boolean) Flag to mark if a field is claimable. Defaults set to `true`
- `consent_refs` (Set of String) List of consents(the ids of the consent in cidaas must be passed) in registration. The data type must be `CONSENT` in this case
- `enabled` (Boolean) Flag to mark if a field is enabled. Defaults set to `true`
- `field_definition` (Attributes) (see [below for nested schema](#nestedatt--field_definition))
- `field_type` (String) Specifies whether the field type is `SYSTEM` or `CUSTOM`. Defaults to `CUSTOM`. This cannot be modified for an existing resource. `SYSTEM` fields cannot be created but can be modified. To modify an existing field import it first and then update.
- `internal` (Boolean) Flag to mark if a field is internal. Defaults set to `false`
- `is_group` (Boolean) Setting is_group to `true` creates a registration field group. Defaults set to `false` The data_type attribute must be set to TEXT when is_group is true.
- `is_list` (Boolean)
- `is_searchable` (Boolean) Flag to mark if a field is searchable. Defaults set to `true`
- `order` (Number) The attribute order is used to set the order of the Field in the UI. Defaults set to `1`
- `overwrite_with_null_value_from_social_provider` (Boolean) Set to true if you want the value should be reset by identity provider. Defaults set to `false`
- `parent_group_id` (String) The ID of the parent registration group. Defaults to `DEFAULT` if not provided.
- `read_only` (Boolean) Flag to mark if a field is read only. Defaults set to `false`
- `required` (Boolean) Flag to mark if a field is required in registration. Defaults set to `false`
- `scopes` (Set of String) The scopes of the registration field.
- `unique` (Boolean) Flag to mark if a field is unique. Defaults set to `false`

### Read-Only

- `base_data_type` (String) The base data type of the field. This is computed property.
- `id` (String) The ID of the resource

<a id="nestedatt--local_texts"></a>
### Nested Schema for `local_texts`

Required:

- `name` (String) The name of the field in the local configured. for example: in **en-US** the name is `Sample Field` in de-DE `Beispielfeld`.

Optional:

- `attributes` (Attributes List) The field attributes must be provided for the data_type SELECT, MULTISELECT and RADIO. it's an array of key value pairs. Example provided in the example section. (see [below for nested schema](#nestedatt--local_texts--attributes))
- `consent_label` (Attributes) required when data_type is CONSENT. Example provided in the example section. (see [below for nested schema](#nestedatt--local_texts--consent_label))
- `locale` (String) The locale of the field. example: de-DE.
- `max_length_msg` (String) warning/error msg to show to the user when user exceeds the maximum character configured. This is applicable only for the attributes of base_data_type string.
- `min_length_msg` (String) warning/error msg to show to the user when user don't provide the minimum character required. This is applicable only for the attributes of base_data_type string.
- `required_msg` (String) When the flag required is set to true the required_msg must be provided. required_msg is shown if user does not provide a required field.

<a id="nestedatt--local_texts--attributes"></a>
### Nested Schema for `local_texts.attributes`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--local_texts--consent_label"></a>
### Nested Schema for `local_texts.consent_label`

Required:

- `label` (String)
- `label_text` (String)



<a id="nestedatt--field_definition"></a>
### Nested Schema for `field_definition`

Optional:

- `initial_date` (String) The initial date. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.
- `initial_date_view` (String) The view of the calender. Applicable only for DATE attributes. Allowed values: `month`, `year` and `multi-year`
- `max_date` (String) The maximum date a user can select. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.
- `max_length` (Number) The maximum length of a string type attribute.
- `min_date` (String) The earliest date a user can select. Applicable only for DATE attributes. Example format: `2024-06-28T18:30:00Z`.
- `min_length` (Number) The minimum length of a string type attribute
- `regex` (String) The regex for max_length and min_length for the data types TEXT and URL.

## Import

Import is supported using the following syntax:

```shell
terraform import cidaas_registration_page_field.resource_name field_key
```