package routes

import (
	"fmt"
	"net/http"

	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/gin-gonic/gin"
)

type SnapshotsBody struct {
	Networks []string `json:"networks"`
}

func GetSnapshots(c *gin.Context, clients *utils.Clients) {
	// Create a slice of strategies.Strategy to store the JSON data
	var body SnapshotsBody

	// Bind the JSON data from the request body to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error invalid request body %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, body)
}
