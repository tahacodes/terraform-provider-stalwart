resource "stalwart_network_listener" "submissions" {
  name     = "submissions"
  protocol = "smtp"
  bind     = ["[::]:465"]
  use_tls  = true
}
