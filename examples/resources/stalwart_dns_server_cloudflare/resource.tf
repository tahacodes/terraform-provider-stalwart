resource "stalwart_dns_server_cloudflare" "primary" {
  description = "Cloudflare DNS for automatic record management"

  secret = {
    type   = "Value"
    secret = var.cloudflare_api_token
  }
}
