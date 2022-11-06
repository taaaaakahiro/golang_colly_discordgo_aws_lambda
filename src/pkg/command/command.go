package command

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"golang-aws-lambda/src/pkg/handler"
	"log"
	"os"
)

const (
	existError = 1
	exitOK     = 0
)

func Run() {
	os.Exit(run(context.Background()))
	//run(context.Background())
}

func run(_ context.Context) int {
	awsExecuteEnv := os.Getenv("EXECUTE_ENV")

	if awsExecuteEnv == "dev" {
		handler.Handler()
		return exitOK
	} else if awsExecuteEnv == "prod" {
		lambda.Start(handler.Handler)
		return exitOK
	} else {
		log.Println("failed to load env")
		log.Fatal()
		return existError

	}

}
