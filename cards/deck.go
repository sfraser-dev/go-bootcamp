package main

import "fmt"

// create a new type of "deck" which is a slice of string

type deck []string
func printDeck(cards []string) {
	for i, card := range cards {
		fmt.Print(i, ": ", card, "\n")
	}
}

