package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tahacodes/terraform-provider-stalwart/internal/client"
)

type dataSourceDescriptor struct {
	Name      string
	JMAPType  string
	Variant   string
	Singleton bool
	Plural    bool
	HasName   bool
	Schema    schema.Schema
}

var (
	_ datasource.DataSource              = (*genericDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*genericDataSource)(nil)
)

type genericDataSource struct {
	descriptor dataSourceDescriptor
	client     *client.Client
}

func newGenericDataSource(d dataSourceDescriptor) func() datasource.DataSource {
	return func() datasource.DataSource {
		return &genericDataSource{descriptor: d}
	}
}

func (d *genericDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.descriptor.Name
}

func (d *genericDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = d.descriptor.Schema
}

func (d *genericDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = c
}

func (d *genericDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.descriptor.Plural {
		d.readPlural(ctx, resp)
		return
	}

	var config types.Object
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	remote, id, err := d.lookup(ctx, config)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read "+d.descriptor.Name, err.Error())
		return
	}

	state, err := remoteToObject(ctx, config.AttributeTypes(ctx), remote, id)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read "+d.descriptor.Name, err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (d *genericDataSource) readPlural(ctx context.Context, resp *datasource.ReadResponse) {
	ids, err := d.client.Query(ctx, d.descriptor.JMAPType)
	if err != nil {
		resp.Diagnostics.AddError("Unable to list "+d.descriptor.JMAPType, err.Error())
		return
	}

	elements := make([]attr.Value, 0, len(ids))
	for _, id := range ids {
		elements = append(elements, types.StringValue(id))
	}
	set, diags := types.SetValue(types.StringType, elements)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ids"), set)...)
}

func (d *genericDataSource) lookup(ctx context.Context, config types.Object) (map[string]any, string, error) {
	if d.descriptor.Singleton {
		return d.byID(ctx, singletonID)
	}

	if id := attributeString(config, "id"); id != "" {
		return d.byID(ctx, id)
	}

	if name := attributeString(config, "name"); d.descriptor.HasName && name != "" {
		return d.byName(ctx, name)
	}

	if d.descriptor.HasName {
		return nil, "", fmt.Errorf("set either id or name")
	}

	return nil, "", fmt.Errorf("set id")
}

func (d *genericDataSource) byID(ctx context.Context, id string) (map[string]any, string, error) {
	objects, err := d.client.Get(ctx, d.descriptor.JMAPType, []string{id})
	if err != nil {
		return nil, "", err
	}
	if len(objects) == 0 {
		return nil, "", fmt.Errorf("%s %q not found", d.descriptor.JMAPType, id)
	}
	if err := verifyVariant(objects[0], d.descriptor.JMAPType, d.descriptor.Variant, id); err != nil {
		return nil, "", err
	}

	return objects[0], id, nil
}

func (d *genericDataSource) byName(ctx context.Context, name string) (map[string]any, string, error) {
	objects, err := d.client.Get(ctx, d.descriptor.JMAPType, nil)
	if err != nil {
		return nil, "", err
	}

	var matches []map[string]any
	for _, object := range objects {
		if object["name"] != name {
			continue
		}
		if d.descriptor.Variant != "" {
			if remoteVariant, ok := object["@type"].(string); ok && remoteVariant != d.descriptor.Variant {
				continue
			}
		}
		matches = append(matches, object)
	}

	switch len(matches) {
	case 0:
		return nil, "", fmt.Errorf("no %s named %q found", d.descriptor.JMAPType, name)
	case 1:
		id, _ := matches[0]["id"].(string)
		if id == "" {
			return nil, "", fmt.Errorf("%s named %q has no id", d.descriptor.JMAPType, name)
		}
		return matches[0], id, nil
	}

	return nil, "", fmt.Errorf("%d %s objects named %q found, use id instead", len(matches), d.descriptor.JMAPType, name)
}
