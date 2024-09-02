package resources_test

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	resourceUserGroup = "cidaas_user_groups.example"
)

var (
	userGroupType = acctest.RandString(10)
	groupID       = acctest.RandString(10)
	userGroupName = acctest.RandString(10)
)

// create, read and update test
func TestUserGroup_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckUserGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserGroupResourceConfig(userGroupType, groupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceUserGroup, "group_type", resourceGroupType, "group_type"),
					resource.TestCheckResourceAttr(resourceUserGroup, "group_id", groupID),
					resource.TestCheckResourceAttr(resourceUserGroup, "group_name", userGroupName),
					resource.TestCheckResourceAttr(resourceUserGroup, "logo_url", "https://cidaas.de/logo"),
					resource.TestCheckResourceAttr(resourceUserGroup, "description", "sample user groups description"),
					resource.TestCheckResourceAttr(resourceUserGroup, "custom_fields.first_name", "cidaas"),
					resource.TestCheckResourceAttr(resourceUserGroup, "custom_fields.family_name", "widas"),
					// default value check
					resource.TestCheckResourceAttr(resourceUserGroup, "make_first_user_admin", strconv.FormatBool(true)),
					resource.TestCheckResourceAttr(resourceUserGroup, "member_profile_visibility", "full"),
					resource.TestCheckResourceAttr(resourceUserGroup, "none_member_profile_visibility", "public"),
					resource.TestCheckResourceAttr(resourceUserGroup, "parent_id", "root"),
					// computed properties check
					resource.TestCheckResourceAttrSet(resourceUserGroup, "id"),
					resource.TestCheckResourceAttrSet(resourceUserGroup, "created_at"),
					resource.TestCheckResourceAttrSet(resourceUserGroup, "updated_at"),
				),
			},
			{
				ResourceName:            resourceUserGroup,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           groupID,
				ImportStateVerifyIgnore: []string{"created_at", "updated_at"},
			},
			{
				Config:      testAccUserGroupResourceConfig(userGroupType, acctest.RandString(10)),
				ExpectError: regexp.MustCompile("Attribute 'group_id' can't be modified"),
			},
			// update works, but with update the order of deletion of both the resources changes group_type first and user groups next
			// and we get conflict. so have these lines commented.
			// TODO: check if the order of destruction can be controlled
			// {
			// 	Config: updatedUserGroupConfig(newRandomStr),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "group_type", userGroupType),
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "group_name", newRandomStr),
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "logo_url", "https://cidaas.de/v2/logo"),
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "description", "updated sample user groups description"),
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "custom_fields.first_name", "rob"),
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "custom_fields.family_name", "pike"),
			// 		// default value check
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "make_first_user_admin", strconv.FormatBool(false)),
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "member_profile_visibility", "public"),
			// 		resource.TestCheckResourceAttr(resourceUserGroup, "none_member_profile_visibility", "none"),
			// 	),
			// },
		},
	})
}

func testAccUserGroupResourceConfig(groupType, groupID string) string {
	return `
		provider "cidaas" {
			base_url = "https://kube-nightlybuild-dev.cidaas.de"
		}
		resource "cidaas_group_type" "example" {
			group_type  = "` + groupType + `"
			role_mode   = "no_roles"
			description = "group type description"
		}
		resource "cidaas_user_groups" "example" {
			group_type                     = cidaas_group_type.example.group_type
			group_id                       = "` + groupID + `"
			group_name                     = "` + userGroupName + `"
			logo_url                       = "https://cidaas.de/logo"
			description                    = "sample user groups description"
			custom_fields = {
				first_name  = "cidaas"
				family_name = "widas"
			}
			make_first_user_admin          = true
			member_profile_visibility      = "full"
			none_member_profile_visibility = "public"
		}		
	`
}

func testCheckUserGroupDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceUserGroup]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceUserGroup)
	}

	ug := cidaas.UserGroup{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := ug.Get(rs.Primary.Attributes["group_id"])
	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}

	rs, ok = s.RootModule().Resources[resourceGroupType]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceGroupType)
	}

	groupType := cidaas.GroupType{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	resp, _ := groupType.Get(rs.Primary.Attributes["group_type"])

	if resp != nil {
		// when resource exists in remote
		return fmt.Errorf("resource %s stil exists", resp.Data.GroupType)
	}
	return nil
}

// missing required fields group_type, group_name and group_id
func TestUserGroup_MissingRequired(t *testing.T) {
	requiredParams := []string{"group_type", "group_name", "group_id"}
	for _, v := range requiredParams {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: `
					provider "cidaas" {
						base_url = "https://kube-nightlybuild-dev.cidaas.de"
					}
					resource "cidaas_user_groups" "example" {}
				`,
					ExpectError: regexp.MustCompile(fmt.Sprintf(`The argument "%s" is required`, v)),
				},
			},
		})
	}
}

// check if group_type, group_name and group_id are empty string
func TestUserGroup_CheckEmptyString(t *testing.T) {
	requiredParams := []string{"group_type", "group_name", "group_id"}
	for _, v := range requiredParams {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: `
					provider "cidaas" {
						base_url = "https://kube-nightlybuild-dev.cidaas.de"
					}
					resource "cidaas_user_groups" "example" {
						group_type  =""
						group_id    = ""
						group_name  = ""
					}
				`,
					ExpectError: regexp.MustCompile(fmt.Sprintf(`Attribute %s string length must be at least 1`, v)),
				},
			},
		})
	}
}
