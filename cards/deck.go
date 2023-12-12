package main

import "fmt"

// create a new type of "deck" which is a slice of string (""like a class"")
type deck []string

func newDeck() deck {
	cards := deck{}
	suits := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	numbers := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			var current string = suits[i] + numbers[j]
			cards = append(cards, current)
		}
	}
	return cards
}

// RECEIVERS SETS UP METHODS ON VARIABLES THAT WE CREATE
// (""like a method"")
// (d deck) is a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Print(i, ": ", card, "\n")
	}
}
