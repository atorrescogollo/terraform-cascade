terraform {
  backend "local" {
    path = "/tmp/cascade/.terraform/ops_base.tfstate"
  }
}
