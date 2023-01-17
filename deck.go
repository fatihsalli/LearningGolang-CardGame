package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is a slice of strings
// Different from OOM we are creating new type which is inherited from [] string - actually equal
type deck []string

// This function gives us deck value as []string with mixing cardSuits and cardValues
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// If we don't want to declare 'i' value we can say that with '_' | Otherwise we will get error
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

// => Return two deck - one of them for handCard and the other for remainCard
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// => To create []byte we need string. We can't convert from []string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// => To save file (Package has changed from 'ioutil' to 'os')
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// => To read from file
// 'nil' value with go using like null or no value
func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program (Chosen)
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// string(data) => Ace of Spades,Two of Spades,Three of Spades ...
	s := strings.Split(string(data), ",")
	return deck(s)
}

// => To mix card (To Shuffle)
func (d deck) shuffle() {
	// Every time we need to new combination of number, so we have to change source to create different rand numbers
	// We use time.Now().UnixNano() to create int-64 value.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//For swap or change two values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
