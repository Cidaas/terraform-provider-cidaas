package resources_test

import (
	"fmt"
	"os"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	resourceApp = "cidaas_app.example"
)

func TestApp_Basic(t *testing.T) {
	clientName := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAppConfig(clientName, "https://cidaas.de"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "client_name", clientName),
					resource.TestCheckResourceAttr(resourceApp, "company_website", "https://cidaas.de"),
					resource.TestCheckResourceAttrSet(resourceApp, "id"),
				),
			},
			{
				Config: testAppConfig(clientName, "https://cidaas.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "company_website", "https://cidaas.com"),
				),
			},
		},
	})
}

func testAppConfig(clientName, companyWebsite string) string {
	return fmt.Sprintf(`
    provider "cidaas" {
      base_url = "%s"
    }
    # The config below has the list of common config and main config
    resource "cidaas_app" "example" {
      client_name                     = "%s"
      client_type                     = "SINGLE_PAGE"
      redirect_uris                   = ["https://cidaas.com"]
      allowed_logout_urls             = ["https://cidaas.com"]
      company_name                    = "Widas ID GmbH"
      company_address                 = "01"
      company_website                 = "%s"
      allowed_scopes                  = ["openid", "cidaas:register", "profile"]
      client_display_name             = "Display Name of the app"
      content_align                   = "CENTER"
      post_logout_redirect_uris       = ["https://cidaas.com"]
      logo_align                      = "CENTER"
      allow_disposable_email          = false
      validate_phone_number           = false
      additional_access_token_payload = ["sample_payload"]
      required_fields                 = ["email"]
      mobile_settings = {
        team_id      = "sample-team-id"
        bundle_id    = "sample-bundle-id"
        package_name = "sample-package-name"
        key_hash     = "sample-key-hash"
      }
      policy_uri                          = "https://cidaas.com"
      tos_uri                             = "https://cidaas.com"
      imprint_uri                         = "https://cidaas.com"
      contacts                            = ["support@cidas.de"]
      token_endpoint_auth_method          = "client_secret_post"
      token_endpoint_auth_signing_alg     = "RS256"
      default_acr_values                  = ["default"]
      web_message_uris                    = ["https://cidaas.com"]
      allowed_fields                      = ["email"]
      smart_mfa                           = false
      captcha_ref                         = "sample-captcha-ref"
      captcha_refs                        = ["sample"]
      consent_refs                        = ["sample"]
      communication_medium_verification   = "email_verification_required_on_usage"
      mobile_number_verification_required = false
      enable_bot_detection                = false
      allow_guest_login_groups = [{
        group_id      = "developer101"
        roles         = ["developer", "qa", "admin"]
        default_roles = ["developer"]
      }]
      is_login_success_page_enabled    = false
      is_register_success_page_enabled = false
      group_ids                        = ["sample"]
      is_group_login_selection_enabled = false
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
    }`,
		os.Getenv("BASE_URL"),
		clientName,
		companyWebsite,
	)
}

func TestApp_CommonConfig(t *testing.T) {
	clientName := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_app" "example" {
					client_type         = "SINGLE_PAGE"
					client_name         = "%s"
					client_display_name = "sample client"
					company_address     = "12 Wimsheim, Germany"
					company_website     = "https://cidaas.de"
					allowed_scopes      = ["profile"]
					company_name        = "Cidaas"
					allowed_logout_urls = ["https://ciddas.com", "https://ciddas.com/en"]
					redirect_uris       = ["https://ciddas.com"]
					group_selection     = {}
					login_spi           = {}
				}		
			`, acctest.BaseURL, clientName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "client_name", clientName),
					resource.TestCheckResourceAttrSet(resourceApp, "id"),
				),
			},
		},
	})
}
