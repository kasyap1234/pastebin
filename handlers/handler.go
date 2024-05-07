package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/pastebin/database"
	"github.com/kasyap1234/pastebin/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)
func CreatePaste(c * gin.Context){
	var pastebin models.Pastebin 

     if err  :=c.ShouldBindJSON(&pastebin); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
}
Collection :=database.GetMongoClient().Database("pastebin").Collection("pastes")

if err := database.InsertOne(Collection,pastebin); err !=nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

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
func deletePasteByID(c *gin.Context){
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
func updatePasteByID(c *gin.Context){
	
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
			{"url", updatePaste.URL},
		}},
	}
    err = database.UpdateOne(collection,filter,update)



	

	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "paste updated successfully"})

}

