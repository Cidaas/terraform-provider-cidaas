package cidaas

import (
	"context"
	"strconv"
	"time"

	"terraform-provider-cidaas/helper_pkg/cidaas_sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomProvider() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCPCreate,
		ReadContext:   resourceCPRead,
		UpdateContext: resourceCPUpdate,
		DeleteContext: resourceCPDelete,

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
				Required: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Required: true,
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
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			// "userinfo_fields": {
			// 	Type:     schema.TypeMap,
			// 	Optional: true,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString,
			// 	},
			// },
			"scopes": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	var diags diag.Diagnostics
	var customProvider cidaas_sdk.CustomProvider

	customProvider.StandardType = d.Get("standard_type").(string)
	customProvider.AuthorizationEndpoint = d.Get("authorization_endpoint").(string)
	customProvider.TokenEndpoint = d.Get("token_endpoint").(string)
	customProvider.ProviderName = d.Get("provider_name").(string)
	customProvider.DisplayName = d.Get("display_name").(string)
	customProvider.LogoUrl = d.Get("logo_url").(string)
	customProvider.ClientId = d.Get("client_id").(string)
	customProvider.ClientSecret = d.Get("client_secret").(string)
	customProvider.UserinfoEndpoint = d.Get("userinfo_endpoint").(string)
	// customProvider.Scopes = interfaceArray2StringArray(d.Get("scopes").([]interface{}))
	customProvider.UserinfoFields.Name = d.Get("username").(string)

	response := cidaas_sdk.CreateCustomProvider(cidaas_client, customProvider)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	if err := d.Set("success", response.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("success_status", int(response.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success_status to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	error_string := ""
	error_code := int(response.Errors.Code)
	error_type := response.Errors.Type

	if err := d.Set("error", error_string); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting error to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("error_code", error_code); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting error_code to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("error_type", error_type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting error_type to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if !response.Success {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create custom provider",
			Detail:   cidaas_client.BaseUrl,
		})
		return diags
	}

	if err := d.Set("_id", response.Data.ID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting _id to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	resourceCPRead(ctx, d, m)
	return diags
}

func resourceCPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas_sdk.CidaasClient)
	provider_name := d.Get("provider_name").(string)
	response := cidaas_sdk.GetCustomProvider(cidaas_client, provider_name)

	if err := d.Set("success", response.Success); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("success_status", int(response.Status)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting success_status to resourceData",
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
	customProvider.ProviderName = d.Get("provider_name").(string)
	customProvider.DisplayName = d.Get("display_name").(string)
	customProvider.LogoUrl = d.Get("logo_url").(string)
	customProvider.ClientId = d.Get("client_id").(string)
	customProvider.ClientSecret = d.Get("client_secret").(string)
	customProvider.UserinfoEndpoint = d.Get("userinfo_endpoint").(string)
	// customProvider.Scopes = interfaceArray2StringArray(d.Get("scopes").([]interface{}))
	customProvider.UserinfoFields.Name = d.Get("username").(string)
	customProvider.ID = d.Get("_id").(string)

	response := cidaas_sdk.UpdateCustomProvider(cidaas_client, customProvider)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

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
			Summary:  "Custom Provider Update Failed",
			Detail:   response.Errors.Error,
		})
	}

	resourceAppRead(ctx, d, m)
	return diags
}

func resourceCPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	cidaas_client := m.(cidaas_sdk.CidaasClient)
	provider_name := d.Get("provider_name").(string)

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
