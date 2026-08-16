package main

import "testing"

func TestResourceName(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"MtaSts":             "mta_sts",
		"DnsServer":          "dns_server",
		"AiModel":            "ai_model",
		"AcmeProvider":       "acme_provider",
		"Dkim1Ed25519Sha256": "dkim1_ed25519_sha256",
		"BluecatV2":          "bluecat_v2",
		"Ipv64":              "ipv64",
		"Route53":            "route53",
		"WebDav":             "web_dav",
	}

	for input, want := range cases {
		if got := resourceName(input); got != want {
			t.Errorf("resourceName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestTerraformName(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"emailAddress": "email_address",
		"connection":   "connection",
		"count":        "count",
		"dnsIpv4":      "dns_ipv4",
		"name":         "name",
	}

	for input, want := range cases {
		if got := terraformName(input); got != want {
			t.Errorf("terraformName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestRootTerraformName(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"connection": "connection_strategy",
		"count":      "count_value",
		"name":       "name",
	}

	for input, want := range cases {
		if got := rootTerraformName(input); got != want {
			t.Errorf("rootTerraformName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestPluralName(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"domain":                  "domains",
		"directory":               "directories",
		"mta_connection_strategy": "mta_connection_strategies",
		"blocked_ip":              "blocked_ips",
	}

	for input, want := range cases {
		if got := pluralName(input); got != want {
			t.Errorf("pluralName(%q) = %q, want %q", input, got, want)
		}
	}
}
