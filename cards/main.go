package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	// non-mutable
	cards = append(cards, "Six of Spades")

	var cards2 deck = newDeck()

	cards.print()
	cards2.print()
}

func newCard() string {
	return "Five of Diamonds"
}
