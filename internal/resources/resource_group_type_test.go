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
	resourceGroupType    = "cidaas_group_type.example"
	groupTypedescription = "Test Group Type Description"
)

var groupType = acctest.RandString(10)

// create, read and update test
func TestAccGroupTypeResource_Basic(t *testing.T) {
	roleMode := "any_roles"
	updatedDescription := "Updated Group Type Description"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckGroupTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupTypeResourceConfig(groupType, roleMode, groupTypedescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceGroupType, "group_type", groupType),
					resource.TestCheckResourceAttr(resourceGroupType, "role_mode", roleMode),
					resource.TestCheckResourceAttr(resourceGroupType, "description", groupTypedescription),
					resource.TestCheckResourceAttrSet(resourceGroupType, "id"),
					resource.TestCheckResourceAttrSet(resourceGroupType, "created_at"),
				),
			},
			{
				ResourceName:      resourceGroupType,
				ImportStateId:     groupType,
				ImportState:       true,
				ImportStateVerify: true,
				// remove ImportStateVerifyIgnore to enhance the result
				ImportStateVerifyIgnore: []string{"updated_at", "created_at"},
			},
			{
				Config: testAccGroupTypeResourceConfig(groupType, roleMode, updatedDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceGroupType, "description", updatedDescription),
					resource.TestCheckResourceAttrSet(resourceGroupType, "updated_at"),
				),
			},
		},
	})
}

func testAccGroupTypeResourceConfig(groupType, roleMode, description string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "%s"
	}
	resource "cidaas_group_type" "example" {
		group_type  = "%s"
		role_mode   = "%s"
		description = "%s"
	}
	`, acctest.BaseURL, groupType, roleMode, description)
}

func testCheckGroupTypeDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceGroupType]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceGroupType)
	}

	groupType := cidaas.GroupType{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := groupType.Get(rs.Primary.Attributes["group_type"])

	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource %s stil exists", res.Data.GroupType)
	}
	return nil
}

// validation test for role_mode
func TestAccGroupTypeResource_InvalidRoleMode(t *testing.T) {
	invalidRoleMode := "invalid_role_mode"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccGroupTypeResourceConfig(groupType, invalidRoleMode, groupTypedescription),
				ExpectError: regexp.MustCompile(`Attribute role_mode value must be one of:`), // TODO: full string comparison
			},
		},
	})
}

// group_type can't be modified
func TestAccGroupTypeResource_UpdateFails(t *testing.T) {
	updatedGroupType := acctest.RandString(10)
	roleMode := "any_roles"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupTypeResourceConfig(groupType, roleMode, groupTypedescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceGroupType, "group_type", groupType),
				),
			},
			{
				Config:      testAccGroupTypeResourceConfig(updatedGroupType, roleMode, groupTypedescription),
				ExpectError: regexp.MustCompile("Attribute 'group_type' can't be modified"), // TODO: full string comparison
			},
		},
	})
}

// allowed_roles must have value when role_mode is allowed_roles or roles_required
func TestAccGroupTypeResource_EmptyAllowedRolesError(t *testing.T) {
	roleMode := "allowed_roles"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_group_type" "example" {
					group_type  = "%s"
					role_mode   = "%s"
					description = "%s"
					allowed_roles = []
				}
				`, acctest.BaseURL, groupType, roleMode, groupTypedescription),
				ExpectError: regexp.MustCompile("The attribute allowed_roles cannot be empty when role_mode is set to"), // TODO: full string comparison
			},
		},
	})
}
