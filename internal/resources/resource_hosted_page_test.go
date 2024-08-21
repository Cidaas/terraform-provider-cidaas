package resources_test

import (
	"regexp"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const resourceHostedPage = "cidaas_hosted_page.test"

func TestAccHostedPageResource_Basic(t *testing.T) {
	hostedPageID := "register_success"
	hostedPageURL := "https://cidaad.de/register_success"
	updatedHostedPageURL := "https://cidaad.de/updated_register_success"
	hostedPageContent := "<html>Success</html>"
	hostedPageGroupName := "Test Hosted Page Group"
	defaultLocale := "en-US"
	hostedPages := []map[string]string{
		{
			"hosted_page_id": hostedPageID,
			"locale":         defaultLocale,
			"url":            hostedPageURL,
			"content":        hostedPageContent,
		},
	}
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
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccHostedPageResourceConfig(
					hostedPageGroupName,
					defaultLocale,
					hostedPages,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_page_group_name", hostedPageGroupName),
					resource.TestCheckResourceAttr(resourceHostedPage, "default_locale", defaultLocale),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.hosted_page_id", hostedPageID),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.locale", defaultLocale),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.url", hostedPageURL),
					resource.TestCheckResourceAttr(resourceHostedPage, "hosted_pages.0.content", hostedPageContent),
				),
			},
			// ImportState testing
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
			// Update
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
	return `
	provider "cidaas" {
		base_url = "https://kube-nightlybuild-dev.cidaas.de"
	}
	resource "cidaas_hosted_page" "test" {
		hosted_page_group_name = "` + hostedPageGroupName + `"
		default_locale = "` + defaultLocale + `"

		hosted_pages =[
			{
			hosted_page_id = "` + hostedPages[0]["hosted_page_id"] + `"
			locale = "` + hostedPages[0]["locale"] + `"
			url = "` + hostedPages[0]["url"] + `"
			content = "` + hostedPages[0]["content"] + `"
		}
	]
}
`
}

func testAccCheckHostedPageResourceDestroy(s *terraform.State) error {
	// Implement your destroy check logic here, usually checking that the resource
	// does not exist in the real system.
	return nil
}

// invalid locale
func TestAccHostedPageResource_InvalidLocale(t *testing.T) {
	hostedPageGroupName := "Test Hosted Page Group"
	invalidLocale := "invalid-locale"
	hostedPages := []map[string]string{
		{
			"hosted_page_id": "register_success",
			"locale":         "en-US",
			"url":            "https://cidaad.de/register_success",
			"content":        "<html>Success</html>",
		},
	}
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
	config1 := `
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_hosted_page" "test" {
			hosted_page_group_name = ""
			default_locale = "en-US"
			hosted_pages =[{
				hosted_page_id = "register_success"
				url = ""
			}]
		}
		`
	config2 := `
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_hosted_page" "test" {
			hosted_page_group_name = ""
			default_locale = "en-US"
		}
		`
	config3 := `
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_hosted_page" "test" {
			hosted_page_group_name = ""
			default_locale = "en-US"
			hosted_pages =[]
		}
		`
	// validation where hosted_page_id and url is required
	// config4 := `
	// 	provider "cidaas" {
	// 		base_url = "https://kube-nightlybuild-dev.cidaas.de"
	// 	}
	// 	resource "cidaas_hosted_page" "test" {
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
