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
	hostedPageURL     = "https://cidaad.de/register_success"
	hostedPageID      = "register_success"
	hostedPageContent = "<html>Register Success</html>"
	defaultLocale     = "en-US"
)

var (
	hostedPages = []map[string]string{
		{
			"hosted_page_id": hostedPageID,
			"locale":         defaultLocale,
			"url":            hostedPageURL,
			"content":        hostedPageContent,
		},
	}
)

// create, read and update test
func TestAccHostedPageResource_Basic(t *testing.T) {
	t.Parallel()

	updatedHostedPageURL := "https://cidaad.de/updated_register_success"
	updatedHostedPages := []map[string]string{
		{
			"hosted_page_id": hostedPageID,
			"locale":         defaultLocale,
			"url":            updatedHostedPageURL,
			"content":        "<html>Updated Success</html>",
		},
	}

	resourceID := acctest.RandString(10)
	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_HOSTED_PAGE, resourceID)

	hostedPageGroupName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckHostedPageDestroyed(testResourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					resourceID,
					hostedPages,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckHostedPageExists(testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "hosted_page_group_name", hostedPageGroupName),
					resource.TestCheckResourceAttr(testResourceName, "default_locale", defaultLocale),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.hosted_page_id", hostedPageID),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.locale", defaultLocale),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.url", hostedPageURL),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.content", hostedPageContent),
				),
			},
			{
				ResourceName:            testResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "updated_at"},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "hosted_page_group_name", hostedPageGroupName),
					resource.TestCheckResourceAttr(testResourceName, "default_locale", defaultLocale),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.hosted_page_id", hostedPageID),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.locale", defaultLocale),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.url", hostedPageURL),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.content", hostedPageContent),
				),
			},
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					resourceID,
					updatedHostedPages,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.url", updatedHostedPageURL),
					resource.TestCheckResourceAttr(testResourceName, "hosted_pages.0.content", "<html>Updated Success</html>"),
				),
			},
		},
	})
}

func testAccHostedPageResourceConfig(hostedPageGroupName, defaultLocale, resourceID string, hostedPages []map[string]string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "%s" {
			hosted_page_group_name = "`+hostedPageGroupName+`"
			default_locale = "`+defaultLocale+`"

			hosted_pages =[
				{
				hosted_page_id = "`+hostedPages[0]["hosted_page_id"]+`"
				locale = "`+hostedPages[0]["locale"]+`"
				url = "`+hostedPages[0]["url"]+`"
				content = "`+hostedPages[0]["content"]+`"
				}
			]
		}
	`, acctest.GetBaseURL(), resourceID)
}

func testCheckHostedPageExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if _, ok := s.RootModule().Resources[resourceName]; !ok {
			return fmt.Errorf("Not found: %s, found resource %s", resourceName, s.RootModule().Resources)
		}
		return nil
	}
}

func testCheckHostedPageDestroyed(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}

		hp := cidaas.HostedPage{
			ClientConfig: cidaas.ClientConfig{
				BaseURL:     os.Getenv("BASE_URL"),
				AccessToken: acctest.TestToken,
			},
		}

		// Add retry logic for eventual consistency
		maxRetries := 5
		for i := 0; i < maxRetries; i++ {
			res, err := hp.Get(context.Background(), rs.Primary.Attributes["hosted_page_group_name"])

			// Check if resource is successfully deleted (nil response)
			if res == nil {
				return nil // Resource successfully deleted
			}

			// Handle other errors
			if err != nil {
				// If error is "not found", that's what we want
				if strings.Contains(err.Error(), "not found") ||
					strings.Contains(err.Error(), "404") ||
					strings.Contains(err.Error(), "204") {
					return nil
				}
				return fmt.Errorf("error checking if hosted page exists: %w", err)
			}

			// If this is the last retry, return error
			if i == maxRetries-1 {
				return fmt.Errorf("hosted page still exists after %d retries: %s", maxRetries, res.Data)
			}

			// Wait before retrying with exponential backoff
			waitTime := time.Duration(i+1) * time.Second * 2
			time.Sleep(waitTime)
		}

		return nil
	}
}

// invalid locale
func TestAccHostedPageResource_InvalidLocale(t *testing.T) {
	t.Parallel()

	invalidLocale := "invalid-locale"
	hostedPageGroupName := acctest.RandString(10)

	resourceID := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      testAccHostedPageResourceConfig(hostedPageGroupName, invalidLocale, resourceID, hostedPages),
				ExpectError: regexp.MustCompile("Attribute default_locale value must be one of"),
			},
		},
	})
}

// missing required fields validation
func TestAccHostedPageResource_MissingRequiredFields(t *testing.T) {
	t.Parallel()

	config1 := fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "%s" {
			hosted_page_group_name = ""
			default_locale = "en-US"
			hosted_pages =[{
				hosted_page_id = "register_success"
				url = ""
			}]
		}
		`, acctest.GetBaseURL(), acctest.RandString(10))
	config2 := fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "%s" {
			hosted_page_group_name = ""
			default_locale = "en-US"
		}
		`, acctest.GetBaseURL(), acctest.RandString(10))
	config3 := fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "%s" {
			hosted_page_group_name = ""
			default_locale = "en-US"
			hosted_pages =[]
		}
		`, acctest.GetBaseURL(), acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      config1,
				ExpectError: regexp.MustCompile("Attribute hosted_page_group_name string length must be at least 1, got: 0"),
			},
			{
				Config:      config2,
				ExpectError: regexp.MustCompile(`The argument "hosted_pages" is required, but no definition was found.`),
			},
			{
				Config:      config3,
				ExpectError: regexp.MustCompile(`Attribute hosted_pages set must contain at least 1 elements, got: 0`),
			},
		},
	})
}

