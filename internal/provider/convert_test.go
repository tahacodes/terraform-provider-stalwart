package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
