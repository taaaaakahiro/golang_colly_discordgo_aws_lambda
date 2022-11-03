pl:
	terraform plan

ap:
	terraform apply -input=false -auto-approve -lock=false

de:
	terraform destroy -input=false -auto-approve -lock=false

run:
	go run ./src/cmd/main.go

test:
	go test ./...