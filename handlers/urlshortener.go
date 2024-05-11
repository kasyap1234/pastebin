package handlers

import (
	"net/http"
	"time"
    "fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/pastebin/database"
	"github.com/kasyap1234/pastebin/models"
	

	

	"github.com/kasyap1234/pastebin/internal"
)
func CreateShortURL (c *gin.Context){

shortURL,err  :=internal.GenerateShortURL(c.Param("longURL"))
if err :=c.ShouldBindJSON(&shortURL); err !=nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return 
}
newURL :=models.URL{
	
	ShortURL: shortURL,
	LongURL: c.Param("longURL"),
	CreatedAt: time.Now(),
}

collection :=database.GetMongoClient().Database("pastebin").Collection("urlshortener"); 

  err =database.InsertOne(collection,shortURL); 
  if err !=nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
  }
  c.JSON(http.StatusOK,newURL.ShortURL)
  fmt.Println("new url created successfully")

  
  
}