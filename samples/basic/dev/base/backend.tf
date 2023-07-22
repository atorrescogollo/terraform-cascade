terraform {
  backend "local" {
    path = "/tmp/.terraform/dev_base.tfstate"
  }
}
