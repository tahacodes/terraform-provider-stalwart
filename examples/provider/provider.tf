terraform {
  required_providers {
    stalwart = {
      source  = "tahacodes/stalwart"
      version = "~> 0.1"
    }
  }
}

provider "stalwart" {
  endpoint = "https://mail.example.com"
  username = "admin"
  password = var.stalwart_admin_password
}
