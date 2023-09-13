package routes

import (
	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(clients *utils.Clients) *gin.Engine {
	r := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Add your allowed origins here
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/snapshots", func(c *gin.Context) {
		GetSnapshots(c, clients)
	})
	r.GET("/strategies", func(c *gin.Context) {
		GetStrategies(c)
	})
	r.GET("/strategies/:strategy_id", func(c *gin.Context) {
		GetStrategy(c)
	})
	r.POST("/scores", func(c *gin.Context) {
		GetScores(c, clients)
	})
	return r
}
