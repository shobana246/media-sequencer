package main

import (
	"media-sequencer/config"
	"media-sequencer/routes"
	"media-sequencer/seed"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	seed.SeedData(db)

	router := gin.Default()

	// add this
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	routes.SetupRoutes(router, db)
	router.Run(":8080")
}
