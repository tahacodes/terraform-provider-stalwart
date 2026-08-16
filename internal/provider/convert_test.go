package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestSetDecodesMembershipMap(t *testing.T) {
	t.Parallel()

	remote := map[string]any{"Dkim1RsaSha256": true, "Dkim1Ed25519Sha256": true, "Disabled": false}

	value, err := toTerraform(context.Background(), types.SetType{ElemType: types.StringType}, remote)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	set, ok := value.(types.Set)
	if !ok {
		t.Fatalf("expected types.Set, got %T", value)
	}

	if got := len(set.Elements()); got != 2 {
		t.Fatalf("expected 2 members, got %d", got)
	}

	first, ok := set.Elements()[0].(types.String)
	if !ok || first.ValueString() != "Dkim1Ed25519Sha256" {
		t.Fatalf("expected sorted members, got %v", set.Elements())
	}
}

func TestSetEncodesMembershipMap(t *testing.T) {
	t.Parallel()

	set, diags := types.SetValue(types.StringType, []attr.Value{
		types.StringValue("alias@example.com"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	encoded, err := toJMAP(context.Background(), set)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	membership, ok := encoded.(map[string]any)
	if !ok {
		t.Fatalf("expected membership map, got %T", encoded)
	}
	if membership["alias@example.com"] != true {
		t.Fatalf("expected member set to true, got %v", membership)
	}
}

func TestSetRoundTrip(t *testing.T) {
	t.Parallel()

	original := map[string]any{"b": true, "a": true}

	decoded, err := toTerraform(context.Background(), types.SetType{ElemType: types.StringType}, original)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	encoded, err := toJMAP(context.Background(), decoded)
	if err != nil {
		t.Fatalf("encode: %v", err)
	}

	membership, ok := encoded.(map[string]any)
	if !ok {
		t.Fatalf("expected map, got %T", encoded)
	}
	if len(membership) != 2 || membership["a"] != true || membership["b"] != true {
		t.Fatalf("round trip lost members: %v", membership)
	}
}

func unionType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"type":             types.StringType,
		"acme_provider_id": types.StringType,
	}}
}

func TestUnionDecodesAtTypeDiscriminator(t *testing.T) {
	t.Parallel()

	remote := map[string]any{"@type": "Automatic", "acmeProviderId": "i3rbkid1acaa"}

	value, err := toTerraform(context.Background(), unionType(), remote)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	object, ok := value.(types.Object)
	if !ok {
		t.Fatalf("expected types.Object, got %T", value)
	}

	discriminator, ok := object.Attributes()["type"].(types.String)
	if !ok || discriminator.ValueString() != "Automatic" {
		t.Fatalf("expected type Automatic, got %v", object.Attributes()["type"])
	}

	provider, ok := object.Attributes()["acme_provider_id"].(types.String)
	if !ok || provider.ValueString() != "i3rbkid1acaa" {
		t.Fatalf("expected camelCase field mapped, got %v", object.Attributes()["acme_provider_id"])
	}
}

func TestUnionEncodesAtTypeDiscriminator(t *testing.T) {
	t.Parallel()

	object, diags := types.ObjectValue(unionType().AttrTypes, map[string]attr.Value{
		"type":             types.StringValue("Manual"),
		"acme_provider_id": types.StringNull(),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	encoded, err := toJMAP(context.Background(), object)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out, ok := encoded.(map[string]any)
	if !ok {
		t.Fatalf("expected map, got %T", encoded)
	}
	if out["@type"] != "Manual" {
		t.Fatalf("expected @type discriminator, got %v", out)
	}
	if _, present := out["type"]; present {
		t.Fatalf("terraform attribute name leaked to the wire: %v", out)
	}
}

func aliasType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"name": types.StringType,
	}}
}

func TestListDecodesIndexedMap(t *testing.T) {
	t.Parallel()

	remote := map[string]any{
		"1": map[string]any{"name": "second"},
		"0": map[string]any{"name": "first"},
	}

	value, err := toTerraform(context.Background(), types.ListType{ElemType: aliasType()}, remote)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	list, ok := value.(types.List)
	if !ok {
		t.Fatalf("expected types.List, got %T", value)
	}
	if got := len(list.Elements()); got != 2 {
		t.Fatalf("expected 2 elements, got %d", got)
	}

	first, ok := list.Elements()[0].(types.Object)
	if !ok {
		t.Fatalf("expected object element, got %T", list.Elements()[0])
	}
	name, ok := first.Attributes()["name"].(types.String)
	if !ok || name.ValueString() != "first" {
		t.Fatalf("expected numeric index ordering, got %v", list.Elements())
	}
}

