package main

import (
	"fmt"
	"os"
)

func main() {
	// => new Deck
	card := newDeck()

	// => to print
	card.print()

	// => to mix card
	card.shuffle()

	// => to save local
	err := card.saveToFile("my_cards")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// => to read from saved file
	newCard := newDeckFromFile("my_cards")

	//=> to print
	newCard.print()
}

/*
func main() {
	// => Define variable
	var card string = "Ace of Spades" => Alternative
	card := newCard()
	card = "Five of Diamonds" => We can change value, but we can't change type

	// => Slice - To create an array if its count can be grown or shrink
	After create new type deck we changed deck instead of [] string
	cards := deck{"Ace of Diamonds", newCard()}

	// => To add new member of array
	cards = append(cards, "Six of Spades")

	// => For loops => for index,card:=range cards{...}
	index:index card: current card we're iterating over - range cards: take the slice of 'cards' and loop over it
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// => To create deck ([]string)
	cards := newDeck()

	// => Deal method give two array first of this for handCard and second one for remainCard
	hand, remainingCards := deal(cards, 5)

	// => Print method shows on the terminal
	hand.print()
	remainingCards.print()

	// => []byte [72 105 32 116 104 101 114 101 33] Ascii code of "Hi there!"
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))
}
*/
