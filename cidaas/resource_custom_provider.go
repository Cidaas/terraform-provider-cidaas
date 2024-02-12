package cidaas

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"terraform-provider-cidaas/helper/cidaas"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCustomProvider() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCPCreate,
		ReadContext:   resourceCPRead,
		UpdateContext: resourceCPUpdate,
		DeleteContext: resourceCPDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logo_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"standard_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"OPENID_CONNECT", "OAUTH2"}, false),
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"authorization_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"userinfo_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scopes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"required": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"recommended": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"scope_display_label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"userinfo_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"family_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"given_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"middle_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"nickname": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"preferred_username": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"profile": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"picture": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"website": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"gender": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"birthdate": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"zoneinfo": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"locale": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"email": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"email_verified": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"mobile_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sub": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_fields": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
					},
				},
			},
		},
	}
}

func resourceCPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	customProvider, err := prepareCpRequestPayload(d)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintln("failed to prepare custom provider payload"),
			Detail:   err.Error(),
		})
	}
	response, err := cidaas_client.CreateCustomProvider(customProvider)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create custom provider %+v", customProvider.ProviderName),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("_id", response.Data.ID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to update _id to custom provider resource",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.ProviderName)
	return diags
}

func resourceCPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	provider_name := d.Id()
	response, err := cidaas_client.GetCustomProvider(strings.ToLower(provider_name))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to read custom provider %+v", provider_name),
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("standard_type", response.Data.StandardType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("authorization_endpoint", response.Data.AuthorizationEndpoint); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("token_endpoint", response.Data.TokenEndpoint); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("provider_name", response.Data.ProviderName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_name", response.Data.DisplayName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("logo_url", response.Data.LogoUrl); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("userinfo_endpoint", response.Data.UserinfoEndpoint); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("_id", response.Data.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scope_display_label", response.Data.Scopes.DisplayLabel); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_id", response.Data.ClientId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_secret", response.Data.ClientSecret); err != nil {
		return diag.FromErr(err)
	}

	scopes := flattenScopes(&response.Data.Scopes.Scopes)

	if err := d.Set("scopes", scopes); err != nil {
		return diag.FromErr(err)
	}

	fields := flattenUserFields(response.Data.UserinfoFields)

	if err := d.Set("userinfo_fields", fields); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to update userinfo_fields in resourceCPRead() %+v", fields...),
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceCPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	customProvider, err := prepareCpRequestPayload(d)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintln("failed to prepare custom provider payload"),
			Detail:   err.Error(),
		})
	}
	_, err = cidaas_client.UpdateCustomProvider(customProvider)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to update custom provider %+v", customProvider.ProviderName),
			Detail:   err.Error(),
		})
	}
	return diags
}

func resourceCPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	provider_name := d.Id()
	_, err := cidaas_client.DeleteCustomProvider(provider_name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete custom provider %+v", provider_name),
			Detail:   err.Error(),
		})
	}
	return diags
}

