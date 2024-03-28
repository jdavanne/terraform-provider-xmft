terraform {
  required_providers {
    axwaycft = {
      source = "hashicorp.com/edu/axway-cft"
      //source = "https://ci8.jda.axwaytest.net:1768"
    }
  }
}

provider "axwaycft" {
  host     = "https://ci8.jda.axwaytest.net:1768"
  username = "admin"
  password = "changeit"
}

data "axwaycft_about" "example" {}

output "about" {
  value = data.axwaycft_about.example
}

resource "axwaycft_cftsend" "zoupi" {
  name    = "ZOUP"
  exec    = ""
  fcode   = "ASCII"
  faction = "NONE"
  parm    = ""
  preexec = ""
  fname   = "/tmpfile2"
}

output "zoupi" {
  value = axwaycft_cftsend.zoupi
}

