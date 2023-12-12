package main

func main() {
	var cards = deck{"Ace of Diamonds", newCard()}
	// non-mutable
	cards = append(cards, "Six of Diamonds")

	printDeck(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
