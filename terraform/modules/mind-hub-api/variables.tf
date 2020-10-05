variable "auth0_audience" {
  type        = string
  description = "Auth0 Audience"
}

variable "auth0_jwks_uri" {
  type        = string
  description = "JWKs URI"
}

variable "auth0_token_issuer" {
  type        = string
  description = "Token Issuer"
}

variable "env" {
  type        = string
  description = "Environment"
}
