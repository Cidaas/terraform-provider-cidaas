package cidaas

import (
	"context"
	"fmt"

	"terraform-provider-cidaas/helper/cidaas"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserGroupUpsert,
		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupUpsert,
		DeleteContext: resourceUserGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"group_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringDoesNotContainAny(" "),
			},
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logo_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"make_first_user_admin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"custom_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"member_profile_visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "public",
				ValidateFunc: validation.StringInSlice([]string{"public", "full"}, true),
			},
			"none_member_profile_visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: validation.StringInSlice([]string{"none", "public"}, true),
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
		},
	}
}

func resourceUserGroupUpsert(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var ug cidaas.UserGroup
	var response *cidaas.UserGroupResponse
	var err error
	var errMsg string

	ug.GroupType = d.Get("group_type").(string)
	ug.GroupId = d.Get("group_id").(string)
	ug.GroupName = d.Get("group_name").(string)
	ug.LogoUrl = d.Get("logo_url").(string)
	ug.Description = d.Get("description").(string)
	ug.MakeFirstUserAdmin = d.Get("make_first_user_admin").(bool)
	ug.CustomFields = d.Get("custom_fields").(map[string]interface{})
	ug.MemberProfileVisibility = d.Get("member_profile_visibility").(string)
	ug.NoneMemberProfileVisibility = d.Get("none_member_profile_visibility").(string)
	ug.ParentId = d.Get("parent_id").(string)

	cidaas_client := m.(cidaas.CidaasClient)
	group_id := d.Id()

	if group_id != "" {
		if group_id != ug.GroupId {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("group_id can't be modified"),
			})
			return diags
		}
		response, err = cidaas_client.UpdateUserGroup(ug)
		errMsg = fmt.Sprintf("failed to update user group %+v", ug.GroupId)
	} else {
		response, err = cidaas_client.CreateUserGroup(ug)
		errMsg = fmt.Sprintf("failed to create user group %+v", ug.GroupId)
	}
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  errMsg,
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(response.Data.GroupId)
	return diags
}

func resourceUserGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	group_id := d.Id()
	response, err := cidaas_client.GetUserGroup(group_id)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to read user group %+v", group_id),
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("group_type", response.Data.GroupType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_id", response.Data.GroupId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_name", response.Data.GroupName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("logo_url", response.Data.LogoUrl); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", response.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("make_first_user_admin", response.Data.MakeFirstUserAdmin); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("custom_fields", response.Data.CustomFields); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("member_profile_visibility", response.Data.MemberProfileVisibility); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("none_member_profile_visibility", response.Data.NoneMemberProfileVisibility); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("parent_id", response.Data.ParentId); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidaas_client := m.(cidaas.CidaasClient)
	group_id := d.Id()
	_, err := cidaas_client.DeleteUserGroup(group_id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to delete user group %+v", group_id),
			Detail:   err.Error(),
		})
	}
	return diags
}
