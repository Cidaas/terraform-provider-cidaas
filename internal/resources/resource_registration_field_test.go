package resources_test

import (
	"fmt"
	"strconv"
	"testing"

	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	resourceRegField = "cidaas_registration_field.example"
)

// create, read and update test
func TestRegistrationField_CheckBoxBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testRegFieldConfig("CHECKBOX", "sample_checkbox_field", true, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceRegField, "field_key", "sample_checkbox_field"),
					resource.TestCheckResourceAttrSet(resourceRegField, "id"),
				),
			},
			{
				ResourceName:      resourceRegField,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "sample_checkbox_field",
			},
			{
				Config: testRegFieldConfig("CHECKBOX", "sample_checkbox_field", false, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceRegField, "id"),
				),
			},
		},
	})
}

func TestRegistrationField_GroupBasic(t *testing.T) {
	fieldKey := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testRegFieldConfig("TEXT", fieldKey, true, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceRegField, "field_key", fieldKey),
					resource.TestCheckResourceAttrSet(resourceRegField, "id"),
				),
			},
			{
				ResourceName:      resourceRegField,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fieldKey,
			},
			{
				Config: testRegFieldConfig("TEXT", fieldKey, false, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceRegField, "id"),
				),
			},
		},
	})
}

func TestRegistrationField_TextBasic(t *testing.T) {
	fiedKey := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "https://automation-test.dev.cidaas.eu"
				}
				resource "cidaas_registration_field" "example" {
					data_type                                      = "TEXT"
					field_key                                      = "%s"
					field_type                                     = "CUSTOM"  // CUSTOM and SYSTEM, SYSTEM can not be created but modified
					internal                                       = true      // Default: false
					required                                       = true      // Default: false
					read_only                                      = true      // Default: false
					is_group                                       = false     // Default: false
					unique                                         = true      // Default: false
					overwrite_with_null_value_from_social_provider = false     // Default: true
					is_searchable                                  = true      // Default: true
					enabled                                        = true      // Default: true
					claimable                                      = true      // Default: true
					order                                          = 1         // Default: 1
					parent_group_id                                = "DEFAULT" // Default: DEFAULT
					scopes                                         = ["profile"]
					local_texts = [
						{
							locale         = "en-US"
							name           = "Sample Field"
							required_msg   = "The field is required"
							max_length_msg = "Maximum 99 chars allowed"
							min_length_msg = "Minimum 99 chars allowed"
						},
						{
							locale         = "de-DE"
							name           = "Beispielfeld"
							required_msg   = "Dieses Feld ist erforderlich"
							max_length_msg = "DE maximum 99 chars allowed"
							min_length_msg = "DE minimum 10 chars allowed"
						}
					]
					field_definition = {
						max_length = 100
						min_length = 10
					}
				}							
			`, fiedKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceRegField, "field_key", fiedKey),
					resource.TestCheckResourceAttrSet(resourceRegField, "id"),
				),
			},
			{
				ResourceName:      resourceRegField,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fiedKey,
			},
		},
	})
}

func testRegFieldConfig(dataType, fieldKey string, internal, isGroup bool) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "https://automation-test.dev.cidaas.eu"
		}
		resource "cidaas_registration_field" "example" {
			data_type                                      = "%s"
			field_key                                      = "%s"
			field_type                                     = "CUSTOM"
			internal                                       = %s
			required                                       = true
			read_only                                      = true
			is_group                                       = %s
			unique                                         = true
			overwrite_with_null_value_from_social_provider = false
			is_searchable                                  = true
			enabled                                        = true
			claimable                                      = true
			order                                          = 1
			parent_group_id                                = "DEFAULT"
			scopes                                         = ["profile"]
			local_texts = [
				{
					locale       = "en-US"
					name         = "Sample Field"
					required_msg = "The field is required"
				},
				{
					locale       = "de-DE"
					name         = "Beispielfeld"
					required_msg = "Dieses Feld ist erforderlich"
				}
			]
		}				
	`, dataType, fieldKey, strconv.FormatBool(internal), strconv.FormatBool(isGroup))
}

func TestRegistrationField_SelectBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				provider "cidaas" {
					base_url = "https://automation-test.dev.cidaas.eu"
				}
				resource "cidaas_registration_field" "example" {
					data_type                                      = "RADIO"
					field_key                                      = "sample_select_field"
					field_type                                     = "CUSTOM"
					internal                                       = false
					required                                       = true
					read_only                                      = false
					is_group                                       = false
					unique                                         = false
					overwrite_with_null_value_from_social_provider = false
					is_searchable                                  = true
					enabled                                        = true
					claimable                                      = true
					order                                          = 1
					parent_group_id                                = "DEFAULT"
					scopes                                         = ["profile"]
					local_texts = [
						{
							locale       = "en-US"
							name         = "Sample Field"
							required_msg = "The field is required"
							attributes = [
								{
									key   = "test_key"
									value = "test_value"
								}
							]
						},
						{
							locale       = "de-DE"
							name         = "Beispielfeld"
							required_msg = "Dieses Feld ist erforderlich"
							attributes = [
								{
									key   = "test_key"
									value = "test_value"
								}
							]
						}
					]
				}
			`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceRegField, "field_key", "sample_select_field"),
					resource.TestCheckResourceAttrSet(resourceRegField, "id"),
				),
			},
			{
				ResourceName:      resourceRegField,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "sample_select_field",
			},
		},
	})
}
