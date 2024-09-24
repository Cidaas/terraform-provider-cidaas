package resources_test

import (
	"fmt"
	"regexp"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const resourceName = "cidaas_role.example"

func TestAccRoleResource(t *testing.T) {
	role := acctest.RandString(10)
	name := "Terraform Admin Role"
	updatedName := "Updated Terraform Admin Role"
	description := "This is a test terraform admin role"
	updatedDescription := "This is a test terraform admin updated role"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// verify check destroy
		CheckDestroy: testAccCheckRoleResourceDestroyed,
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
			// ImportState testing(read after create)
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "role", role),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			// Update and Read testing
			{
				Config: testAccRoleResourceConfig(role, updatedName, updatedDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoleResource_validateAttrSet(t *testing.T) {
	role := acctest.RandString(10)
	name := "Test Terraform Admin Role"
	description := "This is a test terraform admin role"
	updatedDescription := "This is a test terraform admin updated role"
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
				ImportStateIdFunc:                    importStateIDFunc,
				ImportStateVerify:                    true,
			},
			{
				Config: testAccRoleResourceConfig(role, name, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "role"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
				),
			},
		},
	})
}

func TestAccRoleResource_updateRoleFails(t *testing.T) {
	role := acctest.RandString(10)
	updatedRole := acctest.RandString(10)
	name := "Test Role"
	description := "This is a test role"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleResourceConfig(role, name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "role", role),
				),
			},
			{
				Config:      testAccRoleResourceConfig(updatedRole, name, description),
				ExpectError: regexp.MustCompile("Attribute 'role' can't be modified"), // TODO: full string comparison
			},
		},
	})
}

func TestAccRoleResource_readAfterUpdate(t *testing.T) {
	role := acctest.RandString(10)
	name := "Test Role"
	description := "This is a test role"

	updatedName := "Updated Test Role"
	updatedDescription := "This is an updated test role"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleResourceConfig(role, name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "role", role),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			{
				Config: testAccRoleResourceConfig(role, updatedName, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
				),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "role", role),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
				),
			},
		},
	})
}

func TestAccRoleResource_createMissingFields(t *testing.T) {
	missingRoleConfig := fmt.Sprintf(`
	provider "cidaas" {
		base_url = "%s"
	}
	resource "cidaas_role" "example" {
		name = "Test Name"
		description = "Test Description"
	}`, acctest.BaseURL)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      missingRoleConfig,
				ExpectError: regexp.MustCompile(`The argument "role" is required, but no definition was found.`), // TODO: full string comparison
			},
		},
	})
}

func testAccRoleResourceConfig(role, name, description string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "%s"
	}
	resource "cidaas_role" "example" {
		role = "%s"
		name = "%s"
		description = "%s"
	}`, acctest.BaseURL, role, name, description)
}

func importStateIDFunc(s *terraform.State) (string, error) {
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return "", fmt.Errorf("resource not found")
	}
	id, ok := rs.Primary.Attributes["role"]
	if !ok {
		return "", fmt.Errorf("role not set")
	}
	return id, nil
}

func testAccCheckRoleResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Role ID is set")
		}
		return nil
	}
}

func testAccCheckRoleResourceDestroyed(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cidaas_role" {
			continue
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}
		// TODO: check in cidaas by calling get role function if this is required
	}
	return nil
}
