package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	c := big.NewInt(1)
	i := big.NewInt(100)
	limit := big.NewInt(0)

	for i.Cmp(limit) > 0 {
		c = c.Mul(c, i)
		i.Add(i, big.NewInt(-1))
	}
	s := fmt.Sprintf("%s", c)
	total := 0
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		total = total + num
	}
	println(total)

}
