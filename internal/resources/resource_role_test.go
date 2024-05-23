package resources_test

import (
	"fmt"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccRoleResource(t *testing.T) {
	resourceName := "cidaas_role.example"
	role := "terraform_admin_role"
	name := "Test Terraform Admin Role"
	description := "This is a test terraform admin role"
	updatedDescription := "This is a test terraform admin updated role"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccRoleResourceConfig(role, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "role", role),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "id", role),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccRoleResourceConfig(role, name, updatedDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoleResource_Validation(t *testing.T) {
	resourceName := "cidaas_role.example"
	role := "terraform_admin_role"
	name := "Test Terraform Admin Role"
	description := acctest.RandString(16)
	description2 := acctest.RandString(16)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccRoleResourceConfig(role, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "role"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateIdFunc:                    testAccProjectImportID,
				ImportStateVerify:                    true,
			},
			{
				Config: testAccRoleResourceConfig(role, name, description2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "role"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
				),
			},
		},
	})
}

func testAccRoleResourceConfig(role, name, description string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "https://kube-nightlybuild-dev.cidaas.de"
	}
	resource "cidaas_role" "example" {
		role = "%s"
		name = "%s"
		description = "%s"
	}`, role, name, description)
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
