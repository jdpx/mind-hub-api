
resource "auth0_connection" "mind_ui_connection" {
  name     = "mind-${var.env}-users"
  strategy = "auth0"

  options {
    api_enable_users = true
    disable_signup   = true
    password_policy  = "good"

    password_history {
      enable = false
      size   = 3
    }
  }

  enabled_clients = [
    "S1YjAwRqZ0ONFwdAfgU0dXWP4a8rxMBF",
    auth0_client.mind_ui_application.id
  ]
}
