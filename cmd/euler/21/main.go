package main

import (
	"github.com/jeffpignataro/golang/pkg/log"
)

var logger *log.Logger

func init() {
	logger = log.New()
	if logger.Error != nil {
		panic("Failed to initialize logger")
	}
	logger.Info("Logger initialized")
}

func main() {
	logger.Info("Starting Euler 21")
	total := 0
	pairs := []int{}
	for i := 1; i < 10000; i++ {
		skip := false
		for _, v := range pairs {
			if v == i {
				// Skip the other half of the pair
				skip = true
			}
		}
		if skip {
			continue
		}
		divisors := findDivisors(i)
		sum := 0
		for _, v := range divisors {
			sum += v
		}
		pair := findDivisors(sum)
		sum2 := 0
		for _, v := range pair {
			sum2 += v
		}
		if sum2 == i && sum != i {
			logger.Sugar().Infof("Amicable Pair i: %d, sum: %d, sum2: %d", i, sum, sum2)
			pairs = append(pairs, sum)
			total += i
			total += sum
		}
	}
	logger.Sugar().Infof("Total: %d", total)
	logger.Info("Finished Euler 21")
}

func findDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
