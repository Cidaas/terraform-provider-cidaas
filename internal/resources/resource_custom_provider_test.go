package resources_test

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	resourceCustomProvider = "cidaas_custom_provider.example"
	oauth2StandardType     = "OAUTH2"
	displayName            = "Sample Terraform"
	authorizationEndpoint  = "https://cidaas.de/authz-srv/authz"
	tokenEndpoint          = "https://cidaas.de/token-srv/token" //nolint:gosec
	logoURL                = "https://cidaas.de/logo"
	userinfoEndpoint       = "https://cidaas.de/users-srv/userinfo"
	scopeDisplayLabel      = "terraform sample scope display name"
)

var (
	providerName = acctest.RandString(10)
	clientID     = acctest.RandString(10)
	clientSecret = acctest.RandString(10)
)

// create, read and update test
func TestAccCustomProviderResource_Basic(t *testing.T) {
	updatedDisplayName := "Updated Sample Terraform"
	updatedOauth2StandardType := "OPENID_CONNECT"
	updatedAuthorizationEndpoint := "https://cidaas.de/authz-srv/v2/authz"
	updatedTokenEndpoint := "https://cidaas.de/token-srv/v2/token" //nolint:gosec
	updatedLogoURL := "https://cidaas.de/v2/logo"
	updatedUserinfoEndpoint := "https://cidaas.de/users-srv/v2/userinfo"
	updatedScopeDisplayLabel := "updated terraform sample scope display name"
	updatedClientID := acctest.RandString(10)
	updatedClientSecret := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             checkCustomProviderDestroyed,
		Steps: []resource.TestStep{
			{
				Config: resourceCustomProviderConfig(oauth2StandardType, providerName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceCustomProvider, "standard_type", oauth2StandardType),
					resource.TestCheckResourceAttr(resourceCustomProvider, "authorization_endpoint", authorizationEndpoint),
					resource.TestCheckResourceAttr(resourceCustomProvider, "token_endpoint", tokenEndpoint),
					resource.TestCheckResourceAttr(resourceCustomProvider, "provider_name", providerName),
					resource.TestCheckResourceAttr(resourceCustomProvider, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceCustomProvider, "logo_url", logoURL),
					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_endpoint", userinfoEndpoint),
					resource.TestCheckResourceAttr(resourceCustomProvider, "scope_display_label", scopeDisplayLabel),
					resource.TestCheckResourceAttr(resourceCustomProvider, "client_id", clientID),
					resource.TestCheckResourceAttr(resourceCustomProvider, "client_secret", clientSecret),
					resource.TestCheckResourceAttr(resourceCustomProvider, "domains.0", "cidaas.de"),
					resource.TestCheckResourceAttr(resourceCustomProvider, "domains.1", "cidaas.org"),
					resource.TestCheckResourceAttr(resourceCustomProvider, "scopes.0.recommended", strconv.FormatBool(true)),
					resource.TestCheckResourceAttr(resourceCustomProvider, "scopes.0.required", strconv.FormatBool(true)),
					resource.TestCheckResourceAttr(resourceCustomProvider, "scopes.0.scope_name", "email"),
					resource.TestCheckResourceAttrSet(resourceCustomProvider, "id"),
				),
			},
			{
				ResourceName:      resourceCustomProvider,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     providerName,
			},
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_custom_provider" "example" {
					standard_type          = "`+updatedOauth2StandardType+`"
					authorization_endpoint = "`+updatedAuthorizationEndpoint+`"
					token_endpoint         = "`+updatedTokenEndpoint+`"
					provider_name          = "`+providerName+`"
					display_name           = "`+updatedDisplayName+`"
					logo_url               =  "`+updatedLogoURL+`"
					userinfo_endpoint      =  "`+updatedUserinfoEndpoint+`"
					scope_display_label    =  "`+updatedScopeDisplayLabel+`"
					client_id              =  "`+updatedClientID+`"
					client_secret          =  "`+updatedClientSecret+`"
					domains                = ["cidaas.in", "cidaas.com"]

					scopes = [
						{
							scope_name  = "openid"
						}
					]
				}
			`, acctest.BaseURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceCustomProvider, "standard_type", updatedOauth2StandardType),
					resource.TestCheckResourceAttr(resourceCustomProvider, "authorization_endpoint", updatedAuthorizationEndpoint),
					resource.TestCheckResourceAttr(resourceCustomProvider, "token_endpoint", updatedTokenEndpoint),
					resource.TestCheckResourceAttr(resourceCustomProvider, "provider_name", providerName),
					resource.TestCheckResourceAttr(resourceCustomProvider, "display_name", updatedDisplayName),
					resource.TestCheckResourceAttr(resourceCustomProvider, "logo_url", updatedLogoURL),
					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_endpoint", updatedUserinfoEndpoint),
					resource.TestCheckResourceAttr(resourceCustomProvider, "scope_display_label", updatedScopeDisplayLabel),
					resource.TestCheckResourceAttr(resourceCustomProvider, "client_id", updatedClientID),
					resource.TestCheckResourceAttr(resourceCustomProvider, "client_secret", updatedClientSecret),
					resource.TestCheckResourceAttr(resourceCustomProvider, "domains.0", "cidaas.com"),
					resource.TestCheckResourceAttr(resourceCustomProvider, "domains.1", "cidaas.in"),
					resource.TestCheckResourceAttr(resourceCustomProvider, "scopes.0.scope_name", "openid"),
					// default value check scopes[i].recommended & scopes[i].required
					resource.TestCheckResourceAttr(resourceCustomProvider, "scopes.0.recommended", strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(resourceCustomProvider, "scopes.0.required", strconv.FormatBool(false)),
					resource.TestCheckResourceAttrSet(resourceCustomProvider, "id"),
				),
			},
			{
				// provider_name cannot be updated
				Config:      resourceCustomProviderConfig(oauth2StandardType, "new_provider_name"),
				ExpectError: regexp.MustCompile(`Attribute 'provider_name' can't be modified`),
			},
		},
	})
}

