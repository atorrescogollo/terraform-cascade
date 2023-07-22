resource "local_file" "dev" {
  content  = "dev"
  filename = "/tmp/cascade/dev/.account"
}

resource "local_file" "ops" {
  content  = "ops"
  filename = "/tmp/cascade/ops/.account"
}

resource "local_file" "prod" {
  content  = "prod"
  filename = "/tmp/cascade/prod/.account"
}
