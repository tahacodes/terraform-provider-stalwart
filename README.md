# Terraform Provider for Stalwart

Manages configuration of a [Stalwart](https://stalw.art) mail and collaboration server (v0.16+) through its JMAP management API.

Stalwart exposes its entire management surface as JMAP objects under the `x:` namespace, reached by `POST /jmap` with standard `get`, `set` and `query` methods. This provider speaks that protocol directly, so every resource supports real reads and therefore real drift detection.

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
