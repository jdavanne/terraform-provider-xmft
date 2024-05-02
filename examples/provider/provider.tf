terraform {
  required_providers {
    xmft = {
      source = "jdavanne/xmft"
    }
  }
}

provider "xmft" {
  product  = "st"
  alias    = "st1"
  host     = "https://host:8444"
  username = "admin"
  password = "admin*"
}

provider "xmft" {
  product  = "cft"
  alias    = "cft1"
  host     = "https://host:1768"
  username = "admin"
  password = "changeme"
}
