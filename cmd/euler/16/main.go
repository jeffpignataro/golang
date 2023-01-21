package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	n := 2.0
	n = math.Pow(2, 1000)
	fmt.Println(n)
	a := 0
	s := []rune(fmt.Sprintf("%f", n))
	for i := 0; i < len(s); i++ {
		b := string(s[i])
		c, _ := strconv.Atoi(b)
		a = c + a
		fmt.Println(a)
	}
}
