package main

import (
	"fmt"
)

func main() {
	/*
		cards := newDeck()

		//Deal method give two array first of this for hand and second one for remaining Cards
		hand, remainingCards := deal(cards, 5)

		hand.print()
		remainingCards.print()
		//=> []byte [72 105 32 116 104 101 114 101 33] Ascii code of "Hi there!"
		greeting := "Hi there!"
		fmt.Println([]byte(greeting))

	*/

	//=> To write hand and remain cards, we used to 'io-util' library with func WriteFile(...)
	cards := newDeck()
	fmt.Println(cards.toString())

	hand, remainingCards := deal(cards, 5)
	fmt.Println(hand)
	fmt.Println(remainingCards)
}

/*
func main() {
	var card string = "Ace of Spades" => Alternative
	card := newCard()
	card = "Five of Diamonds" => We can change value

	=>Slice - To create an array if its count can grow or shrink
	After create new type deck we changed deck instead of [] string
	cards := deck{"Ace of Diamonds", newCard()}

	=> To add new member of array
	cards = append(cards, "Six of Spades")

	=>For loops => for index,card:=range cards{...}
	index:index card: current card we're iterating over - range cards: take the slice of 'cards' and loop over it
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func => function - newCard() => Define action with method name - string => return value of type
func newCard() string {
	return "Five of Diamonds"
}
*/
