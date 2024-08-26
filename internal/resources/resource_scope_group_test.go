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
	resourceScopeGroup    = "cidaas_scope_group.example"
	scopeGroupdescription = "Test Scope Group Description"
)

var scopeGroupName = acctest.RandString(10)

// create, read and update test
func TestAccScopeGroupResource_Basic(t *testing.T) {
	updatedDescription := "Updated Scope Group Description"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckScopeGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccScopeGroupResourceConfig(scopeGroupName, scopeGroupdescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceScopeGroup, "group_name", scopeGroupName),
					resource.TestCheckResourceAttr(resourceScopeGroup, "description", scopeGroupdescription),
					resource.TestCheckResourceAttrSet(resourceScopeGroup, "id"),
					resource.TestCheckResourceAttrSet(resourceScopeGroup, "created_at"),
					resource.TestCheckResourceAttrSet(resourceScopeGroup, "updated_at"),
				),
			},
			{
				ResourceName:      resourceScopeGroup,
				ImportStateId:     scopeGroupName,
				ImportState:       true,
				ImportStateVerify: true,
				// TODO: remove ImportStateVerifyIgnore
				ImportStateVerifyIgnore: []string{"updated_at", "created_at"},
			},
			{
				Config: testAccScopeGroupResourceConfig(scopeGroupName, updatedDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceScopeGroup, "description", updatedDescription),
					resource.TestCheckResourceAttrSet(resourceScopeGroup, "updated_at"),
				),
			},
		},
	})
}

func testAccScopeGroupResourceConfig(groupType, description string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "https://kube-nightlybuild-dev.cidaas.de"
	}
	resource "cidaas_scope_group" "example" {
		group_name  = "%s"
		description = "%s"
	}
	`, groupType, description)
}

func testCheckScopeGroupDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceScopeGroup]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceScopeGroup)
	}

	scopeGroup := cidaas.ScopeGroup{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := scopeGroup.Get(rs.Primary.Attributes["group_name"])
	if res.Data.ID != "" {
		// when resource exits in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// failed validation on updating immutable proprty group_name
func TestAccScopeGroupResource_GoupNameUpdateFail(t *testing.T) {
	updateGroupName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccScopeGroupResourceConfig(scopeGroupName, scopeGroupdescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceScopeGroup, "group_name", scopeGroupName),
				),
			},
			{
				Config:      testAccScopeGroupResourceConfig(updateGroupName, scopeGroupdescription),
				ExpectError: regexp.MustCompile("Attribute 'group_name' can't be modified"), // TODO: full string comparison
			},
		},
	})
}

// Empty group_name validation test
func TestAccScopeGroupResource_EmptyGroupName(t *testing.T) {
	emptyScopeGroupName := ""

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccScopeGroupResourceConfig(emptyScopeGroupName, scopeGroupdescription),
				ExpectError: regexp.MustCompile(`Attribute group_name string length must be at least 1, got: 0`),
			},
		},
	})
}

// missing required parameter
func TestAccScopeGroupResource_MissingRequired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				provider "cidaas" {
					base_url = "https://kube-nightlybuild-dev.cidaas.de"
				}
				resource "cidaas_scope_group" "example" {
					description = "test description"
				}
				`,
				ExpectError: regexp.MustCompile(`The argument "group_name" is required, but no definition was found.`),
			},
		},
	})
}
