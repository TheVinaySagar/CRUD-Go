package main

import (
	"log"
	"myproject/database"
	"myproject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.ConnectDB()
	// config.Credentials()
	// Create Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)
	routes.RegisterInterview(r)
	routes.CloudinaryRoute(r)

	// Start server
	log.Println("Server running on port 8000")
	r.Run(":8000")
}
