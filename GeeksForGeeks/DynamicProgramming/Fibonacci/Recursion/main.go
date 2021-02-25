package main

import (
	"fmt"
	"time"
)

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {

	fmt.Println("Hello")
	begin := time.Now().UnixNano()

	x := fibo(4)
	end := time.Now().UnixNano()

	fmt.Println(x)
	fmt.Println((end - begin) / int64(time.Microsecond))

}
