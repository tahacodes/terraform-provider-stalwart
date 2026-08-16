data "stalwart_domains" "all" {}

data "stalwart_domain" "each" {
  for_each = data.stalwart_domains.all.ids

  id = each.value
}

output "domain_names" {
  value = [for domain in data.stalwart_domain.each : domain.name]
}
