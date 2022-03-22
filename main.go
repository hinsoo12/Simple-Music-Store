package main

import (
	"simple-music-store/models"

	"github.com/gin-gonic/gin"

	"simple-music-store/controllers"
)

func main() {
	router := gin.Default()

	/*
		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"sample": "Golang Intership Program"})
		})
	*/

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/songs", controllers.FindSongs)
	router.GET("/songs/:id", controllers.FindSong)
	router.POST("/songs", controllers.AddSong)
	router.PATCH("/songs/:id", controllers.UpdateSong)
	router.DELETE("/songs/:id", controllers.DeleteSong)

	// Run the server
	router.Run()

}
