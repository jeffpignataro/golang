package main

import (
	"fmt"
	"math"
)

func main() {
	abundant := make([]int, 0)
	for i := 1; i < 28123; i++ {
		if isAbundant(i) {
			abundant = append(abundant, i)
		}
	}

	sum := 0
	for i := 1; i < 28123; i++ {
		if !isSumOfTwoAbundant(i, abundant) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isSumOfTwoAbundant(n int, abundant []int) bool {
	for _, a := range abundant {
		if a > n {
			return false
		}
		if isAbundant(n - a) {
			return true
		}
	}
	return false
}

func isAbundant(n int) bool {
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum > n
}
