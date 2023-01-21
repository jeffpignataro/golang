package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Builder struct {
	NameArray map[int]string
	Words     []string
}

func main() {

	nameArray := map[int]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}
	counter := 0
	for i := 1; i <= 1000; i++ {
		s := strings.Split(strconv.Itoa(i), "")
		if len(s) == 1 {
			a, _ := strconv.Atoi(s[0])
			fmt.Print(nameArray[a])
			counter += len(nameArray[a])
			fmt.Printf(" %d\n", counter)
			continue
		}

		var words []string
		for i := 0; i < len(s); i++ {
			words = append(words, string(s[i]))
		}

		builder := Builder{
			NameArray: nameArray,
			Words:     words,
		}
		for index, v := range words {
			p := len(words) - index
			n, _ := strconv.Atoi(v)
			result := BuildName(n, p, builder)
			counter += len(strings.ReplaceAll(result, " ", ""))
			fmt.Print(strings.ReplaceAll(result, " ", ""))
		}
		fmt.Printf(" %d\n", counter)
	}
}

func BuildName(n int, p int, builder Builder) string {
	if n == 0 {

		return ""
	}
	if p == 2 && n == 1 {
		// DEAL WITH TEENS
		w, _ := strconv.Atoi(builder.Words[len(builder.Words)-1])
		builder.Words[len(builder.Words)-1] = strconv.Itoa(w + 10)
		if (len(builder.Words)) == 2 {
			return ""
		}
		return "and"
	}
	result := builder.NameArray[n]
	if p == 4 {
		result = result + " thousand"
	} else if p == 3 {
		result = result + " hundred"
	} else if p == 2 || p == 5 {
		result = builder.NameArray[n*10]
		if p == 2 && len(builder.Words) != 2 {
			result = "and " + result
		}
	} else if p == 1 && len(builder.Words) > 2 && builder.Words[1] == "0" {
		result = "and " + result
	} else {
		result = builder.NameArray[n]
	}
	return result
}
