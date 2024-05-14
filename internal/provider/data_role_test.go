package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccRoleDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExampleDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cidaas_role.example", "description", "cidaas provider data source description"),
				),
			},
		},
	})
}

const testAccExampleDataSourceConfig = `
provider "cidaas" {
	base_url = "https://kube-nightlybuild-dev.cidaas.de"
}
resource "cidaas_role" "example_role" {
  name = "data source role name"
	role = "data_source_role_terraform_test"
	description = "cidaas provider data source description"
}
data "cidaas_role" "example" {
  role = cidaas_role.example_role.role
}
`

func TestAccRoleDataSourceExample(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccExampleDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cidaas_role.example", "role"),
					testAccRoleData(t, "data.cidaas_role.example"),
					func(s *terraform.State) error {
						resourceName := "data.cidaas_role.example"
						rs, ok := s.RootModule().Resources[resourceName]
						if !ok {
							return fmt.Errorf("Not found: %s", resourceName)
						}
						_, ok = rs.Primary.Attributes["description"]
						if !ok {
							return fmt.Errorf("Resource %s has no description set", resourceName)
						}

						return nil
					},
				),
			},
		},
	})
}

func testAccRoleData(t *testing.T, resourceName string) resource.TestCheckFunc {
	// just a sample demo
	t.Run("A sample test", func(t *testing.T) {
		if resourceName != "data.cidaas_role.example" {
			t.Errorf("invalid resource name %s", resourceName)
		}
	})

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}
		_, ok = rs.Primary.Attributes["role"]
		if !ok {
			return fmt.Errorf("Resource %s has no role set", resourceName)
		}
		return nil
	}
}
