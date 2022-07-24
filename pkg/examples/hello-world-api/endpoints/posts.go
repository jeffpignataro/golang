package helloworldapi_endpoints

import (
	"encoding/json"
	"net/http"
)

type post struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Text   string `json:"Text"`
}

func Posts(w http.ResponseWriter, r *http.Request) {

	posts := []post{
		{"Post one", "John", "This is first post."},
		{"Post two", "Jane", "This is second post."},
		{"Post three", "John", "This is another post."},
	}

	json.NewEncoder(w).Encode(posts)
}
