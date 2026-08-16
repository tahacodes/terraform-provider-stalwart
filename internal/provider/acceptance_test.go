package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
