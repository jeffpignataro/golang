package helloworldapi

import (
	"net/http"

	helloworldapi_endpoints "github.com/jeffpignataro/golang/pkg/examples/hello-world-api/endpoints"
)

func Run() {

	mux := http.NewServeMux()

	mux.HandleFunc("/welcome", helloworldapi_endpoints.Welcome)
	mux.HandleFunc("/posts", helloworldapi_endpoints.Posts)

	http.ListenAndServe(":5050", mux)
}
