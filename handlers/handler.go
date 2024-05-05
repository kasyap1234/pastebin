package handlers 

import (
"github.com/kasyap1234/pastebin/database"
"github.com/kasyap1234/pastebin/models"

)
func CreatePaste(c *gin.Context){
	var pastebin models.Pastebin 

     if err: =c.ShouldBindJSON(&pastebin); err !=nil {
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
	collection :=database.GetMongoClient().DAtabase("pastebin").Collection("pastes")
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
	var pastebin models.Pastebin
	if err :=c.ShouldBindJSON(&pastebin); err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	filter :=bson.D{{"ID",objID}}
	collection :=database.GetMongoClient().Database("pastebin").Collection("pastes")
	err =database.UpdateOne(collection,filter,pastebin)
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "paste updated successfully"})

}

