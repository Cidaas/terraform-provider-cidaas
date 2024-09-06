# The cidaas API does not require an identifier to import password policy but Terraform's import command does.
# Therefore, you can provide any arbitrary string as the identifier. It will be set to the `id` attribute in the schema.

terraform import cidaas_password_policy.resource_name cidaas
