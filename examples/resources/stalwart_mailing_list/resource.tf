resource "stalwart_mailing_list" "announce" {
  name          = "announce"
  email_address = "announce@example.com"
  domain_id     = stalwart_domain.example.id
  recipients    = ["alice@example.com"]
}
