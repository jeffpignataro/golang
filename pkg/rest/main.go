package rest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Call(url string, method string, body string) ([]byte, error) {
	c := http.Client{}

	var b io.Reader
	if body != "" {
		j := []byte(body)
		b = bytes.NewReader(j)
	}

	r, err := http.NewRequest(method, url, b)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	p, err := c.Do(r)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(p.Body)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return responseData, nil
}

func Get(url string) ([]byte, error) {
	return Call(url, http.MethodGet, "")
}

func Delete(url string) ([]byte, error) {
	return Call(url, http.MethodDelete, "")
}

func Post(url string, body string) ([]byte, error) {
	return Call(url, http.MethodPost, body)
}

func Patch(url string, body string) ([]byte, error) {
	return Call(url, http.MethodPatch, body)
}
