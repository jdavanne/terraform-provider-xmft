resource "xmft_st_account" "account1" {
  provider = xmft.st1
  name     = "account1"
  #type        = "user"
  uid         = "1000"
  gid         = "1000"
  home_folder = "/files/account1"
  user = {
    name = "login1"
    password_credentials = {
      password = "paswword123"
    }
  }
}

resource "tls_private_key" "rsa-1024-example" {
  algorithm = "RSA"
  rsa_bits  = 1024
}

resource "tls_self_signed_cert" "example" {
  private_key_pem = tls_private_key.rsa-1024-example.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "ACME Examples, Inc"
  }

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "xmft_st_certificate" "login-x509" {
  provider = xmft.st1
  name     = "cert1"
  account  = xmft_st_account.account1.name
  type     = "x509"
  usage    = "login"
  content  = tls_self_signed_cert.example.cert_pem
}

resource "tls_private_key" "login-key" {
  algorithm = "RSA"
}

resource "xmft_st_certificate" "login-ssh" {
  provider = xmft.st1
  name     = "` + name + `.ssh"
  account  = xmft_st_account.account1.name
  type     = "ssh"
  usage    = "login"

  content = tls_private_key.login-key.public_key_openssh
  //content            =  tls_private_key.login-key.public_key_pem
  /*content            = <<EOF
---- BEGIN SSH2 PUBLIC KEY ----
AAAAB3NzaC1yc2EAAAADAQABAAABAQCVlGUzWxNd0SAiV5uAINDWbTPF12CZLCAp
5YBhDWtWuSwQrMCy2T2A5SuM0CqdBGHNXWQarxtBIxALhqCDwRhaGEppN8BSEq2O
Wav5mZUQs6zeJthcfGX8/EgQk1HFY4lPFzVKwqa2qO2FSWYwu8sYFfqIGsBq0HJ7
vH32Il9ss9zhhBpfOy693MG8D2F2iYYMwLeQ0zvQBP6dn7BDgiKNXbLS/QbSx21c
d1TkAPTobAooz3XbwO/KAk1B706VTZRC+QSkin/FHwBppeTr2basV1yO3Yavx1P9
O5Kfmmsa7zZmRgFNCQnGRD/39gtoCVznIgbrHbXjpvY+MVbt5FUj
---- END SSH2 PUBLIC KEY ----
EOF*/
  subject         = "CN="
  validity_period = 365
}
