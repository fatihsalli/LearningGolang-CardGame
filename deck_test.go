package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	// To check length of deck
	if len(d) != 16 {
		// => %v filling with len(d) value
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	// To check first value
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades , but got %v", d[0])
	}
	// To check last value
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs , but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
