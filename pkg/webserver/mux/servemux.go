package mux

import (
	"net/http"

	"workspace/pkg/webserver/routing"
)

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(routes []routing.Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}
