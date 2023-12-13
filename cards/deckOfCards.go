package main

import "fmt"

// (""like a class"")
// create a new type of "deck" which is a slice of string
type deckOfCards []string

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

// (""like a method"")
// receivers sets up methods on variables that we create. (d deck) is a receiver
func (dc deckOfCards) print() {
	for i, card := range dc {
		fmt.Printf("%2v: %17s\n", i, card)
	}
	fmt.Print("\n")
}
