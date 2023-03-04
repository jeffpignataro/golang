package main

import (
	"fmt"
	"golang/pkg/okta"
	"os"
)

func main() {
	ctx, client := okta.NewClient(os.Getenv("OKTA_TOKEN"))

	fmt.Print(ctx)
	fmt.Print(client)

}
