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


variable "graph_cms_url" {
  type        = string
  description = "GraphCMS API URL"
}
