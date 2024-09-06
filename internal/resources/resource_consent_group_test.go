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
	resourceConsentGroup = "cidaas_consent_group.example"
	description          = "Test consent Description"
)

var groupName = acctest.RandString(10)

// create, read and update test
func TestAccConsentGroupResource_Basic(t *testing.T) {
	updatedDescription := "Updated consent Description"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckConsentGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccConsentGroupResourceConfig(groupName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceConsentGroup, "group_name", groupName),
					resource.TestCheckResourceAttr(resourceConsentGroup, "description", description),
					resource.TestCheckResourceAttrSet(resourceConsentGroup, "id"),
					resource.TestCheckResourceAttrSet(resourceConsentGroup, "created_at"),
					resource.TestCheckResourceAttrSet(resourceConsentGroup, "updated_at"),
				),
			},
			{
				ResourceName:                         resourceConsentGroup,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportState:                          true,
				ImportStateVerify:                    true,
				// TODO: remove ImportStateVerifyIgnore
				ImportStateVerifyIgnore: []string{"updated_at", "created_at"},
			},
			{
				Config: testAccConsentGroupResourceConfig(groupName, updatedDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceConsentGroup, "description", updatedDescription),
					resource.TestCheckResourceAttrSet(resourceConsentGroup, "updated_at"),
				),
			},
		},
	})
}

func testAccConsentGroupResourceConfig(groupName, description string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "https://kube-nightlybuild-dev.cidaas.de"
	}
	resource "cidaas_consent_group" "example" {
		group_name  = "%s"
		description = "%s"
	}
	`, groupName, description)
}

func testCheckConsentGroupDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceConsentGroup]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceConsentGroup)
	}

	consentGroup := cidaas.ConsentGroup{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := consentGroup.Get(rs.Primary.ID)
	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// failed validation on updating immutable proprty group_name
func TestAccConsentGroupResource_GoupNameUpdateFail(t *testing.T) {
	updateGroupName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConsentGroupResourceConfig(groupName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceConsentGroup, "group_name", groupName),
				),
			},
			{
				Config:      testAccConsentGroupResourceConfig(updateGroupName, description),
				ExpectError: regexp.MustCompile(`Attribute 'group_name' can't be modified.`),
			},
		},
	})
}

// Empty group_name validation test
func TestAccConsentGroupResource_EmptyGroupName(t *testing.T) {
	emptyGroupName := ""

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccConsentGroupResourceConfig(emptyGroupName, description),
				ExpectError: regexp.MustCompile(`Attribute group_name string length must be at least 1, got: 0`),
			},
		},
	})
}

// missing required parameter
func TestAccConsentGroupResource_MissingRequired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				provider "cidaas" {
					base_url = "https://kube-nightlybuild-dev.cidaas.de"
				}
				resource "cidaas_consent_group" "example" {
					description = "test description"
				}
				`,
				ExpectError: regexp.MustCompile(`The argument "group_name" is required, but no definition was found.`),
			},
		},
	})
}
