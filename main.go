package main

func main() {
	//var card string = "Ace of Spades" => Alternative
	//card := newCard()
	//card = "Five of Diamonds" => We can change value

	//=>Slice - To create an array if its count can grow or shrink
	// After create new type deck we changed deck instead of [] string
	//cards := deck{"Ace of Diamonds", newCard()}
	//=> To add new member of array
	//cards = append(cards, "Six of Spades")

	//=>For loops => for index,card:=range cards{...}
	//index:index card: current card we're iterating over - range cards: take the slice of 'cards' and loop over it
	/*for i, card := range cards {
		fmt.Println(i, card)
	}*/
	//cards.print()

	cards := newDeck()

	cards.print()

}

// func => function - newCard() => Define action with method name - string => return value of type
/*func newCard() string {
	return "Five of Diamonds"
}*/
