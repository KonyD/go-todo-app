package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Todo struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool          `json:"completed"`
	Body      string        `json:"body"`
}
