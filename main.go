package main

import (
	"crud/database"
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.POST("/create-pet", routes.CreateAnimal)
	router.GET("/pet/id/:id", routes.FindByID)
	router.GET("/pet/status/:status", routes.FindByStatus)
	router.DELETE("/pet/:id", routes.DeleteByID)
	router.PUT("/pet/fullupdate/:id", routes.FullUpdate)
	router.PATCH("/pet/partialupdate/:id", routes.PartialUpdate)

	router.Run()
}
