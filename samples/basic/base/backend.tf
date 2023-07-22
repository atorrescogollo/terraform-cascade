terraform {
  backend "local" {
    path = "/tmp/.terraform/base.tfstate"
  }
}
