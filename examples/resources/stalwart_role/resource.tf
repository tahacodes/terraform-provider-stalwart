resource "stalwart_role" "auditor" {
  name        = "auditor"
  description = "Read-only access for auditing"

  enabled_permissions = [
    "messageQueueList",
    "reportList",
  ]
}
