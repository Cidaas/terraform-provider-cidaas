package resources_test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const resourcePwdPolicy = "cidaas_password_policy.example"

// create, read and update test
func TestAccPwdPolicyResource_Basic(t *testing.T) {
	policyName := acctest.RandString(10)
	// regex := `^(?=.*[A-Za-z])(?!.*\\s).{6,15}$`
	// updateRegex := `^(?=.*[A-Z])(?=.*[a-z])(?=(.*[0-9]){1})(?!.*\\s).{8,15}$`
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPwdPolicyResourceConfig(policyName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourcePwdPolicy, "policy_name", policyName),
					resource.TestCheckResourceAttr(resourcePwdPolicy, "password_policy.block_compromised", strconv.FormatBool(true)),
					resource.TestCheckResourceAttr(resourcePwdPolicy, "password_policy.deny_usage_count", strconv.Itoa(0)),
					// resource.TestCheckResourceAttr(resourcePwdPolicy, "password_policy.strength_regexes.[0]",  regex), // Bitwise operators are not supported
					resource.TestCheckResourceAttr(resourcePwdPolicy, "password_policy.change_enforcement.expiration_in_days", strconv.Itoa(0)),
					resource.TestCheckResourceAttr(resourcePwdPolicy, "password_policy.change_enforcement.notify_user_before_in_days", strconv.Itoa(0)),
				),
			},
			{
				ResourceName:      resourcePwdPolicy,
				ImportState:       true,
				ImportStateVerify: true,
				// ImportStateId:     "id",
			},
			{
				Config: testAccPwdPolicyResourceConfig(policyName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourcePwdPolicy, "password_policy.block_compromised", strconv.FormatBool(false)),
				),
			},
		},
	})
}

func testAccPwdPolicyResourceConfig(policyName string, blockCompromised bool) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "%s"
	}
	resource "cidaas_password_policy" "example" {
		policy_name 	  = "%s"
		password_policy = {
			block_compromised = "%+v"
			strength_regexes = [
				"^(?=.*[A-Za-z])(?!.*\\s).{6,15}$"
			],
		}
	}
	`, acctest.BaseURL, policyName, blockCompromised)
}

// missing required parameter
func TestAccPwdPolicyResource_MissingRequired(t *testing.T) {
	requiredParams := []string{
		"policy_name",
	}
	// TODO: to add strength_regexes in the list
	for _, param := range requiredParams {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						provider "cidaas" {
							base_url = "%s"
						}
						resource "cidaas_password_policy" "example" {}
					`, acctest.BaseURL),
					ExpectError: regexp.MustCompile(fmt.Sprintf(`The argument "%s" is required, but no definition was found.`, param)),
				},
			},
		})
	}
}
