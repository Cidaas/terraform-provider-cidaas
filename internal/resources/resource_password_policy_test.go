package resources_test

import (
	"fmt"
	"regexp"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const resourcePwdPolicy = "cidaas_password_policy.example"

// create, read and update test
func TestAccPwdPolicyResource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccPwdPolicyResourceConfig(8, 200, 1),
				ExpectError: regexp.MustCompile(`Creating this resource using`),
			},
			{
				ResourceName:       resourcePwdPolicy,
				ImportState:        true,
				ImportStateVerify:  false, // made false as it compares existing state id and imported id. we don't have existing as create is not supported
				ImportStateId:      "cidaas",
				ImportStatePersist: true,
			},
			// update is skipped as password policy resource cannot be deleted in cidaas
			// update might result in unintended changes or incorrect values being applied.
		},
	})
}

func testAccPwdPolicyResourceConfig(minimumLength, maximumLength, reuseLimit int64) string {
	return fmt.Sprintf(`
	provider "cidaas" {
		base_url = "https://automation-test.dev.cidaas.eu"
	}
	resource "cidaas_password_policy" "example" {
		minimum_length       = %d
		maximum_length       = %d
		lower_and_uppercase  = true
		no_of_digits         = 1
		expiration_in_days   = 30
		no_of_special_chars  = 1
		no_of_days_to_remind = 1
		reuse_limit          = %d
	}
	`, minimumLength, maximumLength, reuseLimit)
}

// maximum_length should be greater than minimum_length
func TestAccPwdPolicyResource_MaxMinLengthValidation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// minimumLength 20 & maximumLength 10
				Config: testAccPwdPolicyResourceConfig(20, 10, 1),
				// TODO: add no_of_special_chars + no_of_digits + lower_and_uppercase in validation
				ExpectError: regexp.MustCompile(`Attribute maximum_length value must be at least sum of minimum_length`),
			},
		},
	})
}

// invalid reuse limit
func TestAccPwdPolicyResource_InvalidReuseLimit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// invalid limit cannot be greater than 5
				Config:      testAccPwdPolicyResourceConfig(10, 20, 10),
				ExpectError: regexp.MustCompile(`reuse_limit value must be at most 5`), // improve error and extra oarams to sum
			},
		},
	})
}

// missing required parameter
func TestAccPwdPolicyResource_MissingRequired(t *testing.T) {
	requiredParams := []string{
		"minimum_length", "lower_and_uppercase", "no_of_digits", "expiration_in_days",
		"no_of_special_chars", "no_of_days_to_remind", "reuse_limit", "maximum_length",
	}
	for _, param := range requiredParams {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: `
						provider "cidaas" {
							base_url = "https://automation-test.dev.cidaas.eu"
						}
						resource "cidaas_password_policy" "example" {}
					`,
					ExpectError: regexp.MustCompile(fmt.Sprintf(`The argument "%s" is required, but no definition was found.`, param)),
				},
			},
		})
	}
}
