terraform {
  backend "local" {
    path = "/tmp/cascade/.terraform/base.tfstate"
  }
}
