package resources_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/internal/resources"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGroupTypeResource_Basic(t *testing.T) {
	t.Parallel()

	roleMode := "any_roles"
	updatedDescription := "Updated Group Type Description"
	groupType := acctest.RandString(10)
	description := "Test Group Type Description"

	testResourceID := acctest.RandString(10)
	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_GROUP_TYPE, testResourceID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckGroupTypeDestroyed(testResourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccGroupTypeResourceConfig(groupType, roleMode, description, testResourceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "group_type", groupType),
					resource.TestCheckResourceAttr(testResourceName, "role_mode", roleMode),
					resource.TestCheckResourceAttr(testResourceName, "description", description),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "created_at"),
				),
			},
			{
				ResourceName:      testResourceName,
				ImportStateId:     groupType,
				ImportState:       true,
				ImportStateVerify: true,
				// remove ImportStateVerifyIgnore to enhance the result
				ImportStateVerifyIgnore: []string{"updated_at", "created_at"},
			},
			{
				Config: testAccGroupTypeResourceConfig(groupType, roleMode, updatedDescription, testResourceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "description", updatedDescription),
					resource.TestCheckResourceAttrSet(testResourceName, "updated_at"),
				),
			},
		},
	})
}

func testAccGroupTypeResourceConfig(groupType, roleMode, description, resourceID string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "%s"
	}
	resource "cidaas_group_type" "%s" {
		group_type  = "%s"
		role_mode   = "%s"
		description = "%s"
	}
	`, acctest.GetBaseURL(), resourceID, groupType, roleMode, description)
}

func testCheckGroupTypeDestroyed(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}

		groupType := cidaas.GroupType{
			ClientConfig: cidaas.ClientConfig{
				BaseURL:     os.Getenv("BASE_URL"),
				AccessToken: acctest.TestToken,
			},
		}
		res, _ := groupType.Get(context.Background(), rs.Primary.Attributes["group_type"])

		if res != nil {
			// when resource exists in remote
			return fmt.Errorf("resource %s still exists", res.Data.GroupType)
		}
		return nil
	}
}

// validation test for role_mode
func TestAccGroupTypeResource_InvalidRoleMode(t *testing.T) {
	t.Parallel()

	invalidRoleMode := "invalid_role_mode"
	groupType := acctest.RandString(10)
	description := "Test Group Type Description"

	testResourceID := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccGroupTypeResourceConfig(groupType, invalidRoleMode, description, testResourceID),
				ExpectError: regexp.MustCompile(`Attribute role_mode value must be one of:`), // TODO: full string comparison
			},
		},
	})
}

// group_type can't be modified
func TestAccGroupTypeResource_UpdateFails(t *testing.T) {
	t.Parallel()

	updatedGroupType := acctest.RandString(10)
	roleMode := "any_roles"
	groupType := acctest.RandString(10)
	description := "Test Group Type Description"

	testResourceID := acctest.RandString(10)
	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_GROUP_TYPE, testResourceID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupTypeResourceConfig(groupType, roleMode, description, testResourceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "group_type", groupType),
				),
			},
			{
				Config:      testAccGroupTypeResourceConfig(updatedGroupType, roleMode, description, testResourceID),
				ExpectError: regexp.MustCompile("Attribute 'group_type' can't be modified"), // TODO: full string comparison
			},
		},
	})
}

// allowed_roles must have value when role_mode is allowed_roles or roles_required
func TestAccGroupTypeResource_EmptyAllowedRolesError(t *testing.T) {
	t.Parallel()

	roleMode := "allowed_roles"
	groupType := acctest.RandString(10)
	description := "Test Group Type Description"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_group_type" "%s" {
					group_type  = "%s"
					role_mode   = "%s"
					description = "%s"
					allowed_roles = []
				}
				`, acctest.GetBaseURL(), groupType, groupType, roleMode, description),
				ExpectError: regexp.MustCompile("The attribute allowed_roles cannot be empty when role_mode is set to"), // TODO: full string comparison
			},
		},
	})
}
