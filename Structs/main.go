package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
  var rajesh person
  rajesh.firstName="Rajesh"
	fmt.Println(alex)
  fmt.Printf("%+v", rajesh)


}