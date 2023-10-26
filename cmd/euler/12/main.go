package main

import (
	"math"
	"os"

	"github.com/jeffpignataro/golang/pkg/log"
	"go.uber.org/zap"
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
	sum := 0
	for i := 1; i <= 50000000; i++ {
		sum += i
		findDivisors(sum)
	}
	println(sum)
}

func findDivisors(i int) {
	divisor_counter := 0
	prev_divisor := i / 2
	for d := 1; d <= prev_divisor; d++ {
		if math.Mod(float64(i), float64(d)) == 0 {
			divisor_counter = divisor_counter + 2
			if divisor_counter > 500 {
				logger.Info("Found it!", zap.Int("i", i), zap.Int("d", d))
				os.Exit(1)
			}
		}
		prev_divisor = (i / d)
	}
}
