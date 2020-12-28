package main


func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Five of Spade")
	cards.print()
}

func newCard() string {

	return "Five of Diamonds"

}
