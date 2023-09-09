package main

import (
	"github.com/This-Is-Prince/strategiesGo/api"
	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// All network clients
	clients := utils.NewClients()

	// Create a new Gin router
	r := gin.Default()

	// All routes
	api.SetupRoutes(r, clients)

	// Run the server on port 8080
	r.Run(":8080")
}
