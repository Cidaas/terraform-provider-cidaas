package resources_test

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	resourceWebhook = "cidaas_webhook.example"
	apiKey          = "APIKEY"
	url             = "https://cidaas.de/webhook-srv/webhook"
	apiConfigKey    = "api-key"
	apiPlaceholder  = "key"
	apiPlacement    = "query"
	totp            = "TOTP"
	oauth2          = "CIDAAS_OAUTH2"
)

var (
	events       = []string{"ACCOUNT_MODIFIED"}
	apikeyConfig = map[string]string{
		"key":         apiConfigKey,
		"placeholder": apiPlaceholder,
		"placement":   apiPlacement,
	}
)

// create, read and update test
func TestAccWebhookResource_Basic(t *testing.T) {
	updatedURL := "https://cidaas.de/webhook-srv/v2/webhook"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckWebhookDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccWebhookResourceConfig(apiKey, url, events, apikeyConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceWebhook, "auth_type", apiKey),
					resource.TestCheckResourceAttr(resourceWebhook, "url", url),
					resource.TestCheckResourceAttr(resourceWebhook, "events.0", "ACCOUNT_MODIFIED"),
					resource.TestCheckResourceAttrSet(resourceWebhook, "id"),
					resource.TestCheckResourceAttrSet(resourceWebhook, "disable"),
					resource.TestCheckResourceAttrSet(resourceWebhook, "created_at"),
					resource.TestCheckResourceAttrSet(resourceWebhook, "updated_at"),
				),
			},
			{
				ResourceName:            resourceWebhook,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "updated_at"},
			},
			{
				// url updated
				Config: testAccWebhookResourceConfig(apiKey, updatedURL, events, apikeyConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceWebhook, "url", updatedURL),
				),
			},
		},
	})
}

func testAccWebhookResourceConfig(
	authType, url string,
	events []string,
	apikeyConfig map[string]string,
) string {
	eventsString := "[]"
	if len(events) > 0 {
		eventsString = `["` + strings.Join(events, `", "`) + `"]`
	}

	return `
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_webhook" "example" {
			auth_type = "` + authType + `"
			url = "` + url + `"
			events = ` + eventsString + `
			apikey_config = {
				key = "` + apikeyConfig["key"] + `"
				placeholder = "` + apikeyConfig["placeholder"] + `"
				placement = "` + apikeyConfig["placement"] + `"
			}
		}
	`
}

func testCheckWebhookDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceWebhook]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceWebhook)
	}

	wb := cidaas.Webhook{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := wb.Get(rs.Primary.Attributes["id"])
	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// Invalid auth_type, events and apikey_config placement validation
func TestAccWebhookResource_InvalidAllowedValue(t *testing.T) {
	invalidAuthType := "INVALID"
	invalidEvents := []string{"INVALID"}
	apikeyConfig["placement"] = "body"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccWebhookResourceConfig(invalidAuthType, url, events, apikeyConfig),
				ExpectError: regexp.MustCompile(`Attribute auth_type value must be one of: \["APIKEY" "TOTP" "CIDAAS_OAUTH2"\]`),
			},
			{
				Config:      testAccWebhookResourceConfig(apiKey, url, invalidEvents, apikeyConfig),
				ExpectError: regexp.MustCompile(`value must be one of`), // TODO: full error msg match
			},
			{
				Config:      testAccWebhookResourceConfig(apiKey, url, events, apikeyConfig),
				ExpectError: regexp.MustCompile(`placement value must be one of: \["query" "header"\]`),
			},
		},
	})
}

// apikey_config.placeholder must contain only lowercase alphabets
func TestAccWebhookResource_PlaceholderLowercase(t *testing.T) {
	invalidPlaceholders := []string{"apiKey", "APIKEY", "api-key", "api_KEY"}
	for _, v := range invalidPlaceholders {
		apikeyConfig["placeholder"] = v
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config:      testAccWebhookResourceConfig(apiKey, url, events, apikeyConfig),
					ExpectError: regexp.MustCompile(`Attribute apikey_config.placeholder must contain only lowercase alphabets`),
				},
			},
		})
	}
}

// invalid auth_type and related config(apikey_config, totp_config & cidaas_auth_config) combination
func TestAccWebhookResource_InvalidAuthType(t *testing.T) {
	// apikey_config is reverted to the correct old value after it was updated to invalid ones in previous tests
	apikeyConfig["placement"] = "query"
	apikeyConfig["placeholder"] = "key"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccWebhookResourceConfig(totp, url, events, apikeyConfig),
				ExpectError: regexp.MustCompile(`The attribute totp_config cannot be empty when the auth_type is TOTP`),
			},
			{
				Config:      testAccWebhookResourceConfig(oauth2, url, events, apikeyConfig),
				ExpectError: regexp.MustCompile(`The attribute cidaas_auth_config cannot be empty when the auth_type is`), // TODO: fix why full string match not working
			},
			{
				Config: `
				provider "cidaas" {
					base_url = "https://kube-nightlybuild-dev.cidaas.de"
				}
				resource "cidaas_webhook" "example" {
					auth_type = "APIKEY"
					url = "https://cidaas.de/webhook-srv/webhook"
					events = ["ACCOUNT_MODIFIED"]
					totp_config = {
						key = "api-key"
						placeholder = "key"
						placement = "query"
					}
				}
			`,
				ExpectError: regexp.MustCompile(`The attribute apikey_config cannot be empty when the auth_type is APIKEY`),
			},
		},
	})
}

// create webhook with all 3 auth_type configurations and switch between them
func TestAccWebhookResource_SwitchAuthType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: webhookResouceFullConfig(apiKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceWebhook, "auth_type", apiKey),
				),
			},
			{
				Config: webhookResouceFullConfig(totp),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceWebhook, "auth_type", totp),
				),
			},
			{
				Config: webhookResouceFullConfig(oauth2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceWebhook, "auth_type", oauth2),
				),
			},
		},
	})
}

func webhookResouceFullConfig(authType string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_webhook" "example" {
			auth_type = "%s"
			url = "https://cidaas.de/webhook-srv/webhook"
			events = ["ACCOUNT_MODIFIED"]
			apikey_config = {
				key = "api-key"
				placeholder = "key"
				placement = "query"
			}
			totp_config = {
				key = "totp-key"
				placeholder = "key"
				placement = "header"
			}
			cidaas_auth_config = {
				client_id = "ce90d6ba-9a5a-49b6-9a50"
			}
		}`, authType)
}
