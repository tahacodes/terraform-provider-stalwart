resource "stalwart_http" "this" {
  use_permissive_cors = true
  enable_hsts         = true
  use_x_forwarded     = true
  redirect_root       = "/account"

  response_headers = {
    "X-Frame-Options" = "SAMEORIGIN"
  }
}
