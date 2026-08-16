package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
	"unicode"
)

type property struct {
	Description string          `json:"description"`
	Type        propertyType    `json:"type"`
	Update      string          `json:"update"`
	Enterprise  bool            `json:"enterprise"`
	Raw         json.RawMessage `json:"-"`
}

type propertyType struct {
	Type       string        `json:"type"`
	Format     string        `json:"format"`
	Nullable   bool          `json:"nullable"`
	ObjectName string        `json:"objectName"`
	Class      *propertyType `json:"class"`
}

type fieldGroup struct {
	Properties map[string]property `json:"properties"`
}

type schemaFile struct {
	Fields map[string]fieldGroup `json:"fields"`
}

type target struct {
	Name     string
	JMAPType string
	FieldKey string
	Variant  string
	Required []string
}

var targets = []target{
	{Name: "role", JMAPType: "x:Role", FieldKey: "x:Role", Required: []string{"name"}},
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: generator <schema.json> <output.go>")
		os.Exit(1)
	}

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		fail(err)
	}

	var parsed schemaFile
	if err := json.Unmarshal(raw, &parsed); err != nil {
		fail(err)
	}

	var b bytes.Buffer
	b.WriteString("package provider\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"github.com/hashicorp/terraform-plugin-framework/resource/schema\"\n")
	b.WriteString("\t\"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier\"\n")
	b.WriteString("\t\"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier\"\n")
	b.WriteString("\t\"github.com/hashicorp/terraform-plugin-framework/types\"\n")
	b.WriteString(")\n\n")
	b.WriteString("var generatedResources = []resourceDescriptor{\n")

	skipped := map[string][]string{}

	for _, t := range targets {
		group, ok := parsed.Fields[t.FieldKey]
		if !ok {
			fail(fmt.Errorf("field group %q not found", t.FieldKey))
		}
		emitDescriptor(&b, t, group, skipped)
	}

	b.WriteString("}\n")

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		fail(fmt.Errorf("formatting generated source: %w", err))
	}

	if err := os.WriteFile(os.Args[2], formatted, 0o644); err != nil {
		fail(err)
	}

	for name, fields := range skipped {
		fmt.Printf("skipped in %s: %s\n", name, strings.Join(fields, ", "))
	}
	fmt.Printf("generated %d resources into %s\n", len(targets), os.Args[2])
}

func emitDescriptor(b *bytes.Buffer, t target, group fieldGroup, skipped map[string][]string) {
	fmt.Fprintf(b, "\t{\n\t\tName: %q,\n\t\tJMAPType: %q,\n", t.Name, t.JMAPType)
	if t.Variant != "" {
		fmt.Fprintf(b, "\t\tVariant: %q,\n", t.Variant)
	}
	fmt.Fprintf(b, "\t\tSchema: schema.Schema{\n")
	fmt.Fprintf(b, "\t\t\tMarkdownDescription: %q,\n", fmt.Sprintf("Manages a Stalwart %s (the `%s` object).", t.Name, t.JMAPType))
	fmt.Fprintf(b, "\t\t\tAttributes: map[string]schema.Attribute{\n")

	fmt.Fprintf(b, "\t\t\t\t\"id\": schema.StringAttribute{\n")
	fmt.Fprintf(b, "\t\t\t\t\tMarkdownDescription: %q,\n", "Server-assigned identifier.")
	fmt.Fprintf(b, "\t\t\t\t\tComputed: true,\n")
	fmt.Fprintf(b, "\t\t\t\t\tPlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},\n")
	fmt.Fprintf(b, "\t\t\t\t},\n")

	names := make([]string, 0, len(group.Properties))
	for name := range group.Properties {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		p := group.Properties[name]
		attribute, ok := attributeFor(p)
		if !ok {
			skipped[t.Name] = append(skipped[t.Name], name)
			continue
		}

		fmt.Fprintf(b, "\t\t\t\t%q: %s{\n", terraformName(name), attribute)
		fmt.Fprintf(b, "\t\t\t\t\tMarkdownDescription: %q,\n", p.Description)

		switch {
		case p.Update == "serverSet":
			fmt.Fprintf(b, "\t\t\t\t\tComputed: true,\n")
		case contains(t.Required, name):
			fmt.Fprintf(b, "\t\t\t\t\tRequired: true,\n")
		default:
			fmt.Fprintf(b, "\t\t\t\t\tOptional: true,\n")
		}

		if elem := elementType(p); elem != "" {
			fmt.Fprintf(b, "\t\t\t\t\tElementType: %s,\n", elem)
		}

		fmt.Fprintf(b, "\t\t\t\t},\n")
	}

	fmt.Fprintf(b, "\t\t\t},\n\t\t},\n\t},\n")
}

func attributeFor(p property) (string, bool) {
	switch p.Type.Type {
	case "string", "objectId", "blobId", "utcDateTime", "enum":
		return "schema.StringAttribute", true
	case "boolean":
		return "schema.BoolAttribute", true
	case "number":
		if p.Type.Format == "float" {
			return "schema.Float64Attribute", true
		}
		return "schema.Int64Attribute", true
	case "set":
		if p.Type.Class != nil && scalar(p.Type.Class.Type) {
			return "schema.SetAttribute", true
		}
	case "map":
		if p.Type.Class == nil || scalar(p.Type.Class.Type) {
			return "schema.MapAttribute", true
		}
	}

	return "", false
}

func elementType(p property) string {
	switch p.Type.Type {
	case "set", "map":
		if p.Type.Class == nil {
			return "types.StringType"
		}
		switch p.Type.Class.Type {
		case "boolean":
			return "types.BoolType"
		case "number":
			if p.Type.Class.Format == "float" {
				return "types.Float64Type"
			}
			return "types.Int64Type"
		default:
			return "types.StringType"
		}
	}

	return ""
}

func scalar(kind string) bool {
	switch kind {
	case "string", "objectId", "blobId", "utcDateTime", "enum", "boolean", "number":
		return true
	}

	return false
}

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func terraformName(jmap string) string {
	var b strings.Builder

	for i, r := range jmap {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}

	return b.String()
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "generator:", err)
	os.Exit(1)
}
