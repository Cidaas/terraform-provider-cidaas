package resources_test

import (
	"fmt"
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

const (
	resourceScope       = "cidaas_scope.example"
	scopeSecurityLevel  = "CONFIDENTIAL"
	requiredUserConsent = false
	title               = "scope title in German"
	locale              = "de-DE"
	scopeDescription    = "The description of the scope in German"
)

var (
	scopeKey              = acctest.RandString(10)
	defaultScopeGroupName = []string{"developer"}
	localizedDescriptions = []map[string]string{
		{
			"title":       title,
			"locale":      locale,
			"description": scopeDescription,
		},
	}
)

// TODO: empty groupNameString test fails as it returns a plan to
// update even though there is no change to update

// create, read and update test
func TestAccScopeResource_Basic(t *testing.T) {
	updatedScopeDescription := "Updated description of the scope in German"
	updatedRequiredUserConsent := true
	localizedDescriptions = []map[string]string{
		{
			"title":  title,
			"locale": locale,
			// description is updated to validate update operation
			"description": updatedScopeDescription,
		},
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testCheckScopeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccScopeResourceConfig(scopeSecurityLevel, scopeKey, requiredUserConsent, defaultScopeGroupName, localizedDescriptions),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceScope, "security_level", scopeSecurityLevel),
					resource.TestCheckResourceAttr(resourceScope, "scope_key", scopeKey),
					resource.TestCheckResourceAttr(resourceScope, "required_user_consent", strconv.FormatBool(requiredUserConsent)),
					resource.TestCheckResourceAttr(resourceScope, "group_name.0", "developer"),
					resource.TestCheckResourceAttrSet(resourceScope, "id"),
					resource.TestCheckResourceAttrSet(resourceScope, "scope_owner"),
				),
			},
			{
				ResourceName:      resourceScope,
				ImportStateId:     scopeKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// required_user_consent & description in localized_descriptions updated
				Config: testAccScopeResourceConfig(
					scopeSecurityLevel,
					scopeKey,
					updatedRequiredUserConsent,
					defaultScopeGroupName,
					localizedDescriptions,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceScope, "required_user_consent", strconv.FormatBool(updatedRequiredUserConsent)),
					resource.TestCheckResourceAttr(resourceScope, "localized_descriptions.0.description", updatedScopeDescription),
				),
			},
		},
	})
}

func testAccScopeResourceConfig(
	securityLevel, scopeKey string,
	requiredUserConsent bool,
	groupName []string,
	localizedDescriptions []map[string]string,
) string {
	groupNameString := "[]"
	if len(groupName) > 0 {
		groupNameString = `["` + strings.Join(groupName, `", "`) + `"]`
	}

	return fmt.Sprintf(`
		provider "cidaas" {
			base_url = "%s"
		}
		resource "cidaas_scope" "example" {
			security_level = "`+securityLevel+`"
			scope_key = "`+scopeKey+`"
			required_user_consent = "`+strconv.FormatBool(requiredUserConsent)+`"
			group_name = `+groupNameString+`
			localized_descriptions =[
				{
					title = "`+localizedDescriptions[0]["title"]+`"
					locale = "`+localizedDescriptions[0]["locale"]+`"
					description = "`+localizedDescriptions[0]["description"]+`"
				}
			]
		}
	`, acctest.BaseURL)
}

func testCheckScopeDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources[resourceScope]
	if !ok {
		return fmt.Errorf("resource %s not fround", resourceScope)
	}

	scope := cidaas.ScopeImpl{
		ClientConfig: cidaas.ClientConfig{
			BaseURL:     os.Getenv("BASE_URL"),
			AccessToken: acctest.TestToken,
		},
	}
	res, _ := scope.Get(rs.Primary.Attributes["scope_key"])
	if res != nil {
		// when resource exists in remote
		return fmt.Errorf("resource stil exists %+v", res)
	}
	return nil
}

// failed validation on updating immutable proprty scope_key
func TestAccScopeResource_ImmutableScopeKeyUpdateFail(t *testing.T) {
	updatedScopeKey := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccScopeResourceConfig(scopeSecurityLevel, scopeKey, requiredUserConsent, defaultScopeGroupName, localizedDescriptions),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceScope, "scope_key", scopeKey),
				),
			},
			{
				Config:      testAccScopeResourceConfig(scopeSecurityLevel, updatedScopeKey, requiredUserConsent, defaultScopeGroupName, localizedDescriptions),
				ExpectError: regexp.MustCompile(`Attribute 'scope_key' can't be modified.`),
			},
		},
	})
}

// Invalid security_level validation
func TestAccScopeResource_InvalidSecurityLevel(t *testing.T) {
	invalidSecurityLevel := "INVALID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccScopeResourceConfig(invalidSecurityLevel, scopeKey, requiredUserConsent, defaultScopeGroupName, localizedDescriptions),
				ExpectError: regexp.MustCompile(`Attribute security_level value must be one of: \["PUBLIC" "CONFIDENTIAL"\]`),
			},
		},
	})
}

// missing required parameter scope_key
func TestAccScopeResource_MissingRequired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_scope" "example" {
				}
				`, acctest.BaseURL),
				ExpectError: regexp.MustCompile(`The argument "scope_key" is required, but no definition was found.`),
			},
		},
	})
}

// check default required_user_consent is false
func TestAccScopeResource_DefaultRequiredConsent(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_scope" "example" {
					security_level = "PUBLIC"
					scope_key = "`+scopeKey+`"
					group_name = ["developer"]
					localized_descriptions =[
						{
							title = "`+title+`"
							locale = "`+locale+`"
							description = "`+scopeDescription+`"
						}
					]
				}
				`, acctest.BaseURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceScope, "required_user_consent"),
					resource.TestCheckResourceAttr(resourceScope, "required_user_consent", strconv.FormatBool(false)),
				),
			},
		},
	})
}

// localized_descriptions[i].title is required
func TestAccScopeResource_TitleRequired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_scope" "example" {
					security_level = "PUBLIC"
					scope_key = "`+scopeKey+`"
					group_name = ["developer"]
					localized_descriptions =[
						{
							locale = "`+locale+`"
							description = "`+scopeDescription+`"
						}
					]
				}
				`, acctest.BaseURL),
				ExpectError: regexp.MustCompile(`attribute "title" is required`),
			},
		},
	})
}

// Invalid locale validation
func TestAccScopeResource_InvalidLocale(t *testing.T) {
	invalidLocale := "ab"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				provider "cidaas" {
					base_url = "%s"
				}
				resource "cidaas_scope" "example" {
					security_level = "PUBLIC"
					scope_key = "`+scopeKey+`"
					group_name = ["developer"]
					localized_descriptions =[
						{
							title = "`+title+`"
							locale = "`+invalidLocale+`"
							description = "`+scopeDescription+`"
						}
					]
				}
				`, acctest.BaseURL),
				ExpectError: regexp.MustCompile(`locale value must be one of`), // TODO:full error string comparison
			},
		},
	})
}
