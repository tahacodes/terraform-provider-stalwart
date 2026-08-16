package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
)

type resourceTarget struct {
	ResourceName string
	JMAPType     string
	Variant      string
	Singleton    bool
	ReloadAction string
	Description  string
	Required     map[string]bool
	Attributes   map[string]*attrNode
	TypeValues   []string
}

var excludedObjects = map[string]string{
	"x:Action":              "imperative server actions, not declarative state",
	"x:ApiKey":              "self-service credential view, creation not supported by the server",
	"x:AppPassword":         "self-service credential view, creation not supported by the server",
	"x:ArchivedItem":        "runtime data",
	"x:ArfExternalReport":   "runtime data",
	"x:AccountPassword":     "self-service credential change",
	"x:Bootstrap":           "one-time setup wizard",
	"x:ClusterNode":         "runtime data, creation forbidden by the server",
	"x:DmarcExternalReport": "runtime data",
	"x:DmarcInternalReport": "runtime data",
	"x:Log":                 "runtime data",
	"x:Metric":              "runtime data",
	"x:QueuedMessage":       "runtime data",
	"x:SpamTrainingSample":  "runtime data",
	"x:Task":                "runtime data",
	"x:TlsExternalReport":   "runtime data",
	"x:TlsInternalReport":   "runtime data",
	"x:Trace":               "runtime data",
}

var resourceNameOverrides = map[string]string{
	"x:Account/User":  "user",
	"x:Account/Group": "group",
	"x:OAuthClient":   "oauth_client",
}

var requiredOverrides = map[string][]string{
	"acme_provider":               {"contact"},
	"ai_model":                    {"model", "name", "url"},
	"allowed_ip":                  {"address"},
	"application":                 {"description", "resourceUrl", "urlPrefix"},
	"blocked_ip":                  {"address"},
	"certificate":                 {"certificate", "privateKey"},
	"directory_ldap":              {"baseDn", "description"},
	"directory_oidc":              {"description", "issuerUrl"},
	"directory_sql":               {"description"},
	"dkim_signature_dkim1_ed25519_sha256": {"privateKey"},
	"dkim_signature_dkim1_rsa_sha256":     {"privateKey"},
	"dkim_signature_dkim2_ed25519_sha256": {"privateKey"},
	"dkim_signature_dkim2_rsa_sha256":     {"privateKey"},
	"dns_server_alidns":           {"accessKey", "description", "secretKey"},
	"dns_server_arvan_cloud":      {"description", "secret"},
	"dns_server_azure_dns":        {"clientId", "clientSecret", "description", "resourceGroup", "subscriptionId", "tenantId"},
	"dns_server_bunny":            {"description", "secret"},
	"dns_server_cloudflare":       {"description", "secret"},
	"dns_server_cpanel":           {"baseUrl", "description", "token", "username"},
	"dns_server_de_sec":           {"description", "secret"},
	"dns_server_digital_ocean":    {"description", "secret"},
	"dns_server_dnsimple":         {"accountIdentifier", "description", "secret"},
	"dns_server_domeneshop":       {"authToken", "description", "secret"},
	"dns_server_dreamhost":        {"description", "secret"},
	"dns_server_gcore":            {"description", "secret"},
	"dns_server_glesys":           {"apiKey", "apiUser", "description"},
	"dns_server_google_cloud_dns": {"description", "projectId", "serviceAccountJson"},
	"dns_server_hosting_de":       {"description", "secret"},
	"dns_server_ibm_cloud":        {"apiKey", "description", "username"},
	"dns_server_infomaniak":       {"description", "secret"},
	"dns_server_inwx":             {"description", "password", "username"},
	"dns_server_ionos":            {"description", "secret"},
	"dns_server_lightsail":        {"accessKeyId", "description", "secretAccessKey"},
	"dns_server_linode":           {"description", "secret"},
	"dns_server_lua_dns":          {"authToken", "description", "username"},
	"dns_server_netcup":           {"apiKey", "customerNumber", "description", "password"},
	"dns_server_netlify":          {"description", "secret"},
	"dns_server_ns1":              {"description", "secret"},
	"dns_server_ovh":              {"applicationKey", "applicationSecret", "description"},
	"dns_server_plesk":            {"apiKey", "baseUrl", "description"},
	"dns_server_porkbun":          {"apiKey", "description", "secret"},
	"dns_server_route53":          {"accessKeyId", "description", "secretAccessKey"},
	"dns_server_safedns":          {"description", "secret"},
	"dns_server_scaleway":         {"description", "secret"},
	"dns_server_spaceship":        {"apiKey", "description", "secret"},
	"dns_server_tencent_cloud":    {"description", "secretId", "secretKey"},
	"dns_server_tsig":             {"description", "host", "key", "keyName"},
	"dns_server_ultra_dns":        {"description", "password", "username"},
	"dns_server_vercel":           {"authToken", "description"},
	"dns_server_vultr":            {"description", "secret"},
	"group":                       {"domainId", "name"},
	"http_lookup":                 {"namespace", "url"},
	"mailing_list":                {"domainId", "name"},
	"memory_lookup_key":           {"key", "namespace"},
	"memory_lookup_key_value":     {"key", "namespace", "value"},
	"mta_delivery_schedule":       {"name", "queueId"},
	"mta_hook":                    {"url"},
	"mta_inbound_throttle":        {"description", "rate"},
	"mta_milter":                  {"hostname"},
	"mta_outbound_throttle":       {"description", "rate"},
	"mta_queue_quota":             {"key"},
	"mta_route_relay":             {"address", "name"},
	"network_listener":            {"bind", "name"},
	"oauth_client":                {"clientId"},
	"public_key":                  {"key"},
	"role":                        {"description"},
	"sieve_system_script":         {"contents", "name"},
	"sieve_user_script":           {"contents", "name"},
	"spam_dnsbl_server_any":       {"name", "zone"},
	"spam_dnsbl_server_body":      {"name", "zone"},
	"spam_dnsbl_server_domain":    {"name", "zone"},
	"spam_dnsbl_server_email":     {"name", "zone"},
	"spam_dnsbl_server_header":    {"name", "zone"},
	"spam_dnsbl_server_ip":        {"name", "zone"},
	"spam_dnsbl_server_url":       {"name", "zone"},
	"spam_file_extension":         {"extension"},
	"spam_rule_any":               {"condition", "name"},
	"spam_rule_body":              {"condition", "name"},
	"spam_rule_domain":            {"condition", "name"},
	"spam_rule_email":             {"condition", "name"},
	"spam_rule_header":            {"condition", "name"},
	"spam_rule_ip":                {"condition", "name"},
	"spam_rule_url":               {"condition", "name"},
	"spam_tag_discard":            {"tag"},
	"spam_tag_reject":             {"tag"},
	"spam_tag_score":              {"tag"},
	"store_lookup":                {"namespace", "store"},
	"tracer_log":                  {"path"},
	"tracer_otel_http":            {"endpoint"},
	"user":                        {"domainId", "name"},
	"web_hook":                    {"url"},
}

