package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 123
	fmt.Println(reverse(a))

}

func reverse(x int) int {//	y := 0
	s := strconv.Itoa(x)
	NumOfDigits := len(s)


	return NumOfDigits
}
