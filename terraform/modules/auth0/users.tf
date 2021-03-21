
resource "auth0_user" "mind_ui_test_account" {
  connection_name = auth0_connection.mind_ui_connection.name
  #   username        = "test_account_${var.env}"
  name           = "Test Account"
  given_name     = "Test"
  family_name    = "Account"
  nickname       = "test.account"
  email          = "mind-${var.env}-test@jdpx.co.uk"
  email_verified = true
  password       = var.auth0_user_default_password
  roles = [
    auth0_role.admin_role.id,
    auth0_role.test_organisation_role.id
  ]
}
