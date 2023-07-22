data "local_file" "account" {
  filename = "/tmp/cascade/ops/.account"
}

resource "local_file" "vpc" {
  content  = "vpc"
  filename = "${dirname(abspath(data.local_file.account.filename))}/vpc/.vpc"
}
