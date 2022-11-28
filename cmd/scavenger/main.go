package main

import (
	"fmt"
	"os"

	"github.com/jeffpignataro/golang/pkg/rest"
)

func main() {
	r, err := rest.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Print(string(r))
}
