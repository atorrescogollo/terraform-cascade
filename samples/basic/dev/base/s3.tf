resource "local_file" "s3" {
  content  = "s3"
  filename = "${dirname(abspath(data.local_file.account.filename))}/s3/.s3"
}
