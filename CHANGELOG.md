## Changelog

### 3.5.0
- Added context support for proper HTTP request cancellation and timeout handling
- Enhanced resource `cidaas_app` import to include all schema fields
- Added support for new fields in `cidaas_custom_provider` resource: `groups`, `pkce`, `auth_type`, `apikey_details`, `totp_details`, `cidaas_auth_details`

### 3.4.9

### Bug Fixes

- Added `is_group_login_selection_enabled` flag to the `cidaas_app` resource that removed earlier in v3.4.7, allowing to enable or disable group login selection

### 3.4.8

### Enhancements & Bug Fixes

- Added `enabled` flag to the `cidaas_template` resource, allowing to activate or deactivate a template
- Added support for custom values in `allow_login_with` attribute in `cidaas_app` resource
- Fixed handling of null set/list attributes in `cidaas_app` resource by sending empty arrays in API requests

### 3.4.7

### Enhancements & Bug Fixes

- Unsupported attributes removed from the `cidaas_app` resource schema. The following attributes were removed:

  - always_ask_mfa
  - editable
  - email_verification_required
  - enable_classical_provider
  - fds_enabled
  - is_group_login_selection_enabled
  - mobile_number_verification_required

- Fix provided for the issue in the `cidaas_webhook` resource where `placeholder` attribute was incorrectly rejecting valid values containing dashes (e.g. `test-apikey-placeholder`). It now correctly accepts placeholders using lowercase alphabets and dashes as intended.

- The attribute `group_type` is optional now in the `cidaas_user_groups` resource.

### 3.4.6

#### Enhancements

- The `cidaas_app` resource has been enhanced to behave more accurately based on the `client_type` attribute. With this update, Terraform configurations must now explicitly define values for all relevant attributes, as they are no longer treated as computed or automatically assigned defaults by the provider during resource creation.
For example, the `enabled` attribute was previously defaulted to `true` by the provider when creating an application. With this change, if you do not specify `enabled` in your configuration, the provider will omit it from the API request allowing the server to apply its own default behavior instead.
This ensures a more predictable and transparent configuration experience, aligning the provider behavior more closely with user intent and server-side defaults.

### 3.4.5

#### Bug Fixes

- The attribute `hosted_pages` in the resource `cidaas_hosted_page` has been updated to use an unordered list. This change resolves the issue where Terraform would incorrectly detect changes in the `hosted_pages` attribute, even when there were no actual modifications to the list, apart from reordering.

### 3.4.4

#### Bug Fixes

- Reduced plan time validation from resource `cidaas_template`.

### 3.4.3

#### Enhancements
- The `regex` field has been introduced in `field_definition` for the `cidaas_registration_field` resource (starting from Cidaas version 3.101.5).
- This change **replaces** the `max_length` and `min_length` attributes **for `TEXT` and `URL` data types**.
- Instead of relying on fixed length constraints, validation for these field types will now be handled using **regular expressions (`regex`)**, providing more flexibility.

#### **Example of new regex-based validation**
```python
field_definition = {
    regex = "^(https?:\/\/)?([\da-z.-]+)\.([a-z.]{2,6})([\/\w .-]*)*\/?$"
}
```
#### Bug Fixes

- Fixed state consistency issues in resource `cidaas_template` and `cidaas_template_group`.

### 3.4.2

#### Bugfix

- The`cidaas_password_policy` resource has been updated to support the enhanced password policy introduced in Cidaas version 3.100.x

### 3.4.1

#### Bugfix

- Attribute `scope_display_label` in resource **cidaas_custom_provider** marked optional. This fixes the state inconsistency issue when `scope_display_label` set to empty string.

### 3.4.0

#### Bugfix

- Resource **cidaas_social_provide** bug fix where empty `required_claims` and `optional_claims` provider plan error.
- Fixed the issue in **cidaas_app** resource  where the custom provider was not updated to an empty state after being removed from the config.


### 3.3.9

#### Enhancement

- Schema of the attribute `userinfo_fields` in resource **cidaas_custom_provider** changed to match cidaas api to suppoort external field key and default value.
- Attributses `amr_config` and `userinfo_source` are supported now in resource **cidaas_custom_provider**.

### 3.3.8

#### Bugfix

- Fixed issue where empty `consent_refs` array was being omitted from API requests causing state inconsistency

### 3.3.7

#### Enhancements

- Attribute basic_settings no longer supported in resource cidaas_app.

### 3.3.6

#### Enhancements

- Extend custom provider resource to support custom provider new api contract.

### 3.3.5

#### Enhancements

- Extend custom provider resource to support custom provider new api contract.

### 3.3.4

#### Enhancements

- Attribute `accept_roles_in_the_registration` added to the resource cidaas_app.

### 3.3.3

#### Enhancements

- `match_condition` and `filters` attributes in `group_role_restriction`(cidaas_app) are now required if `group_role_restriction` is declared in the configuration. This helps prevent misconfiguration.
- cidaas_app import now ignore empty `group_role_restriction` objects in the api response fixing schema mismarch issue.

### 3.3.2

#### Enhancements

- Enhanced validation on attributes processing_type and usage_type in resource cidaas_template

### 3.3.1

#### Enhanced Locale Support

The provider now includes additional locales `de-BE`, `id`, `zh-Hans` and `zh-Hant`.

