package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccRoleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccExampleResourceConfig("terraform-acceptance-test-role", "terraform-acceptance-test-role-desciption"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cidaas_role.example", "name", "terraform-acceptance-test-role-name"),
					resource.TestCheckResourceAttr("cidaas_role.example", "role", "terraform-acceptance-test-role"),
					resource.TestCheckResourceAttr("cidaas_role.example", "description", "terraform-acceptance-test-role-desciption"),
					resource.TestCheckResourceAttr("cidaas_role.example", "id", "terraform-acceptance-test-role"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "cidaas_role.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccExampleResourceConfig("terraform-acceptance-test-role", "terraform-acceptance-test-role-updated-desciption"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cidaas_role.example", "description", "terraform-acceptance-test-role-updated-desciption"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoleResource_Validation(t *testing.T) {
	roleName := RandString(16)
	description := RandString(16)
	description2 := RandString(16)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccExampleResourceConfig(roleName, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("cidaas_role.example", "role"),
					resource.TestCheckResourceAttrSet("cidaas_role.example", "name"),
					resource.TestCheckResourceAttrSet("cidaas_role.example", "description"),
					resource.TestCheckResourceAttrSet("cidaas_role.example", "id"),
				),
			},
			{
				ResourceName:                         "cidaas_role.example",
				ImportState:                          true,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateIdFunc:                    testAccProjectImportID,
				ImportStateVerify:                    true,
			},
			{
				Config: testAccExampleResourceConfig(roleName, description2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("cidaas_role.example", "role"),
					resource.TestCheckResourceAttrSet("cidaas_role.example", "name"),
					resource.TestCheckResourceAttrSet("cidaas_role.example", "description"),
				),
			},
		},
	})
}

func testAccExampleResourceConfig(role string, description string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "https://kube-nightlybuild-dev.cidaas.de"
	}
	resource "cidaas_role" "example" {
		role = %[1]q
		name = "terraform-acceptance-test-role-name"
		description = %[2]q
	}`, role, description)
}

func testAccProjectImportID(s *terraform.State) (string, error) {
	rs, ok := s.RootModule().Resources["cidaas_role.example"]
	if !ok {
		return "", fmt.Errorf("resource not found")
	}

	id, ok := rs.Primary.Attributes["role"]
	if !ok {
		return "", fmt.Errorf("role not set")
	}

	return id, nil
}
