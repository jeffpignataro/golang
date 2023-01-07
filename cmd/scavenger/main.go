package main

import (
	"fmt"
	"os"

	"workspace/pkg/rest"
)

func main() {
	r, err := rest.Get("https://www.google.com", nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Print(string(r))
}
