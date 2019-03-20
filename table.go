package models

import "gopkg.in/mgo.v2/bson"

// TablePlayer is the information tracked by each player at the table
type TablePlayer struct {
	ID        bson.ObjectId `json:"_id"`
	Name      string        `json:"Name"`
	Cards     []Card        `json:"Cards"`
	Cash      float64       `json:"Cash"`
	AmountBet float64       `json:"AmountBet"`
	Folded    bool          `json:"folded"`
}

// Table is the structure that holds all the information about the state
// of the poker table
type Table struct {
	Players    []TablePlayer `json:"players"`
	FaceUp     []Card        `json:"faceUp"`
	Pot        float64       `json:"pot"`
	HighestBet float64       `json:"highestBet"`
	Turn       int           `json:"turn"`
	Dealer     int           `json:"dealer"`
	Deck       []Card        `json:"deck"`
}