var reloadActionOverrides = map[string]string{
	"x:Account":     "",
	"x:Domain":      "",
	"x:MailingList": "",
	"x:MaskedEmail": "",
	"x:OAuthClient": "",
	"x:Role":        "",
	"x:Tenant":      "",
	"x:AllowedIp":   "ReloadBlockedIps",
	"x:BlockedIp":   "ReloadBlockedIps",
	"x:Certificate": "ReloadTlsCertificates",
}

func reloadAction(objectName string) string {
	if action, ok := reloadActionOverrides[objectName]; ok {
		return action
	}

	return "ReloadSettings"
}

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "usage: generator <schema.json> <resources.go> <datasources.go>")
		os.Exit(1)
	}

	file, err := loadSchema(os.Args[1])
	if err != nil {
		fail(err)
	}

	r := &resolver{file: file}
	targets := buildTargets(r)
	dataSources := buildDataSourceTargets(r, targets)

	source, err := emit(targets)
	if err != nil {
		fail(err)
	}
	if err := os.WriteFile(os.Args[2], source, 0o644); err != nil {
		fail(err)
	}

	dataSourceCode, err := emitDataSources(dataSources)
	if err != nil {
		fail(err)
	}
	if err := os.WriteFile(os.Args[3], dataSourceCode, 0o644); err != nil {
		fail(err)
	}

	for _, message := range r.skipped {
		fmt.Println("skipped:", message)
	}
	fmt.Printf("generated %d resources into %s\n", len(targets), os.Args[2])
	fmt.Printf("generated %d data sources into %s\n", len(dataSources), os.Args[3])
}

