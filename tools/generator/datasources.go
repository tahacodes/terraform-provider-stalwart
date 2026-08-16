package main

import (
	"bytes"
	"fmt"
	"go/format"
	"sort"
	"strings"
)

type dataSourceTarget struct {
	Name        string
	JMAPType    string
	Variant     string
	Singleton   bool
	Plural      bool
	HasName     bool
	Description string
	Attributes  map[string]*attrNode
	TypeValues  []string
}

var runtimeDataSources = []string{
	"x:ApiKey",
	"x:AppPassword",
	"x:ArchivedItem",
	"x:ArfExternalReport",
	"x:ClusterNode",
	"x:DmarcExternalReport",
	"x:DmarcInternalReport",
	"x:Log",
	"x:Metric",
	"x:QueuedMessage",
	"x:SpamTrainingSample",
	"x:Task",
	"x:TlsExternalReport",
	"x:TlsInternalReport",
}

var pluralNameOverrides = map[string]string{
	"x:Metric":      "metric_data_points",
	"x:OAuthClient": "oauth_clients",
}

func buildDataSourceTargets(r *resolver, resources []resourceTarget) []dataSourceTarget {
	var targets []dataSourceTarget

	for _, t := range resources {
		_, hasName := t.Attributes["name"]
		description := strings.Replace(t.Description, "Manages the", "Reads the", 1)
		targets = append(targets, dataSourceTarget{
			Name:        t.ResourceName,
			JMAPType:    t.JMAPType,
			Variant:     t.Variant,
			Singleton:   t.Singleton,
			HasName:     !t.Singleton && hasName,
			Description: description,
			Attributes:  t.Attributes,
			TypeValues:  t.TypeValues,
		})
	}

	for _, objectName := range runtimeDataSources {
		object, ok := r.file.Objects[objectName]
		if !ok {
			fail(fmt.Errorf("runtime data source %q not found", objectName))
		}
		entry := r.file.Schemas[objectName]

		var attributes map[string]*attrNode
		var typeValues []string
		if entry.Type == "single" {
			attributes = r.resolveGroup(entry.SchemaName)
		} else {
			union := r.resolveUnion(objectName, entry.Variants)
			if union != nil {
				attributes = union.Children
				typeValues = union.TypeValues
			}
		}
		if attributes == nil {
			continue
		}

		_, hasName := attributes["name"]
		targets = append(targets, dataSourceTarget{
			Name:        resourceName(strings.TrimPrefix(objectName, "x:")),
			JMAPType:    objectName,
			HasName:     hasName,
			Description: object.Description,
			Attributes:  attributes,
			TypeValues:  typeValues,
		})
	}

	targets = append(targets, buildPluralTargets(targets)...)

	sort.Slice(targets, func(i, j int) bool { return targets[i].Name < targets[j].Name })
	verifyUniqueNames(targets)

	return targets
}

func buildPluralTargets(singular []dataSourceTarget) []dataSourceTarget {
	seen := map[string]bool{}

	var targets []dataSourceTarget
	for _, t := range singular {
		if t.Singleton || t.Plural || seen[t.JMAPType] {
			continue
		}
		seen[t.JMAPType] = true

		name := pluralNameOverrides[t.JMAPType]
		if name == "" {
			name = pluralName(resourceName(strings.TrimPrefix(t.JMAPType, "x:")))
		}

		targets = append(targets, dataSourceTarget{
			Name:        name,
			JMAPType:    t.JMAPType,
			Plural:      true,
			Description: fmt.Sprintf("Lists the identifiers of all `%s` objects on the server.", t.JMAPType),
		})
	}

	return targets
}

func verifyUniqueNames(targets []dataSourceTarget) {
	seen := map[string]bool{}
	for _, t := range targets {
		if seen[t.Name] {
			fail(fmt.Errorf("duplicate data source name %q", t.Name))
		}
		seen[t.Name] = true
	}
}

type dataSourceEmitter struct {
	imports map[string]bool
	body    bytes.Buffer
}

func emitDataSources(targets []dataSourceTarget) ([]byte, error) {
	e := &dataSourceEmitter{imports: map[string]bool{}}
	e.need("github.com/hashicorp/terraform-plugin-framework/datasource/schema")

	e.body.WriteString("var generatedDataSources = []dataSourceDescriptor{\n")
	for _, t := range targets {
		e.emitTarget(t)
	}
	e.body.WriteString("}\n")

	var out bytes.Buffer
	out.WriteString("package provider\n\nimport (\n")
	for _, path := range sortedKeys(e.imports) {
		fmt.Fprintf(&out, "\t%q\n", path)
	}
	out.WriteString(")\n\n")
	out.Write(e.body.Bytes())

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return nil, fmt.Errorf("formatting generated data sources: %w", err)
	}

	return formatted, nil
}

func (e *dataSourceEmitter) need(path string) {
	e.imports[path] = true
}

