package provider

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var protoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"stalwart": providerserver.NewProtocol6WithError(New("acceptance")()),
}

func testAccPreCheck(t *testing.T) {
	t.Helper()

	for _, key := range []string{"STALWART_URL", "STALWART_USER", "STALWART_PASSWORD"} {
		if os.Getenv(key) == "" {
			t.Fatalf("%s must be set for acceptance tests", key)
		}
	}
}

func TestAccRoleLifecycle(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_role" "test" {
  description = "created by acceptance test"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_role.test", "description", "created by acceptance test"),
					resource.TestCheckResourceAttrSet("stalwart_role.test", "id"),
				),
			},
			{
				Config: `
resource "stalwart_role" "test" {
  description = "updated by acceptance test"
}
`,
				Check: resource.TestCheckResourceAttr("stalwart_role.test", "description", "updated by acceptance test"),
			},
			{
				ResourceName:      "stalwart_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccHTTPSingleton(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_http" "test" {
  enable_hsts = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_http.test", "id", "singleton"),
					resource.TestCheckResourceAttr("stalwart_http.test", "enable_hsts", "true"),
					resource.TestCheckResourceAttrSet("stalwart_http.test", "redirect_root"),
				),
			},
			{
				Config: `
resource "stalwart_http" "test" {
  enable_hsts = false
}
`,
				Check: resource.TestCheckResourceAttr("stalwart_http.test", "enable_hsts", "false"),
			},
			{
				ResourceName:      "stalwart_http.test",
				ImportState:       true,
				ImportStateId:     "singleton",
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSettingsReloadTakesEffect(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_http" "test" {
  response_headers = {
    "X-Acc-Reload" = "applied"
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_http.test", "response_headers.X-Acc-Reload", "applied"),
					checkServerHeader("X-Acc-Reload", "applied"),
				),
			},
		},
	})
}

func checkServerHeader(name, want string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		response, err := http.Get(os.Getenv("STALWART_URL") + "/healthz/live")
		if err != nil {
			return err
		}
		defer func() { _ = response.Body.Close() }()

		if got := response.Header.Get(name); got != want {
			return fmt.Errorf("header %s = %q, want %q", name, got, want)
		}

		return nil
	}
}

func TestAccUserWithAliases(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_domain" "test" {
  name = "acctest-users.com"

  dkim_management = {
    type = "Manual"
  }
}

resource "stalwart_user" "test" {
  name      = "acctest-user"
  domain_id = stalwart_domain.test.id

  aliases = [
    {
      name      = "acctest-alias"
      domain_id = stalwart_domain.test.id
    },
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_user.test", "email_address", "acctest-user@acctest-users.com"),
					resource.TestCheckResourceAttr("stalwart_user.test", "aliases.#", "1"),
					resource.TestCheckResourceAttr("stalwart_user.test", "aliases.0.name", "acctest-alias"),
				),
			},
			{
				Config: `
resource "stalwart_domain" "test" {
  name = "acctest-users.com"

  dkim_management = {
    type = "Manual"
  }
}

resource "stalwart_user" "test" {
  name      = "acctest-user"
  domain_id = stalwart_domain.test.id

  aliases = [
    {
      name      = "acctest-alias"
      domain_id = stalwart_domain.test.id
    },
    {
      name      = "acctest-alias-two"
      domain_id = stalwart_domain.test.id
    },
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_user.test", "aliases.#", "2"),
					resource.TestCheckResourceAttr("stalwart_user.test", "aliases.1.name", "acctest-alias-two"),
				),
			},
		},
	})
}

func TestAccTracerVariant(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_tracer_stdout" "test" {
  enable = true
  level  = "info"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_tracer_stdout.test", "level", "info"),
					resource.TestCheckResourceAttrSet("stalwart_tracer_stdout.test", "id"),
				),
			},
			{
				Config: `
resource "stalwart_tracer_stdout" "test" {
  enable = true
  level  = "debug"
}
`,
				Check: resource.TestCheckResourceAttr("stalwart_tracer_stdout.test", "level", "debug"),
			},
		},
	})
}

func TestAccUnionSingleton(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_asn" "test" {
  type = "Disabled"
}
`,
				Check: resource.TestCheckResourceAttr("stalwart_asn.test", "type", "Disabled"),
			},
			{
				ResourceName:      "stalwart_asn.test",
				ImportState:       true,
				ImportStateId:     "singleton",
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccWriteOnlySecret(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_oauth_client" "test" {
  client_id         = "acctest-client"
  secret_wo         = "first-secret-value"
  secret_wo_version = 1
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("stalwart_oauth_client.test", "id"),
					resource.TestCheckResourceAttr("stalwart_oauth_client.test", "secret_wo_version", "1"),
					resource.TestCheckNoResourceAttr("stalwart_oauth_client.test", "secret_wo"),
				),
			},
			{
				Config: `
resource "stalwart_oauth_client" "test" {
  client_id         = "acctest-client"
  secret_wo         = "second-secret-value"
  secret_wo_version = 2
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_oauth_client.test", "secret_wo_version", "2"),
					resource.TestCheckNoResourceAttr("stalwart_oauth_client.test", "secret_wo"),
				),
			},
		},
	})
}

func TestAccRequiredFieldsCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_network_listener" "test" {
  name     = "acctest-listener"
  protocol = "smtp"
  bind     = ["127.0.0.1:12525"]
}

resource "stalwart_web_hook" "test" {
  url = "https://hooks.example.com/stalwart"
}

resource "stalwart_memory_lookup_key" "test" {
  namespace = "acctest"
  key       = "probe"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("stalwart_network_listener.test", "id"),
					resource.TestCheckResourceAttrSet("stalwart_web_hook.test", "id"),
					resource.TestCheckResourceAttrSet("stalwart_memory_lookup_key.test", "id"),
				),
			},
			{
				ResourceName:      "stalwart_network_listener.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDomainWithUnions(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "stalwart_domain" "test" {
  name        = "acctest-provider.com"
  description = "created by acceptance test"

  dkim_management = {
    type = "Manual"
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stalwart_domain.test", "name", "acctest-provider.com"),
					resource.TestCheckResourceAttr("stalwart_domain.test", "dkim_management.type", "Manual"),
					resource.TestCheckResourceAttrSet("stalwart_domain.test", "id"),
				),
			},
			{
				ResourceName:      "stalwart_domain.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
