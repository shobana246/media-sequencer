package main

import (
	"os"
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

	// Use official cors middleware allowing all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
