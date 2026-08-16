package provider

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func jmapName(terraform string) string {
	parts := strings.Split(terraform, "_")

	var b strings.Builder
	for i, part := range parts {
		if part == "" {
			continue
		}
		if i == 0 {
			b.WriteString(part)
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]))
		b.WriteString(part[1:])
	}

	return b.String()
}

func diagsError(diags diag.Diagnostics) error {
	if !diags.HasError() {
		return nil
	}

	messages := make([]string, 0, len(diags))
	for _, d := range diags.Errors() {
		messages = append(messages, d.Summary()+": "+d.Detail())
	}

	return errors.New(strings.Join(messages, "; "))
}
