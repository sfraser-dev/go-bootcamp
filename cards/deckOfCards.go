package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// (""like a class"")
// create a new data type (deckOfCards) which is a string slice
type deckOfCards []string

// (""like a method"")
// receivers sets up methods on variables that we create. (d deck) is a receiver
func (doc deckOfCards) printDebug(runIt bool) {
	if !runIt {
		return
	}
	fmt.Print("---------------------\n")
	for i, card := range doc {
		fmt.Printf("%02v: %17s\n", i+1, card)
	}
	fmt.Print("\n")
}

// (""like a method"")
func (doc deckOfCards) shuffleTheCards() {
	rand.Shuffle(len(doc), func(i, j int) {
		doc[i], doc[j] = doc[j], doc[i]
	})
}

// (""like a method"")
func (doc deckOfCards) deckOfCardsToSingleString() string {
	stringOut := strings.Join(doc, ",")
	return stringOut
}

// (""like a method"")
func (doc deckOfCards) writeToFile(filename string) {
	err := os.WriteFile(filename, []byte(doc.deckOfCardsToSingleString()), 0666)
	if err != nil {
		fmt.Println("Error writing", filename, ": ", err)
		os.Exit(1)
	}
}

// (""like a method"")
func (doc deckOfCards) readFromFile(filename string) []string {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading ", filename, ": ", err)
		os.Exit(1)
	}
	return strings.Split((string(byteSlice)), ",")
}

// helper funtion
func printDebug(runIt bool, s []string) {
	if !runIt {
		return
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

// helper function
func newDeckOfCards() deckOfCards {
	cards := deckOfCards{}
	suits := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	numbers := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			var current string = numbers[j] + " of " + suits[i]
			cards = append(cards, current)
		}
	}
	return cards
}

// helper function
func dealCards(d deckOfCards, handSize int) (deckOfCards, deckOfCards) {
	return d[:handSize], d[handSize:]
}
