package provider

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"

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
		return listToTerraform(ctx, typed.ElemType, v)

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
	if len(membership) == 0 {
		return types.SetNull(elemType), nil
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

func listToTerraform(ctx context.Context, elemType attr.Type, v any) (attr.Value, error) {
	indexed, ok := v.(map[string]any)
	if !ok {
		return collectionToTerraform(ctx, elemType, v, false)
	}
	if len(indexed) == 0 {
		return types.ListNull(elemType), nil
	}

	keys := make([]int, 0, len(indexed))
	byIndex := make(map[int]any, len(indexed))
	for key, item := range indexed {
		index, err := strconv.Atoi(key)
		if err != nil {
			return nil, fmt.Errorf("list index %q is not a number", key)
		}
		keys = append(keys, index)
		byIndex[index] = item
	}
	sort.Ints(keys)

	elements := make([]attr.Value, 0, len(keys))
	for _, index := range keys {
		converted, err := toTerraform(ctx, elemType, byIndex[index])
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", index, err)
		}
		elements = append(elements, converted)
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
		out := make(map[string]any, len(typed.Elements()))
		for i, item := range typed.Elements() {
			converted, err := toJMAP(ctx, item)
			if err != nil {
				return nil, err
			}
			out[strconv.Itoa(i)] = converted
		}
		return out, nil

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

const maskedSecret = "****"

func fillUnknowns(plan, fresh attr.Value) attr.Value {
	if plan == nil || plan.IsUnknown() {
		return fresh
	}
	if plan.IsNull() || fresh == nil || fresh.IsNull() {
		return plan
	}

	switch planTyped := plan.(type) {
	case basetypes.ObjectValue:
		freshTyped, ok := fresh.(basetypes.ObjectValue)
		if !ok {
			return plan
		}
		attrTypes := planTyped.AttributeTypes(context.Background())
		elements := make(map[string]attr.Value, len(attrTypes))
		for name, item := range planTyped.Attributes() {
			elements[name] = fillUnknowns(item, freshTyped.Attributes()[name])
		}
		merged, diags := types.ObjectValue(attrTypes, elements)
		if diags.HasError() {
			return plan
		}
		return merged

	case basetypes.ListValue:
		freshTyped, ok := fresh.(basetypes.ListValue)
		if !ok || len(freshTyped.Elements()) != len(planTyped.Elements()) {
			return plan
		}
		elements := make([]attr.Value, 0, len(planTyped.Elements()))
		for i, item := range planTyped.Elements() {
			elements = append(elements, fillUnknowns(item, freshTyped.Elements()[i]))
		}
		merged, diags := types.ListValue(planTyped.ElementType(context.Background()), elements)
		if diags.HasError() {
			return plan
		}
		return merged

	case basetypes.MapValue:
		freshTyped, ok := fresh.(basetypes.MapValue)
		if !ok {
			return plan
		}
		elements := make(map[string]attr.Value, len(planTyped.Elements()))
		for key, item := range planTyped.Elements() {
			elements[key] = fillUnknowns(item, freshTyped.Elements()[key])
		}
		merged, diags := types.MapValue(planTyped.ElementType(context.Background()), elements)
		if diags.HasError() {
			return plan
		}
		return merged
	}

	return plan
}

func preserveMasked(reference, fresh attr.Value) attr.Value {
	if reference == nil || reference.IsNull() || reference.IsUnknown() || fresh == nil {
		return fresh
	}

	switch freshTyped := fresh.(type) {
	case basetypes.StringValue:
		if freshTyped.ValueString() != maskedSecret {
			return fresh
		}
		if referenceTyped, ok := reference.(basetypes.StringValue); ok {
			return referenceTyped
		}

	case basetypes.ObjectValue:
		referenceTyped, ok := reference.(basetypes.ObjectValue)
		if !ok {
			return fresh
		}
		attrTypes := freshTyped.AttributeTypes(context.Background())
		elements := make(map[string]attr.Value, len(attrTypes))
		for name, item := range freshTyped.Attributes() {
			elements[name] = preserveMasked(referenceTyped.Attributes()[name], item)
		}
		merged, diags := types.ObjectValue(attrTypes, elements)
		if diags.HasError() {
			return fresh
		}
		return merged

	case basetypes.ListValue:
		referenceTyped, ok := reference.(basetypes.ListValue)
		if !ok || len(referenceTyped.Elements()) != len(freshTyped.Elements()) {
			return fresh
		}
		elements := make([]attr.Value, 0, len(freshTyped.Elements()))
		for i, item := range freshTyped.Elements() {
			elements = append(elements, preserveMasked(referenceTyped.Elements()[i], item))
		}
		merged, diags := types.ListValue(freshTyped.ElementType(context.Background()), elements)
		if diags.HasError() {
			return fresh
		}
		return merged

	case basetypes.MapValue:
		referenceTyped, ok := reference.(basetypes.MapValue)
		if !ok {
			return fresh
		}
		elements := make(map[string]attr.Value, len(freshTyped.Elements()))
		for key, item := range freshTyped.Elements() {
			elements[key] = preserveMasked(referenceTyped.Elements()[key], item)
		}
		merged, diags := types.MapValue(freshTyped.ElementType(context.Background()), elements)
		if diags.HasError() {
			return fresh
		}
		return merged
	}

	return fresh
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
