package main

import (
	"fmt"
	"time"
)

var slice = make([]int, 400)

func fibo(n int) int {
	first, second := 0, 0

	if n <= 1 {
		return n
	}

	if slice[n-1] != -1 {
		first = slice[n-1]
	} else {
		first = fibo(n - 1)
	}

	if slice[n-2] != -1 {
		second = slice[n-2]
	} else {
		second = fibo(n - 2)
	}
	slice[n] = first + second
	return slice[n]
}

func main() {

	for i := range slice {
		slice[i] = -1
	}

	fmt.Println("Hello")
	begin := time.Now().UnixNano()

	x := fibo(19)
	end := time.Now().UnixNano()

	fmt.Println(x, (end-begin)/int64(time.Microsecond))

}
