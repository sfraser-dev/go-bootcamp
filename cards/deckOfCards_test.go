package main

import (
	"regexp"
	"testing"
)

func Test_newDeckOfCards(t *testing.T) {
	d := newDeckOfCards()

	// number of cards generated
	if len(d) != 52 {
		t.Errorf("Expected fifty-two cards in a new deck but got %v", len(d))
	}

	// regex testing
	dSingelString := d.deckOfCardsToSingleString()
	// regex search for /of Spades/
	re := regexp.MustCompile(`\b(of){1}\b\s+S[a-z]{4}s`)
	found := re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected thirteen Spades in the deck but got %v", len(found))
	}
	// regex search for /of Hearts/
	re = regexp.MustCompile(`\b(of){1}\b\s+H[a-z]{4}s`)
	found = re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected thirteen Hearts in the deck but got %v", len(found))
	}
	// regex search for /of Diamonds/
	re = regexp.MustCompile(`\b(of){1}\b\s+D[a-z]{6}s`)
	found = re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected thirteen Diamonds in the deck but got %v", len(found))
	}
	// regex search for /of Clubs/
	re = regexp.MustCompile(`\b(of){1}\b\s+C[a-z]{3}s`)
	found = re.FindAllString(dSingelString, -1)
	if len(found) != 13 {
		t.Errorf("Expected thirteen Clubs in the deck but got %v", len(found))
	}

	// regex search for /Ace of/, /Two of/ ... /King of/, should be 4 of each
	var search string
	for i := 1; i <= 13; i++ {
		switch i {
		case 1:
			search = "Ace of"
		case 2:
			search = "Two of"
		case 3:
			search = "Three of"
		case 4:
			search = "Four of"
		case 5:
			search = "Five of"
		case 6:
			search = "Six of"
		case 7:
			search = "Seven of"
		case 8:
			search = "Eight of"
		case 9:
			search = "Nine of"
		case 10:
			search = "Ten of"
		case 11:
			search = "Jack of"
		case 12:
			search = "Queen of"
		case 13:
			search = "King of"
		default:
			t.Error("Error, switch statement shouldn't reach here")
		}
		re = regexp.MustCompile(search)
		found = re.FindAllString(dSingelString, -1)
		if len(found) != 4 {
			t.Errorf("Expected four '%s' in the deck but got %v", search, len(found))
		}
	}
}
