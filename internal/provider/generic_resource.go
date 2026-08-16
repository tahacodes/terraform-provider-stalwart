package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tahacodes/terraform-provider-stalwart/internal/client"
)

type resourceDescriptor struct {
	Name     string
	JMAPType string
	Variant  string
	Schema   schema.Schema
}

var (
	_ resource.Resource                = (*genericResource)(nil)
	_ resource.ResourceWithConfigure   = (*genericResource)(nil)
	_ resource.ResourceWithImportState = (*genericResource)(nil)
)

type genericResource struct {
	descriptor resourceDescriptor
	client     *client.Client
}

func newGenericResource(d resourceDescriptor) func() resource.Resource {
	return func() resource.Resource {
		return &genericResource{descriptor: d}
	}
}

func (r *genericResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.descriptor.Name
}

func (r *genericResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.descriptor.Schema
}

func (r *genericResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *genericResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	object, err := r.payload(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Unable to build request", err.Error())
		return
	}

	created, err := r.client.Create(ctx, r.descriptor.JMAPType, object)
	if err != nil {
		resp.Diagnostics.AddError("Unable to create "+r.descriptor.Name, err.Error())
		return
	}

	id, _ := created["id"].(string)
	if id == "" {
		resp.Diagnostics.AddError("Unable to create "+r.descriptor.Name, "server returned no id")
		return
	}

	state, err := r.read(ctx, id, plan)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read back "+r.descriptor.Name, err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *genericResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var current types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &current)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := attributeString(current, "id")
	if id == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	state, err := r.read(ctx, id, current)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read "+r.descriptor.Name, err.Error())
		return
	}
	if state.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *genericResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, current types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &current)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := attributeString(current, "id")
	patch, err := r.payload(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Unable to build request", err.Error())
		return
	}

	if err := r.client.Update(ctx, r.descriptor.JMAPType, id, patch); err != nil {
		resp.Diagnostics.AddError("Unable to update "+r.descriptor.Name, err.Error())
		return
	}

	state, err := r.read(ctx, id, plan)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read back "+r.descriptor.Name, err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *genericResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var current types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &current)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := attributeString(current, "id")
	if id == "" {
		return
	}

	if err := r.client.Destroy(ctx, r.descriptor.JMAPType, id); err != nil {
		resp.Diagnostics.AddError("Unable to delete "+r.descriptor.Name, err.Error())
	}
}

func (r *genericResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *genericResource) payload(ctx context.Context, plan types.Object) (map[string]any, error) {
	out := make(map[string]any)

	for name, value := range plan.Attributes() {
		if name == "id" {
			continue
		}
		converted, err := toJMAP(ctx, value)
		if err != nil {
			return nil, fmt.Errorf("attribute %q: %w", name, err)
		}
		if converted == nil {
			continue
		}
		out[jmapName(name)] = converted
	}

	if r.descriptor.Variant != "" {
		out["_type"] = r.descriptor.Variant
	}

	return out, nil
}

func (r *genericResource) read(ctx context.Context, id string, reference types.Object) (types.Object, error) {
	attrTypes := reference.AttributeTypes(ctx)

	objects, err := r.client.Get(ctx, r.descriptor.JMAPType, []string{id})
	if err != nil {
		return types.ObjectNull(attrTypes), err
	}
	if len(objects) == 0 {
		return types.ObjectNull(attrTypes), nil
	}

	remote := objects[0]
	elements := make(map[string]attr.Value, len(attrTypes))

	for name, attrType := range attrTypes {
		if name == "id" {
			elements[name] = types.StringValue(id)
			continue
		}
		value, err := toTerraform(ctx, attrType, remote[jmapName(name)])
		if err != nil {
			return types.ObjectNull(attrTypes), fmt.Errorf("attribute %q: %w", name, err)
		}
		elements[name] = value
	}

	object, diags := types.ObjectValue(attrTypes, elements)

	return object, diagsError(diags)
}

func attributeString(object types.Object, name string) string {
	value, ok := object.Attributes()[name]
	if !ok {
		return ""
	}

	s, ok := value.(types.String)
	if !ok || s.IsNull() || s.IsUnknown() {
		return ""
	}

	return s.ValueString()
}
