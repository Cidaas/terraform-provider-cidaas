package cidaas

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"terraform-provider-cidaas/helper/cidaas"
)

func resourceScopeGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: resourceScopeGroupCreateOrUpdate,
		ReadContext:   resourceScopeGroupRead,
		UpdateContext: resourceScopeGroupCreateOrUpdate,
		DeleteContext: resourceScopeGroupDelete,
	}

}

func resourceScopeGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	var scopeGroupConfig cidaas.ScopeGroupConfig
	scopeGroupConfig.GroupName = d.Get("group_name").(string)
	scopeGroupConfig.Description = d.Get("description").(string)
	response, err := cidaasClient.CreateOrUpdateScopeGroup(scopeGroupConfig)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create scope %+v", scopeGroupConfig.GroupName),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("group_name", response.Data.GroupName); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while setting group_name to scope resource",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.GroupName)
	return diags
}

func resourceScopeGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	scopeGroupName := d.Id()
	response, err := cidaasClient.GetScopeGroup(strings.ToLower(scopeGroupName))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to read scope group %+v", scopeGroupName),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("description", response.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_name", response.Data.GroupName); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceScopeGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	scopeGroupName := d.Id()
	_, err := cidaasClient.DeleteScopeGroup(scopeGroupName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete scope group %+v", scopeGroupName),
			Detail:   err.Error(),
		})
	}
	return diags
}
