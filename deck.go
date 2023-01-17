package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// Which is a slice of strings
// Different from OOM we are creating new type which is inherited from [] string - actually equal
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// If we don't declare 'i' value we can say that with '_' - otherwise we will get error
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// => Any variable of type "deck" now gets access to the "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// => Return two deck - one of them for handCard and the other for remain
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// => To create []byte we need string rather than []string.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
