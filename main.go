package main

import (
	"crud/database"
	"crud/routes"
	_"crud/docs"


	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

)

func init() {
	database.ConnectToDB()
}

// @title Pet API documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /
func main() {

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/create-pet", routes.CreateAnimal)
	router.GET("/pet/id/:id", routes.FindByID)
	router.GET("/pet/status/:status", routes.FindByStatus)
	router.DELETE("/pet/:id", routes.DeleteByID)
	router.PUT("/pet/fullupdate/:id", routes.FullUpdate)
	router.PATCH("/pet/partialupdate/:id", routes.PartialUpdate)

	router.Run()
}
