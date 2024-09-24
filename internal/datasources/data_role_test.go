package datasources_test

// import (
// 	"fmt"
// 	"testing"

// 	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
// 	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
// 	"github.com/hashicorp/terraform-plugin-testing/terraform"
// )

// // to test data sources, existing role in cidaas must be provided
// // this can be tested by creating a role prior to the test run and cleaning up after the test
// // here test_role12 is an existing cidaas_role
// // TODO: implement presetup & cleanup
// func testAccDataSourceRoleConfig(role string) string {
// 	return fmt.Sprintf(`
// 	provider "cidaas" {
// 		base_url = "https://automation-test.dev.cidaas.eu"
// 	}
// 	data "cidaas_role" "example" {
// 	role = "%s"
// 	}`, role)
// }

// func TestAccRoleDataSource(t *testing.T) {
// 	role := "test_role12"
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Read testing
// 			{
// 				Config: testAccDataSourceRoleConfig(role),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("data.cidaas_role.example", "id", role),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccRoleDataSource_validateAttrSet(t *testing.T) {
// 	role := "test_role12"
// 	resource.Test(t, resource.TestCase{
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccDataSourceRoleConfig(role),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrSet("data.cidaas_role.example", "role"),
// 					resource.TestCheckResourceAttrSet("data.cidaas_role.example", "id"),
// 					testAccRoleData(t, "data.cidaas_role.example"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccRoleData(t *testing.T, resourceName string) resource.TestCheckFunc {
// 	// just a sample demo for reference
// 	t.Run("A sample test", func(t *testing.T) {
// 		if resourceName != "data.cidaas_role.example" {
// 			t.Errorf("invalid resource name %s", resourceName)
// 		}
// 	})
// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[resourceName]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", resourceName)
// 		}
// 		_, ok = rs.Primary.Attributes["role"]
// 		if !ok {
// 			return fmt.Errorf("Resource %s has no role set", resourceName)
// 		}
// 		return nil
// 	}
// }