### 3.3.0

#### Removed common_configs from resource app

The attribute `common_configs` is removed from the resource cidaas_app as we introduce [terraform-cidaas-app](https://github.com/Cidaas/terraform-cidaas-app) module.

### 3.2.0

#### Addition of datasources

This release includes the following datasources:

- cidaas_consent
- cidaas_custom_provider
- cidaas_group_type
- cidaas_registration_field
- cidaas_role
- cidaas_scope_group
- cidaas_scope
- cidaas_social_provider
- cidaas_system_template_option

#### Additional attribute support in resource cidaas_app

The following attributes are added to the resource `cidaas_app`:

- require_auth_time
- enable_login_spi
- backchannel_logout_session_required
- suggest_verification_methods
- group_role_restriction
- basic_settings

#### Bug Fix

- Fixed the issue **Consent Not Found** when the name of the consent resource is in uppercase during update & destroy

### 3.1.2

#### Enhancements

- **Multiple Password Policy Support:** Password Policy resource changed to support multiple policies

### 3.1.1

#### Enhancements

- **Locale Support for Template Resource:** Added support for the **rm** & **rm-CH** language code in the template resource.

### 3.1.0

This release includes the below new resources

- cidaas_social_provider
- cidaas_password_policy
- cidaas_consent_group
- cidaas_consent
- cidaas_consent_version

Please find the readme [here](https://github.com/Cidaas/terraform-provider-cidaas/blob/v3.1.0/README.md) to explore more on the new resources.

### 3.0.5

#### Bug Fix

- password_policy_ref empty string or null values can be passed as "" when configured.
- Addressed the issue where computed attributes group_selection, login_spi & mobile_settings are not known after terraform apply, a default value {} is assigned to them.

### 3.0.4

#### Bug Fix

- **custom provider schema fix:** The issue with the sub attribute not aligning with the schema of the custom provider has been resolved.
- **app schema fix**: The app resource's list nested attributes are now updated to align with the Cidaas API response.

### 3.0.3

#### Enhancements

- **Enhanced State Management:** Fixed state inconsistencies for attributes computed by Cidaas APIs due to dependencies or API support changes.

### 3.0.2

#### Enhancements

- **Validation Update:** `group_id` is now required when `is_system_template=true` in resource cidaas_template.

#### Fixes

- **Validation Removed:** Removed the validation that checked the availability of template group by `group_id` in Cidaas before creating a template as the api sometimes may fail to fetch the template group immediately after its creation.

### 3.0.1

#### Removed

- **URL Validation**: Removed strict URL validation that enforced URLs to start with `https://`.

### 3.0.0

This new release is based on Terraform Plugin Framework and is designed to be mostly backwards compatible with existing implementations. It offers several benefits including enhanced performance, improved debugging capabilities and streamlined development processes. Specific advantages include:

- **Simplified Resource Management**: More straightforward management through enhanced schemas.
- **Improved Error Handling**: Error handling has been revamped. Errors now includes suggestions that should assist you to manage your resources.
- **Enhanced Performance**: Custom plan-time validations across all the resources that helps faster plugin operations and dynamic provider configurations in resource app.

Despite these improvements, some breaking changes are present. Users need to be aware of the following modifications:

- **Schema Changes**: The new release includes changes in some of the existing schemas. Review and update schema definitions to align with the new structure.
- **Change in Import Statement**: The import statement of resource cidaas_template has been changed.
- **Resource Name Update**: The resource name of cidaas_user_group_category and cidaas_registration_page_field changed.

#### Additions

- A new resource `cidaas_template_group` has been added to support template groups,which are required for creating system templates.
- **SYSTEM** templates can now be created using the provider. Refer to the template section in the documentation for more details.
- Added support for internationalization in `cidaas_registration_field` and `cidaas_scope` with multi-language capabilities.
- `cidaas_registration_field` now supports all the datatypes that Cidaas supports.

### 2.5.8

#### Additions

- resource cidaas_registration_page_field schema update to toggle `overwrite_with_null_value_from_social_provider` in SYSTEM fields.

### 2.5.7

#### Additions

- app_resource schema update to support `is_provider_visible` in customProviders, socialProviders & adProviders.

### 2.5.6

#### Additions

- `application_meta_data` added to support custom fields in cidaas_app resource.

- Validations added to prevent users from updating the locale and template_type of an existing `cidaas_template` state. This ensures data integrity and consistency.

- Enhanced error messages when users provide an incorrect locale,

- Updated the format of the `cidaas_template` ID from `templateKey_templateType` to `templateKey_templateType_locale`. Added validation checks for incorrect format types.

#### Bug Fixes

- Resolved the issue causing the error message `failed to unmarshal JSON body, EOF` during template deletion.

### 2.5.5

#### Bug Fixes

- Fixed the issue **subject can not be empty for template_key EMAIL** even though subject is available in the terraform config file

- app_key marked sensitive

- README updated with the instructions to guide Windows users to set env variables and scopes required for templates are added

### 2.5.4

#### Bug Fixes

- Fix added to address the issue where updating an existing cidaas_app without the `client_id` attribute throws error **client id is missing**.

- Improved error handling in terraform cidaas_app destroy. This solves the issue **invalid memory address or nil pointer dereference** while deleting client in cidaas.
