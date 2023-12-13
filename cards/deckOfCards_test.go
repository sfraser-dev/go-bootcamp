package main

import (
	"regexp"
	"testing"
)

func Test_newDeckOfCards(t *testing.T) {
	d := newDeckOfCards()

	// number of cards generated
	if len(d) != 52 {
		t.Errorf("Expected 52 cards in a new deck but got %v", len(d))
	}

	// regex testing
	dSingelString := d.deckOfCardsToSingleString()
	// regex search for /of Spades/
	re := regexp.MustCompile(`\b(of){1}\b\s+S[a-z]{4}s`)
	found := re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected 13 Spades in the deck but got %v", len(found))
	}
	// regex search for /of Hearts/
	re = regexp.MustCompile(`\b(of){1}\b\s+H[a-z]{4}s`)
	found = re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected 13 Hearts in the deck but got %v", len(found))
	}
	// regex search for /of Diamonds/
	re = regexp.MustCompile(`\b(of){1}\b\s+D[a-z]{6}s`)
	found = re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected 13 Diamonds in the deck but got %v", len(found))
	}
	// regex search for /of Clubs/
	re = regexp.MustCompile(`\b(of){1}\b\s+C[a-z]{3}s`)
	found = re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected 13 Clubs in the deck but got %v", len(found))
	}

	// regex search for /Ace of/, /Two of/ ... /King of/, should be 4 of each
	var search string
	for i := 1; i <= 13; i++ {
		switch i {
		case 1:
			search = "Ace\\s+of"
		case 2:
			search = "Two\\s+of"
		case 3:
			search = "Three\\s+of"
		case 4:
			search = "Four\\s+of"
		case 5:
			search = "Five\\s+of"
		case 6:
			search = "Six\\s+of"
		case 7:
			search = "Seven\\s+of"
		case 8:
			search = "Eight\\s+of"
		case 9:
			search = "Nine\\s+of"
		case 10:
			search = "Ten\\s+of"
		case 11:
			search = "Jack\\s+of"
		case 12:
			search = "Queen\\s+of"
		case 13:
			search = "King\\s+of"
		}
		re = regexp.MustCompile(search)
		found = re.FindAllString(dSingelString, -1)
		if len(found) != 4 {
			t.Errorf("Expected 4 '%s' in the deck but got %v", search, len(found))
		}
	}
}
