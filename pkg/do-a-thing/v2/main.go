package doathing

import "fmt"

func Doathing(s string) string {
	fmt.Print("This is a breaking change!!\n")
	return fmt.Sprintf("Doing a thing...%s", s)
}
