data "stalwart_domain" "main" {
  name = "example.com"
}

resource "stalwart_user" "info" {
  name      = "info"
  domain_id = data.stalwart_domain.main.id
}