func TestListEncodesIndexedMap(t *testing.T) {
	t.Parallel()

	element, diags := types.ObjectValue(aliasType().AttrTypes, map[string]attr.Value{
		"name": types.StringValue("alias"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	list, diags := types.ListValue(aliasType(), []attr.Value{element})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	encoded, err := toJMAP(context.Background(), list)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	indexed, ok := encoded.(map[string]any)
	if !ok {
		t.Fatalf("expected indexed map, got %T", encoded)
	}
	entry, ok := indexed["0"].(map[string]any)
	if !ok || entry["name"] != "alias" {
		t.Fatalf("expected element under key \"0\", got %v", indexed)
	}
}

func TestEmptyIndexedMapDecodesToNullList(t *testing.T) {
	t.Parallel()

	value, err := toTerraform(context.Background(), types.ListType{ElemType: aliasType()}, map[string]any{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !value.IsNull() {
		t.Fatalf("expected null list, got %v", value)
	}
}

func TestPreserveMaskedKeepsReferenceSecret(t *testing.T) {
	t.Parallel()

	secretType := types.ObjectType{AttrTypes: map[string]attr.Type{
		"type":   types.StringType,
		"secret": types.StringType,
	}}
	reference, diags := types.ObjectValue(secretType.AttrTypes, map[string]attr.Value{
		"type":   types.StringValue("Value"),
		"secret": types.StringValue("hunter2"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	fresh, diags := types.ObjectValue(secretType.AttrTypes, map[string]attr.Value{
		"type":   types.StringValue("Value"),
		"secret": types.StringValue("****"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	merged, ok := preserveMasked(reference, fresh).(types.Object)
	if !ok {
		t.Fatalf("expected object, got %T", preserveMasked(reference, fresh))
	}
	secret, ok := merged.Attributes()["secret"].(types.String)
	if !ok || secret.ValueString() != "hunter2" {
		t.Fatalf("expected reference secret kept, got %v", merged.Attributes()["secret"])
	}
}

func TestPreserveMaskedKeepsFreshValueWithoutReference(t *testing.T) {
	t.Parallel()

	fresh := types.StringValue("****")
	if got := preserveMasked(types.StringNull(), fresh); !got.Equal(fresh) {
		t.Fatalf("expected masked value kept when reference is null, got %v", got)
	}
}

func TestNullValuesDecodeToNull(t *testing.T) {
	t.Parallel()

	value, err := toTerraform(context.Background(), types.StringType, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !value.IsNull() {
		t.Fatalf("expected null, got %v", value)
	}
}

func TestJMAPNameMapping(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"email_address":    "emailAddress",
		"domain_id":        "domainId",
		"name":             "name",
		"use_x_forwarded":  "useXForwarded",
		"member_group_ids": "memberGroupIds",
	}

	for input, want := range cases {
		if got := jmapName(input); got != want {
			t.Errorf("jmapName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestPayloadMergesWriteOnlySecrets(t *testing.T) {
	t.Parallel()

	r := &genericResource{descriptor: resourceDescriptor{
		Name:     "oauth_client",
		JMAPType: "x:OAuthClient",
		Schema: rschema.Schema{Attributes: map[string]rschema.Attribute{
			"id":                rschema.StringAttribute{Computed: true},
			"client_id":         rschema.StringAttribute{Required: true},
			"secret":            rschema.StringAttribute{Optional: true, Computed: true, Sensitive: true},
			"secret_wo":         rschema.StringAttribute{Optional: true, WriteOnly: true, Sensitive: true},
			"secret_wo_version": rschema.Int64Attribute{Optional: true},
		}},
	}}

	attrTypes := map[string]attr.Type{
		"id":                types.StringType,
		"client_id":         types.StringType,
		"secret":            types.StringType,
		"secret_wo":         types.StringType,
		"secret_wo_version": types.Int64Type,
	}
	plan, diags := types.ObjectValue(attrTypes, map[string]attr.Value{
		"id":                types.StringUnknown(),
		"client_id":         types.StringValue("client"),
		"secret":            types.StringUnknown(),
		"secret_wo":         types.StringNull(),
		"secret_wo_version": types.Int64Value(1),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	config, diags := types.ObjectValue(attrTypes, map[string]attr.Value{
		"id":                types.StringNull(),
		"client_id":         types.StringValue("client"),
		"secret":            types.StringNull(),
		"secret_wo":         types.StringValue("hunter2"),
		"secret_wo_version": types.Int64Value(1),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	payload, err := r.payload(context.Background(), plan, config, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if payload["secret"] != "hunter2" {
		t.Fatalf("expected write-only value on the wire field, got %v", payload)
	}
	if _, present := payload["secretWoVersion"]; present {
		t.Fatalf("version attribute leaked to the wire: %v", payload)
	}
	if _, present := payload["secretWo"]; present {
		t.Fatalf("write-only attribute name leaked to the wire: %v", payload)
	}
}