func buildTargets(r *resolver) []resourceTarget {
	var targets []resourceTarget

	for _, objectName := range sortedKeys(r.file.Objects) {
		object := r.file.Objects[objectName]
		if object.Type == "view" || !strings.HasPrefix(objectName, "x:") {
			continue
		}
		if _, excluded := excludedObjects[objectName]; excluded {
			continue
		}

		entry, ok := r.file.Schemas[objectName]
		if !ok {
			r.note(objectName, "no schema entry")
			continue
		}

		base := resourceNameOverrides[objectName]
		if base == "" {
			base = resourceName(strings.TrimPrefix(objectName, "x:"))
		}
		singleton := object.Type == "singleton"

		switch {
		case entry.Type == "single":
			attributes := r.resolveGroup(entry.SchemaName)
			if attributes == nil {
				continue
			}
			targets = append(targets, newTarget(base, objectName, "", singleton, object, attributes, nil))

		case singleton:
			union := r.resolveUnion(objectName, entry.Variants)
			if union == nil {
				continue
			}
			targets = append(targets, newTarget(base, objectName, "", true, object, union.Children, union.TypeValues))

		default:
			for _, v := range entry.Variants {
				if deprecated(v.Name) {
					continue
				}
				name := resourceNameOverrides[objectName+"/"+v.Name]
				if name == "" {
					name = base + "_" + resourceName(v.Name)
				}
				attributes := map[string]*attrNode{}
				if v.SchemaName != "" {
					attributes = r.resolveGroup(v.SchemaName)
					if attributes == nil {
						continue
					}
				}
				target := newTarget(name, objectName, v.Name, false, object, attributes, nil)
				target.Description = variantDescription(object, v)
				targets = append(targets, target)
			}
		}
	}

	sort.Slice(targets, func(i, j int) bool { return targets[i].ResourceName < targets[j].ResourceName })

	return targets
}

func newTarget(name, jmapType, variantName string, singleton bool, object objectEntry, attributes map[string]*attrNode, typeValues []string) resourceTarget {
	required := map[string]bool{}
	for _, field := range requiredOverrides[name] {
		node, ok := attributes[field]
		if !ok {
			fail(fmt.Errorf("required override %s.%s does not exist", name, field))
		}
		if node.ServerSet {
			fail(fmt.Errorf("required override %s.%s is server set", name, field))
		}
		required[field] = true
	}
	if len(required) == 0 && !singleton {
		if node, ok := attributes["name"]; ok && !node.ServerSet {
			required["name"] = true
		}
	}

	description := object.Description
	if singleton {
		description += " This singleton object always exists: creating the resource adopts the current server settings and destroying it only removes the resource from state."
	}
	if object.Enterprise {
		description += " Requires an Enterprise license."
	}

	return resourceTarget{
		ResourceName: name,
		JMAPType:     jmapType,
		Variant:      variantName,
		Singleton:    singleton,
		ReloadAction: reloadAction(jmapType),
		Description:  description,
		Required:     required,
		Attributes:   attributes,
		TypeValues:   typeValues,
	}
}

func variantDescription(object objectEntry, v variant) string {
	description := fmt.Sprintf("%s Manages the %s variant.", object.Description, v.Label)
	if object.Enterprise {
		description += " Requires an Enterprise license."
	}

	return description
}

type emitter struct {
	imports map[string]bool
	body    bytes.Buffer
}

func emit(targets []resourceTarget) ([]byte, error) {
	e := &emitter{imports: map[string]bool{}}

	e.body.WriteString("var generatedResources = []resourceDescriptor{\n")
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
		return nil, fmt.Errorf("formatting generated source: %w", err)
	}

	return formatted, nil
}

func (e *emitter) need(path string) {
	e.imports[path] = true
}

func (e *emitter) emitTarget(t resourceTarget) {
	e.need("github.com/hashicorp/terraform-plugin-framework/resource/schema")
	e.need("github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier")
	e.need("github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier")

	fmt.Fprintf(&e.body, "\t{\n\t\tName: %q,\n\t\tJMAPType: %q,\n", t.ResourceName, t.JMAPType)
	if t.Variant != "" {
		fmt.Fprintf(&e.body, "\t\tVariant: %q,\n", t.Variant)
	}
	if t.Singleton {
		fmt.Fprintf(&e.body, "\t\tSingleton: true,\n")
	}
	if t.ReloadAction != "" {
		fmt.Fprintf(&e.body, "\t\tReloadAction: %q,\n", t.ReloadAction)
	}
	fmt.Fprintf(&e.body, "\t\tSchema: schema.Schema{\n")
	fmt.Fprintf(&e.body, "\t\t\tMarkdownDescription: %q,\n", t.Description)
	fmt.Fprintf(&e.body, "\t\t\tAttributes: map[string]schema.Attribute{\n")

	fmt.Fprintf(&e.body, "\t\t\t\t\"id\": schema.StringAttribute{\n")
	fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: \"Server-assigned identifier.\",\n")
	fmt.Fprintf(&e.body, "\t\t\t\t\tComputed: true,\n")
	fmt.Fprintf(&e.body, "\t\t\t\t\tPlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},\n")
	fmt.Fprintf(&e.body, "\t\t\t\t},\n")

	if len(t.TypeValues) > 0 {
		fmt.Fprintf(&e.body, "\t\t\t\t\"type\": %s,\n", e.renderDiscriminator(t.TypeValues, "\t\t\t\t"))
	}

	for _, name := range sortedKeys(t.Attributes) {
		node := t.Attributes[name]
		node.Required = t.Required[name]
		fmt.Fprintf(&e.body, "\t\t\t\t%q: %s,\n", rootTerraformName(name), e.render(node, "\t\t\t\t"))
		if writeOnlyCandidate(node) {
			e.emitWriteOnlyPair(rootTerraformName(name))
		}
	}

	fmt.Fprintf(&e.body, "\t\t\t},\n\t\t},\n\t},\n")
}

