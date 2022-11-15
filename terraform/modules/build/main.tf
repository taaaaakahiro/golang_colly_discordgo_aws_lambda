resource "null_resource" "go_build" {
  triggers = {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o ../src/bin/main ../src/cmd && zip main.zip ../src/bin/main"
  }
}

data "archive_file" "go" {
  depends_on  = [null_resource.go_build]
  type        = "zip"
  source_file = "../src/cmd/main.go"
  output_path = "../src/bin/main.zip"
}