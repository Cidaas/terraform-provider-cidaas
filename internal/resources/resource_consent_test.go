package resources_test

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const (
	resourceConsent         = "cidaas_consent.example"
	refResourceConsentGroup = "cidaas_consent_group.example"
)

// create, read and update test
func TestAccConsentResource_Basic(t *testing.T) {
	groupName := acctest.RandString(10)
	name := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckConsentDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccConsentResourceConfig(groupName, name, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceConsent, "consent_group_id", refResourceConsentGroup, "id"),
					resource.TestCheckResourceAttr(resourceConsent, "name", name),
					resource.TestCheckResourceAttrSet(resourceConsent, "id"),
					resource.TestCheckResourceAttrSet(resourceConsent, "enabled"),
					resource.TestCheckResourceAttrSet(resourceConsent, "created_at"),
					resource.TestCheckResourceAttrSet(resourceConsent, "updated_at"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cidaas_consent.example",
						tfjsonpath.New("enabled"),
						knownvalue.Bool(true),
					),
				},
			},
			{
				ResourceName: resourceConsent,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[refResourceConsentGroup]
					if !ok {
						return "", fmt.Errorf("Not found: %s", refResourceConsentGroup)
					}
					return rs.Primary.ID + ":" + name, nil
				},
				ImportState:       true,
				ImportStateVerify: true,
				// TODO: remove ImportStateVerifyIgnore
				ImportStateVerifyIgnore: []string{"updated_at", "created_at"},
			},
			{
				Config: testAccConsentResourceConfig(groupName, name, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceConsent, "updated_at"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cidaas_consent.example",
						tfjsonpath.New("enabled"),
						knownvalue.Bool(false),
					),
				},
			},
		},
	})
}

func testAccConsentResourceConfig(groupName, name string, enabled bool) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "%s"
	}
	resource "cidaas_consent_group" "example" {
		group_name  = "%s"
	}
	resource "cidaas_consent" "example" {
		consent_group_id  = cidaas_consent_group.example.id
		name = "%s"
		enabled = "%v"
	}
	`, acctest.BaseURL, groupName, name, enabled)
}

func testCheckConsentDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[refResourceConsentGroup]
	if !ok {
		return fmt.Errorf("resource %s not fround", refResourceConsentGroup)
	}

	consent := cidaas.ConsentClient{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := consent.GetConsentInstances(rs.Primary.ID)
	if res != nil && res.Status != http.StatusNoContent && len(res.Data) > 0 {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// failed validation on updating immutable proprty group_name
func TestAccConsentResource_GoupNameUpdateFail(t *testing.T) {
	groupName := acctest.RandString(10)
	name := acctest.RandString(10)
	updatedName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConsentResourceConfig(groupName, name, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceConsent, "consent_group_id", refResourceConsentGroup, "id"),
				),
			},
			{
				Config:      testAccConsentResourceConfig(groupName, updatedName, true),
				ExpectError: regexp.MustCompile(`Attribute 'name' can't be modified.`),
			},
		},
	})
}

// empty consent_group_id & group_name validation test
func TestAccConsentResource_EmptyGroupName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_consent" "example" {
					consent_group_id  = ""
					name = ""
				}
				`, acctest.BaseURL),
				ExpectError: regexp.MustCompile(`Attribute consent_group_id string length must be at least 1, got: 0`),
			},
		},
	})
}

// missing required parameters
func TestAccConsentResource_MissingRequired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_consent" "example" {
				}
				`, acctest.BaseURL),
				ExpectError: regexp.MustCompile(`The argument "name" is required, but no definition was found.`),
			},
		},
	})
}
