terraform {
  backend "local" {
    path = "/tmp/.terraform/prod_base.tfstate"
  }
}
