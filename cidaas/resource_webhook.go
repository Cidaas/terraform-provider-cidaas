package cidaas

import (
	"context"
	"fmt"
	"terraform-provider-cidaas/helper/cidaas"
	"terraform-provider-cidaas/helper/util"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebhookCreate,
		ReadContext:   resourceWebhookRead,
		UpdateContext: resourceWebhookUpdate,
		DeleteContext: resourceWebhookDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"events": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"apikey_placeholder": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"apikey_placement": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"apikey": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"totp_placeholder": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"totp_placement": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"totpkey": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceWebhookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	webhook, diags := prepareWebhookRequestPayload(d)
	if diags != nil {
		return diags
	}
	cidaas_client := m.(cidaas.CidaasClient)
	response, err := cidaas_client.CreateOrUpdateWebhook(webhook)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create webhook"),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("_id", response.Data.ID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while setting param _id to webhook resource",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("disable", response.Data.Disable); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while setting param disable to webhook resource",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.ID)
	resourceWebhookRead(ctx, d, m)
	return diags
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	wb_id := d.Id()
	response, err := cidaas_client.GetWebhook(wb_id)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to read webhook id %+v", wb_id),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("_id", response.Data.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", response.Data.AuthType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", response.Data.Url); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("events", response.Data.Events); err != nil {
		return diag.FromErr(err)
	}
	if response.Data.AuthType == "APIKEY" {
		if err := d.Set("apikey_placeholder", response.Data.ApiKeyDetails.ApikeyPlaceholder); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("apikey_placement", response.Data.ApiKeyDetails.ApikeyPlacement); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("apikey", response.Data.ApiKeyDetails.Apikey); err != nil {
			return diag.FromErr(err)
		}
	}
	if response.Data.AuthType == "TOTP" {
		if err := d.Set("totp_placeholder", response.Data.TotpDetails.TotpPlaceholder); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("totp_placement", response.Data.TotpDetails.TotpPlacement); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("totpkey", response.Data.TotpDetails.TotpKey); err != nil {
			return diag.FromErr(err)
		}
	}
	if response.Data.AuthType == "CIDAAS_OAUTH2" {
		if err := d.Set("client_id", response.Data.CidaasAuthDetails.ClientId); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("disable", response.Data.Disable); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	webhook, diags := prepareWebhookRequestPayload(d)
	if diags != nil {
		return diags
	}
	webhook.ID = d.Get("_id").(string)
	cidaas_client := m.(cidaas.CidaasClient)
	_, err := cidaas_client.CreateOrUpdateWebhook(webhook)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to update webhook. webhook id %+v", webhook.ID),
			Detail:   err.Error(),
		})
		return diags
	}
	resourceWebhookRead(ctx, d, m)
	return diags
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	wb_id := d.Id()
	_, err := cidaas_client.DeleteWebhook(wb_id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete webhook. webhook id %+v", wb_id),
			Detail:   err.Error(),
		})
	}
	return diags
}

func isValidAllowedValue(allowed_value []string, auth_type string) bool {
	for _, v := range allowed_value {
		if v == auth_type {
			return true
		}
	}
	return false
}

func prepareWebhookRequestPayload(d *schema.ResourceData) (*cidaas.WebhookRequestPayload, diag.Diagnostics) {
	var webhook cidaas.WebhookRequestPayload
	var diags diag.Diagnostics
	webhook.AuthType = d.Get("auth_type").(string)
	if !isValidAllowedValue(cidaas.AllowedAuthType, webhook.AuthType) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("invlaid auth_type"),
			Detail:   fmt.Sprintf("auth_type must match one of these %+v", cidaas.AllowedAuthType),
		})
		return nil, diags
	}
	if webhook.AuthType == "APIKEY" {
		webhook.ApiKeyDetails.Apikey = d.Get("apikey").(string)
		webhook.ApiKeyDetails.ApikeyPlaceholder = d.Get("apikey_placeholder").(string)
		webhook.ApiKeyDetails.ApikeyPlacement = d.Get("apikey_placement").(string)
		if webhook.ApiKeyDetails.Apikey == "" || webhook.ApiKeyDetails.ApikeyPlaceholder == "" || webhook.ApiKeyDetails.ApikeyPlacement == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("missing required field"),
				Detail:   fmt.Sprintf("the fields apikey_placeholder, apikey_placement and apikey must be provided for auth type %+v", webhook.AuthType),
			})
			return nil, diags
		}
		if !isValidAllowedValue(cidaas.AllowedKeyPlacementValue, webhook.ApiKeyDetails.ApikeyPlacement) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("invlaid apikey_placement"),
				Detail:   fmt.Sprintf("apikey_placement must match one of these %+v", cidaas.AllowedKeyPlacementValue),
			})
			return nil, diags
		}
	}
	if webhook.AuthType == "TOTP" {
		webhook.TotpDetails.TotpPlaceholder = d.Get("totp_placeholder").(string)
		webhook.TotpDetails.TotpPlacement = d.Get("totp_placement").(string)
		webhook.TotpDetails.TotpKey = d.Get("totpkey").(string)
		if webhook.TotpDetails.TotpPlaceholder == "" || webhook.TotpDetails.TotpPlacement == "" || webhook.TotpDetails.TotpKey == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("missing required field"),
				Detail:   fmt.Sprintf("the fields totp_placeholder, totp_placement and totpkey must be provided for auth_type %+v", webhook.AuthType),
			})
			return nil, diags
		}
		if !isValidAllowedValue(cidaas.AllowedKeyPlacementValue, webhook.TotpDetails.TotpPlacement) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("invlaid totp_placement"),
				Detail:   fmt.Sprintf("totp_placement must match one of these %+v", cidaas.AllowedKeyPlacementValue),
			})
			return nil, diags
		}
	}
	if webhook.AuthType == "CIDAAS_OAUTH2" {
		webhook.CidaasAuthDetails.ClientId = d.Get("client_id").(string)
		if webhook.CidaasAuthDetails.ClientId == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("missing required field"),
				Detail:   fmt.Sprintf("client_id must be provided for auth_type %+v", webhook.AuthType),
			})
			return nil, diags
		}
	}
	webhook.Url = d.Get("url").(string)
	webhook.Events = util.InterfaceArray2StringArray(d.Get("events").([]interface{}))
	return &webhook, nil
}
