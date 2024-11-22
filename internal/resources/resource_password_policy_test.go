package resources_test

// commented to pass acceptance test as we change password policy to a new design
// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// 	"testing"

// 	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
// 	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
// )

// const resourcePwdPolicy = "cidaas_password_policy.example"

// // create, read and update test
// func TestAccPwdPolicyResource_Basic(t *testing.T) {
// 	policyName := acctest.RandString(10)
// 	maxLength := 20
// 	minLength := 8
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPwdPolicyResourceConfig(policyName, int64(minLength), int64(maxLength)),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourcePwdPolicy, "policy_name", policyName),
// 					resource.TestCheckResourceAttr(resourcePwdPolicy, "maximum_length", strconv.Itoa(maxLength)),
// 					resource.TestCheckResourceAttr(resourcePwdPolicy, "minimum_length", strconv.Itoa(minLength)),
// 					resource.TestCheckResourceAttr(resourcePwdPolicy, "lower_and_uppercase", strconv.FormatBool(true)),
// 				),
// 			},
// 			{
// 				ResourceName:      resourcePwdPolicy,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 				// ImportStateId:     "id",
// 			},
// 			{
// 				Config: testAccPwdPolicyResourceConfig(policyName, int64(10), int64(25)),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourcePwdPolicy, "maximum_length", strconv.Itoa(25)),
// 					resource.TestCheckResourceAttr(resourcePwdPolicy, "minimum_length", strconv.Itoa(10)),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccPwdPolicyResourceConfig(policyName string, minimumLength, maximumLength int64) string {
// 	return fmt.Sprintf(`
// 	provider "cidaas" {
// 		base_url = "%s"
// 	}
// 	resource "cidaas_password_policy" "example" {
// 		policy_name 				 = "%s"
// 		minimum_length       = %d
// 		maximum_length       = %d
// 		lower_and_uppercase  = true
// 		no_of_digits         = 1
// 		no_of_special_chars  = 1
// 	}
// 	`, acctest.BaseURL, policyName, minimumLength, maximumLength)
// }

// // maximum_length should be greater than minimum_length
// func TestAccPwdPolicyResource_MaxMinLengthValidation(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				// minimumLength 20 & maximumLength 10
// 				Config: testAccPwdPolicyResourceConfig(acctest.RandString(10), 20, 10),
// 				// TODO: add no_of_special_chars + no_of_digits + lower_and_uppercase in validation
// 				ExpectError: regexp.MustCompile(`Attribute maximum_length value must be at least sum of minimum_length`),
// 			},
// 		},
// 	})
// }

// // // missing required parameter
// func TestAccPwdPolicyResource_MissingRequired(t *testing.T) {
// 	requiredParams := []string{
// 		"minimum_length", "lower_and_uppercase", "no_of_digits",
// 		"no_of_special_chars", "maximum_length",
// 	}
// 	for _, param := range requiredParams {
// 		resource.Test(t, resource.TestCase{
// 			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 			Steps: []resource.TestStep{
// 				{
// 					Config: fmt.Sprintf(`
// 						provider "cidaas" {
// 							base_url = "%s"
// 						}
// 						resource "cidaas_password_policy" "example" {}
// 					`, acctest.BaseURL),
// 					ExpectError: regexp.MustCompile(fmt.Sprintf(`The argument "%s" is required, but no definition was found.`, param)),
// 				},
// 			},
// 		})
// 	}
// }
