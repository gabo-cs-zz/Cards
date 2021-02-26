package main

import (
	"os"
	"testing"
)

const expectedDeckLength = 16
const expectedFirstCard = "Ace of Spades"
const expectedLastCard = "Four of Clubs"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card of %v, but got %v", expectedFirstCard, d[0])
	}

	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card of %v, but got %v", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	_ = os.Remove("_decktesting")

	d := newDeck()
	_ = d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected %v cards in deck, but got %v", expectedDeckLength, len(loadedDeck))
	}

	_ = os.Remove("_decktesting")
}


