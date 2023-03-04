package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	thirteen "golang/pkg/euler/13"
	"strconv"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	f := thirteen.Run()
	s := strconv.FormatFloat(f, 'E', -1, 64)

	return s, nil
}

func main() {
	lambda.Start(HandleRequest)
}
