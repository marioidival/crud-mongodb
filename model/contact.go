package model

import "github.com/globalsign/mgo/bson"

// Address model to address ...
type Address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	Number  int    `bson:"number" json:"number"`
	Country string `bson:"country" json:"country"`
}

// Contact model to contact ...
type Contact struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"name" json:"name"`
	Phone   string        `bson:"phone" json:"phone"`
	Email   string        `bson:"email" json:"email"`
	Adrress Address       `bson:"address" json:"address,omitempty"`
}
