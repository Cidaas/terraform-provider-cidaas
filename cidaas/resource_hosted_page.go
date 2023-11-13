package cidaas

import (
	"context"
	"fmt"

	"terraform-provider-cidaas/helper/cidaas"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHostedPage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHostedPageCreate,
		ReadContext:   resourceHostedPageRead,
		UpdateContext: resourceHostedPageUpdate,
		DeleteContext: resourceHostedPageDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"hosted_page_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_locale": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hosted_pages": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hosted_page_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"locale": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceHostedPageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var hpRequestPayload cidaas.HostedPagePayload

	hpRequestPayload.GroupOwner = "client"
	hpRequestPayload.ID = d.Get("hosted_page_group_name").(string)
	hpRequestPayload.DefaultLocale = d.Get("default_locale").(string)

	hosted_pages := d.Get("hosted_pages").([]interface{})
	hps := []cidaas.HostedPage{}
	for _, hosted_page := range hosted_pages {
		temp := hosted_page.(map[string]interface{})
		hp := cidaas.HostedPage{
			HostedPageId: temp["hosted_page_id"].(string),
			Locale:       temp["locale"].(string),
			Url:          temp["url"].(string),
		}
		hps = append(hps, hp)
	}
	hpRequestPayload.HostedPages = hps
	cidaas_client := m.(cidaas.CidaasClient)
	response, err := cidaas_client.CreateOrUpdateHostedPage(hpRequestPayload)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create hosted page %+v", hpRequestPayload.ID),
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.ID)
	resourceHostedPageRead(ctx, d, m)
	return diags
}

func resourceHostedPageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	hp_group_name := d.Id()
	response, err := cidaas_client.GetHostedPage(hp_group_name)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to read hosted page %+v", hp_group_name),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("hosted_page_group_name", response.Data.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_locale", response.Data.DefaultLocale); err != nil {
		return diag.FromErr(err)
	}

	hps := flattenHostedPages(&response.Data.HostedPages)
	if err := d.Set("hosted_pages", hps); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceHostedPageUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var hpRequestPayload cidaas.HostedPagePayload

	hpRequestPayload.GroupOwner = "client"
	hpRequestPayload.ID = d.Get("hosted_page_group_name").(string)
	if hpRequestPayload.ID != d.Id() {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("hosted_page_group_name can not be modified"),
		})
		return diags
	}
	hpRequestPayload.DefaultLocale = d.Get("default_locale").(string)

	hosted_pages := d.Get("hosted_pages").([]interface{})
	hps := []cidaas.HostedPage{}
	for _, hosted_page := range hosted_pages {
		temp := hosted_page.(map[string]interface{})
		hp := cidaas.HostedPage{
			HostedPageId: temp["hosted_page_id"].(string),
			Locale:       temp["locale"].(string),
			Url:          temp["url"].(string),
		}
		hps = append(hps, hp)
	}
	hpRequestPayload.HostedPages = hps
	cidaas_client := m.(cidaas.CidaasClient)
	_, err := cidaas_client.CreateOrUpdateHostedPage(hpRequestPayload)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to update hosted page %+v", hpRequestPayload.ID),
			Detail:   err.Error(),
		})
	}
	resourceHostedPageRead(ctx, d, m)
	return diags
}

func resourceHostedPageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	hp_group_name := d.Id()
	_, err := cidaas_client.DeleteHostedPage(hp_group_name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete hosted page %+v", hp_group_name),
			Detail:   err.Error(),
		})
	}
	return diags
}

func flattenHostedPages(hps *[]cidaas.HostedPage) []interface{} {
	if hps != nil {
		ois := make([]interface{}, len(*hps), len(*hps))
		for i, hp := range *hps {
			oi := make(map[string]interface{})
			oi["hosted_page_id"] = hp.HostedPageId
			oi["locale"] = hp.Locale
			oi["url"] = hp.Url
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}
