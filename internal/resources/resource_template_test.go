package resources_test

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	acctest "github.com/Cidaas/terraform-provider-cidaas/internal/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// only covers custom temaplates tests
// locale validation is already done in other resources, hence skipped in template resource

const resourceTemplate = "cidaas_template.example"

var (
	templateLocale  = "de-de"
	templateKey     = strings.ToUpper(acctest.RandString(10))
	templateType    = "SMS"
	templateContent = acctest.RandString(256)
)

// create, read and update test
func TestTemplate_Basic(t *testing.T) {
	updatedTemplateContent := acctest.RandString(256)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             checkTemplateDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testTemplateConfig(templateLocale, templateKey, templateType, templateContent),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTemplate, "locale", templateLocale),
					resource.TestCheckResourceAttr(resourceTemplate, "template_key", templateKey),
					resource.TestCheckResourceAttr(resourceTemplate, "template_type", templateType),
					resource.TestCheckResourceAttr(resourceTemplate, "content", templateContent),

					resource.TestCheckResourceAttrSet(resourceTemplate, "id"),
					resource.TestCheckResourceAttrSet(resourceTemplate, "template_owner"),
					resource.TestCheckResourceAttrSet(resourceTemplate, "group_id"),
					resource.TestCheckResourceAttrSet(resourceTemplate, "is_system_template"),
				),
			},
			{
				ResourceName:      resourceTemplate,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     templateKey + ":" + templateType + ":" + templateLocale,
			},
			{
				Config: testTemplateConfig(templateLocale, templateKey, templateType, updatedTemplateContent),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTemplate, "content", updatedTemplateContent),
					// check default value
					resource.TestCheckResourceAttr(resourceTemplate, "is_system_template", strconv.FormatBool(false)),
				),
			},
			// locale, template_key and template type can't be modified
			{
				Config:      testTemplateConfig("en-us", strings.ToUpper(acctest.RandString(10)), "IVR", updatedTemplateContent),
				ExpectError: regexp.MustCompile("can't be modified"),
			},
		},
	})
}

func testTemplateConfig(locale, templateKey, templateType, content string) string {
	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_template" "example" {
			locale        = "%s"
			template_key  = "%s"
			template_type = "%s"
			content       = "%s"
		}
		`, "https://automation-test.dev.cidaas.eu", locale, templateKey, templateType, content)
}

func checkTemplateDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceTemplate]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceTemplate)
	}

	template := cidaas.Template{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}

	templatePayload := cidaas.TemplateModel{
		Locale:       rs.Primary.Attributes["locale"],
		TemplateKey:  rs.Primary.Attributes["temaplte_key"],
		TemplateType: rs.Primary.Attributes["temaplte_type"],
	}

	res, _ := template.Get(templatePayload, false)
	if res != nil && res.Status != http.StatusNoContent {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// subject can not be empty when template type is SMS
func TestTemplate_EmailSubjectCheck(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testTemplateConfig(templateLocale, templateKey, "EMAIL", templateContent),
				ExpectError: regexp.MustCompile("subject can not be empty when template_type is EMAIL"),
			},
		},
	})
}

// template_key must be a valid string consisting only of uppercase letters,
// digits (0-9), underscores (_), and hyphens (-)
func TestTemplate_TemplateKeyValidation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testTemplateConfig(templateLocale, acctest.RandString(10), templateType, templateContent),
				ExpectError: regexp.MustCompile("template_key must be a valid string consisting"), // TODO: full string validation
			},
		},
	})
}

// template_type must be one of "EMAIL", "SMS", "IVR" and "PUSH"
func TestTemplate_TemplateTypeValidation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testTemplateConfig(templateLocale, acctest.RandString(10), "INVALID", templateContent),
				ExpectError: regexp.MustCompile("template_key must be a valid string consisting"), // TODO: full string validation
			},
		},
	})
}

// required params locale, template_key, teamplte_type and content
func TestTemplate_MissingRequired(t *testing.T) {
	requiredParams := []string{"locale", "template_key", "template_type", "content"}
	for _, v := range requiredParams {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						provider "cidaas" {
							base_url = "%s"
						}
						resource "cidaas_template" "example" {}
					`, "https://automation-test.dev.cidaas.eu"),
					ExpectError: regexp.MustCompile(fmt.Sprintf(`"%s" is required`, v)), // TODO: full string validation
				},
			},
		})
	}
}

// System Template basic create update and delete, system template can not be imported
func TestTemplate_SystemTemplateBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             checkTemplateDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_template" "example" {
					locale             = "en-us"
					template_key       = "VERIFY_USER"
					template_type      = "SMS"
					content            = "Hi {{name}}, here is the {{code}} to verify the user"
					is_system_template = true
					group_id           = "sample_group"
					processing_type    = "GENERAL"
					verification_type  = "SMS"
					usage_type         = "VERIFICATION_CONFIGURATION"
				}
				`, "https://automation-test.dev.cidaas.eu"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTemplate, "is_system_template", strconv.FormatBool(true)),
					resource.TestCheckResourceAttrSet(resourceTemplate, "id"),
				),
			},
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_template" "example" {
					locale             = "en-us"
					template_key       = "VERIFY_USER"
					template_type      = "SMS"
					content            = "Hi {{name}}, here is the {{code}} to verify the user updated"
					is_system_template = true
					group_id           = "sample_group"
					processing_type    = "GENERAL"
					verification_type  = "SMS"
					usage_type         = "VERIFICATION_CONFIGURATION"
				}
				`, "https://automation-test.dev.cidaas.eu"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTemplate, "content", "Hi {{name}}, here is the {{code}} to verify the user updated"),
				),
			},
			// templated reverted back to the old state
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_template" "example" {
					locale             = "en-us"
					template_key       = "VERIFY_USER"
					template_type      = "SMS"
					content            = "Hi {{name}}, here is the {{code}} to verify the user"
					is_system_template = true
					group_id           = "sample_group"
					processing_type    = "GENERAL"
					verification_type  = "SMS"
					usage_type         = "VERIFICATION_CONFIGURATION"
				}
				`, "https://automation-test.dev.cidaas.eu"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTemplate, "content", "Hi {{name}}, here is the {{code}} to verify the user"),
				),
			},
		},
	})
}
