package main 
import (
	"fmt"
    "github.com/kasyap1234/pastebin/handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/kasyap1234/pastebin/database"
	"github.com/swaggo/gin-swagger"

)

// @title TODO APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /api/v1
// @schemes http

// @schemes http
func main(){
 	database.ConnectDB();
 r :=gin.Default(); 
 r.GET("/getPasteBYID/:ID",handlers.GetPasteByID)
 r.POST("/createPaste",handlers.CreatePaste)
 r.PUT("/updatePaste/:ID",handlers.UpdatePasteByID)
 r.DELETE("/deletePaste/:ID",handlers.DeletePasteByID)
 r.POST("/createShortURL/:longURL",handlers.CreateShortURL)
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

 r.Run(":8080")
  
fmt.Println("server started successfully and is running on port 8080  ");

}
