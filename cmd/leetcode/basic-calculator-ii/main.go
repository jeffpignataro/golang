package main

import (
	"fmt"
	"strings"
)

func main() {
	a := calculate("3+2*2")
	fmt.Println(a)
}

func calculate(s string) []string {
	operators := []string{"*", "/", "+", "-"}

	var str []string
	slices := []string{s}
	for _, operator := range operators {
		slices = strings.Split(slices[0], operator)
		if len(slices) > 1 {
			str = append(str, operator)
			str = append(str, slices...)
		}
	}
	return str
}
