resource "cidaas_password_policy" "sample" {
  minimum_length       = 8
  lower_and_uppercase  = true
  no_of_digits         = 1
  expiration_in_days   = 30
  no_of_special_chars  = 1
  no_of_days_to_remind = 1
  reuse_limit          = 1
  maximum_length       = 20
}
