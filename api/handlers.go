package api

import (
	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, clients *utils.Clients) {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.POST("/scores", func(c *gin.Context) {
		GetScores(c, clients)
	})
}
