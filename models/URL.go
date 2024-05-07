package models 
import ("time"

)
type URL struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`
	LongURL string `json:"longURL" bson:"longURL"`
	ShortURL string `json:"shortURL" bson:"shortURL"`
    CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
