package cidaas

import (
	"context"
	"fmt"
	"terraform-provider-cidaas/helper/cidaas"
	"terraform-provider-cidaas/helper/util"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserGroupCategory() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role_mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"any_roles", "no_roles", "roles_required", "allowed_roles"}, false),
			},
			"group_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"allowed_roles": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: resourceUserGroupCategoryCreate,
		ReadContext:   resourceUserGroupCategoryRead,
		UpdateContext: resourceUserGroupCategoryUpdate,
		DeleteContext: resourceUserGroupCategoryDelete,
	}

}

func resourceUserGroupCategoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	var ugc cidaas.UserGroupCategory
	ugc.RoleMode = d.Get("role_mode").(string)
	ugc.GroupType = d.Get("group_type").(string)
	ugc.Description = d.Get("description").(string)
	ugc.AllowedRoles = util.InterfaceArray2StringArray(d.Get("allowed_roles").([]interface{}))
	ugc.ObjectOwner = "client"
	response, err := cidaasClient.CreateUserGroupCategory(ugc)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create user group category %+v", ugc.GroupType),
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.GroupType)
	return diags
}

func resourceUserGroupCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	group_type := d.Id()
	response, err := cidaasClient.GetUserGroupCategory(group_type)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to read user group category %+v", group_type),
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("role_mode", response.Data.RoleMode); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_type", response.Data.GroupType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", response.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_roles", response.Data.AllowedRoles); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserGroupCategoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	var ugc cidaas.UserGroupCategory
	if d.Id() != d.Get("group_type").(string) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("group_type can't be modified"),
		})
		return diags
	}
	ugc.RoleMode = d.Get("role_mode").(string)
	ugc.GroupType = d.Id()
	ugc.Description = d.Get("description").(string)
	ugc.AllowedRoles = util.InterfaceArray2StringArray(d.Get("allowed_roles").([]interface{}))
	ugc.ObjectOwner = "client"
	_, err := cidaasClient.UpdateUserGroupCategory(ugc)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to update user group category %+v", ugc.GroupType),
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}
func resourceUserGroupCategoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaasClient := m.(cidaas.CidaasClient)
	group_type := d.Id()
	_, err := cidaasClient.DeleteUserGroupCategory(group_type)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete user group category %+v", group_type),
			Detail:   err.Error(),
		})
	}
	return diags
}
