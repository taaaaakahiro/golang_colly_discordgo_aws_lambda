package command

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"golang-aws-lambda/src/pkg/handler"
	"os"
)

const (
	existError = 1
	exitOK     = 0
)

var (
	asins = []string{
		"B08HST559L",
		"B085CZST7T",
	}
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	if os.Getenv("AWS_EXECUTE_ENV") == "dev" {
		hook := os.Getenv("HOOK_ESC_KEY")
		for _, a := range asins {
			handler.Handler(hook, os.Getenv("AMAZON_URL")+a)
		}

	} else {
		lambda.Start(handler.Handler)
	}
	return exitOK
}
