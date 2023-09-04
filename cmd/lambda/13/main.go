package main

import (
	"context"

	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	thirteen "github.com/jeffpignataro/golang/pkg/euler/13"
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
