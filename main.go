package main

import (
	"github.com/This-Is-Prince/strategiesGo/routes"
	"github.com/This-Is-Prince/strategiesGo/utils"
)

func main() {
	// All network clients
	clients := utils.NewClients()

	// All routes
	r := routes.SetupRoutes(clients)

	// Run the server on port 8080
	r.Run(":8080")
}
