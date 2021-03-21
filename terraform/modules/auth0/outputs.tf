
output "api_audience" {
  value       = auth0_resource_server.mind_ui_api.identifier
  description = "Mind API Audience"
}

output "tenant_id" {
  value       = auth0_tenant.mind_tenant.id
  description = "Mind API Audience"
}
