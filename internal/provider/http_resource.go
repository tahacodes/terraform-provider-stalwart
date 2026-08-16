package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tahacodes/terraform-provider-stalwart/internal/client"
)

const (
	httpObjectType = "x:Http"
	httpSingleton  = "singleton"
)

var (
	_ resource.Resource                = (*httpResource)(nil)
	_ resource.ResourceWithConfigure   = (*httpResource)(nil)
	_ resource.ResourceWithImportState = (*httpResource)(nil)
)

type httpResource struct {
	client *client.Client
}

type httpModel struct {
	ID                types.String `tfsdk:"id"`
	UsePermissiveCors types.Bool   `tfsdk:"use_permissive_cors"`
	EnableHsts        types.Bool   `tfsdk:"enable_hsts"`
	UseXForwarded     types.Bool   `tfsdk:"use_x_forwarded"`
	RedirectRoot      types.String `tfsdk:"redirect_root"`
	ResponseHeaders   types.Map    `tfsdk:"response_headers"`
}

func NewHTTPResource() resource.Resource {
	return &httpResource{}
}

func (r *httpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_http"
}

func (r *httpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages the Stalwart HTTP server settings (the `x:Http` singleton object). This object always exists, so creating the resource adopts the existing settings rather than provisioning anything new, and destroying it removes the resource from state without changing the server.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Always `singleton`.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"use_permissive_cors": schema.BoolAttribute{
				MarkdownDescription: "Allow all origins in the CORS policy for the HTTP server. Required when a browser-based client is served from a different origin than the server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_hsts": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Strict Transport Security.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_x_forwarded": schema.BoolAttribute{
				MarkdownDescription: "Use the `Forwarded` or `X-Forwarded-For` header to determine the client IP address.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"redirect_root": schema.StringAttribute{
				MarkdownDescription: "URL to redirect to when a client requests the root path. When unset the server returns 404.",
				Optional:            true,
			},
			"response_headers": schema.MapAttribute{
				MarkdownDescription: "Additional headers to include in HTTP responses.",
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *httpResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected provider data",
			fmt.Sprintf("Expected *client.Client, got %T.", req.ProviderData),
		)
		return
	}

	r.client = c
}

func (r *httpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan httpModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.apply(ctx, plan); err != nil {
		resp.Diagnostics.AddError("Unable to update Stalwart HTTP settings", err.Error())
		return
	}

	plan.ID = types.StringValue(httpSingleton)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *httpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state httpModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	objects, err := r.client.Get(ctx, httpObjectType, []string{httpSingleton})
	if err != nil {
		resp.Diagnostics.AddError("Unable to read Stalwart HTTP settings", err.Error())
		return
	}
	if len(objects) == 0 {
		resp.State.RemoveResource(ctx)
		return
	}

	object := objects[0]
	state.ID = types.StringValue(httpSingleton)
	state.UsePermissiveCors = types.BoolValue(boolField(object, "usePermissiveCors"))
	state.EnableHsts = types.BoolValue(boolField(object, "enableHsts"))
	state.UseXForwarded = types.BoolValue(boolField(object, "useXForwarded"))

	if v, ok := object["redirectRoot"].(string); ok && v != "" {
		state.RedirectRoot = types.StringValue(v)
	} else {
		state.RedirectRoot = types.StringNull()
	}

	headers, diags := stringMap(ctx, object["responseHeaders"])
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !state.ResponseHeaders.IsNull() || len(headers.Elements()) > 0 {
		state.ResponseHeaders = headers
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *httpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan httpModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.apply(ctx, plan); err != nil {
		resp.Diagnostics.AddError("Unable to update Stalwart HTTP settings", err.Error())
		return
	}

	plan.ID = types.StringValue(httpSingleton)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *httpResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}

func (r *httpResource) ImportState(ctx context.Context, _ resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, pathRoot("id"), httpSingleton)...)
}

func (r *httpResource) apply(ctx context.Context, plan httpModel) error {
	patch := map[string]any{
		"usePermissiveCors": plan.UsePermissiveCors.ValueBool(),
		"enableHsts":        plan.EnableHsts.ValueBool(),
		"useXForwarded":     plan.UseXForwarded.ValueBool(),
	}

	if plan.RedirectRoot.IsNull() {
		patch["redirectRoot"] = nil
	} else {
		patch["redirectRoot"] = plan.RedirectRoot.ValueString()
	}

	headers := map[string]any{}
	for k, v := range plan.ResponseHeaders.Elements() {
		s, ok := v.(types.String)
		if !ok {
			return fmt.Errorf("response_headers value for %q is not a string", k)
		}
		headers[k] = s.ValueString()
	}
	patch["responseHeaders"] = headers

	return r.client.Update(ctx, httpObjectType, httpSingleton, patch)
}
