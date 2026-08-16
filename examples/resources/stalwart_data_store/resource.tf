resource "stalwart_data_store" "postgres" {
  type          = "PostgreSql"
  host          = "db.example.com"
  port          = 5432
  database      = "stalwart"
  auth_username = "stalwart"

  auth_secret = {
    type   = "Value"
    secret = var.database_password
  }
}