func resourceCustomProviderConfig(standardType, providerName string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_custom_provider" "example" {
			standard_type          = "%s"
			authorization_endpoint = "`+authorizationEndpoint+`"
			token_endpoint         =  "`+tokenEndpoint+`"
			provider_name          = "%s"
			display_name           = "`+displayName+`"
			logo_url               =  "`+logoURL+`"
			userinfo_endpoint      =  "`+userinfoEndpoint+`"
			scope_display_label    =  "`+scopeDisplayLabel+`"
			client_id              =  "`+clientID+`"
			client_secret          =  "`+clientSecret+`"
			domains                = ["cidaas.de", "cidaas.org"]
		
			scopes = [
				{
					recommended = true
					required    = true
					scope_name  = "email"
				}
			]
		}
	`, acctest.BaseURL, standardType, providerName)
}

func checkCustomProviderDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceCustomProvider]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceCustomProvider)
	}

	cp := cidaas.CustomProvider{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := cp.GetCustomProvider(rs.Primary.Attributes["provider_name"])
	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// Invalid standard_type validation
func TestAccCustomProviderResource_InvalidStandardType(t *testing.T) {
	invalidStandardType := "OAUTH1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      resourceCustomProviderConfig(invalidStandardType, providerName),
				ExpectError: regexp.MustCompile(`Attribute standard_type value must be one of: \["OPENID_CONNECT" "OAUTH2"\]`),
			},
		},
	})
}

// missing required parameter
func TestAccCustomProviderResource_MissingRequired(t *testing.T) {
	requiredParams := []string{
		"provider_name", "display_name", "client_id", "client_secret",
		"authorization_endpoint", "token_endpoint", "userinfo_endpoint",
	}
	for _, param := range requiredParams {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						provider "cidaas" {
							base_url = "%s"
						}
						resource "cidaas_custom_provider" "example" {}
					`, acctest.BaseURL),
					ExpectError: regexp.MustCompile(fmt.Sprintf(`The argument "%s" is required`, param)),
				},
			},
		})
	}
}

// check userinfo_fields parameters
// func TestAccCustomProviderResource_UserinfoFieldsCheck(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: fmt.Sprintf(`
// 				provider "cidaas" {
// 					base_url = "%s"
// 				}
// 				resource "cidaas_custom_provider" "example" {
// 					standard_type          = "`+oauth2StandardType+`"
// 					authorization_endpoint = "`+authorizationEndpoint+`"
// 					token_endpoint         = "`+tokenEndpoint+`"
// 					provider_name          = "`+providerName+`"
// 					display_name           = "`+displayName+`"
// 					userinfo_endpoint      = "`+userinfoEndpoint+`"
// 					scope_display_label    = "`+scopeDisplayLabel+`"
// 					client_id              = "`+clientID+`"
// 					client_secret          = "`+clientSecret+`"
// 					scopes = [
// 						{
// 							scope_name  = "email"
// 						}
// 					]
// 					userinfo_fields = {
// 						family_name        = "cp_family_name"
// 						address            = "cp_address"
// 						birthdate          = "01-01-2000"
// 						email              = "cp@cidaas.de"
// 						email_verified     = "true"
// 						gender             = "male"
// 						given_name         = "cp_given_name"
// 						locale             = "cp_locale"
// 						middle_name        = "cp_middle_name"
// 						mobile_number      = "100000000"
// 						phone_number       = "10000000"
// 						picture            = "https://cidaas.de/image.jpg"
// 						preferred_username = "cp_preferred_username"
// 						profile            = "cp_profile"
// 						updated_at         = "01-01-01"
// 						website            = "https://cidaas.de"
// 						zoneinfo           = "cp_zone_info"
// 						custom_fields = {
// 							zipcode         = "123456"
// 							alternate_phone = "1234567890"
// 						}
// 					}
// 				}`, acctest.BaseURL),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					// default value check
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.family_name", "cp_family_name"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.address", "cp_address"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.birthdate", "01-01-2000"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.email", "cp@cidaas.de"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.email_verified", "true"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.gender", "male"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.given_name", "cp_given_name"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.locale", "cp_locale"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.middle_name", "cp_middle_name"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.mobile_number", "100000000"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.phone_number", "10000000"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.picture", "https://cidaas.de/image.jpg"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.preferred_username", "cp_preferred_username"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.profile", "cp_profile"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.updated_at", "01-01-01"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.website", "https://cidaas.de"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.zoneinfo", "cp_zone_info"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.custom_fields.zipcode", "123456"),
// 					resource.TestCheckResourceAttr(resourceCustomProvider, "userinfo_fields.custom_fields.alternate_phone", "1234567890"),
// 				),
// 			},
// 		},
// 	})
// }
