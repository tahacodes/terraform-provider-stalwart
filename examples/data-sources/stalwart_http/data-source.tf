data "stalwart_http" "current" {}

output "redirect_root" {
  value = data.stalwart_http.current.redirect_root
}
