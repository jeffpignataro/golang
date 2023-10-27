package main

import (
	"embed"
	"sort"
	"strings"

	"github.com/jeffpignataro/golang/pkg/log"
)

var (
	//go:embed names.txt
	res embed.FS

	logger *log.Logger
)

var alphabet = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func init() {
	logger = log.New()
	if logger.Error != nil {
		panic("Failed to initialize logger")
	}
	logger.Info("Logger initialized")
}

func main() {
	logger.Info("Starting Euler 22")

	f, err := res.ReadFile("names.txt")
	if err != nil {
		logger.Sugar().Fatalf("Failed to open file: %s", err)
	}

	totalValue := 0
	nameArray := strings.Split(string(f), ",")
	sort.Slice(nameArray, func(i, j int) bool {
		return nameArray[i] < nameArray[j]
	})
	logger.Sugar().Infof("Number of names: %d", len(nameArray))

	for i, name := range nameArray {
		nameValue := 0
		for _, letter := range name {
			if alphabet[string(letter)] == 0 {
				continue
			}
			nameValue += alphabet[string(letter)]
		}
		totalValue += nameValue * (i + 1)
		logger.Sugar().Debugf("%s - %d, At: %d, Name Value: %d, Total: %d", name, nameValue, i+1, nameValue*(i+1), totalValue)
	}
	logger.Sugar().Infof("Total value: %d", totalValue)
	logger.Info("Finished Euler 22")
}
