package resources_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	resourceHostedPage = "cidaas_hosted_page.example"
	hostedPageURL      = "https://cidaad.de/register_success"
	hostedPageID       = "register_success"
	hostedPageContent  = "<html>Register Success</html>"
	defaultLocale      = "en-US"
)

var (
	hostedPageGroupName = acctest.RandString(10)
	hostedPages         = []map[string]string{
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
	updatedHostedPageURL := "https://cidaad.de/updated_register_success"
	updatedHostedPages := []map[string]string{
		{
			"hosted_page_id": hostedPageID,
			"locale":         defaultLocale,
			"url":            updatedHostedPageURL,
			"content":        "<html>Updated Success</html>",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckHostedPageDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					hostedPages,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckHostedPageExists(resourceHostedPage),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_page_group_name", hostedPageGroupName),
					resource.TestCheckResourceAttr(resourceHostedPage, "default_locale", defaultLocale),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.hosted_page_id", hostedPageID),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.locale", defaultLocale),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.url", hostedPageURL),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.content", hostedPageContent),
				),
			},
			{
				ResourceName:            resourceHostedPage,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "updated_at"},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_page_group_name", hostedPageGroupName),
					resource.TestCheckResourceAttr(resourceHostedPage, "default_locale", defaultLocale),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.hosted_page_id", hostedPageID),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.locale", defaultLocale),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.url", hostedPageURL),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.content", hostedPageContent),
				),
			},
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					updatedHostedPages,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.url", updatedHostedPageURL),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.content", "<html>Updated Success</html>"),
				),
			},
		},
	})
}

func testAccHostedPageResourceConfig(hostedPageGroupName, defaultLocale string, hostedPages []map[string]string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "example" {
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
	`, acctest.BaseURL)
}

func testCheckHostedPageExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if _, ok := s.RootModule().Resources[resourceName]; !ok {
			return fmt.Errorf("Not found: %s, found resource %s", resourceName, s.RootModule().Resources)
		}
		return nil
	}
}

func testCheckHostedPageDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceHostedPage]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceHostedPage)
	}

	hp := cidaas.HostedPage{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := hp.Get(rs.Primary.Attributes["hosted_page_group_name"])

	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource %s stil exists", res.Data)
	}
	return nil
}

// invalid locale
func TestAccHostedPageResource_InvalidLocale(t *testing.T) {
	invalidLocale := "invalid-locale"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      testAccHostedPageResourceConfig(hostedPageGroupName, invalidLocale, hostedPages),
				ExpectError: regexp.MustCompile("Attribute default_locale value must be one of"),
			},
		},
	})
}

// missing required fields validation
func TestAccHostedPageResource_MissingRequiredFields(t *testing.T) {
	config1 := fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "example" {
			hosted_page_group_name = ""
			default_locale = "en-US"
			hosted_pages =[{
				hosted_page_id = "register_success"
				url = ""
			}]
		}
		`, acctest.BaseURL)
	config2 := fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "example" {
			hosted_page_group_name = ""
			default_locale = "en-US"
		}
		`, acctest.BaseURL)
	config3 := fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_hosted_page" "example" {
			hosted_page_group_name = ""
			default_locale = "en-US"
			hosted_pages =[]
		}
		`, acctest.BaseURL)
	// validation where hosted_page_id and url is required
	// config4 := `
	// 	provider "cidaas" {
	// 		base_url = "https://automation-test.dev.cidaas.eu"
	// 	}
	// 	resource "cidaas_hosted_page" "example" {
	// 		hosted_page_group_name = ""
	// 		default_locale = "en-US"
	// 		hosted_pages =[{}]
	// 	}
	// 	`
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
				ExpectError: regexp.MustCompile(`Attribute hosted_pages list must contain at least 1 elements, got: 0`),
			},
		},
	})
}

// Immutable attribute hosted_page_group_name validation
func TestAccHostedPageResource_UniqueIdentifier(t *testing.T) {
	updatedHostedPageGroupName := "Updated Hosted Page Group"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					hostedPages,
				),
			},
			{
				Config: testAccHostedPageResourceConfig(
					updatedHostedPageGroupName,
					defaultLocale,
					hostedPages,
				),
				ExpectError: regexp.MustCompile("Attribute 'hosted_page_group_name' can't be modified"),
			},
		},
	})
}

// Invalid hosted_page_id
func TestAccHostedPageResource_InvalidHostedPageID(t *testing.T) {
	InvalidHostedPageID := "invalid"
	hostedPages := []map[string]string{
		{
			"hosted_page_id": InvalidHostedPageID,
			"locale":         defaultLocale,
			"url":            hostedPageURL,
			"content":        hostedPageContent,
		},
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					hostedPages,
				),
				ExpectError: regexp.MustCompile("hosted_page_id value must be one of"), // TODO: full string comparison
			},
		},
	})
}

// validate multiple hosted pages
func TestAccHostedPageResource_MultipleHostedPages(t *testing.T) {
	hostedPageID2 := "login_success"
	hostedPageURL2 := "https://cidaad.de/login_success"
	hostedPageContent2 := "<html>Login Success</html>"

	config := fmt.Sprintf(`
	provider "cidaas" {
		base_url = "%s"
	}
	resource "cidaas_hosted_page" "example" {
		hosted_page_group_name = "`+hostedPageGroupName+`"
		default_locale = "`+defaultLocale+`"

		hosted_pages =[
			{
				hosted_page_id = "`+hostedPageID+`"
				locale = "`+defaultLocale+`"
				url = "`+hostedPageURL+`"
				content = "`+hostedPageContent+`"
		 },
		 {
				hosted_page_id = "`+hostedPageID2+`"
				locale = "en-IN"
				url = "`+hostedPageURL2+`"
				content = "`+hostedPageContent2+`"
		 }
		]
	}
	`, acctest.BaseURL)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.hosted_page_id", hostedPageID),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.url", hostedPageURL),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.content", hostedPageContent),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.1.hosted_page_id", hostedPageID2),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.1.url", hostedPageURL2),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.1.content", hostedPageContent2),
				),
			},
		},
	})
}
