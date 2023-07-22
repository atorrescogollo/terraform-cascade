terraform {
  backend "local" {
    path = "/tmp/.terraform/ops_base.tfstate"
  }
}
