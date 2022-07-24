package helloworldapi_endpoints

import (
	"io"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "I'm alive!!")
}
