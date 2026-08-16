package main

import (
	"fmt"
	"slices"
	"sort"
)

type attrNode struct {
	Kind        string
	Description string
	Sensitive   bool
	Required    bool
	ServerSet   bool
	Immutable   bool
	ElemKind    string
	EnumValues  []string
	TypeValues  []string
	Children    map[string]*attrNode
}

type resolver struct {
	file    schemaFile
	skipped []string
	stack   []string
}

func (r *resolver) note(context, reason string) {
	r.skipped = append(r.skipped, fmt.Sprintf("%s: %s", context, reason))
}

func (r *resolver) inStack(group string) bool {
	return slices.Contains(r.stack, group)
}

func (r *resolver) resolveGroup(groupKey string) map[string]*attrNode {
	group, ok := r.file.Fields[groupKey]
	if !ok {
		r.note(groupKey, "field group not found")
		return nil
	}
	if r.inStack(groupKey) {
		r.note(groupKey, "recursive field group")
		return nil
	}

	r.stack = append(r.stack, groupKey)
	defer func() { r.stack = r.stack[:len(r.stack)-1] }()

	out := make(map[string]*attrNode, len(group.Properties))
	for name, p := range group.Properties {
		node := r.resolveProperty(groupKey+"."+name, p)
		if node == nil {
			continue
		}
		out[name] = node
	}

	return out
}

func (r *resolver) resolveProperty(context string, p property) *attrNode {
	node := r.resolveType(context, p.Type)
	if node == nil {
		return nil
	}

	node.Description = p.Description
	if p.Enterprise {
		node.Description += " Requires an Enterprise license."
	}
	node.ServerSet = p.Update == "serverSet"
	node.Immutable = p.Update == "immutable"

	return node
}

func (r *resolver) resolveType(context string, t propertyType) *attrNode {
	switch t.Type {
	case "string", "objectId", "blobId", "utcDateTime":
		return &attrNode{Kind: "string", Sensitive: sensitiveFormat(t.Format)}

	case "enum":
		return &attrNode{Kind: "string", EnumValues: r.enumValues(t.EnumName)}

	case "boolean":
		return &attrNode{Kind: "bool"}

	case "number":
		if t.Format == "float" {
			return &attrNode{Kind: "float"}
		}
		return &attrNode{Kind: "int"}

	case "set":
		elem := "string"
		var values []string
		if t.Class != nil {
			if !scalar(t.Class.Type) {
				r.note(context, "set of non-scalar values")
				return nil
			}
			elem = scalarKind(*t.Class)
			values = r.enumValues(t.Class.EnumName)
		}
		return &attrNode{Kind: "set", ElemKind: elem, EnumValues: values, Sensitive: t.Class != nil && sensitiveFormat(t.Class.Format)}

	case "map":
		elem := "string"
		if t.Class != nil {
			if !scalar(t.Class.Type) {
				r.note(context, "map of non-scalar values")
				return nil
			}
			elem = scalarKind(*t.Class)
		}
		return &attrNode{Kind: "map", ElemKind: elem}

	case "object":
		return r.resolveObject(context, t.ObjectName)

	case "objectList":
		return r.resolveObjectList(context, t.ObjectName)
	}

	r.note(context, "unsupported type "+t.Type)

	return nil
}

func (r *resolver) resolveObject(context, objectName string) *attrNode {
	entry, ok := r.file.Schemas[objectName]
	if !ok {
		r.note(context, "schema entry not found for "+objectName)
		return nil
	}

	switch entry.Type {
	case "single":
		children := r.resolveGroup(entry.SchemaName)
		if children == nil {
			return nil
		}
		return &attrNode{Kind: "object", Children: children}

	case "multiple":
		return r.resolveUnion(context, entry.Variants)
	}

	r.note(context, "unsupported schema entry type "+entry.Type)

	return nil
}

func (r *resolver) resolveObjectList(context, objectName string) *attrNode {
	var element *attrNode

	if _, ok := r.file.Fields[objectName]; ok {
		children := r.resolveGroup(objectName)
		if children == nil {
			return nil
		}
		element = &attrNode{Kind: "object", Children: children}
	} else {
		element = r.resolveObject(context, objectName)
	}

	if element == nil {
		return nil
	}

	return &attrNode{Kind: "list", Children: element.Children, TypeValues: element.TypeValues}
}

func (r *resolver) resolveUnion(context string, variants []variant) *attrNode {
	node := &attrNode{Kind: "union", Children: map[string]*attrNode{}}

	for _, v := range variants {
		if deprecated(v.Name) {
			continue
		}
		node.TypeValues = append(node.TypeValues, v.Name)
		if v.SchemaName == "" {
			continue
		}
		children := r.resolveGroup(v.SchemaName)
		for _, name := range sortedKeys(children) {
			merged := r.merge(context+"."+name, node.Children[name], children[name])
			if merged != nil {
				node.Children[name] = merged
			}
		}
	}

	if len(node.TypeValues) == 0 {
		r.note(context, "union with no usable variants")
		return nil
	}

	return node
}

func (r *resolver) merge(context string, a, b *attrNode) *attrNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Kind != b.Kind {
		r.note(context, fmt.Sprintf("variant field kinds conflict (%s vs %s)", a.Kind, b.Kind))
		return preferred(a, b)
	}

	out := *a
	out.Sensitive = a.Sensitive || b.Sensitive
	out.EnumValues = unionStrings(a.EnumValues, b.EnumValues)
	out.TypeValues = unionStrings(a.TypeValues, b.TypeValues)
	if out.Description == "" {
		out.Description = b.Description
	}

	if a.Children != nil || b.Children != nil {
		out.Children = map[string]*attrNode{}
		for _, name := range sortedKeys(a.Children) {
			out.Children[name] = a.Children[name]
		}
		for _, name := range sortedKeys(b.Children) {
			out.Children[name] = r.merge(context+"."+name, out.Children[name], b.Children[name])
		}
	}

	return &out
}

func preferred(a, b *attrNode) *attrNode {
	if a.Kind == "union" || a.Kind == "object" {
		return a
	}
	if b.Kind == "union" || b.Kind == "object" {
		return b
	}

	return a
}

func (r *resolver) enumValues(name string) []string {
	if name == "" {
		return nil
	}

	values := make([]string, 0, len(r.file.Enums[name]))
	for _, v := range r.file.Enums[name] {
		values = append(values, v.Name)
	}

	return values
}

func scalar(kind string) bool {
	switch kind {
	case "string", "objectId", "blobId", "utcDateTime", "enum", "boolean", "number":
		return true
	}

	return false
}

func scalarKind(t propertyType) string {
	switch t.Type {
	case "boolean":
		return "bool"
	case "number":
		if t.Format == "float" {
			return "float"
		}
		return "int"
	}

	return "string"
}

func sensitiveFormat(format string) bool {
	return format == "secret" || format == "secretText"
}

func deprecated(name string) bool {
	return len(name) >= 10 && name[:10] == "Deprecated"
}

func unionStrings(a, b []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(a)+len(b))
	for _, v := range append(append([]string{}, a...), b...) {
		if seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}

	return out
}

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
