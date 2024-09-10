package datasources_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var groupType = acctest.RandString(10)

func testGroupTypeConfig(groupType string) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "https://kube-nightlybuild-dev.cidaas.de"
	}
	data "cidaas_group_type" "example" {
		group_type = "%s"
	}`, groupType)
}

func TestAccGroupTypeDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             destroyGroupTypeIfExists,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					gt := cidaas.GroupType{
						ClientConfig: cidaas.ClientConfig{
							BaseURL:     os.Getenv("BASE_URL"),
							AccessToken: acctest.TestToken,
						},
					}
					payload := cidaas.GroupTypeData{
						RoleMode:    "no_roles",
						GroupType:   groupType,
						Description: "terraform user group type description",
					}
					res, _ := gt.Create(payload)
					if res != nil && res.Status != http.StatusNoContent {
						log.Print("failed to complete pre config")
					}
				},
				Config: testGroupTypeConfig(groupType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cidaas_group_type.example", "role_mode", "no_roles"),
				),
			},
		},
	})
}

func destroyGroupTypeIfExists(s *terraform.State) error {
	var groupTypeInState string
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cidaas_group_type" {
			continue
		}
		groupTypeInState = rs.Primary.Attributes["group_type"]
	}

	if groupTypeInState != groupType {
		return fmt.Errorf("resource not found with by the role created in preconfig step")
	}

	// resource found and destroy from remote
	client := cidaas.GroupType{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}

	err := client.Delete(groupTypeInState)
	if err != nil {
		return fmt.Errorf("failed to destroy resource in remote created in the preconfig step")
	}
	return nil
}
