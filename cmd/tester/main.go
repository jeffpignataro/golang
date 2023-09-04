package main

import (
	"fmt"
	"net/http"

	"github.com/jeffpignataro/golang/pkg/rest"
)

func main() {
	r, _ := rest.Call("https://www.google.com", http.MethodGet, nil, nil, nil)
	fmt.Print(r)
}
