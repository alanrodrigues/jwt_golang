package models

import "gopkg.in/mgo.v2/bson"

type (
	Offer struct {
		ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name      string        `json:"name" bson:"name"`
	}
)
