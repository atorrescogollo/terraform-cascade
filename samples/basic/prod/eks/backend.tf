terraform {
  backend "local" {
    path = "/tmp/.terraform/prod_eks.tfstate"
  }
}
