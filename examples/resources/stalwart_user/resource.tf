resource "stalwart_user" "alice" {
  name          = "alice"
  description   = "Alice Example"
  email_address = "alice@example.com"
  domain_id     = stalwart_domain.example.id
}
