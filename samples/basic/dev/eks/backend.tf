terraform {
  backend "local" {
    path = "/tmp/.terraform/dev_eks.tfstate"
  }
}
