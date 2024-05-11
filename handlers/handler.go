package handlers

import (
	"net/http"
    //  "fmt"
	// "time"
	"time"
	"github.com/kasyap1234/pastebin/internal"
	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/pastebin/database"
	"github.com/kasyap1234/pastebin/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func CreatePaste(c * gin.Context){
	var pastebin models.Pastebin 

     if err  :=c.ShouldBindJSON(&pastebin); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body "})
		return 
}
Collection :=database.GetMongoClient().Database("pastebin").Collection("pastes")

if err := database.InsertOne(Collection,pastebin); err !=nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
pastebin.LongURL=internal.CreateLongURL(); 
var shortURL string 
pastebin.ShortURL =shortURL; 
shortURL,err :=internal.GenerateShortURL(pastebin.LongURL); 
if err !=nil{
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}


pastebin.Expires = time.Now().Add(7 * 24 * time.Hour).Unix()
database.UpdateOne(Collection,bson.M{"ID":pastebin.ID},bson.M{"$set":bson.M{"shortURL":shortURL,"expires":pastebin.Expires}})

c.JSON(http.StatusOK,pastebin)

}
func GetPasteByID(c *gin.Context){
	pasteID :=c.Param("ID")
	objID,err :=primitive.ObjectIDFromHex(pasteID); 
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	collection :=database.GetMongoClient().Database("pastebin").Collection("pastes")
	filter :=bson.D{{"ID",objID}}
	paste,err := database.FindOneByID(collection,filter); 
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return 
	}
	c.JSON(http.StatusOK,paste)


}
func DeletePasteByID(c *gin.Context){
	pasteID :=c.Param("ID")
	objID,err :=primitive.ObjectIDFromHex(pasteID);
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	collection :=database.GetMongoClient().Database("pastebin").Collection("pastes")
	filter :=bson.D{{"ID",objID}}
	err =database.DeleteOne(collection,filter); 
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "paste deleted successfully"})
}
func UpdatePasteByID(c *gin.Context){
	
	pasteID :=c.Param("ID")
	objID,err :=primitive.ObjectIDFromHex(pasteID);
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	var updatePaste models.Pastebin
	if err :=c.ShouldBindJSON(&updatePaste); err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	filter :=bson.D{{"ID",objID}}
	collection :=database.GetMongoClient().Database("pastebin").Collection("pastes")
	update := bson.D{
		{"$set", bson.D{
			{"content", updatePaste.Content},
			{"language", updatePaste.Language},
			{"expires", updatePaste.Expires},
			{"views", updatePaste.Views},
			{"owner", updatePaste.Owner},
			{"password", updatePaste.Password},
			{"url", updatePaste.LongURL},
		}},
	}
    err = database.UpdateOne(collection,filter,update)

 

	

	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "paste updated successfully"})

}

