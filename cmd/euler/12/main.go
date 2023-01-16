package main

import (
	"fmt"
	"math"
)

func main() {
	count := 1000000
	for i := 1; i < count; i++ {
		divisor_counter := 0
		for d := 1; d < ((count / 2) - 1); d++ {
			if math.Mod(float64(i), float64(d)) == 0 {
				divisor_counter++
				if divisor_counter > 500 {
					fmt.Printf("%d %d divisor\n", d, i)
				}
			}
		}
	}
}
