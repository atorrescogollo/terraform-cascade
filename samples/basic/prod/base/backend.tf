terraform {
  backend "local" {
    path = "/tmp/cascade/.terraform/prod_base.tfstate"
  }
}
