package cidaas

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"terraform-provider-cidaas/helper/cidaas"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: resourceRoleUpsert,
		ReadContext:   resourceRoleRead,
		UpdateContext: resourceRoleUpsert,
		DeleteContext: resourceRoleDelete,
	}

}

func resourceRoleUpsert(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	var roleConfig cidaas.RoleConfig
	isCreate := d.Id() == ""
	role := d.Id()
	if !isCreate && role != "" && role != d.Get("role").(string) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Role %v does not exist, cannot update. Please create one first", d.Get("role").(string)),
		})
		return diags
	}
	roleConfig.Name = d.Get("name").(string)
	roleConfig.Role = d.Get("role").(string)
	roleConfig.Description = d.Get("description").(string)
	response, err := cidaasClient.UpsertRole(roleConfig)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create role %+v", roleConfig.Role),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("role", response.Data.Role); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error while setting role to Role resource",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.Role)
	return diags
}

func resourceRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	role := d.Id()
	response, err := cidaasClient.GetRole(role)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to read role %+v", role),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("description", response.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role", response.Data.Role); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", response.Data.Name); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRoleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	role := d.Id()
	_, err := cidaasClient.DeleteRole(role)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete role %+v", role),
			Detail:   err.Error(),
		})
	}
	return diags
}
