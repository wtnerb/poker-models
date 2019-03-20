package models

import (
	"fmt"
	"testing"
)

func TestStringifyCards(t *testing.T) {
	testTable := []struct {
		c        Card
		expected string
	}{
		{Card{two, spade}, "2♠"},
		{Card{ACE, heart}, "A♥"},
		{Card{king, club}, "K♣"},
		{Card{queen, diamond}, "Q♦"},
		{Card{jack, 0}, "J" + unknown + " suit"},
		{Card{0, spade}, unknown + " value♠"},
		{Card{ten, spade}, "10♠"},
	}

	for _, test := range testTable {
		s := fmt.Sprint(test.c)
		if s != test.expected {
			t.Error("Expected:", test.expected, ", recieved:", s)
		}
	}
}

func TestValidateCards(t *testing.T) {
	tests := []struct {
		card  string
		valid bool
	}{
		{"2♠", true},
		{"A♥", true},
		{"K♣", true},
		{"Q♦", true},
		{"10♠", true},
		{"J♠", true},
		{"♠", false},
		{"unknown1♠", false},
		{"1♠", false},
		{"3H", false},
		{"10", false},
		{"1K", false},
	}

	for _, v := range tests {
		if !(ValidCard(v.card) == v.valid) {
			t.Error("Unexpected result.", v.card, "was expected to be", v.valid, "and wasn't")
		}
	}
}