// Immutable attribute hosted_page_group_name validation
func TestAccHostedPageResource_UniqueIdentifier(t *testing.T) {
	t.Parallel()

	updatedHostedPageGroupName := "Updated Hosted Page Group"

	resourceID := acctest.RandString(10)
	hostedPageGroupName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					resourceID,
					hostedPages,
				),
			},
			{
				Config: testAccHostedPageResourceConfig(
					updatedHostedPageGroupName,
					defaultLocale,
					resourceID,
					hostedPages,
				),
				ExpectError: regexp.MustCompile("Attribute 'hosted_page_group_name' can't be modified"),
			},
		},
	})
}

// Invalid hosted_page_id
// func TestAccHostedPageResource_InvalidHostedPageID(t *testing.T) {
// 	InvalidHostedPageID := "invalid"
// 	hostedPages := []map[string]string{
// 		{
// 			"hosted_page_id": InvalidHostedPageID,
// 			"locale":         defaultLocale,
// 			"url":            hostedPageURL,
// 			"content":        hostedPageContent,
// 		},
// 	}
// 	resource.Test(t, resource.TestCase{
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccHostedPageResourceConfig(
// 					hostedPageGroupName,
// 					defaultLocale,
// 					hostedPages,
// 				),
// 				ExpectError: regexp.MustCompile("hosted_page_id value must be one of"), // TODO: full string comparison
// 			},
// 		},
// 	})
// }

// validate multiple hosted pages
// func TestAccHostedPageResource_MultipleHostedPages(t *testing.T) {
// 	hostedPageID2 := "login_success"
// 	hostedPageURL2 := "https://cidaad.de/login_success"
// 	hostedPageContent2 := "<html>Login Success</html>"

// 	config := fmt.Sprintf(`
// 	provider "cidaas" {
// 		base_url = "%s"
// 	}
// 	resource "cidaas_hosted_page" "example" {
// 		hosted_page_group_name = "`+hostedPageGroupName+`"
// 		default_locale = "`+defaultLocale+`"

// 		hosted_pages =[
// 			{
// 				hosted_page_id = "`+hostedPageID+`"
// 				locale = "`+defaultLocale+`"
// 				url = "`+hostedPageURL+`"
// 				content = "`+hostedPageContent+`"
// 		 },
// 		 {
// 				hosted_page_id = "`+hostedPageID2+`"
// 				locale = "en-IN"
// 				url = "`+hostedPageURL2+`"
// 				content = "`+hostedPageContent2+`"
// 		 }
// 		]
// 	}
// 	`, acctest.GetBaseURL())
// 	resource.Test(t, resource.TestCase{
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		Steps: []resource.TestStep{
// 			{
// 				Config: config,
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.hosted_page_id", hostedPageID),
// 					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.url", hostedPageURL),
// 					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.content", hostedPageContent),
// 					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.1.hosted_page_id", hostedPageID2),
// 					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.1.url", hostedPageURL2),
// 					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.1.content", hostedPageContent2),
// 				),
// 			},
// 		},
// 	})
// }
