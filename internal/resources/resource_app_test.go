package resources_test

import (
	"fmt"
	"os"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	resourceApp = "cidaas_app.example"
)

func TestApp_Basic(t *testing.T) {
	clientName := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAppConfig(clientName, "https://cidaas.de"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "client_name", clientName),
					resource.TestCheckResourceAttr(resourceApp, "company_website", "https://cidaas.de"),
					resource.TestCheckResourceAttrSet(resourceApp, "id"),
				),
			},
			{
				Config: testAppConfig(clientName, "https://cidaas.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceApp, "company_website", "https://cidaas.com"),
				),
			},
		},
	})
}

func testAppConfig(clientName, companyWebsite string) string {
	return fmt.Sprintf(`
    provider "cidaas" {
      base_url = "%s"
    }
    # The config below has the list of common config and main config
    resource "cidaas_app" "example" {
      client_name         = "%s"
      client_type         = "SINGLE_PAGE"
      company_address     = "01"
      company_website     = "%s"
      company_name        = "Widas ID GmbH"
      redirect_uris       = ["https://cidaas.com"]
      allow_login_with    = ["EMAIL", "MOBILE"]
      allowed_logout_urls = ["https://cidaas.com"]
      allowed_scopes      = ["openid"]
      response_types      = ["code"]
      grant_types         = ["authorization_code", "implicit", "refresh_token"]
    }`,
		os.Getenv("BASE_URL"),
		clientName,
		companyWebsite,
	)
}
