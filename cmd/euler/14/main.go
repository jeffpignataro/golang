package main

import (
	"fmt"
)

func main() {
	high := 0
	for i := 0; i < 1000000; i++ {
		a := i
		n := 1
		for a > 1 {
			if a%2 == 0 {
				a = Even(a)
			} else {
				a = Odd(a)
			}
			n++
		}
		if n > high {
			high = n
			fmt.Println(i)
		}
	}
	fmt.Printf("Highest iterations %d\n", high)
}

func Even(n int) int {
	return (n / 2)
}

func Odd(n int) int {
	return (3 * n) + 1
}
