package resources_test

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/internal/resources"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccConsentResource_Basic(t *testing.T) {
	t.Parallel()

	groupName := acctest.RandString(10)
	name := acctest.RandString(10)

	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_CONSENT, groupName)
	consentGroupResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_CONSENT_GROUP, "example")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckConsentDestroyed(consentGroupResourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccConsentResourceConfig(groupName, name, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(testResourceName, "consent_group_id", consentGroupResourceName, "id"),
					resource.TestCheckResourceAttr(testResourceName, "name", name),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "enabled"),
					resource.TestCheckResourceAttrSet(testResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(testResourceName, "updated_at"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						testResourceName,
						tfjsonpath.New("enabled"),
						knownvalue.Bool(true),
					),
				},
			},
			{
				ResourceName: testResourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[consentGroupResourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", consentGroupResourceName)
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
					resource.TestCheckResourceAttrSet(testResourceName, "updated_at"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						testResourceName,
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
	resource "cidaas_consent" "%s" {
		consent_group_id  = cidaas_consent_group.example.id
		name = "%s"
		enabled = "%v"
	}
	`, acctest.GetBaseURL(), groupName, groupName, name, enabled)
}

func testCheckConsentDestroyed(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}

		consent := cidaas.Consent{
			ClientConfig: cidaas.ClientConfig{
				BaseURL:     acctest.GetBaseURL(),
				AccessToken: acctest.TestToken,
			},
		}
		res, _ := consent.GetConsentInstances(context.Background(), rs.Primary.ID)
		if res != nil && res.Status != http.StatusNoContent && len(res.Data) > 0 {
			// when resource exists in remote
			return fmt.Errorf("resource still exists %+v", res)
		}
		return nil
	}
}

// failed validation on updating immutable proprty group_name
func TestAccConsentResource_GroupNameUpdateFail(t *testing.T) {
	t.Parallel()
	groupName := acctest.RandString(10)
	name := acctest.RandString(10)
	updatedName := acctest.RandString(10)

	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_CONSENT, groupName)
	consentGroupResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_CONSENT_GROUP, "example")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConsentResourceConfig(groupName, name, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(testResourceName, "consent_group_id", consentGroupResourceName, "id"),
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
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_consent" "%s" {
					consent_group_id  = ""
					name = ""
				}
				`, acctest.GetBaseURL(), acctest.RandString(10)),
				ExpectError: regexp.MustCompile(`Attribute consent_group_id string length must be at least 1, got: 0`),
			},
		},
	})
}

// missing required parameters
func TestAccConsentResource_MissingRequired(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_consent" "%s" {
				}
				`, acctest.GetBaseURL(), acctest.RandString(10)),
				ExpectError: regexp.MustCompile(`The argument "name" is required, but no definition was found.`),
			},
		},
	})
}
