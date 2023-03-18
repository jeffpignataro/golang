package main

import (
	"fmt"
	"golang/pkg/rest"
	"net/http"
)

func main() {
	r, _ := rest.Call("https://www.google.com", http.MethodGet, nil, nil, nil)
	fmt.Print(r)
}
