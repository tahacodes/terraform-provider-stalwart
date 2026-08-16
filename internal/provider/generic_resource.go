package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tahacodes/terraform-provider-stalwart/internal/client"
)

type resourceDescriptor struct {
	Name         string
	JMAPType     string
	Variant      string
	Singleton    bool
	ReloadAction string
	Schema       schema.Schema
}

const singletonID = "singleton"

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
	var plan, config types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	object, err := r.payload(ctx, plan, config, false)
	if err != nil {
		resp.Diagnostics.AddError("Unable to build request", err.Error())
		return
	}

	var id string
	if r.descriptor.Singleton {
		id = singletonID
		if err := r.client.Update(ctx, r.descriptor.JMAPType, id, object); err != nil {
			resp.Diagnostics.AddError("Unable to apply "+r.descriptor.Name+" settings", err.Error())
			return
		}
	} else {
		created, err := r.client.Create(ctx, r.descriptor.JMAPType, object)
		if err != nil {
			resp.Diagnostics.AddError("Unable to create "+r.descriptor.Name, err.Error())
			return
		}

		id, _ = created["id"].(string)
		if id == "" {
			resp.Diagnostics.AddError("Unable to create "+r.descriptor.Name, "server returned no id")
			return
		}
	}

	if err := r.client.Reload(ctx, r.descriptor.ReloadAction); err != nil {
		resp.Diagnostics.AddWarning("Unable to reload server configuration", err.Error())
	}

	state, err := r.read(ctx, id, plan)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read back "+r.descriptor.Name, err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, fillUnknowns(plan, state))...)
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
	var plan, current, config types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &current)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := attributeString(current, "id")
	patch, err := r.payload(ctx, plan, config, true)
	if err != nil {
		resp.Diagnostics.AddError("Unable to build request", err.Error())
		return
	}

	if err := r.client.Update(ctx, r.descriptor.JMAPType, id, patch); err != nil {
		resp.Diagnostics.AddError("Unable to update "+r.descriptor.Name, err.Error())
		return
	}

	if err := r.client.Reload(ctx, r.descriptor.ReloadAction); err != nil {
		resp.Diagnostics.AddWarning("Unable to reload server configuration", err.Error())
	}

	state, err := r.read(ctx, id, plan)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read back "+r.descriptor.Name, err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, fillUnknowns(plan, state))...)
}

func (r *genericResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var current types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &current)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.descriptor.Singleton {
		return
	}

	id := attributeString(current, "id")
	if id == "" {
		return
	}

	if err := r.client.Destroy(ctx, r.descriptor.JMAPType, id); err != nil {
		resp.Diagnostics.AddError("Unable to delete "+r.descriptor.Name, err.Error())
		return
	}

	if err := r.client.Reload(ctx, r.descriptor.ReloadAction); err != nil {
		resp.Diagnostics.AddWarning("Unable to reload server configuration", err.Error())
	}
}

func (r *genericResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if r.descriptor.Singleton {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), singletonID)...)
		return
	}

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *genericResource) payload(ctx context.Context, plan, config types.Object, clearOmitted bool) (map[string]any, error) {
	out := make(map[string]any)

	for name, value := range plan.Attributes() {
		if name == "id" || r.serverOwned(name) || writeOnlyBookkeeping(name) || value.IsUnknown() {
			continue
		}

		if value.IsNull() {
			if clearOmitted {
				out[rootFieldName(name)] = emptyValue(value)
			}
			continue
		}

		attribute, declared := r.descriptor.Schema.Attributes[name]
		converted, err := encodeChild(ctx, attribute, declared, value)
		if err != nil {
			return nil, fmt.Errorf("attribute %q: %w", name, err)
		}
		out[rootFieldName(name)] = converted
	}

	for name, value := range config.Attributes() {
		if !strings.HasSuffix(name, writeOnlySuffix) || value.IsNull() || value.IsUnknown() {
			continue
		}

		converted, err := toJMAP(ctx, value)
		if err != nil {
			return nil, fmt.Errorf("attribute %q: %w", name, err)
		}
		out[rootFieldName(strings.TrimSuffix(name, writeOnlySuffix))] = converted
	}

	if r.descriptor.Variant != "" {
		out["@type"] = r.descriptor.Variant
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
	if err := verifyVariant(remote, r.descriptor.JMAPType, r.descriptor.Variant, id); err != nil {
		return types.ObjectNull(attrTypes), err
	}

	object, err := remoteToObject(ctx, attrTypes, remote, id)
	if err != nil {
		return types.ObjectNull(attrTypes), err
	}

	elements := make(map[string]attr.Value, len(attrTypes))
	for name, value := range object.Attributes() {
		if writeOnlyBookkeeping(name) {
			elements[name] = referenceOrNull(reference, name, attrTypes[name])
			continue
		}
		elements[name] = preserveMasked(reference.Attributes()[name], value)
	}

	merged, diags := types.ObjectValue(attrTypes, elements)

	return merged, diagsError(diags)
}

func referenceOrNull(reference types.Object, name string, attrType attr.Type) attr.Value {
	if value, ok := reference.Attributes()[name]; ok && !value.IsUnknown() {
		return value
	}

	null, err := nullOf(attrType)
	if err != nil {
		return types.StringNull()
	}

	return null
}

func verifyVariant(remote map[string]any, jmapType, variant, id string) error {
	if variant == "" {
		return nil
	}
	if remoteVariant, ok := remote["@type"].(string); ok && remoteVariant != variant {
		return fmt.Errorf("object %q is a %s of variant %q, not %q", id, jmapType, remoteVariant, variant)
	}

	return nil
}

func remoteToObject(ctx context.Context, attrTypes map[string]attr.Type, remote map[string]any, id string) (types.Object, error) {
	elements := make(map[string]attr.Value, len(attrTypes))
	for name, attrType := range attrTypes {
		if name == "id" {
			elements[name] = types.StringValue(id)
			continue
		}
		value, err := toTerraform(ctx, attrType, remote[rootFieldName(name)])
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

func (r *genericResource) serverOwned(name string) bool {
	attribute, ok := r.descriptor.Schema.Attributes[name]
	if !ok {
		return false
	}

	return attribute.IsComputed() && !attribute.IsOptional()
}

func emptyValue(value attr.Value) any {
	switch value.(type) {
	case types.Set, types.Map, types.List:
		return map[string]any{}
	}

	return nil
}
