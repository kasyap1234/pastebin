package main 
import (
	"fmt"
    "github.com/kasyap1234/pastebin/handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/kasyap1234/pastebin/database"
	

)

func main(){
 	database.ConnectDB();
 r :=gin.Default(); 
 r.GET("/getPasteBYID/:ID",handlers.GetPasteByID)
 r.POST("/createPaste",handlers.CreatePaste)
 r.PUT("/updatePaste/:ID",handlers.UpdatePasteByID)
 r.DELETE("/deletePaste/:ID",handlers.DeletePasteByID)
 r.POST("/createShortURL/:longURL",handlers.CreateShortURL)


 r.Run(":8080")
  
fmt.Println("server started successfully and is running on port 8080  ");

}
