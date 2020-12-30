package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "Jim@beam.com",
			zipCode: 12345,
		},
	}
jim.print()
jim.updateName("Jam")
jim.print()
}

func (p person) print(){
  fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string){
p.firstName=newFirstName

}
