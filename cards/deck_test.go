package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	print("Length is ", len(d))
	if len(d) != 53 {
		t.Errorf("Expected length of 53, but got %v", len(d))
	}

}
