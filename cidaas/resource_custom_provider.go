package cidaas

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"terraform-provider-cidaas/helper_pkg/cidaas_sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/google/uuid"
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"scope_names": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scope_display_label": {
				Type:     schema.TypeString,
				Required: true,
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
					},
				},
			},
			"success_status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"success": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_code": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"error_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	var customProvider cidaas_sdk.CustomProvider

	customProvider.StandardType = d.Get("standard_type").(string)
	customProvider.AuthorizationEndpoint = d.Get("authorization_endpoint").(string)
	customProvider.TokenEndpoint = d.Get("token_endpoint").(string)
	customProvider.ProviderName = d.Get("provider_name").(string)
	customProvider.DisplayName = d.Get("display_name").(string)
	customProvider.LogoUrl = d.Get("logo_url").(string)
	customProvider.UserinfoEndpoint = d.Get("userinfo_endpoint").(string)
	customProvider.Scopes.DisplayLabel = d.Get("scope_display_label").(string)
	customProvider.Scopes.Scopes = arrayOfInterface(d.Get("scope_names").([]interface{}))

	ufs, ok := d.GetOk("userinfo_fields")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while parsing userinfo_fields",
			Detail:   "error while parsing userinfo_fields",
		})
		return diags
	}
	fileds := ufs.([]interface{})
	for _, templateConfigBlock := range fileds {
		customProvider.UserinfoFields = templateConfigBlock.(map[string]interface{})
	}

	customProvider.ClientId = uuid.New().String()
	customProvider.ClientSecret = uuid.New().String()

	if err := d.Set("client_id", customProvider.ClientId); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while updating client_id in custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	cidaas_client := m.(cidaas_sdk.CidaasClient)
	response := cidaas_sdk.CreateCustomProvider(cidaas_client, customProvider)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	if err := d.Set("success", response.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("success_status", int(response.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success_status to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("error", response.Errors.Error); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting error to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("error_code", int(response.Errors.Code)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting error_code to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("error_type", response.Errors.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting error_type to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	if !response.Success {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to create custom provider %+v", response.Errors.Error),
			Detail:   response.Errors.Error,
		})
		return diags
	}

	if err := d.Set("_id", response.Data.ID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting _id to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(response.Data.ProviderName)
	return diags
}

func resourceCPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas_sdk.CidaasClient)
	provider_name := d.Id()
	response := cidaas_sdk.GetCustomProvider(cidaas_client, strings.ToLower(provider_name))

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

	var scopes []string

	for _, value := range response.Data.Scopes.Scopes {
		scopes = append(scopes, value["scope_name"])
	}

	if err := d.Set("scope_names", scopes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scope_display_label", response.Data.Scopes.DisplayLabel); err != nil {
		return diag.FromErr(err)
	}

	// read userinfo_fields and set
	// var fileds []interface{}

	// for _, value := range response.Data.UserinfoFields.(map[string]interface{}) {
	// 	fileds = append(fileds, value)
	// }

	// if err := d.Set("userinfo_fields", fileds); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  fmt.Sprintf("Error Occured while setting success to custom provider %+v", fileds...),
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	if err := d.Set("success", response.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("success_status", int(response.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success_status to custom provider",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceCPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var customProvider cidaas_sdk.CustomProvider

	customProvider.StandardType = d.Get("standard_type").(string)
	customProvider.AuthorizationEndpoint = d.Get("authorization_endpoint").(string)
	customProvider.TokenEndpoint = d.Get("token_endpoint").(string)
	customProvider.DisplayName = d.Get("display_name").(string)
	customProvider.LogoUrl = d.Get("logo_url").(string)
	customProvider.UserinfoEndpoint = d.Get("userinfo_endpoint").(string)
	customProvider.ID = d.Get("_id").(string)
	customProvider.ProviderName = strings.ToLower(d.Get("provider_name").(string))
	customProvider.Scopes.DisplayLabel = d.Get("scope_display_label").(string)
	customProvider.Scopes.Scopes = arrayOfInterface(d.Get("scope_names").([]interface{}))

	ufs, ok := d.GetOk("userinfo_fields")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while parsing userinfo_fields",
			Detail:   "error while parsing userinfo_fields",
		})
		return diags
	}
	fileds := ufs.([]interface{})
	for _, templateConfigBlock := range fileds {
		customProvider.UserinfoFields = templateConfigBlock.(map[string]interface{})
	}

	json_payload, _ := json.Marshal(customProvider)

	payload_string := string(json_payload)
	response := cidaas_sdk.UpdateCustomProvider(cidaas_client, customProvider)

	d.SetId(response.Data.ProviderName)

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Custom Provider Update Success",
		Detail:   strconv.FormatBool(response.Success),
	})

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Custom Provider Update Status",
		Detail:   strconv.Itoa(response.Status),
	})

	if !response.Success {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Custom Provider Update Failed %+v", payload_string),
			Detail:   response.Errors.Error,
		})
	}

	return diags
}

func resourceCPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	cidaas_client := m.(cidaas_sdk.CidaasClient)
	provider_name := d.Id()

	response := cidaas_sdk.DeleteCustomProvider(cidaas_client, provider_name)

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Custom provider Deletion Success",
		Detail:   strconv.FormatBool(response.Success),
	})

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Custom provider Deletion Status",
		Detail:   strconv.Itoa(response.Status),
	})

	if !response.Success {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Custom provider Deletion Failed",
			Detail:   response.Errors.Error,
		})
	}
	return diags
}
