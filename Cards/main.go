package main

//import "fmt"

func main() {
	cards:=newDeckFromFile("my_cards")
cards.print()
}

func newCard() string {

	return "Five of Diamonds"

}
