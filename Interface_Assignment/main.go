package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{height float64; base float64}
type square struct{sideLength float64}

func main() {

	t := triangle{1,5}
	s := square{2}

	printArea(t)
	printArea(s)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
func (t triangle) getArea() float64 {
	return t.base*t.height*0.5
}

func (s square) getArea() float64 {
	return s.sideLength*s.sideLength
}
