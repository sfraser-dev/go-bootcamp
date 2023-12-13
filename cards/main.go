package main

import (
	"fmt"

	"bitbucket.org/weebucket/gobootcamp"
)

func main() {
	// of type deckOfCards
	cards := newDeckOfCards()
	// cards.print()
	
	// hand, remainingCards := dealCards(cards, 5)
	// hand.print()
	// remainingCards.print()

	// save to file
	cards.writeToFile("mycards.log")
	strRead:= cards.readFromFile("mycards.log")
	fmt.Println(strRead)

	// testing
	gobootcamp.Saytop()
	greetingStr := "Hello"
	byteSlice := ([]byte(greetingStr))
	fmt.Println(byteSlice)
	fmt.Println(string(byteSlice))

}
