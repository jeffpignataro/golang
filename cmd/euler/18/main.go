package main

import (
	"fmt"
	"math"
)

func main() {
	var arrays = [][]int{
		{75, 0},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 3, 34},
		{88, 2, 77, 73, 7, 63, 67},
		{99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
	}

	val := 0.0
	idx := 0
	for _, array := range arrays {
		i, v := findBiggestNumber(array, idx)
		val += float64(v)
		idx = i
	}

	fmt.Printf("Total %f\n", val)

	val = 0
	for idx := len(arrays) - 2; idx > 0; idx-- {
		current := arrays[idx+1]
		next := arrays[idx]
		for i := 0; i < len(next)-1; i++ {
			v := current[i]
			first := v + next[i]
			second := v + next[i+1]
			val += math.Max(float64(first), float64(second))
		}
	}

	fmt.Printf("Total %f\n", val)
}

func findBiggestNumber(a []int, idx int) (index int, value int) {
	if a[idx] > a[idx+1] {
		return idx, a[idx]
	} else {
		return idx + 1, a[idx+1]
	}
}
