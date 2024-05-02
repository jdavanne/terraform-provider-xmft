resource "xmft_st_advanced_routing_application" "ar1" {
  provider       = xmft.st1
  name           = "ar1"
  type           = "AdvancedRouting"
  notes          = "mynotes"
  business_units = []
}

resource "xmft_st_route_template" "template1" {
  provider       = xmft.st1
  name           = "template1"
  type           = "TEMPLATE"
  description    = "mydescription"
  business_units = []
}

resource "xmft_st_account" "account1" {
  provider    = xmft.st1
  name        = "account1"
  uid         = "1000"
  gid         = "1000"
  home_folder = "/files/account1"
  user = {
    name = "login1"
    password_credentials = {
      password = "password1"
    }
  }
}

resource "xmft_st_subscription_ar" "sub1" {
  provider    = xmft.st1
  account     = xmft_st_account.account1.name
  folder      = "/folder1"
  application = xmft_st_advanced_routing_application.ar1.name
}

resource "xmft_st_site_ssh" "ssh1" {
  provider      = xmft.st1
  name          = "ssh2"
  account       = xmft_st_account.account1.name
  host          = "host"
  port          = "8022"
  user_name     = xmft_st_account.account1.user.name
  password      = xmft_st_account.account1.user.password_credentials.password
  upload_folder = "/"
}

resource "xmft_st_route_simple" "simple1" {
  provider = xmft.st1
  name     = "simple1"
  #type           = "SIMPLE"
  description = "mydescription"

  steps = [{
    type                     = "SendToPartner"
    status                   = "ENABLED"
    transfer_site_expression = "ssh1#!#CVD#!#"
    }
  ]
}

resource "xmft_st_route_composite" "route_composite1" {
  provider       = xmft.st1
  name           = "route1"
  description    = "mydescription"
  route_template = xmft_st_route_template.template1.id
  subscriptions  = [xmft_st_subscription_ar.sub1.id]
  account        = xmft_st_account.account1.name

  steps = [
    { execute_route = xmft_st_route_simple.simple1.id }
  ]
}
