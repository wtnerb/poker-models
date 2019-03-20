package models

import "regexp"

// suit is the Suit of a Card
type suit int8

// value is the Value of a Card
type value int8

// Card is a playing card
type Card struct {
	Value value `bson:"Value"`
	Suit  suit  `bson:"Suit"`
}

// Allowed values of cards are constants
const (
	two value = iota + 2
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ACE
)

// Allowed Suits a Card can have are constants
const (
	spade suit = iota + 1
	heart
	club
	diamond
)

// this string flags a suit or
const unknown = "unknown"

// This renders a card value to the characters on the card
func (v value) str() string {
	name := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	min := int8(two)
	max := int8(ACE)
	val := int8(v)
	if val < min || val > max {
		return unknown + " value"
	}
	return name[val-2]
}

// This renders the suit to a unicode depiction of the card's suit
func (s suit) str() string {
	suits := []string{"♠", "♥", "♣", "♦"}
	min := int8(spade)
	max := int8(diamond)
	val := int8(s)
	if val < min || val > max {
		return unknown + " suit"
	}
	return suits[val-1]
}

// String renders a card to a string, eg K♣
// Implements the Stringer interface from fmt
func (c Card) String() string {
	return c.Value.str() + c.Suit.str()
}

// ValidCard tests a string that might be produced by the
// card.String() method and checks if the result is a valid card.
func ValidCard(s string) bool {
	ok, err := regexp.MatchString("^((10)|[2-9AKQJ])[♠♥♣♦]$", s)
	if err != nil {
		panic("card cannot check for its own validity")
	}
	return ok
}

// NewCard Creates a new card given a value and a suit
func NewCard(v int8, s int8) Card {
	return Card{value(v), suit(s)}
}
