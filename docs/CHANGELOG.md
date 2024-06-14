## Changelog

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
