resource "stalwart_domain" "example" {
  name        = "example.com"
  description = "Primary mail domain"
  is_enabled  = true

  dkim_management = {
    type = "Automatic"
  }

  certificate_management = {
    type = "Manual"
  }
}
