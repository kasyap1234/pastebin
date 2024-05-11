package models

import "gopkg.in/mgo.v2/bson"

type Pastebin struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Content  string        `json:"content" bson:"content"`
	Language string        `json:"language" bson:"language"`
	Expires  int64         `json:"expires" bson:"expires"`
	Views    string        `json:"views" bson:"views"`
	Owner    string        `json:"owner" bson:"owner"`
	Password string        `json:"password" bson:"password"`
	LongURL      string        `json:"url" bson:"url"`
	ShortURL  string        `json:"shortURL" bson:"shortURL"`

}

