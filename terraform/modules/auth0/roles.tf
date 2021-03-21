
resource "auth0_role" "test_organisation_role" {
  name        = "OrganisationTest${var.env}"
  description = "Member Of Test Organisation ${var.env}"

  permissions {
    resource_server_identifier = auth0_resource_server.mind_ui_api.identifier
    name                       = "read:organisation:1ohTDUgTsyvWet5lJFCCqnA2F1S"
  }
}

resource "auth0_role" "admin_role" {
  name        = "Admin User ${var.env}"
  description = "Admin User ${var.env}"

  permissions {
    resource_server_identifier = auth0_resource_server.mind_ui_api.identifier
    name                       = "mind:admin"
  }
}

