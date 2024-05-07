package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)
type URL struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	LongURL string `json:"longURL" bson:"longURL"`
	ShortURL string `json:"shortURL" bson:"shortURL"`
    CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
