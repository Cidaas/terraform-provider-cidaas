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
			"api_key_details": {
				Type:     schema.TypeMap,
				Required: true,
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
	var webhook cidaas.WebhookRequestPayload

	webhook.AuthType = d.Get("auth_type").(string)
	webhook.Url = d.Get("url").(string)
	webhook.Events = util.InterfaceArray2StringArray(d.Get("events").([]interface{}))
	webhook.ApiKeyDetails = d.Get("api_key_details").(map[string]interface{})

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
	if err := d.Set("api_key_details", response.Data.ApiKeyDetails); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("disable", response.Data.Disable); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var webhook cidaas.WebhookRequestPayload

	webhook.AuthType = d.Get("auth_type").(string)
	webhook.Url = d.Get("url").(string)
	webhook.Events = util.InterfaceArray2StringArray(d.Get("events").([]interface{}))
	webhook.ApiKeyDetails = d.Get("api_key_details").(map[string]interface{})
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
