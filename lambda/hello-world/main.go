package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	helloworld "github.com/jeffpignataro/golang/pkg/hello-world"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return helloworld.Hello(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