func writeOnlyCandidate(node *attrNode) bool {
	return node.Kind == "string" && node.Sensitive && !node.ServerSet && !node.Required
}

func (e *emitter) emitWriteOnlyPair(base string) {
	fmt.Fprintf(&e.body, "\t\t\t\t%q: schema.StringAttribute{\n", base+"_wo")
	fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: %q,\n",
		fmt.Sprintf("Write-only variant of `%s`. The value is sent to the server but never stored in state. Set `%s_wo_version` and change it to roll the secret. Requires Terraform 1.11 or later.", base, base))
	fmt.Fprintf(&e.body, "\t\t\t\t\tOptional: true,\n")
	fmt.Fprintf(&e.body, "\t\t\t\t\tWriteOnly: true,\n")
	fmt.Fprintf(&e.body, "\t\t\t\t\tSensitive: true,\n")
	fmt.Fprintf(&e.body, "\t\t\t\t},\n")

	fmt.Fprintf(&e.body, "\t\t\t\t%q: schema.Int64Attribute{\n", base+"_wo_version")
	fmt.Fprintf(&e.body, "\t\t\t\t\tMarkdownDescription: %q,\n",
		fmt.Sprintf("Version counter for `%s_wo`. Change this value to send the current `%s_wo` value to the server again.", base, base))
	fmt.Fprintf(&e.body, "\t\t\t\t\tOptional: true,\n")
	fmt.Fprintf(&e.body, "\t\t\t\t},\n")
}

func (e *emitter) renderDiscriminator(values []string, indent string) string {
	e.need("github.com/hashicorp/terraform-plugin-framework/schema/validator")
	e.need("github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator")

	inner := indent + "\t"

	var b strings.Builder
	fmt.Fprintf(&b, "schema.StringAttribute{\n")
	fmt.Fprintf(&b, "%sMarkdownDescription: \"Variant discriminator.\",\n", inner)
	fmt.Fprintf(&b, "%sRequired: true,\n", inner)
	fmt.Fprintf(&b, "%sValidators: []validator.String{stringvalidator.OneOf(%s)},\n", inner, quotedList(values))
	fmt.Fprintf(&b, "%s}", indent)

	return b.String()
}

