package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CountBits(1234))
}

func CountBits(n uint) int {
	s := strconv.FormatUint(uint64(n), 2)
	a := 0
	for i := 0; i < len(s); i++ {
		b, _ := strconv.Atoi(string(s[i]))
		a += b
	}
	return a
}
