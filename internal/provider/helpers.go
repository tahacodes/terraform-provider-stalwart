package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func pathRoot(name string) path.Path {
	return path.Root(name)
}

func boolField(object map[string]any, key string) bool {
	v, _ := object[key].(bool)

	return v
}

func stringMap(ctx context.Context, value any) (types.Map, diag.Diagnostics) {
	raw, ok := value.(map[string]any)
	if !ok || len(raw) == 0 {
		return types.MapNull(types.StringType), nil
	}

	elements := make(map[string]string, len(raw))
	for k, v := range raw {
		s, ok := v.(string)
		if !ok {
			continue
		}
		elements[k] = s
	}

	return types.MapValueFrom(ctx, types.StringType, elements)
}
