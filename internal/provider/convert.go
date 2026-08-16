package provider

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func toTerraform(ctx context.Context, t attr.Type, v any) (attr.Value, error) {
	if v == nil {
		return nullOf(t)
	}

	switch typed := t.(type) {
	case basetypes.StringType:
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("expected string, got %T", v)
		}
		return types.StringValue(s), nil

	case basetypes.BoolType:
		b, ok := v.(bool)
		if !ok {
			return nil, fmt.Errorf("expected bool, got %T", v)
		}
		return types.BoolValue(b), nil

	case basetypes.Int64Type:
		f, ok := v.(float64)
		if !ok {
			return nil, fmt.Errorf("expected number, got %T", v)
		}
		return types.Int64Value(int64(math.Round(f))), nil

	case basetypes.Float64Type:
		f, ok := v.(float64)
		if !ok {
			return nil, fmt.Errorf("expected number, got %T", v)
		}
		return types.Float64Value(f), nil

	case basetypes.SetType:
		return setToTerraform(ctx, typed.ElemType, v)

	case basetypes.ListType:
		return collectionToTerraform(ctx, typed.ElemType, v, false)

	case basetypes.MapType:
		raw, ok := v.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("expected object for map, got %T", v)
		}
		elements := make(map[string]attr.Value, len(raw))
		for k, item := range raw {
			converted, err := toTerraform(ctx, typed.ElemType, item)
			if err != nil {
				return nil, fmt.Errorf("map key %q: %w", k, err)
			}
			elements[k] = converted
		}
		value, diags := types.MapValue(typed.ElemType, elements)
		return value, diagsError(diags)

	case basetypes.ObjectType:
		raw, ok := v.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("expected object, got %T", v)
		}
		attrTypes := typed.AttributeTypes()
		elements := make(map[string]attr.Value, len(attrTypes))
		for name, attrType := range attrTypes {
			converted, err := toTerraform(ctx, attrType, raw[objectFieldName(name)])
			if err != nil {
				return nil, fmt.Errorf("attribute %q: %w", name, err)
			}
			elements[name] = converted
		}
		value, diags := types.ObjectValue(attrTypes, elements)
		return value, diagsError(diags)
	}

	return nil, fmt.Errorf("unsupported attribute type %s", t.String())
}

func setToTerraform(ctx context.Context, elemType attr.Type, v any) (attr.Value, error) {
	membership, ok := v.(map[string]any)
	if !ok {
		return collectionToTerraform(ctx, elemType, v, true)
	}

	members := make([]string, 0, len(membership))
	for key, present := range membership {
		if enabled, ok := present.(bool); ok && !enabled {
			continue
		}
		members = append(members, key)
	}
	sort.Strings(members)

	elements := make([]attr.Value, 0, len(members))
	for _, member := range members {
		converted, err := toTerraform(ctx, elemType, member)
		if err != nil {
			return nil, err
		}
		elements = append(elements, converted)
	}

	value, diags := types.SetValue(elemType, elements)

	return value, diagsError(diags)
}

func collectionToTerraform(ctx context.Context, elemType attr.Type, v any, isSet bool) (attr.Value, error) {
	raw, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("expected array, got %T", v)
	}

	elements := make([]attr.Value, 0, len(raw))
	for i, item := range raw {
		converted, err := toTerraform(ctx, elemType, item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}
		elements = append(elements, converted)
	}

	if isSet {
		value, diags := types.SetValue(elemType, elements)
		return value, diagsError(diags)
	}

	value, diags := types.ListValue(elemType, elements)
	return value, diagsError(diags)
}

func toJMAP(ctx context.Context, v attr.Value) (any, error) {
	if v == nil || v.IsNull() || v.IsUnknown() {
		return nil, nil
	}

	switch typed := v.(type) {
	case basetypes.StringValue:
		return typed.ValueString(), nil
	case basetypes.BoolValue:
		return typed.ValueBool(), nil
	case basetypes.Int64Value:
		return typed.ValueInt64(), nil
	case basetypes.Float64Value:
		return typed.ValueFloat64(), nil

	case basetypes.SetValue:
		membership := make(map[string]any, len(typed.Elements()))
		for _, item := range typed.Elements() {
			converted, err := toJMAP(ctx, item)
			if err != nil {
				return nil, err
			}
			key, ok := converted.(string)
			if !ok {
				return nil, fmt.Errorf("set members must be strings, got %T", converted)
			}
			membership[key] = true
		}
		return membership, nil

	case basetypes.ListValue:
		return elementsToJMAP(ctx, typed.Elements())

	case basetypes.MapValue:
		out := make(map[string]any, len(typed.Elements()))
		for k, item := range typed.Elements() {
			converted, err := toJMAP(ctx, item)
			if err != nil {
				return nil, err
			}
			out[k] = converted
		}
		return out, nil

	case basetypes.ObjectValue:
		out := make(map[string]any, len(typed.Attributes()))
		for name, item := range typed.Attributes() {
			converted, err := toJMAP(ctx, item)
			if err != nil {
				return nil, err
			}
			if converted == nil {
				continue
			}
			out[objectFieldName(name)] = converted
		}
		return out, nil
	}

	return nil, fmt.Errorf("unsupported value type %T", v)
}

func elementsToJMAP(ctx context.Context, elements []attr.Value) (any, error) {
	out := make([]any, 0, len(elements))
	for _, item := range elements {
		converted, err := toJMAP(ctx, item)
		if err != nil {
			return nil, err
		}
		out = append(out, converted)
	}

	return out, nil
}

func nullOf(t attr.Type) (attr.Value, error) {
	switch typed := t.(type) {
	case basetypes.StringType:
		return types.StringNull(), nil
	case basetypes.BoolType:
		return types.BoolNull(), nil
	case basetypes.Int64Type:
		return types.Int64Null(), nil
	case basetypes.Float64Type:
		return types.Float64Null(), nil
	case basetypes.SetType:
		return types.SetNull(typed.ElemType), nil
	case basetypes.ListType:
		return types.ListNull(typed.ElemType), nil
	case basetypes.MapType:
		return types.MapNull(typed.ElemType), nil
	case basetypes.ObjectType:
		return types.ObjectNull(typed.AttributeTypes()), nil
	}

	return nil, fmt.Errorf("unsupported attribute type %s", t.String())
}
