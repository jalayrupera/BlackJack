package model

import "testing"

func TestNewDeck(t *testing.T){
	deck := New_deck("")

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 16, but got length of  %d", len(deck))
	}
}