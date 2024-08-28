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
	authType        = "APIKEY"
	url             = "https://cidaas.de/webhook-srv/webhook"
	apiConfigKey    = "api-key"
	apiPlaceholder  = "key"
	apiPlacement    = "query"
)

var (
	events        = []string{"ACCOUNT_MODIFIED"}
	apikey_config = map[string]string{
		"key":         apiConfigKey,
		"placeholder": apiPlaceholder,
		"placement":   apiPlacement,
	}
)

// test scenarios
// apikey_config.placeholder must contain only lowercase alphabets

// create, read and update test
func TestAccWebhookResource_Basic(t *testing.T) {
	updatedUrl := "https://cidaas.de/webhook-srv/v2/webhook"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckWebhookDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccWebhookResourceConfig(authType, url, events, apikey_config),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceWebhook, "auth_type", authType),
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
				Config: testAccWebhookResourceConfig(authType, updatedUrl, events, apikey_config),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceWebhook, "url", updatedUrl),
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

// Invalid auth_type validation
func TestAccWebhookResource_InvalidAllowedValue(t *testing.T) {
	invalidAuthType := "INVALID"
	invalidEvents := []string{"INVALID"}
	apikey_config["placement"] = "body"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccWebhookResourceConfig(invalidAuthType, url, events, apikey_config),
				ExpectError: regexp.MustCompile(`Attribute auth_type value must be one of: \["APIKEY" "TOTP" "CIDAAS_OAUTH2"\]`),
			},
			{
				Config:      testAccWebhookResourceConfig(invalidAuthType, url, invalidEvents, apikey_config),
				ExpectError: regexp.MustCompile(`value must be one of`), // TODO: full erro msg match
			},
			{
				Config:      testAccWebhookResourceConfig(authType, url, events, apikey_config),
				ExpectError: regexp.MustCompile(`placement value must be one of: \["query" "header"\]`),
			},
		},
	})
}
