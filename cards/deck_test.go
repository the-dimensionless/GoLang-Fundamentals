package main

import (
	"os"
	"testing"
)

const filename string = "my_cards"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(d))
	}

	if d[0] != "Joker" {
		t.Errorf("Expected Joker but got %v", d[0])
	}

}

func TestSaveToFile(t *testing.T) {
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := readDeckFromFile(filename)

	if len(loadedDeck) != 53 {
		t.Errorf("Expected 53 cards in deck but got %v", len(loadedDeck))
	}

	os.Remove(filename)

}