func flattenScopes(scs *[]cidaas.ScopesChild) []interface{} {
	if scs != nil {
		ois := make([]interface{}, len(*scs), len(*scs))
		for i, sc := range *scs {
			oi := make(map[string]interface{})
			oi["scope_name"] = sc.ScopeName
			oi["recommended"] = sc.Recommened
			oi["required"] = sc.Required
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenUserFields(userinfo map[string]interface{}) []interface{} {
	fields := make(map[string]interface{})
	fields["name"] = userinfo["name"]
	fields["family_name"] = userinfo["family_name"]
	fields["given_name"] = userinfo["given_name"]
	fields["middle_name"] = userinfo["middle_name"]
	fields["nickname"] = userinfo["nickname"]
	fields["preferred_username"] = userinfo["preferred_username"]
	fields["profile"] = userinfo["profile"]
	fields["picture"] = userinfo["picture"]
	fields["website"] = userinfo["website"]
	fields["gender"] = userinfo["gender"]
	fields["birthdate"] = userinfo["birthdate"]
	fields["zoneinfo"] = userinfo["zoneinfo"]
	fields["locale"] = userinfo["locale"]
	fields["updated_at"] = userinfo["updated_at"]
	fields["email"] = userinfo["email"]
	fields["email_verified"] = userinfo["email_verified"]
	fields["phone_number"] = userinfo["phone_number"]
	fields["mobile_number"] = userinfo["mobile_number"]
	fields["address"] = userinfo["address"]
	fields["sub"] = userinfo["sub"]

	var temp []interface{}
	var keys []string

	for k := range userinfo {
		keys = append(keys, k)
	}

	for _, str := range keys {
		if strings.HasPrefix(str, "customFields.") {
			result := strings.TrimPrefix(str, "customFields.")
			temp = append(temp, map[string]interface{}{"key": result, "value": userinfo[str]})
		}
	}
	fields["custom_fields"] = temp
	return []interface{}{fields}
}

func prepareCpRequestPayload(d *schema.ResourceData) (*cidaas.CustomProvider, error) {
	var customProvider cidaas.CustomProvider

	customProvider.StandardType = d.Get("standard_type").(string)
	customProvider.AuthorizationEndpoint = d.Get("authorization_endpoint").(string)
	customProvider.TokenEndpoint = d.Get("token_endpoint").(string)
	customProvider.ProviderName = d.Get("provider_name").(string)
	customProvider.DisplayName = d.Get("display_name").(string)
	customProvider.LogoUrl = d.Get("logo_url").(string)
	customProvider.UserinfoEndpoint = d.Get("userinfo_endpoint").(string)
	customProvider.Scopes.DisplayLabel = d.Get("scope_display_label").(string)
	customProvider.ClientId = d.Get("client_id").(string)
	customProvider.ClientSecret = d.Get("client_secret").(string)

	if d.Get("_id").(string) != "" {
		customProvider.ID = d.Get("_id").(string)
	}

	scopes := d.Get("scopes").([]interface{})
	scs := []cidaas.ScopesChild{}

	for _, scope := range scopes {
		temp := scope.(map[string]interface{})
		sc := cidaas.ScopesChild{
			ScopeName:  temp["scope_name"].(string),
			Recommened: temp["recommended"].(bool),
			Required:   temp["required"].(bool),
		}
		scs = append(scs, sc)
	}

	customProvider.Scopes.Scopes = scs

	ufs := d.Get("userinfo_fields").([]interface{})
	fileds := cidaas.UserInfo{}

	for _, uf := range ufs {
		field := uf.(map[string]interface{})

		fileds = cidaas.UserInfo{
			Name:              field["name"].(string),
			FamilyName:        field["family_name"].(string),
			GivenName:         field["given_name"].(string),
			MiddleName:        field["middle_name"].(string),
			Nickname:          field["nickname"].(string),
			PreferredUsername: field["preferred_username"].(string),
			Profile:           field["profile"].(string),
			Picture:           field["picture"].(string),
			Website:           field["website"].(string),
			Gender:            field["gender"].(string),
			Birthdate:         field["birthdate"].(string),
			Zoneinfo:          field["zoneinfo"].(string),
			Locale:            field["locale"].(string),
			Updated_at:        field["updated_at"].(string),
			Email:             field["email"].(string),
			EmailVerified:     field["email_verified"].(string),
			PhoneNumber:       field["phone_number"].(string),
			MobileNumber:      field["mobile_number"].(string),
			Address:           field["address"].(string),
			CustomFields:      field["custom_fields"].([]interface{}),
			Sub:               field["sub"].(string),
		}
	}

	newVar := make(map[string]interface{})
	newVar["name"] = fileds.Name
	newVar["family_name"] = fileds.FamilyName
	newVar["given_name"] = fileds.GivenName
	newVar["middle_name"] = fileds.MiddleName
	newVar["nickname"] = fileds.Nickname
	newVar["preferred_username"] = fileds.PreferredUsername
	newVar["profile"] = fileds.Profile
	newVar["picture"] = fileds.Picture
	newVar["website"] = fileds.Website
	newVar["gender"] = fileds.Gender
	newVar["birthdate"] = fileds.Birthdate
	newVar["zoneinfo"] = fileds.Zoneinfo
	newVar["locale"] = fileds.Locale
	newVar["updated_at"] = fileds.Updated_at
	newVar["email"] = fileds.Email
	newVar["email_verified"] = fileds.EmailVerified
	newVar["phone_number"] = fileds.PhoneNumber
	newVar["mobile_number"] = fileds.MobileNumber
	newVar["address"] = fileds.Address
	newVar["sub"] = fileds.Sub

	for _, item := range fileds.CustomFields {
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		var data cidaas.CustomFields
		if err := json.Unmarshal(b, &data); err != nil {
			return nil, err
		}
		newVar["customFields."+data.Key] = data.Value
	}

	customProvider.UserinfoFields = newVar
	return &customProvider, nil
}
