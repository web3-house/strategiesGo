package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStrategies(c *gin.Context) {
	data, err := ioutil.ReadFile("strategies.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error %s", err.Error())})
		return
	}

	// Parse JSON into Go struct
	var strategiesJSON []interface{}
	err = json.Unmarshal(data, &strategiesJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, strategiesJSON)
}
