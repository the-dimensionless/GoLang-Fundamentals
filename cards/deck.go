package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// fx with receiver to turn deck into String
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	// 0666 -> anyone can read & write this file
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log error & return newDeck or
		// log error & quit program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(int64(time.Now().UnixNano())) // Create new source/seed
	r := rand.New(source)

	for i := range d {
		newIndex := r.Intn(len(d) - 1)
		// one line swapping
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
