package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Max with 15 decimal places of precision: ")
	f := math.Max(0.000000000000001, 1.000000000000001)
	fmt.Println(f)
	fmt.Printf("Min with 15 decimal places of precision: ")
	f = math.Min(0.000000000000001, 1.000000000000001)
	fmt.Println(f)
	fmt.Println()

	fmt.Printf("Max with 16 decimal places of precision: ")
	f = math.Max(0.0000000000000001, 1.0000000000000001)
	fmt.Println(f) // IS THIS A BUG?
	fmt.Printf("Min with 16 decimal places of precision: ")
	f = math.Min(0.0000000000000001, 1.0000000000000001)
	fmt.Println(f)
	fmt.Println()

	fmt.Println("Max and min with over 300 decimal precision")
	max := math.MaxFloat64
	min := (0 - math.MaxFloat64) // equivalent of math.MinFloat64
	f = math.Max(max, min)
	fmt.Println(f)
	f = math.Min(max, min)
	fmt.Println(f)
}
