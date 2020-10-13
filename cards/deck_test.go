package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length 52, got %v", len(d))
	}
}

func TestItems(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Two of Clubs" {
		t.Errorf("Expected last card Two of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()
	d.saveToFile("_decktesting.txt")

	loadedDeck := readFromFile("_decktesting.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length 52, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")

}
