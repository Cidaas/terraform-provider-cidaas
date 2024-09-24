package resources_test

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const resourceTemplateGroup = "cidaas_template_group.example"

var templateGroupID = acctest.RandString(10)

// create, read and update test
func TestTemplateGroup_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             CheckTemplateGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					provider "cidaas" {
						base_url = "%s"
					}
					resource "cidaas_template_group" "example" {
						group_id                       = "`+templateGroupID+`"
					}		
				`, acctest.BaseURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTemplateGroup, "group_id", templateGroupID),

					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "id"),
					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.from_email"),
					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.from_name"),
					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.reply_to"),
					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.sender_names.#"),
					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "ivr_sender_config.sender_names.#"),
					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "push_sender_config.sender_names.#"),
					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "sms_sender_config.sender_names.#"),
				),
			},
			{
				ResourceName:      resourceTemplateGroup,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     templateGroupID,
			},
			// if you update only from_email in cidaas_template_group then the test fails
			// TODO: fix refresh plan not empty issye terraform apply(update)
			{
				Config: fmt.Sprintf(`
					provider "cidaas" {
						base_url = "%s"
					}
					resource "cidaas_template_group" "example" {
						group_id	= "`+templateGroupID+`"
						email_sender_config = {
							from_email = "noreply@cidaas.eu"
							from_name  = "Kube-dev"
							reply_to   = "noreply@cidaas.de"
							sender_names = [
								"SYSTEM",
							]
						}
					}		
				`, acctest.BaseURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTemplateGroup, "email_sender_config.from_email", "noreply@cidaas.eu"),
				),
			},
		},
	})
}

func CheckTemplateGroupDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceTemplateGroup]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceTemplateGroup)
	}

	tg := cidaas.TemplateGroup{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := tg.Get(rs.Primary.Attributes["group_id"])
	if res != nil && res.Status != http.StatusNoContent {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// group_id length should not be greater than 15
func TestTemplateGroup_GourpIDLenghCheck(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             CheckTemplateGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					provider "cidaas" {
						base_url = "%s"
					}
					resource "cidaas_template_group" "example" {
						group_id  = "`+acctest.RandString(16)+`"
					}		
				`, acctest.BaseURL),
				ExpectError: regexp.MustCompile("group_id string length must be at most 15"),
			},
		},
	})
}
