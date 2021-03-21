
resource "auth0_client" "mind_ui_application" {
  name     = "mind-ui-${var.env}"
  app_type = "spa"

  allowed_logout_urls = [
    "https://www.${var.env}.mind.jdpx.co.uk",
    "https://${var.env}.mind.jdpx.co.uk",
  ]
  allowed_origins = [
    "https://www.${var.env}.mind.jdpx.co.uk",
    "https://${var.env}.mind.jdpx.co.uk",
  ]
  callbacks = [
    "https://www.${var.env}.mind.jdpx.co.uk",
    "https://${var.env}.mind.jdpx.co.uk",
  ]
  web_origins = [
    "https://www.${var.env}.mind.jdpx.co.uk",
    "https://${var.env}.mind.jdpx.co.uk",
  ]
}
