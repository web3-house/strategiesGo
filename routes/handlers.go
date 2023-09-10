package routes

import (
	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(clients *utils.Clients) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/scores", func(c *gin.Context) {
		GetScores(c, clients)
	})
	return r
}
