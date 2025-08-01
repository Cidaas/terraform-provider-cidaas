package resources_test

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/Cidaas/terraform-provider-cidaas/internal/resources"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// var templateGroupID = acctest.RandString(10)

// create, read and update test
// func TestTemplateGroup_Basic(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		CheckDestroy:             CheckTemplateGroupDestroyed,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: fmt.Sprintf(`
// 					provider "cidaas" {
// 						base_url = "%s"
// 					}
// 					resource "cidaas_template_group" "example" {
// 						group_id                       = "`+templateGroupID+`"
// 					}
// 				`, acctest.GetBaseURL()),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceTemplateGroup, "group_id", templateGroupID),

// 					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "id"),
// 					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.from_email"),
// 					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.from_name"),
// 					// resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.reply_to"),
// 					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "email_sender_config.sender_names.#"),
// 					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "ivr_sender_config.sender_names.#"),
// 					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "push_sender_config.sender_names.#"),
// 					resource.TestCheckResourceAttrSet(resourceTemplateGroup, "sms_sender_config.sender_names.#"),
// 				),
// 			},
// 			{
// 				ResourceName:      resourceTemplateGroup,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 				ImportStateId:     templateGroupID,
// 			},
// 			// if you update only from_email in cidaas_template_group then the test fails
// 			// TODO: fix refresh plan not empty issue terraform apply(update)
// 			{
// 				Config: fmt.Sprintf(`
// 					provider "cidaas" {
// 						base_url = "%s"
// 					}
// 					resource "cidaas_template_group" "example" {
// 						group_id	= "`+templateGroupID+`"
// 						email_sender_config = {
// 							from_email = "noreply@cidaas.eu"
// 							from_name  = "Kube-dev"
// 							reply_to   = "noreply@cidaas.de"
// 							sender_names = [
// 								"SYSTEM",
// 							]
// 						}
// 					}
// 				`, acctest.GetBaseURL()),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceTemplateGroup, "email_sender_config.from_email", "noreply@cidaas.eu"),
// 				),
// 			},
// 		},
// 	})
// }

func checkTemplateGroupDestroyed(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}

		tg := cidaas.TemplateGroup{
			ClientConfig: cidaas.ClientConfig{
				BaseURL:     acctest.GetBaseURL(),
				AccessToken: acctest.TestToken,
			},
		}

		// Add retry logic for eventual consistency
		maxRetries := 5
		for i := 0; i < maxRetries; i++ {
			res, err := tg.Get(context.Background(), rs.Primary.Attributes["group_id"])

			// Check if resource is successfully deleted (nil or NoContent status)
			if res == nil || res.Status == http.StatusNoContent {
				return nil // Resource successfully deleted
			}

			// Handle other errors
			if err != nil {
				// If error is "not found", that's what we want
				if strings.Contains(err.Error(), "not found") ||
					strings.Contains(err.Error(), "404") ||
					strings.Contains(err.Error(), "204") {
					return nil
				}
				return fmt.Errorf("error checking if template group exists: %w", err)
			}

			// If this is the last retry, return error
			if i == maxRetries-1 {
				return fmt.Errorf("template group still exists after %d retries: %+v", maxRetries, res)
			}

			// Wait before retrying with exponential backoff
			waitTime := time.Duration(i+1) * time.Second * 2
			time.Sleep(waitTime)
		}

		return nil
	}
}

// group_id length should not be greater than 15
func TestTemplateGroup_GourpIDLenghCheck(t *testing.T) {
	t.Parallel()

	testResourceID := acctest.RandString(10)
	testResourceName := fmt.Sprintf("%s.%s", resources.RESOURCE_TEMPLATE_GROUP, testResourceID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             checkTemplateGroupDestroyed(testResourceName),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					provider "cidaas" {
						base_url = "%s"
					}
					resource "cidaas_template_group" "%s" {
						group_id  = "`+acctest.RandString(16)+`"
					}		
				`, acctest.GetBaseURL(), testResourceID),
				ExpectError: regexp.MustCompile("group_id string length must be at most 15"),
			},
		},
	})
}
