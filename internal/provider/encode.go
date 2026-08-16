package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func serverSet(attribute schema.Attribute) bool {
	return attribute.IsComputed() && !attribute.IsOptional()
}

func nestedAttributes(attribute schema.Attribute) map[string]schema.Attribute {
	switch typed := attribute.(type) {
	case schema.SingleNestedAttribute:
		return typed.Attributes
	case schema.ListNestedAttribute:
		return typed.NestedObject.Attributes
	case schema.SetNestedAttribute:
		return typed.NestedObject.Attributes
	case schema.MapNestedAttribute:
		return typed.NestedObject.Attributes
	}

	return nil
}

func encodeAttribute(ctx context.Context, attribute schema.Attribute, value attr.Value) (any, error) {
	nested := nestedAttributes(attribute)
	if nested == nil {
		return toJMAP(ctx, value)
	}

	return encodeNested(ctx, nested, value)
}

func encodeNested(ctx context.Context, nested map[string]schema.Attribute, value attr.Value) (any, error) {
	if value == nil || value.IsNull() || value.IsUnknown() {
		return nil, nil
	}

	switch typed := value.(type) {
	case basetypes.ObjectValue:
		out := make(map[string]any, len(typed.Attributes()))
		for name, item := range typed.Attributes() {
			attribute, declared := nested[name]
			if declared && serverSet(attribute) {
				continue
			}

			converted, err := encodeChild(ctx, attribute, declared, item)
			if err != nil {
				return nil, err
			}
			if converted == nil {
				continue
			}

			out[objectFieldName(name)] = converted
		}
		return out, nil

	case basetypes.ListValue:
		out := make(map[string]any, len(typed.Elements()))
		for index, item := range typed.Elements() {
			converted, err := encodeNested(ctx, nested, item)
			if err != nil {
				return nil, err
			}
			out[strconv.Itoa(index)] = converted
		}
		return out, nil

	case basetypes.MapValue:
		out := make(map[string]any, len(typed.Elements()))
		for key, item := range typed.Elements() {
			converted, err := encodeNested(ctx, nested, item)
			if err != nil {
				return nil, err
			}
			out[key] = converted
		}
		return out, nil
	}

	return toJMAP(ctx, value)
}

func encodeChild(ctx context.Context, attribute schema.Attribute, declared bool, value attr.Value) (any, error) {
	if !declared {
		return toJMAP(ctx, value)
	}

	return encodeAttribute(ctx, attribute, value)
}
