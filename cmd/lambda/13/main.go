package main

import (
	"context"

	thirteen "workspace/pkg/euler/13"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return thirteen.Run()
}

func main() {
	lambda.Start(HandleRequest)
}
