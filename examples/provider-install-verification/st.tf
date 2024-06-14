locals {
  host = "ci8.jda.axwaytest.net"
  rnd  = "21"
}

provider "xmft" {
  product = "st"
  alias   = "st1"
  host    = "https://${local.host}:8444"
  #host     = "https://localhost:1769"
  username = "admin"
  password = "admin*"
}


data "xmft_st_version" "version1" {
  provider = xmft.st1
}

output "version1" {
  value = data.xmft_st_version.version1
}

#resource "random_string" "name" {
#  length = 8
#}

resource "xmft_st_account" "account_src" {
  provider    = xmft.st1
  name        = "account${local.rnd}" #+ random_string.name.result
  home_folder = "/files/account${local.rnd}"
  user = {
    name = "login${local.rnd}"
    password_credentials = {
      password = "password${local.rnd}"
    }
  }
}

resource "xmft_st_account" "account3" {
  provider    = xmft.st1
  name        = "account-target${local.rnd}"
  home_folder = "/files/account-target-${local.rnd}"
  user = {
    name = "logintarget${local.rnd}"
    password_credentials = {
      password = "passwordt${local.rnd}"
    }
  }
}

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
  description    = "mydescription"
  business_units = []
}

resource "xmft_st_subscription_ar" "sub2" {
  provider    = xmft.st1
  account     = xmft_st_account.account_src.name
  folder      = "/folder2"
  application = xmft_st_advanced_routing_application.ar1.name
}

resource "xmft_st_route_composite" "route_composite1" {
  provider       = xmft.st1
  name           = "route1"
  description    = "mydescription"
  route_template = xmft_st_route_template.template1.id
  subscriptions  = [xmft_st_subscription_ar.sub2.id]
  account        = xmft_st_account.account_src.name

  steps = [
    { execute_route = xmft_st_route_simple.simple1.id }
  ]
}

resource "xmft_st_route_simple" "simple1" {
  name     = "simple31"
  provider = xmft.st1
  steps = [{
    type                     = "SendToPartner"
    transfer_site_expression = "${xmft_st_site_ssh.ssh2.name}#!#CVD#!#"
    }
  ]
}

resource "xmft_st_site_ssh" "ssh2" {
  provider = xmft.st1
  name     = "ssh2"
  account  = xmft_st_account.account_src.name
  #type		     = "ssh"
  host      = local.host
  port      = "8022"
  user_name = xmft_st_account.account3.user.name
  password  = xmft_st_account.account3.user.password_credentials.password
  #download_folder  = "/download"
  upload_folder = "/"
}

output "cmd1" {
  value = <<EOF
SSHPASS=${xmft_st_account.account_src.user.password_credentials.password} sshpass -e sftp -P ${xmft_st_site_ssh.ssh2.port} -oBatchMode=no -b - ${xmft_st_account.account_src.user.name}@${local.host} << !
   cd ${xmft_st_subscription_ar.sub2.folder}
   ls
   put main.tf main.tf-$(date +'%Y%m%d_%H%M%S')
   ls
   bye
!

SSHPASS=${xmft_st_account.account3.user.password_credentials.password} sshpass -e sftp -P ${xmft_st_site_ssh.ssh2.port} -oBatchMode=no -b - ${xmft_st_account.account3.user.name}@${xmft_st_site_ssh.ssh2.host} << !
   ls
   bye
!
EOF
}
