package main

import (
	"context"
	"fmt"
	"os"

	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
	ctx, client, err := okta.NewClient(
		context.TODO(),
		okta.WithOrgUrl("https://dev-20308777.okta.com"),
		okta.WithToken(os.Getenv("OKTA_TOKEN")),
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Context: %+v\n Client: %+v\n", ctx, client)
}
