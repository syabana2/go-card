package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}


	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First Card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Cubes" {
		t.Errorf("Expected Four Card of Four of Cubes, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {

	d := newDeck()
	d.saveToFile("_decktesting")

	loadFile := newDeckFromFile("_decktesting")

	if len(loadFile) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadFile))
	}

	os.Remove("_decktesting")
}