package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}

func newCard() string {

	return "Five of Diamonds"

}