func (e *dataSourceEmitter) emitTarget(t dataSourceTarget) {
	fmt.Fprintf(&e.body, "\t{\n\t\tName: %q,\n\t\tJMAPType: %q,\n", t.Name, t.JMAPType)
	if t.Variant != "" {
		fmt.Fprintf(&e.body, "\t\tVariant: %q,\n", t.Variant)
	}
	if t.Singleton {
		fmt.Fprintf(&e.body, "\t\tSingleton: true,\n")
	}
	if t.Plural {
		fmt.Fprintf(&e.body, "\t\tPlural: true,\n")
	}
	if t.HasName {
		fmt.Fprintf(&e.body, "\t\tHasName: true,\n")
	}
	fmt.Fprintf(&e.body, "\t\tSchema: schema.Schema{\n")
	fmt.Fprintf(&e.body, "\t\t\tMarkdownDescription: %q,\n", t.Description)
	fmt.Fprintf(&e.body, "\t\t\tAttributes: map[string]schema.Attribute{\n")

	if t.Plural {
		e.need("github.com/hashicorp/terraform-plugin-framework/types")
		fmt.Fprintf(&e.body, "\t\t\t\t\"ids\": schema.SetAttribute{\n")
		fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: \"Identifiers of all objects of this type.\",\n")
		fmt.Fprintf(&e.body, "\t\t\t\t\tComputed: true,\n")
		fmt.Fprintf(&e.body, "\t\t\t\t\tElementType: types.StringType,\n")
		fmt.Fprintf(&e.body, "\t\t\t\t},\n")
	} else {
		e.emitID(t)
		if len(t.TypeValues) > 0 {
			fmt.Fprintf(&e.body, "\t\t\t\t\"type\": schema.StringAttribute{\n")
			fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: \"Variant discriminator.\",\n")
			fmt.Fprintf(&e.body, "\t\t\t\t\tComputed: true,\n")
			fmt.Fprintf(&e.body, "\t\t\t\t},\n")
		}
		for _, name := range sortedKeys(t.Attributes) {
			lookup := t.HasName && name == "name"
			fmt.Fprintf(&e.body, "\t\t\t\t%q: %s,\n", rootTerraformName(name), e.render(t.Attributes[name], lookup, "\t\t\t\t"))
		}
	}

	fmt.Fprintf(&e.body, "\t\t\t},\n\t\t},\n\t},\n")
}

func (e *dataSourceEmitter) emitID(t dataSourceTarget) {
	fmt.Fprintf(&e.body, "\t\t\t\t\"id\": schema.StringAttribute{\n")
	switch {
	case t.Singleton:
		fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: \"Always `singleton`.\",\n")
		fmt.Fprintf(&e.body, "\t\t\t\t\tComputed: true,\n")
	case t.HasName:
		fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: \"Identifier of the object. Set either `id` or `name`.\",\n")
		fmt.Fprintf(&e.body, "\t\t\t\t\tOptional: true,\n\t\t\t\t\tComputed: true,\n")
	default:
		fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: \"Identifier of the object.\",\n")
		fmt.Fprintf(&e.body, "\t\t\t\t\tRequired: true,\n")
	}
	fmt.Fprintf(&e.body, "\t\t\t\t},\n")
}

func (e *dataSourceEmitter) render(node *attrNode, lookup bool, indent string) string {
	inner := indent + "\t"
	kind := attributeKind(node)

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", kind)
	fmt.Fprintf(&b, "%sMarkdownDescription: %q,\n", inner, node.Description)
	if lookup {
		fmt.Fprintf(&b, "%sOptional: true,\n", inner)
	}
	fmt.Fprintf(&b, "%sComputed: true,\n", inner)
	if node.Sensitive {
		fmt.Fprintf(&b, "%sSensitive: true,\n", inner)
	}

	if node.Kind == "set" || node.Kind == "map" {
		e.need("github.com/hashicorp/terraform-plugin-framework/types")
		fmt.Fprintf(&b, "%sElementType: %s,\n", inner, elementType(node.ElemKind))
	}

	switch node.Kind {
	case "object", "union":
		fmt.Fprintf(&b, "%sAttributes: map[string]schema.Attribute{\n", inner)
		if node.Kind == "union" {
			fmt.Fprintf(&b, "%s\t\"type\": schema.StringAttribute{\n", inner)
			fmt.Fprintf(&b, "%s\t\tMarkdownDescription: \"Variant discriminator.\",\n", inner)
			fmt.Fprintf(&b, "%s\t\tComputed: true,\n", inner)
			fmt.Fprintf(&b, "%s\t},\n", inner)
		}
		for _, name := range sortedKeys(node.Children) {
			fmt.Fprintf(&b, "%s\t%q: %s,\n", inner, terraformName(name), e.render(node.Children[name], false, inner+"\t"))
		}
		fmt.Fprintf(&b, "%s},\n", inner)

	case "list":
		fmt.Fprintf(&b, "%sNestedObject: schema.NestedAttributeObject{\n", inner)
		fmt.Fprintf(&b, "%s\tAttributes: map[string]schema.Attribute{\n", inner)
		if len(node.TypeValues) > 0 {
			fmt.Fprintf(&b, "%s\t\t\"type\": schema.StringAttribute{\n", inner)
			fmt.Fprintf(&b, "%s\t\t\tMarkdownDescription: \"Variant discriminator.\",\n", inner)
			fmt.Fprintf(&b, "%s\t\t\tComputed: true,\n", inner)
			fmt.Fprintf(&b, "%s\t\t},\n", inner)
		}
		for _, name := range sortedKeys(node.Children) {
			fmt.Fprintf(&b, "%s\t\t%q: %s,\n", inner, terraformName(name), e.render(node.Children[name], false, inner+"\t\t"))
		}
		fmt.Fprintf(&b, "%s\t},\n%s},\n", inner, inner)
	}

	fmt.Fprintf(&b, "%s}", indent)

	return b.String()
}
