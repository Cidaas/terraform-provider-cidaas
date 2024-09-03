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
	resourceSocialPorvider = "cidaas_social_provider.example"
)

var (
	spName         = acctest.RandString(10)
	spProviderName = "google"
	spClientID     = acctest.RandString(10)
	spClientSecret = acctest.RandString(10)
)

// create, read and update test
func TestSocialProvider_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             checksocialProviderDestroyed,
		Steps: []resource.TestStep{
			{
				Config: socialProviderConfig(spName, spProviderName, spClientID, spClientSecret),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceSocialPorvider, "name", spName),
					resource.TestCheckResourceAttr(resourceSocialPorvider, "provider_name", spProviderName),
					resource.TestCheckResourceAttr(resourceSocialPorvider, "client_id", spClientID),
					resource.TestCheckResourceAttr(resourceSocialPorvider, "client_secret", spClientSecret),
					resource.TestCheckResourceAttr(resourceSocialPorvider, "claims.required_claims.user_info.0", "name"),

					// default value check
					resource.TestCheckResourceAttr(resourceSocialPorvider, "enabled", strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(resourceSocialPorvider, "enabled_for_admin_portal", strconv.FormatBool(false)),

					resource.TestCheckResourceAttrSet(resourceSocialPorvider, "id"),
				),
			},
			{
				ResourceName:      resourceSocialPorvider,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceSocialPorvider]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceSocialPorvider)
					}
					return rs.Primary.Attributes["provider_name"] + ":" + rs.Primary.ID, nil
				},
			},
			{
				Config: socialProviderConfig(spName, spProviderName, spClientID, spClientSecret),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceSocialPorvider, "id"),
				),
			},
			// immutable name and provider_name validation
			{
				Config:      socialProviderConfig(spName, "facebook", spClientID, spClientSecret),
				ExpectError: regexp.MustCompile(`Attribute 'provider_name' can't be modified`),
			},
			{
				Config:      socialProviderConfig(acctest.RandString(5), spProviderName, spClientID, spClientSecret),
				ExpectError: regexp.MustCompile(`Attribute 'name' can't be modified`),
			},
		},
	})
}

func socialProviderConfig(name, providerName, clientID, clientSecret string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_social_provider" "example" {
			name                     = "%s"
			provider_name            = "%s"
			client_id                = "%s"
			client_secret            = "%s"
			scopes                   = ["profile", "email"]
			claims = {
				required_claims = {
					user_info = ["name"]
					id_token  = ["phone_number"]
				}
				optional_claims = {
					user_info = ["website"]
					id_token  = ["street_address"]
				}
			}
			userinfo_fields = [
				{
					inner_key       = "sample_custom_field"
					external_key    = "external_sample_cf"
					is_custom_field = true
					is_system_field = false
				},
				{
					inner_key       = "sample_system_field"
					external_key    = "external_sample_sf"
					is_custom_field = false
					is_system_field = true
				}
			]
		}
	`, name, providerName, clientID, clientSecret)
}

func checksocialProviderDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceSocialPorvider]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceSocialPorvider)
	}

	sp := cidaas.SocialProvider{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := sp.Get(rs.Primary.Attributes["provider_name"], rs.Primary.Attributes["id"])
	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// Invalid provider_name validation
func TestSocialProvider_InvalidProviderName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      socialProviderConfig(spName, acctest.RandString(10), spClientID, spClientSecret),
				ExpectError: regexp.MustCompile(`Attribute provider_name value must be one of:`),
			},
		},
	})
}

// missing required parameter
func TestSocialProvider_MissingRequired(t *testing.T) {
	requiredParams := []string{
		"name", "provider_name", "client_id", "client_secret",
	}
	for _, param := range requiredParams {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: `
						provider "cidaas" {
							base_url = "https://kube-nightlybuild-dev.cidaas.de"
						}
						resource "cidaas_social_provider" "example" {}
					`,
					ExpectError: regexp.MustCompile(fmt.Sprintf(`The argument "%s" is required`, param)),
				},
			},
		})
	}
}
