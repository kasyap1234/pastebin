package models

type Pastebin struct {
    ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Content     string             `json:"content,omitempty" bson:"content,omitempty"`
    Language    string             `json:"language,omitempty" bson:"language,omitempty"`
    Expires     int64              `json:"expires,omitempty" bson:"expires,omitempty"`
    Views       string             `json:"views,omitempty" bson:"views,omitempty"`
    Owner       string             `json:"owner,omitempty" bson:"owner,omitempty"`
    Password    string             `json:"password,omitempty" bson:"password,omitempty"`
    LongURL     string             `json:"url,omitempty" bson:"url,omitempty"`
    ShortURL    string             `json:"shortURL,omitempty" bson:"shortURL,omitempty"`
}
