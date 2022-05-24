package main

import "fmt"

// Create a new type of deck which is a slice of strings
type deck []string

// Create a cards slice/array of type deck
func newDeck() deck {
	cards := deck{
		"Joker",
	}

	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// fx with Receiver to print cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// fx with arguments & multiple return types
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