func (e *emitter) render(node *attrNode, indent string) string {
	inner := indent + "\t"
	kind := attributeKind(node)

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", kind)
	fmt.Fprintf(&b, "%sMarkdownDescription: %q,\n", inner, node.Description)

	switch {
	case node.ServerSet:
		fmt.Fprintf(&b, "%sComputed: true,\n", inner)
		e.writeModifiers(&b, kind, "UseStateForUnknown()", inner)
	case node.Required:
		fmt.Fprintf(&b, "%sRequired: true,\n", inner)
		if node.Immutable {
			e.writeModifiers(&b, kind, "RequiresReplace()", inner)
		}
	default:
		fmt.Fprintf(&b, "%sOptional: true,\n%sComputed: true,\n", inner, inner)
		if node.Immutable {
			e.writeModifiers(&b, kind, "UseStateForUnknown()", "RequiresReplace()", inner)
		} else {
			e.writeModifiers(&b, kind, "UseStateForUnknown()", inner)
		}
	}

	if node.Sensitive {
		fmt.Fprintf(&b, "%sSensitive: true,\n", inner)
	}

	if node.Kind == "set" || node.Kind == "map" {
		e.need("github.com/hashicorp/terraform-plugin-framework/types")
		fmt.Fprintf(&b, "%sElementType: %s,\n", inner, elementType(node.ElemKind))
	}

	if node.Kind == "string" && len(node.EnumValues) > 0 {
		e.need("github.com/hashicorp/terraform-plugin-framework/schema/validator")
		e.need("github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator")
		fmt.Fprintf(&b, "%sValidators: []validator.String{stringvalidator.OneOf(%s)},\n", inner, quotedList(node.EnumValues))
	}
	if node.Kind == "set" && len(node.EnumValues) > 0 {
		e.need("github.com/hashicorp/terraform-plugin-framework/schema/validator")
		e.need("github.com/hashicorp/terraform-plugin-framework-validators/setvalidator")
		e.need("github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator")
		fmt.Fprintf(&b, "%sValidators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf(%s))},\n", inner, quotedList(node.EnumValues))
	}

	switch node.Kind {
	case "object", "union":
		fmt.Fprintf(&b, "%sAttributes: map[string]schema.Attribute{\n", inner)
		if node.Kind == "union" {
			fmt.Fprintf(&b, "%s\t\"type\": %s,\n", inner, e.renderDiscriminator(node.TypeValues, inner+"\t"))
		}
		for _, name := range sortedKeys(node.Children) {
			fmt.Fprintf(&b, "%s\t%q: %s,\n", inner, terraformName(name), e.render(node.Children[name], inner+"\t"))
		}
		fmt.Fprintf(&b, "%s},\n", inner)

	case "list":
		fmt.Fprintf(&b, "%sNestedObject: schema.NestedAttributeObject{\n", inner)
		fmt.Fprintf(&b, "%s\tAttributes: map[string]schema.Attribute{\n", inner)
		if len(node.TypeValues) > 0 {
			fmt.Fprintf(&b, "%s\t\t\"type\": %s,\n", inner, e.renderDiscriminator(node.TypeValues, inner+"\t\t"))
		}
		for _, name := range sortedKeys(node.Children) {
			fmt.Fprintf(&b, "%s\t\t%q: %s,\n", inner, terraformName(name), e.render(node.Children[name], inner+"\t\t"))
		}
		fmt.Fprintf(&b, "%s\t},\n%s},\n", inner, inner)
	}

	fmt.Fprintf(&b, "%s}", indent)

	return b.String()
}

func (e *emitter) writeModifiers(b *strings.Builder, kind string, parts ...string) {
	indent := parts[len(parts)-1]
	modifiers := parts[:len(parts)-1]

	suffix := strings.TrimPrefix(kind, "schema.")
	suffix = strings.TrimSuffix(suffix, "Attribute")
	var planType, pkg string
	switch suffix {
	case "String":
		planType, pkg = "String", "stringplanmodifier"
	case "Bool":
		planType, pkg = "Bool", "boolplanmodifier"
	case "Int64":
		planType, pkg = "Int64", "int64planmodifier"
	case "Float64":
		planType, pkg = "Float64", "float64planmodifier"
	case "Set":
		planType, pkg = "Set", "setplanmodifier"
	case "Map":
		planType, pkg = "Map", "mapplanmodifier"
	case "SingleNested":
		planType, pkg = "Object", "objectplanmodifier"
	case "ListNested":
		planType, pkg = "List", "listplanmodifier"
	default:
		return
	}

	e.need("github.com/hashicorp/terraform-plugin-framework/resource/schema/" + pkg)
	calls := make([]string, 0, len(modifiers))
	for _, m := range modifiers {
		calls = append(calls, pkg+"."+m)
	}
	fmt.Fprintf(b, "%sPlanModifiers: []planmodifier.%s{%s},\n", indent, planType, strings.Join(calls, ", "))
}

func attributeKind(node *attrNode) string {
	switch node.Kind {
	case "string":
		return "schema.StringAttribute"
	case "bool":
		return "schema.BoolAttribute"
	case "int":
		return "schema.Int64Attribute"
	case "float":
		return "schema.Float64Attribute"
	case "set":
		return "schema.SetAttribute"
	case "map":
		return "schema.MapAttribute"
	case "object", "union":
		return "schema.SingleNestedAttribute"
	case "list":
		return "schema.ListNestedAttribute"
	}

	panic("unknown attribute kind " + node.Kind)
}

func elementType(kind string) string {
	switch kind {
	case "bool":
		return "types.BoolType"
	case "int":
		return "types.Int64Type"
	case "float":
		return "types.Float64Type"
	}

	return "types.StringType"
}

func quotedList(values []string) string {
	quoted := make([]string, 0, len(values))
	for _, v := range values {
		quoted = append(quoted, fmt.Sprintf("%q", v))
	}

	return strings.Join(quoted, ", ")
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "generator:", err)
	os.Exit(1)
}
