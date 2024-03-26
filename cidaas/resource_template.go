package cidaas

import (
	"context"
	"fmt"

	"terraform-provider-cidaas/helper/cidaas"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateUpsert,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpsert,
		DeleteContext: resourceTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"template_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"EMAIL", "SMS", "IVR", "PUSH"}, false),
			},
			"content": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"subject": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"usage_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTemplateUpsert(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var template cidaas.Template
	id := d.Id()
	template.TemplateKey = d.Get("template_key").(string)
	if id != "" {
		if d.HasChange("template_key") {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "template_key can't be modified",
			})
			return diags
		}
		if d.HasChange("template_type") {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "template_type can't be modified",
			})
			return diags
		}
		template.ID = d.Get("_id").(string)
	}
	template.Locale = d.Get("locale").(string)
	template.TemplateType = d.Get("template_type").(string)
	template.Content = d.Get("content").(string)
	template.Subject = d.Get("subject").(string)

	if template.TemplateType == "EMAIL" && template.Subject == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "subject can not be empty for template_key EMAIL",
		})
		return diags
	}

	cidaas_client := m.(cidaas.CidaasClient)
	response, err := cidaas_client.UpsertTemplate(template)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create template %+v", template.TemplateKey),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("_id", response.Data.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template_owner", response.Data.TemplateOwner); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("usage_type", response.Data.UsageType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("language", response.Data.Language); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_id", response.Data.GroupId); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(response.Data.TemplateKey + "_" + response.Data.TemplateType)
	return diags
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	id := d.Id()
	template := cidaas.PrepareTemplate(id)
	response, err := cidaas_client.GetTemplate(template)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to read template",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("_id", response.Data.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("locale", response.Data.Locale); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template_key", response.Data.TemplateKey); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template_type", response.Data.TemplateType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("content", response.Data.Content); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subject", response.Data.Subject); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template_owner", response.Data.TemplateOwner); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("usage_type", response.Data.UsageType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("language", response.Data.Language); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_id", response.Data.GroupId); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(response.Data.TemplateKey + "_" + response.Data.TemplateType)
	return diags
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	id := d.Id()
	template := cidaas.PrepareTemplate(id)
	_, err := cidaas_client.DeleteTemplate(template)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete template %+v", template.TemplateKey),
			Detail:   err.Error(),
		})
	}
	return diags
}
