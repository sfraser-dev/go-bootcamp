package main

import (
	"bitbucket.org/weebucket/gobootcamp"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// type deckOfCards
	var cards deckOfCards = newDeckOfCards()

	var (
		currentHand           deckOfCards
		currentRemainingCards deckOfCards
	)

	currentHand, currentRemainingCards = dealCards(cards, 5)

	// write to file and read from file
	cards.writeToFile("mycards.log")
	cardsReadFromFile := cards.readFromFile("mycards.log")

	// testing
	gobootcamp.SayTopLevel()
	greetingStr := "Hello"
	byteSlice := ([]byte(greetingStr))
	fmt.Println(byteSlice)
	fmt.Println(string(byteSlice))
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)

	// assigning an unused variable to the blank identifier _ will silence the unused variable error
	_, _, _ = currentHand, currentRemainingCards, cardsReadFromFile
	printDebug(false, []string{""})
	cards.printDebug(false)
	cards.shuffleTheCards()
	cards.printDebug(true)
}
