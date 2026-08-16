# Terraform Provider for Stalwart

Manages configuration of a [Stalwart](https://stalw.art) mail and collaboration server (v0.16+) through its JMAP management API.

Stalwart exposes its entire management surface as JMAP objects under the `x:` namespace, reached by `POST /jmap` with standard `get`, `set` and `query` methods. This provider speaks that protocol directly, so every resource supports real reads and therefore real drift detection.

The provider covers the full management surface: 196 resources generated from the server's own schema. That includes principals (domains, users, groups, mailing lists, tenants, roles), security (API keys, OAuth clients, allowed and blocked IPs, certificates, ACME providers), storage backends (data, blob, in-memory, search, metrics and tracing stores), directories (LDAP, SQL, OIDC), the full MTA pipeline (stages, routes, throttles, queues, strategies, milters, hooks), spam filtering, DKIM, DNS automation for about 70 DNS providers, telemetry, and every server settings singleton (HTTP, IMAP, JMAP, WebDAV, calendar, sieve, and more). Runtime data such as queued messages, tasks, traces and received reports is intentionally not exposed as resources.

## Resource semantics

- Server settings such as `stalwart_http` or `stalwart_imap` are singletons: they always exist on the server, so creating the resource adopts the current settings, and destroying it only removes the resource from state.
- Backend choices are variant unions. Singleton unions such as `stalwart_data_store` take a `type` attribute (for example `PostgreSql`), while multi-instance unions map to one resource per variant, such as `stalwart_directory_ldap` or `stalwart_dns_server_cloudflare`.
- Secret values are marked sensitive. The server returns them masked (`****`) on reads, and the provider keeps the configured value in state instead of the mask.
- Stalwart persists settings changes without applying them to the running server. The provider therefore triggers the matching reload action (`ReloadSettings`, `ReloadTlsCertificates` or `ReloadBlockedIps`) after each apply, with coalescing so parallel applies do not cause redundant reloads. Directory-backed resources such as users, domains and roles take effect immediately and never trigger a reload. Set `auto_reload = false` on the provider to manage reloads yourself.
- Every resource has a matching data source with the same name, looked up by `id` or, where the object has one, by `name`. Plural data sources such as `stalwart_domains` list all identifiers of a type. Read-only runtime objects that are deliberately not resources are also available as data sources: `stalwart_queued_message`, `stalwart_task`, `stalwart_log`, `stalwart_metric`, `stalwart_archived_item`, `stalwart_cluster_node`, `stalwart_api_key`, `stalwart_app_password`, the DMARC/TLS/ARF report objects and their plural forms. Secrets read through data sources stay masked as `****`.
- Required attributes match the server's own validation rules, verified against a live server for every creatable type, so missing fields fail at plan time instead of at apply time.

## Keeping secrets out of state

Two mechanisms keep secret values out of the Terraform state file:

- Most credentials are secret unions with three variants. `type = "Value"` embeds the secret in configuration and state; `type = "EnvironmentVariable"` and `type = "File"` make the server read the secret from its own environment or filesystem, so the value never touches Terraform at all:

  ```hcl
  secret = {
    type          = "EnvironmentVariable"
    variable_name = "CLOUDFLARE_API_TOKEN"
  }
  ```

- Plain string secrets (`stalwart_oauth_client.secret`, `stalwart_acme_provider.eab_hmac_key`) have a write-only companion. Set `secret_wo` instead of `secret` and the value is sent to the server but never stored in state or plan; change `secret_wo_version` to roll the secret. Requires Terraform 1.11 or later:

  ```hcl
  resource "stalwart_oauth_client" "app" {
    client_id         = "my-app"
    secret_wo         = var.oauth_secret
    secret_wo_version = 1
  }
  ```

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.9
- [Go](https://go.dev/doc/install) >= 1.24 to build from source
- Stalwart >= 0.16

## Using the provider

```hcl
terraform {
  required_providers {
    stalwart = {
      source  = "tahacodes/stalwart"
      version = "~> 0.1"
    }
  }
}

provider "stalwart" {
  endpoint = "https://mail.example.com"
  username = "admin"
  password = var.stalwart_admin_password
}

resource "stalwart_http" "this" {
  use_permissive_cors = true
  enable_hsts         = true
}
```

Credentials may also come from the environment, using the same variables as `stalwart-cli`:

| Variable | Purpose |
| --- | --- |
| `STALWART_URL` | Base URL of the server |
| `STALWART_USER` | Basic-auth username |
| `STALWART_PASSWORD` | Basic-auth password |
| `STALWART_TOKEN` | Bearer token, instead of username and password |

## Schema-driven design

Stalwart publishes a machine-readable description of its own configuration at `GET /api/schema`: 150 object types, 316 field groups and 153 enums, each carrying names, types, descriptions and defaults. A snapshot is vendored under `schema/` per Stalwart version so builds are reproducible and CI needs no live server.

Refresh a snapshot against a running server with:

```sh
make schema STALWART_URL=https://mail.example.com STALWART_USER=admin STALWART_PASSWORD=... STALWART_VERSION=v0.16.17
```

## Development

```sh
make build      # compile
make test       # unit tests
make testacc    # acceptance tests, requires a reachable Stalwart server
make docs       # regenerate docs/ from schema descriptions and examples
make lint
```

Acceptance tests create real objects on the server they target. Never point them at production.

## License

MIT. See [LICENSE](LICENSE).
