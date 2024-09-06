package resources_test

import (
	"fmt"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	resourceApp = "cidaas_app.example"
)

// create, read and update test
func TestApp_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAppConfig("https://cidaas.de"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "client_name", "aiPqVcrYQQ3WzW8J"),
					resource.TestCheckResourceAttr(resourceApp, "company_website", "https://cidaas.de"),
					resource.TestCheckResourceAttrSet(resourceApp, "id"),
				),
			},
			{
				Config: testAppConfig("https://cidaas.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "company_website", "https://cidaas.com"),
				),
			},
		},
	})
}

func testAppConfig(companyWebsite string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		# The config below has the list of common config and main config
resource "cidaas_app" "example" {
  client_name                     = "aiPqVcrYQQ3WzW8J" // unique
  client_display_name             = "Display Name of the app"    // unique
  content_align                   = "CENTER"                     // Default: CENTER
  post_logout_redirect_uris       = ["https://cidaas.com"]
  logo_align                      = "CENTER" // Default: CENTER
  allow_disposable_email          = false    // Default: false
  validate_phone_number           = false    // Default: false
  additional_access_token_payload = ["sample_payload"]
  required_fields                 = ["email"]
  mobile_settings = {
    team_id      = "sample-team-id"
    bundle_id    = "sample-bundle-id"
    package_name = "sample-package-name"
    key_hash     = "sample-key-hash"
  }
  // for custom client credentials use client_id and client_secret, you can leave blank if you want cidaas to create a set for you
  # client_id                       = ""
  # client_secret                   = ""
  policy_uri                          = "https://cidaas.com"
  tos_uri                             = "https://cidaas.com"
  imprint_uri                         = "https://cidaas.com"
  contacts                            = ["support@cidas.de"]
  token_endpoint_auth_method          = "client_secret_post" // Default: client_secret_post
  token_endpoint_auth_signing_alg     = "RS256"              // Default: RS256
  default_acr_values                  = ["default"]
  web_message_uris                    = ["https://cidaas.com"]
  allowed_fields                      = ["email"]
  smart_mfa                           = false // Default: false
  captcha_ref                         = "sample-captcha-ref"
  captcha_refs                        = ["sample"]
  consent_refs                        = ["sample"]
  communication_medium_verification   = "email_verification_required_on_usage"
  mobile_number_verification_required = false // Default: false
  enable_bot_detection                = false // Default: false
  allow_guest_login_groups = [{
    group_id      = "developer101"
    roles         = ["developer", "qa", "admin"]
    default_roles = ["developer"]
  }]
  is_login_success_page_enabled    = false // Default: false
  is_register_success_page_enabled = false // Default: false
  group_ids                        = ["sample"]
  is_group_login_selection_enabled = false // Default: false
  group_selection = {
    selectable_groups      = ["developer-users"]
    selectable_group_types = ["sample"]
  }
  group_types                     = ["sample"]
  logo_uri                        = "https://cidaas.com"
  initiate_login_uri              = "https://cidaas.com"
  registration_client_uri         = "https://cidaas.com"
  registration_access_token       = "registration access token"
  client_uri                      = "https://cidaas.com"
  jwks_uri                        = "https://cidaas.com"
  jwks                            = "https://cidaas.com/jwks"
  sector_identifier_uri           = "https://cidaas.com"
  subject_type                    = "sample subject type"
  id_token_signed_response_alg    = "RS256"
  id_token_encrypted_response_alg = "RS256"
  id_token_encrypted_response_enc = "example"
  userinfo_signed_response_alg    = "RS256"
  userinfo_encrypted_response_alg = "RS256"
  userinfo_encrypted_response_enc = "example"
  request_object_signing_alg      = "RS256"
  request_object_encryption_alg   = "RS256"
  request_object_encryption_enc   = "userinfo_encrypted_response_enc"
  request_uris                    = ["sample"]
  description                     = "app description"
  consent_page_group              = "sample-consent-page-group"
  password_policy_ref             = "password-policy-ref"
  blocking_mechanism_ref          = "blocking-mechanism-ref"
  sub                             = "sample-sub"
  role                            = "sample-role"
  mfa_configuration               = "sample-configuration"
  suggest_mfa                     = ["OFF"]
  login_spi = {
    oauth_client_id = "bcb-4a6b-9777-8a64abe6af"
    spi_url         = "https://cidaas.com/spi-url"
  }
  background_uri  = "https://cidaas.com"
  video_url       = "https://cidaas.com"
  bot_captcha_ref = "sample-bot-captcha-ref"
  application_meta_data = {
    status : "active"
    version : "1.0.0"
  }
  // common config starts here. The attributes from common config can be part of main config
  // if an attribute is available both common_config and main config then attribute from the main config will be considered to create an app
  common_configs = {
    client_type                       = "SINGLE_PAGE"
    accent_color                      = "#ef4923"                        // Default: #ef4923
    primary_color                     = "#ef4923"                        // Default: #f7941d
    media_type                        = "IMAGE"                          // Default: IMAGE
    allow_login_with                  = ["EMAIL", "MOBILE", "USER_NAME"] // Default: ["EMAIL", "MOBILE", "USER_NAME"]
    redirect_uris                     = ["https://cidaas.com"]
    allowed_logout_urls               = ["https://cidaas.com"]
    enable_deduplication              = true      // Default: false
    auto_login_after_register         = true      // Default: false
    enable_passwordless_auth          = false     // Default: true
    register_with_login_information   = false     // Default: false
    fds_enabled                       = false     // Default: true
    hosted_page_group                 = "default" // Default: default
    company_name                      = "Widas ID GmbH"
    company_address                   = "01"
    company_website                   = "%s"
    allowed_scopes                    = ["openid", "cidaas:register", "profile"]
    response_types                    = ["code", "token", "id_token"] // Default: ["code", "token", "id_token"]
    grant_types                       = ["client_credentials"]        // Default: ["implicit", "authorization_code", "password", "refresh_token"]
    login_providers                   = ["login_provider1", "login_provider2"]
    is_hybrid_app                     = true // Default: false
    allowed_web_origins               = ["https://cidaas.com"]
    allowed_origins                   = ["https://cidaas.com"]
    default_max_age                   = 86400      // Default: 86400
    token_lifetime_in_seconds         = 86400      // Default: 86400
    id_token_lifetime_in_seconds      = 86400      // Default: 86400
    refresh_token_lifetime_in_seconds = 15780000   // Default: 15780000
    template_group_id                 = "custtemp" // Default: default
    editable                          = true       // Default: true
    social_providers = [{
      provider_name = "cidaas social provider"
      social_id     = "fdc63bd0-6044-4fa0-abff"
      display_name  = "cidaas"
    }]
    custom_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-custom-provider"
      display_name  = "sample-custom-provider"
      type          = "CUSTOM_OPENID_CONNECT"
    }]
    saml_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-sampl-provider"
      display_name  = "sample-sampl-provider"
      type          = "SAMPL_IDP_PROVIDER"
    }]
    ad_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-ad-provider"
      display_name  = "sample-ad-provider"
      type          = "ADD_PROVIDER"
    }]
    jwe_enabled  = true // Default: false
    user_consent = true // Default: false
    allowed_groups = [{
      group_id      = "developer101"
      roles         = ["developer", "qa", "admin"]
      default_roles = ["developer"]
    }]
    operations_allowed_groups = [{
      group_id      = "developer101"
      roles         = ["developer", "qa", "admin"]
      default_roles = ["developer"]
    }]
    enabled                     = true // Default: true
    always_ask_mfa              = true // Default: false
    allowed_mfa                 = ["OFF"]
    email_verification_required = true // Default: true
    allowed_roles               = ["sample"]
    default_roles               = ["sample"]
    enable_classical_provider   = true     // Default: true
    is_remember_me_selected     = true     // Default: true
    bot_provider                = "CIDAAS" // Default: CIDAAS
    allow_guest_login           = true     // Default: false
    #  mfa Default:
    #  {
    #   setting = "OFF"
    #  }
    mfa = {
      setting = "OFF"
    }
    webfinger      = "no_redirection"
    default_scopes = ["sample"]
    pending_scopes = ["sample"]
  }
}		
	`, companyWebsite)
}

func TestApp_CommonConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				provider "cidaas" {
					base_url = "https://kube-nightlybuild-dev.cidaas.de"
				}
				resource "cidaas_app" "example" {
					client_type         = "SINGLE_PAGE"
					client_name         = "aiPqVcrYQQ3WzW8J"
					client_display_name = "The client Terraform Example App is a sample application designed to demonstrate the configuration of the terraform cidaas_app resource."
					company_address     = "12 Wimsheim, Germany"
					company_website     = "https://cidaas.de"
					allowed_scopes      = ["profile"]
					company_name        = "Cidaas"
					allowed_logout_urls = ["https://ciddas.com", "https://ciddas.com/en"]
					redirect_uris       = ["https://ciddas.com"]
					group_selection     = {}
					login_spi           = {}
					common_configs = {}
				}		
			`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "client_name", "aiPqVcrYQQ3WzW8J"),
					resource.TestCheckResourceAttrSet(resourceApp, "id"),
				),
			},
		},
	})
}
