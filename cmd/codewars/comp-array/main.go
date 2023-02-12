package main

import (
	"fmt"
	"math"
)

func dotest(array1 []int, array2 []int, exp bool) {
	var ans = Comp(array1, array2)
	if ans == exp {
		fmt.Println("Correct")
	} else {
		fmt.Println("Wrong")
	}
}

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	dotest(a1, a2, true)
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	dotest(a1, a2, false)
	a1 = nil
	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	dotest(a1, a2, false)
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	dotest(a1, a2, true)
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{231, 14641, 20736, 361, 25921, 361, 20736, 361}
	dotest(a1, a2, false)
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}
	dotest(a1, a2, false)
	a1 = []int{}
	a2 = []int{}
	dotest(a1, a2, true)
	a1 = []int{}
	a2 = []int{25, 49}
	dotest(a1, a2, false)
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11, 1008}
	a2 = []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}
	dotest(a1, a2, false)
	a1 = []int{-10000000, 100000000}
	a2 = []int{100000000000000, 10000000000000000}
	dotest(a1, a2, false)
}

func Comp(array1 []int, array2 []int) bool {
	test := make([]bool, len(array2))
	for i, w := range array2 {
		for _, v := range array1 {
			n := math.Sqrt(float64(w))
			if float64(v) == n {
				test[i] = true
				break
			} else {
				test[i] = false
			}
		}
	}
	r := true
	for _, a := range test {
		if !a {
			r = a
			break
		}
	}
	return r
}
