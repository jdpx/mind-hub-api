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

variable "api_stage_name" {
  type        = string
  description = "API Gateway Stage Name"
  default     = "v1"
}

variable "graph_cms_url_mapping" {
  type        = string
  description = "A string representation for the mapping from organisation id to GraphCMS API URL"
}
