package routes

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/This-Is-Prince/strategiesGo/strategies"
	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/gin-gonic/gin"
)

type VotingPowerBody struct {
	Address    string                `json:"address"`
	Strategies []strategies.Strategy `json:"strategies"`
}

func GetVotingPower(c *gin.Context, clients *utils.Clients) {
	// Create a slice of strategies.Strategy to store the JSON data
	var body VotingPowerBody

	// Bind the JSON data from the request body to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error invalid request body %s", err.Error())})
		return
	}

	votingPower := []*big.Float{}
	for _, strategy := range body.Strategies {
		votingPower = append(votingPower, strategy.Score(clients, body.Address))
	}

	c.JSON(http.StatusOK, votingPower)
}
