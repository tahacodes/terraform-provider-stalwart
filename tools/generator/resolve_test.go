package main

import "testing"

func testFile() schemaFile {
	return schemaFile{
		Fields: map[string]fieldGroup{
			"x:First": {Properties: map[string]property{
				"shared": {Type: propertyType{Type: "string"}},
				"only":   {Type: propertyType{Type: "boolean"}},
			}},
			"x:Second": {Properties: map[string]property{
				"shared": {Type: propertyType{Type: "string"}},
				"other":  {Type: propertyType{Type: "number"}},
			}},
			"x:Loop": {Properties: map[string]property{
				"nested": {Type: propertyType{Type: "object", ObjectName: "x:Loop"}},
				"value":  {Type: propertyType{Type: "string"}},
			}},
			"x:Secret": {Properties: map[string]property{
				"secret": {Type: propertyType{Type: "string", Format: "secret"}},
			}},
		},
		Schemas: map[string]schemaEntry{
			"x:Loop": {Type: "single", SchemaName: "x:Loop"},
		},
	}
}

func TestResolveUnionMergesVariants(t *testing.T) {
	t.Parallel()

	r := &resolver{file: testFile()}
	node := r.resolveUnion("test", []variant{
		{Name: "A", SchemaName: "x:First"},
		{Name: "B", SchemaName: "x:Second"},
		{Name: "C"},
		{Name: "Deprecated1"},
	})

	if node == nil {
		t.Fatal("expected union node")
	}
	if len(node.TypeValues) != 3 {
		t.Fatalf("expected variants A, B and C, got %v", node.TypeValues)
	}
	if len(node.Children) != 3 {
		t.Fatalf("expected merged fields shared, only and other, got %v", sortedKeys(node.Children))
	}
	if node.Children["shared"].Kind != "string" {
		t.Fatalf("expected shared field to stay a string, got %q", node.Children["shared"].Kind)
	}
}

func TestResolveGroupStopsOnRecursion(t *testing.T) {
	t.Parallel()

	r := &resolver{file: testFile()}
	children := r.resolveGroup("x:Loop")

	if children == nil {
		t.Fatal("expected resolved group")
	}
	if _, ok := children["nested"]; ok {
		t.Fatal("expected recursive property to be dropped")
	}
	if _, ok := children["value"]; !ok {
		t.Fatal("expected scalar property to survive")
	}
	if len(r.skipped) == 0 {
		t.Fatal("expected recursion to be reported")
	}
}

func TestResolveMarksSecretsSensitive(t *testing.T) {
	t.Parallel()

	r := &resolver{file: testFile()}
	children := r.resolveGroup("x:Secret")

	if children == nil || children["secret"] == nil {
		t.Fatal("expected resolved secret field")
	}
	if !children["secret"].Sensitive {
		t.Fatal("expected secret format to be sensitive")
	}
}

func TestMergePrefersStructuredKindOnConflict(t *testing.T) {
	t.Parallel()

	r := &resolver{file: testFile()}
	union := &attrNode{Kind: "union", TypeValues: []string{"Value"}}
	plain := &attrNode{Kind: "string"}

	if got := r.merge("test", plain, union); got.Kind != "union" {
		t.Fatalf("expected union kind to win, got %q", got.Kind)
	}
	if len(r.skipped) == 0 {
		t.Fatal("expected conflict to be reported")
	}
}
