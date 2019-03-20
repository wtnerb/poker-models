package models

import "gopkg.in/mgo.v2/bson"

// Player represents what a player looks like in the database
type Player struct {
	ID     bson.ObjectId `json:"_id"`
	Played int           `json:"Games_Played"`
	Wins   int           `json:"Games_Won"`
	Money  float64       `json:"Cash_Available"`
	Name   string        `json:"Name"`
}
