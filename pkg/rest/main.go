package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Authentication struct {
	Authenticate bool
	Type         *string
	Key          *string
}

func NoAuth() Authentication {
	return Authentication{
		Authenticate: false,
	}
}

func NewAuthentication(a bool, t *string, k *string) Authentication {
	return Authentication{
		Authenticate: a,
		Type:         t,
		Key:          k,
	}
}

func Call(url string, method string, body *string, auth *Authentication, headers *map[string]string) ([]byte, error) {
	c := http.Client{}

	var reader io.Reader
	if body != nil {
		j := []byte(*body)
		reader = bytes.NewReader(j)
	}

	r, err := http.NewRequest(method, url, reader)

	if auth != nil && auth.Authenticate {
		r.Header.Add("Authorization", fmt.Sprintf("%s %s", *auth.Type, *auth.Key))
	}

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	if headers != nil {
		for k, v := range *headers {
			r.Header.Add(k, v)
		}
	}

	p, err := c.Do(r)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	responseData, err := io.ReadAll(p.Body)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return responseData, nil
}

func Get(url string, auth *Authentication) ([]byte, error) {
	return Call(url, http.MethodGet, nil, auth, nil)
}

func Delete(url string, auth *Authentication) ([]byte, error) {
	return Call(url, http.MethodDelete, nil, auth, nil)
}

func Post(url string, body string, auth *Authentication) ([]byte, error) {
	return Call(url, http.MethodPost, &body, auth, nil)
}

func Patch(url string, body string, auth *Authentication) ([]byte, error) {
	return Call(url, http.MethodPatch, &body, auth, nil)
}

func GetWithHeaders(url string, auth *Authentication, headers *map[string]string) ([]byte, error) {
	return Call(url, http.MethodGet, nil, auth, headers)
}

func DeleteWithHeaders(url string, auth *Authentication, headers *map[string]string) ([]byte, error) {
	return Call(url, http.MethodDelete, nil, auth, headers)
}

func PostWithHeaders(url string, body string, auth *Authentication, headers *map[string]string) ([]byte, error) {
	return Call(url, http.MethodPost, &body, auth, headers)
}

func PatchWithHeaders(url string, body string, auth *Authentication, headers *map[string]string) ([]byte, error) {
	return Call(url, http.MethodPatch, &body, auth, headers)
}
