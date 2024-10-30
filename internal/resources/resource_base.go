package resources

import (
	"context"
	"fmt"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/cidaas"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// nolint:revive
const (
	RESOURCE_APP                = "cidaas_app"                // nolint:stylecheck
	RESOURCE_CONSENT_GROUP      = "cidaas_consent_group"      // nolint:stylecheck
	RESOURCE_CONSENT_VERSION    = "cidaas_consent_version"    // nolint:stylecheck
	RESOURCE_CONSENT            = "cidaas_consent"            // nolint:stylecheck
	RESOURCE_CUSTOM_PROVIDER    = "cidaas_custom_provider"    // nolint:stylecheck
	RESOURCE_GROUP_TYPE         = "cidaas_group_type"         // nolint:stylecheck
	RESOURCE_HOSTED_PAGE        = "cidaas_hosted_page"        // nolint:stylecheck
	RESOURCE_PASSWORD_POLICY    = "cidaas_password_policy"    // nolint:stylecheck
	RESOURCE_REGISTRATION_FIELD = "cidaas_registration_field" // nolint:stylecheck
	RESOURCE_ROLE               = "cidaas_role"               // nolint:stylecheck
	RESOURCE_SCOPE_GROUP        = "cidaas_scope_group"        // nolint:stylecheck
	RESOURCE_SCOPE              = "cidaas_scope"              // nolint:stylecheck
	RESOURCE_SOCIAL_PROVIDER    = "cidaas_social_provider"    // nolint:stylecheck
	RESOURCE_TEMPLATE_GROUP     = "cidaas_template_group"     // nolint:stylecheck
	RESOURCE_TEMPLATE           = "cidaas_template"           // nolint:stylecheck
	RESOURCE_USER_GROUP         = "cidaas_user_groups"        // nolint:stylecheck
	RESOURCE_WEBHOOK            = "cidaas_webhook"            // nolint:stylecheck
)

type BaseResourceConfig struct {
	Name   string
	Schema *schema.Schema
}

type BaseResource struct {
	Config       BaseResourceConfig
	cidaasClient *cidaas.Client
}

func NewBaseResource(cfg BaseResourceConfig) BaseResource {
	return BaseResource{
		Config: cfg,
	}
}

func (r *BaseResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured
	if req.ProviderData == nil {
		return
	}

	r.cidaasClient = GetResourceMeta(req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
}

func GetResourceMeta(
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) *cidaas.Client {
	client, ok := req.ProviderData.(*cidaas.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected cidaas.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}
	return client
}

func (r *BaseResource) Metadata(
	_ context.Context,
	_ resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = r.Config.Name
}

func (r *BaseResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	if r.Config.Schema == nil {
		resp.Diagnostics.AddError(
			"Missing Schema",
			"Base resource was not provided a schema. "+
				"Please provide a Schema config attribute or implement, the Schema(...) function.",
		)
		return
	}
	resp.Schema = *r.Config.Schema
}
