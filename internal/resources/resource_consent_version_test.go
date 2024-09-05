package resources_test

import (
	"fmt"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	resourceConsentVersion = "cidaas_consent_version.example"
)

// create, read and update test
func TestConsentVersion_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConsentVersionConfig("consent version in German"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceConsentVersion, "consent_type", "SCOPES"),
					resource.TestCheckResourceAttrSet(resourceConsentVersion, "id"),
				),
			},
			{
				ResourceName:      resourceConsentVersion,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceConsentVersion]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceConsentVersion)
					}
					return rs.Primary.Attributes["consent_id"] + ":" + rs.Primary.ID + ":de:en", nil
				},
			},
			{
				Config: testConsentVersionConfig("updated consent version in German"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceConsentVersion, "consent_type", "SCOPES"),
				),
			},
		},
	})
}

func testConsentVersionConfig(content string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_consent_group" "sample" {
			group_name  = "sample_consent_group"
			description = "sample description"
		}
		resource "cidaas_consent" "sample" {
			consent_group_id = cidaas_consent_group.sample.id
			name             = "sample_consent"
			enabled          = true
		}
		resource "cidaas_consent_version" "example" {
			version         = 1
			consent_id      = cidaas_consent.sample.id
			consent_type    = "SCOPES"
			scopes          = ["developer"]
			required_fields = ["name"]
			consent_locales = [
				{
					content = "%s"
					locale  = "de"
				},
				{
					content = "consent version in English"
					locale  = "en"
				}
			]
		}		
	`, content)
}
