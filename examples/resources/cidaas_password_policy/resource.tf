resource "cidaas_password_policy" "sample" {
  policy_name         = "sample_terraform_policy"
  minimum_length      = 8
  maximum_length      = 20
  lower_and_uppercase = true
  no_of_digits        = 1
  no_of_special_chars = 1
}
