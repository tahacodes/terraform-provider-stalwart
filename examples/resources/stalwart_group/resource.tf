resource "stalwart_group" "admins" {
  name          = "admins"
  description   = "Platform administrators"
  email_address = "admins@example.com"
  domain_id     = stalwart_domain.example.id
}
