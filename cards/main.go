package main

import (
	"fmt"

	"bitbucket.org/weebucket/gobootcamp"
)

func main() {
	// type deckOfCards
	var cards deckOfCards = newDeckOfCards()

	var (
		hand           deckOfCards
		remainingCards deckOfCards
	)

	hand, remainingCards = dealCards(cards, 5)
	// assigning an unused variable to the blank identifier _ will silence the unused variable error
	var _ = hand
	var _ = remainingCards

	// save to file
	cards.writeToFile("mycards.log")
	cardRead := cards.readFromFile("mycards.log")
	var _ = cardRead

	// testing
	gobootcamp.Saytop()
	greetingStr := "Hello"
	byteSlice := ([]byte(greetingStr))
	fmt.Println(byteSlice)
	fmt.Println(string(byteSlice))

}
