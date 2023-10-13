package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStrategy(c *gin.Context) {
	strategyID := c.Param("strategy_id")
	isSubstrate := c.Query("substrate")
	strategiesFileName := "strategies.json"
	if isSubstrate == "true" {
		strategiesFileName = "substrate-strategies.json"
	}
	data, err := ioutil.ReadFile(strategiesFileName)
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

	for _, strategyJSON := range strategiesJSON {
		strategy := strategyJSON.(map[string]interface{})
		if strategy["id"] == strategyID {
			c.JSON(http.StatusOK, strategy)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintln("Strategy not found")})
}
