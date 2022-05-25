package main

func main() {
	//cards := newDeck()
	//cards.print()
	//fmt.Println(cards.toString())

	/* hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print() */

	//cards.saveToFile("my_cards")
	cards := readDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
