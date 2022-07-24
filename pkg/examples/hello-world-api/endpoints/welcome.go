package helloworldapi_endpoints

import (
	"io"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "Welcome!")
}
