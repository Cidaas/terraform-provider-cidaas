package cidaas

import (
	"context"
	"fmt"

	"terraform-provider-cidaas/helper/cidaas"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var allowedLocals = []string{"ar", "ar-ae", "ar-bh", "ar-dz", "ar-eg", "ar-iq", "ar-jo", "ar-kw", "ar-lb", "ar-ly", "ar-ma", "ar-om", "ar-qa", "ar-sa",
	"ar-sd", "ar-sy", "ar-tn", "ar-ye", "be", "be-by", "bg", "bg-bg", "ca", "ca-es", "cs", "cs-cz", "da", "da-dk", "de", "de-at",
	"de-ch", "de-de", "de-lu", "el", "el-gr", "en", "en-au", "en-ca", "en-gb", "en-ie", "en-in", "en-nz", "en-us", "en-za",
	"es", "es-ar", "es-bo", "es-cl", "es-co", "es-cr", "es-do", "es-ec", "es-es", "es-gt", "es-hn", "es-mx", "es-ni", "es-pa",
	"es-pe", "es-pr", "es-py", "es-sv", "es-uy", "es-ve", "et", "et-ee", "fi", "fi-fi", "fr", "fr-be", "fr-ca", "fr-ch", "fr-fr",
	"fr-lu", "hi-in", "hr", "hr-hr", "hu", "hu-hu", "is", "is-is", "it", "it-ch", "it-it", "iw", "iw-il", "ja", "ja-jp", "ko",
	"ko-kr", "lt", "lt-lt", "lv", "lv-lv", "mk", "mk-mk", "nl", "nl-be", "nl-nl", "no", "no-no", "no-no-ny", "pl", "pl-pl",
	"pt", "pt-br", "pt-pt", "ro", "ro-ro", "ru", "ru-ru", "sk", "sk-sk", "sl", "sl-si", "sq", "sq-al", "sr", "sr-ba", "sr-cs",
	"sv", "sv-se", "th", "th-th", "th-th-th", "tr", "tr-tr", "uk", "uk-ua", "vi", "vi-vn", "zh", "zh-cn", "zh-hk", "zh-tw",
}

var allowedTemplateTypes = []string{"EMAIL", "SMS", "IVR", "PUSH"}

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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.NoZeroValues, validation.StringInSlice(allowedLocals, false)),
			},
			"template_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"template_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(allowedTemplateTypes, false),
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
		if d.HasChange("locale") {
			oldValue, _ := d.GetChange("locale")
			d.Set("locale", oldValue)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "locale can't be modified. if you wish to create a template for a different locale with the same templateKey please create a new cidaas_template resource",
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
	d.SetId(response.Data.TemplateKey + "_" + response.Data.TemplateType + "_" + response.Data.Locale)
	return diags
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	id := d.Id()
	template := cidaas.PrepareTemplate(id)

	if !containsString(allowedTemplateTypes, template.TemplateType) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Invalid ID format provided. The correct format should follow this structure: 'templateKey_templateType_locale'. For example: 'Sample_EMAIL_en-us'. Please note that template types are case-sensitive. Allowed template types are: %+v.", allowedTemplateTypes),
		})
		return diags
	}

	if !containsString(allowedLocals, template.Locale) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Invalid ID format provided. The correct format should follow this structure: 'templateKey_templateType_locale'. For example: 'Sample_EMAIL_en-us'. Allowed locales are: %+v.", allowedLocals),
		})
		return diags
	}
	locale := d.Get("locale").(string)
	if locale != "" {
		template.Locale = locale
	}

	response, err := cidaas_client.GetTemplate(template)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to fetch template",
			Detail:   fmt.Sprintf("error: %+v", err.Error()),
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
	d.SetId(response.Data.TemplateKey + "_" + response.Data.TemplateType + "_" + response.Data.Locale)
	return diags
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	id := d.Id()
	template := cidaas.PrepareTemplate(id)
	err := cidaas_client.DeleteTemplate(template)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete template %+v", template.TemplateKey),
			Detail:   err.Error(),
		})
	}
	return diags
}

func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
