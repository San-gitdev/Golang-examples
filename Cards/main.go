package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Diamonds"}
	fmt.Println(cards)
	cards = append(cards, "Five of Spade")
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {

	return "Five of Diamonds"

}
