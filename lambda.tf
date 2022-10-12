resource "null_resource" "go_build" {
  triggers = {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "cd src && GOOS=linux GOARCH=amd64 go build -o main main.go && zip main.zip main"
  }
}

data "archive_file" "go" {
  depends_on  = [null_resource.go_build]
  type        = "zip"
  source_file = "src/main.go"
  output_path = "src/main.zip"
}