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

type propertyType struct {
	Type       string        `json:"type"`
	Format     string        `json:"format"`
	Nullable   bool          `json:"nullable"`
	ObjectName string        `json:"objectName"`
	EnumName   string        `json:"enumName"`
	Class      *propertyType `json:"class"`
}

type property struct {
	Description string       `json:"description"`
	Type        propertyType `json:"type"`
	Update      string       `json:"update"`
}

type fieldGroup struct {
	Properties map[string]property `json:"properties"`
	Defaults   map[string]any      `json:"defaults"`
}

type variant struct {
	Name       string `json:"name"`
	Label      string `json:"label"`
	SchemaName string `json:"schemaName"`
}

type schemaEntry struct {
	Type     string    `json:"type"`
	Variants []variant `json:"variants"`
}

type enumValue struct {
	Name string `json:"name"`
}

type schemaFile struct {
	Fields  map[string]fieldGroup  `json:"fields"`
	Schemas map[string]schemaEntry `json:"schemas"`
	Enums   map[string][]enumValue `json:"enums"`
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
	{Name: "domain", JMAPType: "x:Domain", FieldKey: "x:Domain", Required: []string{"name"}},
	{Name: "mailing_list", JMAPType: "x:MailingList", FieldKey: "x:MailingList", Required: []string{"name", "emailAddress"}},
	{Name: "tenant", JMAPType: "x:Tenant", FieldKey: "x:Tenant", Required: []string{"name"}},
	{Name: "user", JMAPType: "x:Account", FieldKey: "x:UserAccount", Variant: "User", Required: []string{"name"}},
	{Name: "group", JMAPType: "x:Account", FieldKey: "x:GroupAccount", Variant: "Group", Required: []string{"name"}},
}

var file schemaFile

var imports = map[string]string{}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: generator <schema.json> <output.go>")
		os.Exit(1)
	}

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		fail(err)
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		fail(err)
	}

	skipped := map[string][]string{}

	var body bytes.Buffer
	body.WriteString("var generatedResources = []resourceDescriptor{\n")
	for _, t := range targets {
		group, ok := file.Fields[t.FieldKey]
		if !ok {
			fail(fmt.Errorf("field group %q not found", t.FieldKey))
		}
		emitDescriptor(&body, t, group, skipped)
	}
	body.WriteString("}\n")

	var out bytes.Buffer
	out.WriteString("package provider\n\nimport (\n")
	paths := make([]string, 0, len(imports))
	for path := range imports {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Fprintf(&out, "\t%q\n", path)
	}
	out.WriteString(")\n\n")
	out.Write(body.Bytes())

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		fail(fmt.Errorf("formatting generated source: %w", err))
	}
	if err := os.WriteFile(os.Args[2], formatted, 0o644); err != nil {
		fail(err)
	}

	for _, t := range targets {
		if fields := skipped[t.Name]; len(fields) > 0 {
			fmt.Printf("skipped in %s: %s\n", t.Name, strings.Join(fields, ", "))
		}
	}
	fmt.Printf("generated %d resources into %s\n", len(targets), os.Args[2])
}

func need(path string) {
	imports[path] = ""
}

func emitDescriptor(b *bytes.Buffer, t target, group fieldGroup, skipped map[string][]string) {
	need("github.com/hashicorp/terraform-plugin-framework/resource/schema")
	need("github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier")
	need("github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier")

	fmt.Fprintf(b, "\t{\n\t\tName: %q,\n\t\tJMAPType: %q,\n", t.Name, t.JMAPType)
	if t.Variant != "" {
		fmt.Fprintf(b, "\t\tVariant: %q,\n", t.Variant)
	}
	fmt.Fprintf(b, "\t\tSchema: schema.Schema{\n")
	fmt.Fprintf(b, "\t\t\tMarkdownDescription: %q,\n",
		fmt.Sprintf("Manages a Stalwart %s (the `%s` object).", strings.ReplaceAll(t.Name, "_", " "), t.JMAPType))
	fmt.Fprintf(b, "\t\t\tAttributes: map[string]schema.Attribute{\n")

	fmt.Fprintf(b, "\t\t\t\t\"id\": schema.StringAttribute{\n")
	fmt.Fprintf(b, "\t\t\t\t\tMarkdownDescription: \"Server-assigned identifier.\",\n")
	fmt.Fprintf(b, "\t\t\t\t\tComputed: true,\n")
	fmt.Fprintf(b, "\t\t\t\t\tPlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},\n")
	fmt.Fprintf(b, "\t\t\t\t},\n")

	for _, name := range sortedKeys(group.Properties) {
		p := group.Properties[name]
		rendered, ok := renderAttribute(p, contains(t.Required, name), group.Defaults[name], "\t\t\t\t")
		if !ok {
			skipped[t.Name] = append(skipped[t.Name], name)
			continue
		}
		fmt.Fprintf(b, "\t\t\t\t%q: %s,\n", terraformName(name), rendered)
	}

	fmt.Fprintf(b, "\t\t\t},\n\t\t},\n\t},\n")
}

