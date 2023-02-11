package euler

import (
	"encoding/json"
	"fmt"
	"workspace/pkg/rest"
)

type Answer struct {
	Status int         `json:"status"`
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
}

func Submit(question int, answer string) (s *string, err error) {
	r, err := rest.Get(fmt.Sprintf("http://www.projecteulerapi.com/api/problem/%d/check/%s", question, answer), nil)
	if err != nil {
		return nil, err
	}
	a := Answer{}
	json.Unmarshal(r, &a)
	return &a.Result, nil
}
