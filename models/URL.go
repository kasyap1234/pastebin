package models 

type URL struct {
	LongURL string `json:"longURL" bson:"longURL"`
	ShortURL string `json:"shortURL" bson:"shortURL"`

}
