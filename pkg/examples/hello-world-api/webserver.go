package helloworldapi

import (
	"net/http"

	helloworldapi_endpoints "workspace/pkg/examples/hello-world-api/endpoints"
)

// Handler for http requests
type Handler struct {
	mux *http.ServeMux
}

func New(s *http.ServeMux) *Handler {
	h := Handler{s}
	s.HandleFunc("/welcome", helloworldapi_endpoints.Welcome)
	s.HandleFunc("/posts", helloworldapi_endpoints.Posts)
	s.HandleFunc("/healthz", helloworldapi_endpoints.Healthz)

	return &h
}