func renderAttribute(p property, required bool, def any, indent string) (string, bool) {
	inner := indent + "\t"

	kind, ok := attributeKind(p)
	if !ok {
		return "", false
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", kind)
	fmt.Fprintf(&b, "%sMarkdownDescription: %q,\n", inner, p.Description)

	switch {
	case p.Update == "serverSet":
		fmt.Fprintf(&b, "%sComputed: true,\n", inner)
	case required:
		fmt.Fprintf(&b, "%sRequired: true,\n", inner)
	default:
		fmt.Fprintf(&b, "%sOptional: true,\n", inner)
	}

	if p.Update == "immutable" {
		if modifier := replaceModifier(kind); modifier != "" {
			fmt.Fprintf(&b, "%s%s\n", inner, modifier)
		}
	}

	if elem := elementType(p); elem != "" {
		need("github.com/hashicorp/terraform-plugin-framework/types")
		fmt.Fprintf(&b, "%sElementType: %s,\n", inner, elem)
	}

	if values := enumValues(p); len(values) > 0 {
		need("github.com/hashicorp/terraform-plugin-framework/schema/validator")
		need("github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator")
		quoted := make([]string, 0, len(values))
		for _, v := range values {
			quoted = append(quoted, fmt.Sprintf("%q", v))
		}
		fmt.Fprintf(&b, "%sValidators: []validator.String{stringvalidator.OneOf(%s)},\n", inner, strings.Join(quoted, ", "))
	}

	if literal := defaultLiteral(p, def); literal != "" {
		fmt.Fprintf(&b, "%sComputed: true,\n%sDefault: %s,\n", inner, inner, literal)
	}

	if nested := renderNested(p, inner); nested != "" {
		b.WriteString(nested)
	}

	fmt.Fprintf(&b, "%s}", indent)

	return b.String(), true
}

func renderNested(p property, indent string) string {
	inner := indent + "\t"

	switch p.Type.Type {
	case "object":
		entry, ok := file.Schemas[p.Type.ObjectName]
		if !ok || entry.Type != "multiple" {
			return ""
		}

		var b strings.Builder
		fmt.Fprintf(&b, "%sAttributes: map[string]schema.Attribute{\n", indent)

		need("github.com/hashicorp/terraform-plugin-framework/schema/validator")
		need("github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator")
		names := make([]string, 0, len(entry.Variants))
		for _, v := range entry.Variants {
			names = append(names, fmt.Sprintf("%q", v.Name))
		}
		fmt.Fprintf(&b, "%s\"type\": schema.StringAttribute{\n", inner)
		fmt.Fprintf(&b, "%s\tMarkdownDescription: %q,\n", inner, "Variant discriminator.")
		fmt.Fprintf(&b, "%s\tRequired: true,\n", inner)
		fmt.Fprintf(&b, "%s\tValidators: []validator.String{stringvalidator.OneOf(%s)},\n", inner, strings.Join(names, ", "))
		fmt.Fprintf(&b, "%s},\n", inner)

		merged := map[string]property{}
		for _, v := range entry.Variants {
			if v.SchemaName == "" {
				continue
			}
			for name, vp := range file.Fields[v.SchemaName].Properties {
				merged[name] = vp
			}
		}
		for _, name := range sortedKeys(merged) {
			rendered, ok := renderAttribute(merged[name], false, nil, inner)
			if !ok {
				continue
			}
			fmt.Fprintf(&b, "%s%q: %s,\n", inner, terraformName(name), rendered)
		}

		fmt.Fprintf(&b, "%s},\n", indent)

		return b.String()

	case "objectList":
		group, ok := file.Fields[p.Type.ObjectName]
		if !ok || len(group.Properties) == 0 {
			return ""
		}

		var b strings.Builder
		fmt.Fprintf(&b, "%sNestedObject: schema.NestedAttributeObject{\n", indent)
		fmt.Fprintf(&b, "%s\tAttributes: map[string]schema.Attribute{\n", indent)
		for _, name := range sortedKeys(group.Properties) {
			rendered, ok := renderAttribute(group.Properties[name], false, nil, inner+"\t")
			if !ok {
				continue
			}
			fmt.Fprintf(&b, "%s\t\t%q: %s,\n", indent, terraformName(name), rendered)
		}
		fmt.Fprintf(&b, "%s\t},\n%s},\n", indent, indent)

		return b.String()
	}

	return ""
}

func attributeKind(p property) (string, bool) {
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
	case "object":
		if entry, ok := file.Schemas[p.Type.ObjectName]; ok && entry.Type == "multiple" {
			return "schema.SingleNestedAttribute", true
		}
	case "objectList":
		if group, ok := file.Fields[p.Type.ObjectName]; ok && len(group.Properties) > 0 {
			return "schema.ListNestedAttribute", true
		}
	}

	return "", false
}

func replaceModifier(kind string) string {
	switch kind {
	case "schema.StringAttribute":
		need("github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier")
		return "PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},"
	case "schema.BoolAttribute":
		need("github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier")
		return "PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},"
	case "schema.Int64Attribute":
		need("github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier")
		return "PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},"
	}

	return ""
}

func defaultLiteral(p property, def any) string {
	if def == nil {
		return ""
	}

	switch p.Type.Type {
	case "boolean":
		b, ok := def.(bool)
		if !ok {
			return ""
		}
		need("github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault")
		return fmt.Sprintf("booldefault.StaticBool(%t)", b)
	case "string":
		s, ok := def.(string)
		if !ok {
			return ""
		}
		need("github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault")
		return fmt.Sprintf("stringdefault.StaticString(%q)", s)
	}

	return ""
}

func enumValues(p property) []string {
	name := p.Type.EnumName
	if name == "" {
		return nil
	}

	values := make([]string, 0)
	for _, v := range file.Enums[name] {
		values = append(values, v.Name)
	}

	return values
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

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
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
