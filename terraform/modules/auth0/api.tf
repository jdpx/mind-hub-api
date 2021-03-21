
resource "auth0_resource_server" "mind_ui_api" {
  name        = "mind-api-${var.env}"
  identifier  = "https://api.${var.env}.mind.jdpx.co.uk"
  signing_alg = "RS256"

  allow_offline_access                            = false
  enforce_policies                                = true
  skip_consent_for_verifiable_first_party_clients = true
  token_dialect                                   = "access_token" #tfsec:ignore:GEN003
  token_lifetime                                  = 86400

  scopes {
    value       = "read:organisation:1ohTDUgTsyvWet5lJFCCqnA2F1S"
    description = "Allow Test Organisation Access"
  }

  scopes {
    value       = "mind:admin"
    description = "Admin Access"
  }
}
