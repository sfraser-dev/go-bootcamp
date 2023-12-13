package main

import "bitbucket.org/weebucket/gobootcamp"

func main() {
	var cards deckOfCards = newDeckOfCards()
	cards.print()

	gobootcamp.Saytop()
}