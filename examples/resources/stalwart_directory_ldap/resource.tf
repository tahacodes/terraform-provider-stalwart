resource "stalwart_directory_ldap" "corporate" {
  description = "Corporate LDAP directory"
  url         = "ldaps://ldap.example.com:636"
  base_dn     = "dc=example,dc=com"
  bind_dn     = "cn=stalwart,ou=services,dc=example,dc=com"

  bind_secret = {
    type   = "Value"
    secret = var.ldap_bind_password
  }

  filter_login   = "(&(objectClass=posixAccount)(uid=?))"
  filter_mailbox = "(&(objectClass=posixAccount)(mail=?))"
}
