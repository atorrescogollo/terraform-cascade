data "local_file" "account" {
  filename = "/tmp/cascade/dev/.account"
}

data "local_file" "vpc" {
  filename = "${dirname(abspath(data.local_file.account.filename))}/vpc/.vpc"
}

resource "local_file" "eks" {
  content  = "eks"
  filename = "${dirname(abspath(data.local_file.vpc.filename))}/eks/.eks"
}
