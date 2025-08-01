package resources_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/internal/resources"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	apiKey         = "APIKEY"
	url            = "https://cidaas.de/webhook-srv/webhook"
	apiConfigKey   = "api-key"
	apiPlaceholder = "key"
	apiPlacement   = "query"
	totp           = "TOTP"
	oauth2         = "CIDAAS_OAUTH2"
)

var events = []string{"ACCOUNT_MODIFIED"}

func getDefaultAPIKeyConfig() map[string]string {
	return map[string]string{
		"key":         apiConfigKey,
		"placeholder": apiPlaceholder,
		"placement":   apiPlacement,
	}
}

func TestAccWebhookResource_Basic(t *testing.T) {
	t.Parallel()
	updatedURL := "https://cidaas.de/webhook-srv/v2/webhook"

	testResourceID := acctest.RandString(10)
	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_WEBHOOK, testResourceID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckWebhookDestroyed(testResourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccWebhookResourceConfig(apiKey, url, testResourceID, events, getDefaultAPIKeyConfig()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "auth_type", apiKey),
					resource.TestCheckResourceAttr(testResourceName, "url", url),
					resource.TestCheckResourceAttr(testResourceName, "events.0", "ACCOUNT_MODIFIED"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "disable"),
					resource.TestCheckResourceAttrSet(testResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(testResourceName, "updated_at"),
				),
			},
			{
				ResourceName:            testResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "updated_at"},
			},
			{
				// url updated
				Config: testAccWebhookResourceConfig(apiKey, updatedURL, testResourceID, events, getDefaultAPIKeyConfig()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "url", updatedURL),
				),
			},
		},
	})
}

func testAccWebhookResourceConfig(
	authType, url, resourceID string,
	events []string,
	apikeyConfig map[string]string,
) string {
	eventsString := "[]"
	if len(events) > 0 {
		eventsString = `["` + strings.Join(events, `", "`) + `"]`
	}

	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_webhook" "%s" {
			auth_type = "`+authType+`"
			url = "`+url+`"
			events = `+eventsString+`
			apikey_config = {
				key = "`+apikeyConfig["key"]+`"
				placeholder = "`+apikeyConfig["placeholder"]+`"
				placement = "`+apikeyConfig["placement"]+`"
			}
		}
	`, acctest.GetBaseURL(), resourceID)
}

func testCheckWebhookDestroyed(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}

		wb := cidaas.Webhook{
			ClientConfig: cidaas.ClientConfig{
				BaseURL:     os.Getenv("BASE_URL"),
				AccessToken: acctest.TestToken,
			},
		}

		// Add retry logic for eventual consistency
		maxRetries := 5
		for i := 0; i < maxRetries; i++ {
			res, err := wb.Get(context.Background(), rs.Primary.Attributes["id"])
			if err != nil {
				// If error is "not found", that's what we want
				if strings.Contains(err.Error(), "not found") ||
					strings.Contains(err.Error(), "404") ||
					strings.Contains(err.Error(), "204") {
					return nil
				}
				return fmt.Errorf("error checking if resource exists: %w", err)
			}

			// Check if resource is nil or marked as unsuccessful
			if res == nil {
				return nil
			}

			// If this is the last retry, return error
			if i == maxRetries-1 {
				return fmt.Errorf("resource still exists after %d retries: %+v", maxRetries, res)
			}

			// Wait before retrying with exponential backoff
			waitTime := time.Duration(i+1) * time.Second * 2
			time.Sleep(waitTime)
		}

		return nil
	}
}

// Invalid auth_type, events and apikey_config placement validation
func TestAccWebhookResource_InvalidAllowedValue(t *testing.T) {
	t.Parallel()
	invalidAuthType := "INVALID"
	invalidEvents := []string{"INVALID"}
	localApiKeyConfig := getDefaultAPIKeyConfig()
	localApiKeyConfig["placement"] = "body"

	testResourceID := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccWebhookResourceConfig(invalidAuthType, url, testResourceID, events, localApiKeyConfig),
				ExpectError: regexp.MustCompile(`Attribute auth_type value must be one of: \["APIKEY" "TOTP" "CIDAAS_OAUTH2"\]`),
			},
			{
				Config:      testAccWebhookResourceConfig(apiKey, url, testResourceID, invalidEvents, localApiKeyConfig),
				ExpectError: regexp.MustCompile(`value must be one of`), // TODO: full error msg match
			},
			{
				Config:      testAccWebhookResourceConfig(apiKey, url, testResourceID, events, localApiKeyConfig),
				ExpectError: regexp.MustCompile(`placement value must be one of: \["query" "header"\]`),
			},
		},
	})
}

// apikey_config.placeholder must contain only lowercase alphabets
func TestAccWebhookResource_PlaceholderLowercase(t *testing.T) {
	t.Parallel()
	invalidPlaceholders := []string{"apiKey", "APIKEY", "api_KEY"}
	testResourceID := acctest.RandString(10)

	for _, v := range invalidPlaceholders {
		localApiKeyConfig := getDefaultAPIKeyConfig()
		localApiKeyConfig["placeholder"] = v
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config:      testAccWebhookResourceConfig(apiKey, url, testResourceID, events, localApiKeyConfig),
					ExpectError: regexp.MustCompile(`Attribute apikey_config.placeholder must contain only lowercase alphabets`),
				},
			},
		})
	}
}

// invalid auth_type and related config(apikey_config, totp_config & cidaas_auth_config) combination
func TestAccWebhookResource_InvalidAuthType(t *testing.T) {
	t.Parallel()
	testResourceID := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccWebhookResourceConfig(totp, url, testResourceID, events, getDefaultAPIKeyConfig()),
				ExpectError: regexp.MustCompile(`The attribute totp_config cannot be empty when the auth_type is TOTP`),
			},
			{
				Config:      testAccWebhookResourceConfig(oauth2, url, testResourceID, events, getDefaultAPIKeyConfig()),
				ExpectError: regexp.MustCompile(`The attribute cidaas_auth_config cannot be empty when the auth_type is`), // TODO: fix why full string match not working
			},
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_webhook" "%s" {
					auth_type = "APIKEY"
					url = "https://cidaas.de/webhook-srv/webhook"
					events = ["ACCOUNT_MODIFIED"]
					totp_config = {
						key = "api-key"
						placeholder = "key"
						placement = "query"
					}
				}
			`, acctest.GetBaseURL(), testResourceID),
				ExpectError: regexp.MustCompile(`The attribute apikey_config cannot be empty when the auth_type is APIKEY`),
			},
		},
	})
}

// create webhook with all 3 auth_type configurations and switch between them
func TestAccWebhookResource_SwitchAuthType(t *testing.T) {
	t.Parallel()
	testResourceID := acctest.RandString(10)
	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_WEBHOOK, testResourceID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: webhookResouceFullConfig(apiKey, testResourceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "auth_type", apiKey),
				),
			},
			{
				Config: webhookResouceFullConfig(totp, testResourceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "auth_type", totp),
				),
			},
			{
				Config: webhookResouceFullConfig(oauth2, testResourceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "auth_type", oauth2),
				),
			},
		},
	})
}

func webhookResouceFullConfig(authType, resourceID string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_webhook" "%s" {
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
		}`, acctest.GetBaseURL(), resourceID, authType)
}
