## Changelog

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
