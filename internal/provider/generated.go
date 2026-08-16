package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var generatedResources = []resourceDescriptor{
	{
		Name:     "role",
		JMAPType: "x:Role",
		Schema: schema.Schema{
			MarkdownDescription: "Manages a Stalwart role (the `x:Role` object).",
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					MarkdownDescription: "Server-assigned identifier.",
					Computed:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				},
				"description": schema.StringAttribute{
					MarkdownDescription: "Description of the role",
					Optional:            true,
				},
				"disabled_permissions": schema.SetAttribute{
					MarkdownDescription: "List of permissions that are explicitly disabled, even if they would be inherited through other roles or groups. This takes precedence over enabled permissions.",
					Optional:            true,
					ElementType:         types.StringType,
				},
				"enabled_permissions": schema.SetAttribute{
					MarkdownDescription: "List of permissions that are explicitly enabled.",
					Optional:            true,
					ElementType:         types.StringType,
				},
				"member_tenant_id": schema.StringAttribute{
					MarkdownDescription: "Identifier for the tenant this role belongs to",
					Optional:            true,
				},
				"role_ids": schema.SetAttribute{
					MarkdownDescription: "List of roles this role extends",
					Optional:            true,
					ElementType:         types.StringType,
				},
			},
		},
	},
}
