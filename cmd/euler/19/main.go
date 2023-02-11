package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	layout := "2006/01/02"
	t1 := time.Date(1901, 1, 1, 12, 00, 00, 00, &time.Location{})
	t2 := time.Date(2000, 12, 31, 12, 00, 00, 00, &time.Location{})
	if len(os.Args) > 2 {
		t, err := time.Parse(layout, os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		t1 = t
		t, err = time.Parse(layout, os.Args[2])
		if err != nil {
			fmt.Println(err.Error())
		}
		t2 = t
	}
	between := t2.Sub(t1)
	days := int(between.Hours() / 24)

	counter := 0

	for i := 0; i < days; i++ {
		x, _ := time.ParseDuration(fmt.Sprintf("%dh", i*24))
		d := t1.Add(x)
		if int(d.Weekday()) == 0 && d.Day() == 1 {
			counter++
		}
	}
	fmt.Printf("%d Sundays which fell on the 1st of a month between the dates.\n", counter)
}
