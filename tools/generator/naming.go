package main

import (
	"strings"
	"unicode"
)

var rootAttributeRenames = map[string]string{
	"connection": "connection_strategy",
	"count":      "count_value",
}

func rootTerraformName(jmap string) string {
	if renamed, ok := rootAttributeRenames[jmap]; ok {
		return renamed
	}

	return terraformName(jmap)
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

func pluralName(singular string) string {
	if strings.HasSuffix(singular, "y") && len(singular) > 1 && !strings.ContainsRune("aeiou", rune(singular[len(singular)-2])) {
		return singular[:len(singular)-1] + "ies"
	}

	return singular + "s"
}

func resourceName(camel string) string {
	runes := []rune(camel)

	var b strings.Builder
	for i, r := range runes {
		if unicode.IsUpper(r) {
			previousLower := i > 0 && (unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1]))
			nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if i > 0 && (previousLower || (nextLower && unicode.IsUpper(runes[i-1]))) {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}

	return b.String()
}
